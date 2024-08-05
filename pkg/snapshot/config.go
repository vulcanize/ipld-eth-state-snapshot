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
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/sirupsen/logrus"

	"github.com/cerc-io/plugeth-statediff/indexer/database/file"
	"github.com/cerc-io/plugeth-statediff/indexer/database/sql/postgres"
	ethNode "github.com/cerc-io/plugeth-statediff/indexer/node"
	"github.com/spf13/viper"
)

// SnapshotMode specifies the snapshot data output method
type SnapshotMode string

const (
	PgSnapshot   SnapshotMode = "postgres"
	FileSnapshot SnapshotMode = "file"

	defaultOutputDir = "./snapshot_output"
)

// Config contains params for both databases the service uses
type Config struct {
	Eth     *EthDBConfig
	DB      *DBConfig
	File    *FileConfig
	Service *ServiceConfig
}

// EthDBConfig is config parameters for the chain DB.
type EthDBConfig struct {
	DBPath        string
	AncientDBPath string
	NodeInfo      ethNode.Info
}

// DBConfig contains options for DB output mode.
type DBConfig = postgres.Config

// FileConfig contains options for file output mode.  Note that this service currently only supports
// CSV output, and does not record watched addresses, so not all fields are used.
type FileConfig = file.Config

type ServiceConfig struct {
	AllowedAccounts []common.Address
}

func NewConfig(mode SnapshotMode) (*Config, error) {
	ret := &Config{
		&EthDBConfig{},
		&DBConfig{},
		&FileConfig{},
		&ServiceConfig{},
	}
	return ret, ret.Init(mode)
}

func NewInPlaceSnapshotConfig() *Config {
	ret := &Config{
		&EthDBConfig{},
		&DBConfig{},
		&FileConfig{},
		&ServiceConfig{},
	}
	InitDB(ret.DB)

	return ret
}

// Init Initialises config
func (c *Config) Init(mode SnapshotMode) error {
	viper.BindEnv(LOG_FILE_TOML, LOG_FILE)
	viper.BindEnv(ETH_NODE_ID_TOML, ETH_NODE_ID)
	viper.BindEnv(ETH_CLIENT_NAME_TOML, ETH_CLIENT_NAME)
	viper.BindEnv(ETH_GENESIS_BLOCK_TOML, ETH_GENESIS_BLOCK)
	viper.BindEnv(ETH_NETWORK_ID_TOML, ETH_NETWORK_ID)
	viper.BindEnv(ETH_CHAIN_ID_TOML, ETH_CHAIN_ID)

	c.Eth.NodeInfo = ethNode.Info{
		ID:           viper.GetString(ETH_NODE_ID_TOML),
		ClientName:   viper.GetString(ETH_CLIENT_NAME_TOML),
		GenesisBlock: viper.GetString(ETH_GENESIS_BLOCK_TOML),
		NetworkID:    viper.GetString(ETH_NETWORK_ID_TOML),
		ChainID:      viper.GetUint64(ETH_CHAIN_ID_TOML),
	}

	viper.BindEnv(ETHDB_ANCIENT_TOML, ETHDB_ANCIENT)
	viper.BindEnv(ETHDB_PATH_TOML, ETHDB_PATH)

	c.Eth.DBPath = viper.GetString(ETHDB_PATH_TOML)
	c.Eth.AncientDBPath = viper.GetString(ETHDB_ANCIENT_TOML)
	if len(c.Eth.AncientDBPath) == 0 {
		c.Eth.AncientDBPath = c.Eth.DBPath + "/ancient"
	}

	switch mode {
	case FileSnapshot:
		InitFile(c.File)
	case PgSnapshot:
		InitDB(c.DB)
	default:
		return fmt.Errorf("no output mode specified")
	}
	return c.Service.Init()
}

func InitDB(c *DBConfig) {
	viper.BindEnv(DATABASE_NAME_TOML, DATABASE_NAME)
	viper.BindEnv(DATABASE_HOSTNAME_TOML, DATABASE_HOSTNAME)
	viper.BindEnv(DATABASE_PORT_TOML, DATABASE_PORT)
	viper.BindEnv(DATABASE_USER_TOML, DATABASE_USER)
	viper.BindEnv(DATABASE_PASSWORD_TOML, DATABASE_PASSWORD)
	viper.BindEnv(DATABASE_MAX_IDLE_CONNECTIONS_TOML, DATABASE_MAX_IDLE_CONNECTIONS)
	viper.BindEnv(DATABASE_MAX_OPEN_CONNECTIONS_TOML, DATABASE_MAX_OPEN_CONNECTIONS)
	viper.BindEnv(DATABASE_MAX_CONN_LIFETIME_TOML, DATABASE_MAX_CONN_LIFETIME)

	// DB params
	c.DatabaseName = viper.GetString(DATABASE_NAME_TOML)
	c.Hostname = viper.GetString(DATABASE_HOSTNAME_TOML)
	c.Port = viper.GetInt(DATABASE_PORT_TOML)
	c.Username = viper.GetString(DATABASE_USER_TOML)
	c.Password = viper.GetString(DATABASE_PASSWORD_TOML)
	// Connection config
	c.MaxIdle = viper.GetInt(DATABASE_MAX_IDLE_CONNECTIONS_TOML)
	c.MaxConns = viper.GetInt(DATABASE_MAX_OPEN_CONNECTIONS_TOML)
	c.MaxConnLifetime = time.Duration(viper.GetInt(DATABASE_MAX_CONN_LIFETIME_TOML)) * time.Second

	c.Driver = postgres.SQLX
}

func InitFile(c *FileConfig) error {
	viper.BindEnv(FILE_OUTPUT_DIR_TOML, FILE_OUTPUT_DIR)
	c.OutputDir = viper.GetString(FILE_OUTPUT_DIR_TOML)
	if c.OutputDir == "" {
		logrus.Infof("no output directory set, using default: %s", defaultOutputDir)
		c.OutputDir = defaultOutputDir
	}
	// Only support CSV for now
	c.Mode = file.CSV
	return nil
}

func (c *ServiceConfig) Init() error {
	viper.BindEnv(SNAPSHOT_BLOCK_HEIGHT_TOML, SNAPSHOT_BLOCK_HEIGHT)
	viper.BindEnv(SNAPSHOT_MODE_TOML, SNAPSHOT_MODE)
	viper.BindEnv(SNAPSHOT_WORKERS_TOML, SNAPSHOT_WORKERS)
	viper.BindEnv(SNAPSHOT_RECOVERY_FILE_TOML, SNAPSHOT_RECOVERY_FILE)

	viper.BindEnv(PROM_DB_STATS_TOML, PROM_DB_STATS)
	viper.BindEnv(PROM_HTTP_TOML, PROM_HTTP)
	viper.BindEnv(PROM_HTTP_ADDR_TOML, PROM_HTTP_ADDR)
	viper.BindEnv(PROM_HTTP_PORT_TOML, PROM_HTTP_PORT)
	viper.BindEnv(PROM_METRICS_TOML, PROM_METRICS)

	viper.BindEnv(SNAPSHOT_ACCOUNTS_TOML, SNAPSHOT_ACCOUNTS)
	var allowedAccounts []string
	viper.UnmarshalKey(SNAPSHOT_ACCOUNTS_TOML, &allowedAccounts)
	accountsLen := len(allowedAccounts)
	if accountsLen != 0 {
		c.AllowedAccounts = make([]common.Address, 0, accountsLen)
		for _, allowedAccount := range allowedAccounts {
			c.AllowedAccounts = append(c.AllowedAccounts, common.HexToAddress(allowedAccount))
		}
	} else {
		logrus.Infof("no snapshot addresses specified, will perform snapshot of entire trie(s)")
	}
	return nil
}
