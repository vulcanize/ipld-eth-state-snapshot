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

// ENV variables
const (
	SNAPSHOT_BLOCK_HEIGHT  = "SNAPSHOT_BLOCK_HEIGHT"
	SNAPSHOT_WORKERS       = "SNAPSHOT_WORKERS"
	SNAPSHOT_RECOVERY_FILE = "SNAPSHOT_RECOVERY_FILE"
	SNAPSHOT_MODE          = "SNAPSHOT_MODE"
	SNAPSHOT_ACCOUNTS      = "SNAPSHOT_ACCOUNTS"

	LOG_LEVEL = "LOG_LEVEL"
	LOG_FILE  = "LOG_FILE"

	PROM_METRICS   = "PROM_METRICS"
	PROM_HTTP      = "PROM_HTTP"
	PROM_HTTP_ADDR = "PROM_HTTP_ADDR"
	PROM_HTTP_PORT = "PROM_HTTP_PORT"
	PROM_DB_STATS  = "PROM_DB_STATS"

	FILE_OUTPUT_DIR = "FILE_OUTPUT_DIR"

	ETHDB_ANCIENT = "ETHDB_ANCIENT"
	ETHDB_PATH    = "ETHDB_PATH"

	ETH_CLIENT_NAME   = "ETH_CLIENT_NAME"
	ETH_GENESIS_BLOCK = "ETH_GENESIS_BLOCK"
	ETH_NETWORK_ID    = "ETH_NETWORK_ID"
	ETH_NODE_ID       = "ETH_NODE_ID"
	ETH_CHAIN_ID      = "ETH_CHAIN_ID"

	DATABASE_NAME                 = "DATABASE_NAME"
	DATABASE_HOSTNAME             = "DATABASE_HOSTNAME"
	DATABASE_PORT                 = "DATABASE_PORT"
	DATABASE_USER                 = "DATABASE_USER"
	DATABASE_PASSWORD             = "DATABASE_PASSWORD"
	DATABASE_MAX_IDLE_CONNECTIONS = "DATABASE_MAX_IDLE_CONNECTIONS"
	DATABASE_MAX_OPEN_CONNECTIONS = "DATABASE_MAX_OPEN_CONNECTIONS"
	DATABASE_MAX_CONN_LIFETIME    = "DATABASE_MAX_CONN_LIFETIME"
)

// TOML bindings
const (
	SNAPSHOT_BLOCK_HEIGHT_TOML  = "snapshot.blockHeight"
	SNAPSHOT_WORKERS_TOML       = "snapshot.workers"
	SNAPSHOT_RECOVERY_FILE_TOML = "snapshot.recoveryFile"
	SNAPSHOT_MODE_TOML          = "snapshot.mode"
	SNAPSHOT_ACCOUNTS_TOML      = "snapshot.accounts"

	LOG_LEVEL_TOML = "log.level"
	LOG_FILE_TOML  = "log.file"

	PROM_METRICS_TOML   = "prom.metrics"
	PROM_HTTP_TOML      = "prom.http"
	PROM_HTTP_ADDR_TOML = "prom.httpAddr"
	PROM_HTTP_PORT_TOML = "prom.httpPort"
	PROM_DB_STATS_TOML  = "prom.dbStats"

	FILE_OUTPUT_DIR_TOML = "file.outputDir"

	ETHDB_ANCIENT_TOML = "ethdb.ancient"
	ETHDB_PATH_TOML    = "ethdb.path"

	ETH_CLIENT_NAME_TOML   = "ethereum.clientName"
	ETH_GENESIS_BLOCK_TOML = "ethereum.genesisBlock"
	ETH_NETWORK_ID_TOML    = "ethereum.networkID"
	ETH_NODE_ID_TOML       = "ethereum.nodeID"
	ETH_CHAIN_ID_TOML      = "ethereum.chainID"

	DATABASE_NAME_TOML                 = "database.name"
	DATABASE_HOSTNAME_TOML             = "database.hostname"
	DATABASE_PORT_TOML                 = "database.port"
	DATABASE_USER_TOML                 = "database.user"
	DATABASE_PASSWORD_TOML             = "database.password"
	DATABASE_MAX_IDLE_CONNECTIONS_TOML = "database.maxIdle"
	DATABASE_MAX_OPEN_CONNECTIONS_TOML = "database.maxOpen"
	DATABASE_MAX_CONN_LIFETIME_TOML    = "database.maxLifetime"
)

// CLI flags
const (
	SNAPSHOT_BLOCK_HEIGHT_CLI  = "block-height"
	SNAPSHOT_WORKERS_CLI       = "workers"
	SNAPSHOT_RECOVERY_FILE_CLI = "recovery-file"
	SNAPSHOT_MODE_CLI          = "snapshot-mode"
	SNAPSHOT_ACCOUNTS_CLI      = "snapshot-accounts"

	LOG_LEVEL_CLI = "log-level"
	LOG_FILE_CLI  = "log-file"

	PROM_METRICS_CLI   = "prom-metrics"
	PROM_HTTP_CLI      = "prom-http"
	PROM_HTTP_ADDR_CLI = "prom-httpAddr"
	PROM_HTTP_PORT_CLI = "prom-httpPort"
	PROM_DB_STATS_CLI  = "prom-dbStats"

	FILE_OUTPUT_DIR_CLI = "output-dir"

	ETHDB_ANCIENT_CLI = "ancient-path"
	ETHDB_PATH_CLI    = "ethdb-path"

	ETH_CLIENT_NAME_CLI   = "ethereum-client-name"
	ETH_GENESIS_BLOCK_CLI = "ethereum-genesis-block"
	ETH_NETWORK_ID_CLI    = "ethereum-network-id"
	ETH_NODE_ID_CLI       = "ethereum-node-id"
	ETH_CHAIN_ID_CLI      = "ethereum-chain-id"

	DATABASE_NAME_CLI                 = "database-name"
	DATABASE_HOSTNAME_CLI             = "database-hostname"
	DATABASE_PORT_CLI                 = "database-port"
	DATABASE_USER_CLI                 = "database-user"
	DATABASE_PASSWORD_CLI             = "database-password"
	DATABASE_MAX_IDLE_CONNECTIONS_CLI = "database-max-idle"
	DATABASE_MAX_OPEN_CONNECTIONS_CLI = "database-max-open"
	DATABASE_MAX_CONN_LIFETIME_CLI    = "database-max-lifetime"
)
