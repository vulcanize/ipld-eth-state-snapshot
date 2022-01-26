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
	ethNode "github.com/ethereum/go-ethereum/statediff/indexer/node"
	"github.com/ethereum/go-ethereum/statediff/indexer/postgres"
	"github.com/spf13/viper"
)

const (
	ANCIENT_DB_PATH   = "ANCIENT_DB_PATH"
	ETH_CLIENT_NAME   = "ETH_CLIENT_NAME"
	ETH_GENESIS_BLOCK = "ETH_GENESIS_BLOCK"
	ETH_NETWORK_ID    = "ETH_NETWORK_ID"
	ETH_NODE_ID       = "ETH_NODE_ID"
	LVL_DB_PATH       = "LVL_DB_PATH"
)

// Config contains params for both databases the service uses
type Config struct {
	DB  *DBConfig
	Eth *EthConfig
}

// DBConfig is config parameters for DB.
type DBConfig struct {
	Node       ethNode.Info
	URI        string
	ConnConfig postgres.ConnectionConfig
}

// EthConfig is config parameters for the chain.
type EthConfig struct {
	LevelDBPath   string
	AncientDBPath string
}

func NewConfig() *Config {
	ret := &Config{&DBConfig{}, &EthConfig{}}
	ret.Init()
	return ret
}

// Init Initialises config
func (c *Config) Init() {
	c.DB.dbInit()
	viper.BindEnv("ethereum.nodeID", ETH_NODE_ID)
	viper.BindEnv("ethereum.clientName", ETH_CLIENT_NAME)
	viper.BindEnv("ethereum.genesisBlock", ETH_GENESIS_BLOCK)
	viper.BindEnv("ethereum.networkID", ETH_NETWORK_ID)

	c.DB.Node = ethNode.Info{
		ID:           viper.GetString("ethereum.nodeID"),
		ClientName:   viper.GetString("ethereum.clientName"),
		GenesisBlock: viper.GetString("ethereum.genesisBlock"),
		NetworkID:    viper.GetString("ethereum.networkID"),
		ChainID:      viper.GetUint64("ethereum.chainID"),
	}

	viper.BindEnv("leveldb.ancient", ANCIENT_DB_PATH)
	viper.BindEnv("leveldb.path", LVL_DB_PATH)

	c.Eth.AncientDBPath = viper.GetString("leveldb.ancient")
	c.Eth.LevelDBPath = viper.GetString("leveldb.path")
}

func (c *DBConfig) dbInit() {
	viper.BindEnv("database.name", postgres.DATABASE_NAME)
	viper.BindEnv("database.hostname", postgres.DATABASE_HOSTNAME)
	viper.BindEnv("database.port", postgres.DATABASE_PORT)
	viper.BindEnv("database.user", postgres.DATABASE_USER)
	viper.BindEnv("database.password", postgres.DATABASE_PASSWORD)
	viper.BindEnv("database.maxIdle", postgres.DATABASE_MAX_IDLE_CONNECTIONS)
	viper.BindEnv("database.maxOpen", postgres.DATABASE_MAX_OPEN_CONNECTIONS)
	viper.BindEnv("database.maxLifetime", postgres.DATABASE_MAX_CONN_LIFETIME)

	dbParams := postgres.ConnectionParams{}
	// DB params
	dbParams.Name = viper.GetString("database.name")
	dbParams.Hostname = viper.GetString("database.hostname")
	dbParams.Port = viper.GetInt("database.port")
	dbParams.User = viper.GetString("database.user")
	dbParams.Password = viper.GetString("database.password")

	c.URI = postgres.DbConnectionString(dbParams)
	// Connection config
	c.ConnConfig.MaxIdle = viper.GetInt("database.maxIdle")
	c.ConnConfig.MaxOpen = viper.GetInt("database.maxOpen")
	c.ConnConfig.MaxLifetime = viper.GetInt("database.maxLifetime")
}
