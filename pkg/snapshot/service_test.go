package snapshot_test

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"testing"
	"time"

	"github.com/cerc-io/eth-testing/chains"
	"github.com/cerc-io/plugeth-statediff/indexer/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/cerc-io/ipld-eth-state-snapshot/internal/mocks"
	. "github.com/cerc-io/ipld-eth-state-snapshot/pkg/snapshot"
	fixture "github.com/cerc-io/ipld-eth-state-snapshot/test"
)

var (
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Note: block 1 doesn't have storage nodes. TODO: add fixtures with storage nodes
	// chainAblock1StateKeys = sliceToSet(fixture.ChainA_Block1_StateNodeLeafKeys)
	chainAblock1IpldCids = sliceToSet(fixture.ChainA_Block1_IpldCids)

	subtrieWorkerCases = []uint{1, 4, 8, 16, 32}
)

type selectiveData struct {
	StateNodes   map[string]*models.StateNodeModel
	StorageNodes map[string]map[string]*models.StorageNodeModel
}

func testConfig(ethdbpath, ancientdbpath string) *Config {
	return &Config{
		Eth: &EthDBConfig{
			DBPath:        ethdbpath,
			AncientDBPath: ancientdbpath,
			NodeInfo:      DefaultNodeInfo,
		},
		DB: &DefaultPgConfig,
	}
}

func TestSnapshot(t *testing.T) {
	runCase := func(t *testing.T, workers uint) {
		params := SnapshotParams{Height: 1, Workers: workers}
		data := doSnapshot(t, fixture.ChainA, params)
		verify_chainAblock1(t, data)
	}

	for _, tc := range subtrieWorkerCases {
		t.Run(fmt.Sprintf("with %d subtries", tc), func(t *testing.T) { runCase(t, tc) })
	}
}

func TestAccountSelectiveSnapshot(t *testing.T) {
	height := uint64(32)
	watchedAddresses, expected := watchedAccountData_chainBblock32()

	runCase := func(t *testing.T, workers uint) {
		params := SnapshotParams{
			Height:           height,
			Workers:          workers,
			WatchedAddresses: watchedAddresses,
		}
		data := doSnapshot(t, fixture.ChainB, params)
		expected.verify(t, data)
	}

	for _, tc := range subtrieWorkerCases {
		t.Run(fmt.Sprintf("with %d subtries", tc), func(t *testing.T) { runCase(t, tc) })
	}
}

func TestSnapshotRecovery(t *testing.T) {
	runCase := func(t *testing.T, workers uint, interruptAt uint) {
		params := SnapshotParams{Height: 1, Workers: workers}
		data := doSnapshotWithRecovery(t, fixture.ChainA, params, interruptAt)
		verify_chainAblock1(t, data)
	}

	interrupts := make([]uint, 4)
	for i := 0; i < len(interrupts); i++ {
		N := len(fixture.ChainA_Block1_StateNodeLeafKeys)
		interrupts[i] = uint(rand.Intn(N/2) + N/4)
	}

	for _, tc := range subtrieWorkerCases {
		for i, interrupt := range interrupts {
			t.Run(
				fmt.Sprintf("with %d subtries %d", tc, i),
				func(t *testing.T) { runCase(t, tc, interrupt) },
			)
		}
	}
}

func TestAccountSelectiveSnapshotRecovery(t *testing.T) {
	height := uint64(32)
	watchedAddresses, expected := watchedAccountData_chainBblock32()

	runCase := func(t *testing.T, workers uint, interruptAt uint) {
		params := SnapshotParams{
			Height:           height,
			Workers:          workers,
			WatchedAddresses: watchedAddresses,
		}
		data := doSnapshotWithRecovery(t, fixture.ChainB, params, interruptAt)
		expected.verify(t, data)
	}

	for _, tc := range subtrieWorkerCases {
		t.Run(
			fmt.Sprintf("with %d subtries", tc),
			func(t *testing.T) { runCase(t, tc, 1) },
		)
	}
}

func verify_chainAblock1(t *testing.T, data mocks.IndexerData) {
	// Extract indexed keys and sort them for comparison
	var indexedStateKeys []string
	for _, stateNode := range data.StateNodes {
		stateKey := common.BytesToHash(stateNode.AccountWrapper.LeafKey).String()
		indexedStateKeys = append(indexedStateKeys, stateKey)
	}
	require.ElementsMatch(t, fixture.ChainA_Block1_StateNodeLeafKeys, indexedStateKeys)

	ipldCids := make(map[string]struct{})
	for _, ipld := range data.IPLDs {
		ipldCids[ipld.CID] = struct{}{}
	}
	require.Equal(t, chainAblock1IpldCids, ipldCids)
}

func watchedAccountData_chainBblock32() ([]common.Address, selectiveData) {
	watchedAddresses := []common.Address{
		// hash 0xcabc5edb305583e33f66322ceee43088aa99277da772feb5053512d03a0a702b
		common.HexToAddress("0x825a6eec09e44Cb0fa19b84353ad0f7858d7F61a"),
		// hash 0x33153abc667e873b6036c8a46bdd847e2ade3f89b9331c78ef2553fea194c50d
		common.HexToAddress("0x0616F59D291a898e796a1FAD044C5926ed2103eC"),
	}
	var expected selectiveData
	expected.StateNodes = make(map[string]*models.StateNodeModel)
	for _, index := range []int{0, 4} {
		node := &fixture.ChainB_Block32_StateNodes[index]
		expected.StateNodes[node.StateKey] = node
	}

	// Map account leaf keys to corresponding storage
	expectedStorageNodeIndexes := []struct {
		address common.Address
		indexes []int
	}{
		{watchedAddresses[0], []int{9, 11}},
		{watchedAddresses[1], []int{0, 1, 2, 4, 6}},
	}
	expected.StorageNodes = make(map[string]map[string]*models.StorageNodeModel)
	for _, account := range expectedStorageNodeIndexes {
		leafKey := crypto.Keccak256Hash(account.address[:]).String()
		storageNodes := make(map[string]*models.StorageNodeModel)
		for _, index := range account.indexes {
			node := &fixture.ChainB_Block32_StorageNodes[index]
			storageNodes[node.StorageKey] = node
		}
		expected.StorageNodes[leafKey] = storageNodes
	}
	return watchedAddresses, expected
}

func (expected selectiveData) verify(t *testing.T, data mocks.IndexerData) {
	// check that all indexed nodes are expected and correct
	indexedStateKeys := make(map[string]struct{})
	for _, stateNode := range data.StateNodes {
		stateKey := common.BytesToHash(stateNode.AccountWrapper.LeafKey).String()
		indexedStateKeys[stateKey] = struct{}{}
		require.Contains(t, expected.StateNodes, stateKey, "unexpected state node")

		model := expected.StateNodes[stateKey]
		require.Equal(t, model.CID, stateNode.AccountWrapper.CID)
		require.Equal(t, model.Balance, stateNode.AccountWrapper.Account.Balance.String())
		require.Equal(t, model.StorageRoot, stateNode.AccountWrapper.Account.Root.String())

		expectedStorage := expected.StorageNodes[stateKey]
		indexedStorageKeys := make(map[string]struct{})
		for _, storageNode := range stateNode.StorageDiff {
			storageKey := common.BytesToHash(storageNode.LeafKey).String()
			indexedStorageKeys[storageKey] = struct{}{}
			require.Contains(t, expectedStorage, storageKey, "unexpected storage node")

			require.Equal(t, expectedStorage[storageKey].CID, storageNode.CID)
			require.Equal(t, expectedStorage[storageKey].Value, storageNode.Value)
		}
		// check for completeness
		for storageNode := range expectedStorage {
			require.Contains(t, indexedStorageKeys, storageNode, "missing storage node")
		}
	}
	// check for completeness
	for stateNode := range expected.StateNodes {
		require.Contains(t, indexedStateKeys, stateNode, "missing state node")
	}
}

func doSnapshot(t *testing.T, chain *chains.Paths, params SnapshotParams) mocks.IndexerData {
	chainDataPath, ancientDataPath := chain.ChainData, chain.Ancient
	config := testConfig(chainDataPath, ancientDataPath)
	edb, err := NewEthDB(config.Eth)
	require.NoError(t, err)
	defer edb.Close()

	idx := mocks.NewIndexer(t)
	recovery := filepath.Join(t.TempDir(), "recover.csv")
	service, err := NewSnapshotService(edb, idx, recovery)
	require.NoError(t, err)

	err = service.CreateSnapshot(params)
	require.NoError(t, err)
	return idx.IndexerData
}

func doSnapshotWithRecovery(
	t *testing.T,
	chain *chains.Paths,
	params SnapshotParams,
	failAfter uint,
) mocks.IndexerData {
	chainDataPath, ancientDataPath := chain.ChainData, chain.Ancient
	config := testConfig(chainDataPath, ancientDataPath)
	edb, err := NewEthDB(config.Eth)
	require.NoError(t, err)
	defer edb.Close()

	indexer := &mocks.InterruptingIndexer{
		Indexer:        mocks.NewIndexer(t),
		InterruptAfter: failAfter,
	}
	t.Logf("Will interrupt after %d state nodes", failAfter)

	recoveryFile := filepath.Join(t.TempDir(), "recover.csv")
	service, err := NewSnapshotService(edb, indexer, recoveryFile)
	require.NoError(t, err)
	err = service.CreateSnapshot(params)
	require.Error(t, err)

	require.FileExists(t, recoveryFile)
	// We should only have processed nodes up to the break, plus an extra node per worker
	require.LessOrEqual(t, len(indexer.StateNodes), int(indexer.InterruptAfter+params.Workers))

	// use the nested mock indexer, to continue where it left off
	recoveryIndexer := indexer.Indexer
	service, err = NewSnapshotService(edb, recoveryIndexer, recoveryFile)
	require.NoError(t, err)
	err = service.CreateSnapshot(params)
	require.NoError(t, err)

	return recoveryIndexer.IndexerData
}

func sliceToSet[T comparable](slice []T) map[T]struct{} {
	set := make(map[T]struct{})
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}
