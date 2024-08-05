// Copyright © 2020 Vulcanize, Inc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package snapshot

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/cerc-io/ipld-eth-state-snapshot/pkg/prom"
	statediff "github.com/cerc-io/plugeth-statediff"
	"github.com/cerc-io/plugeth-statediff/adapt"
	"github.com/cerc-io/plugeth-statediff/indexer"
	"github.com/cerc-io/plugeth-statediff/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
)

var (
	emptyNode, _      = rlp.EncodeToBytes(&[]byte{})
	emptyCodeHash     = crypto.Keccak256([]byte{})
	emptyContractRoot = crypto.Keccak256Hash(emptyNode)

	defaultBatchSize = uint(100)
)

// Service holds ethDB and stateDB to read data from lvldb and Publisher
// to publish trie in postgres DB.
type Service struct {
	ethDB        ethdb.Database
	stateDB      state.Database
	indexer      indexer.Indexer
	maxBatchSize uint
	recoveryFile string
}

func NewEthDB(con *EthDBConfig) (ethdb.Database, error) {
	return rawdb.Open(rawdb.OpenOptions{
		Directory:         con.DBPath,
		AncientsDirectory: con.AncientDBPath,
		Namespace:         "ipld-eth-state-snapshot",
		Cache:             1024,
		Handles:           256,
		ReadOnly:          true,
	})
}

// NewSnapshotService creates Service.
func NewSnapshotService(edb ethdb.Database, indexer indexer.Indexer, recoveryFile string) (*Service, error) {
	return &Service{
		ethDB:        edb,
		stateDB:      state.NewDatabase(edb),
		indexer:      indexer,
		maxBatchSize: defaultBatchSize,
		recoveryFile: recoveryFile,
	}, nil
}

type SnapshotParams struct {
	WatchedAddresses []common.Address
	Height           uint64
	Workers          uint
}

func (s *Service) CreateSnapshot(params SnapshotParams) error {
	// extract header from lvldb and publish to PG-IPFS
	// hold onto the headerID so that we can link the state nodes to this header
	hash := rawdb.ReadCanonicalHash(s.ethDB, params.Height)
	header := rawdb.ReadHeader(s.ethDB, hash, params.Height)
	if header == nil {
		return fmt.Errorf("unable to read canonical header at height %d", params.Height)
	}
	log.WithField("height", params.Height).WithField("hash", hash).Info("Creating snapshot")

	// Context for snapshot work
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()
	// Cancel context on receiving a signal. On cancellation, all tracked iterators complete
	// processing of their current node before stopping.
	captureSignal(cancelCtx)

	var err error
	tx := s.indexer.BeginTx(header.Number, ctx)
	defer tx.RollbackOnFailure(err)

	var headerid string
	headerid, err = s.indexer.PushHeader(tx, header, big.NewInt(0), big.NewInt(0))
	if err != nil {
		return err
	}

	tr := prom.NewTracker(s.recoveryFile, params.Workers)
	defer func() {
		err := tr.CloseAndSave()
		if err != nil {
			log.Errorf("failed to write recovery file: %v", err)
		}
	}()

	var nodeMtx, ipldMtx sync.Mutex
	nodeSink := func(node types.StateLeafNode) error {
		nodeMtx.Lock()
		defer nodeMtx.Unlock()
		prom.IncStateNodeCount()
		prom.AddStorageNodeCount(len(node.StorageDiff))
		return s.indexer.PushStateNode(tx, node, headerid)
	}
	ipldSink := func(c types.IPLD) error {
		ipldMtx.Lock()
		defer ipldMtx.Unlock()
		return s.indexer.PushIPLD(tx, c)
	}

	sdparams := statediff.Params{
		WatchedAddresses: params.WatchedAddresses,
	}
	sdparams.ComputeWatchedAddressesLeafPaths()
	builder := statediff.NewBuilder(adapt.GethStateView(s.stateDB))
	builder.SetSubtrieWorkers(params.Workers)
	if err = builder.WriteStateSnapshot(ctx, header.Root, sdparams, nodeSink, ipldSink, tr); err != nil {
		return err
	}

	if err = tx.Submit(); err != nil {
		return fmt.Errorf("batch transaction submission failed: %w", err)
	}
	return err
}

// CreateLatestSnapshot snapshot at head (ignores height param)
func (s *Service) CreateLatestSnapshot(workers uint, watchedAddresses []common.Address) error {
	log.Info("Creating snapshot at head")
	hash := rawdb.ReadHeadHeaderHash(s.ethDB)
	height := rawdb.ReadHeaderNumber(s.ethDB, hash)
	if height == nil {
		return fmt.Errorf("unable to read header height for header hash %s", hash)
	}
	return s.CreateSnapshot(SnapshotParams{Height: *height, Workers: workers, WatchedAddresses: watchedAddresses})
}

func captureSignal(cb func()) {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		log.Errorf("Signal received (%v), stopping", sig)
		cb()
	}()
}
