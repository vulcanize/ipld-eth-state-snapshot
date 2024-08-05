# ipld-eth-state-snapshot

> Tool for extracting the entire Ethereum state at a particular block height from a cold database into Postgres-backed IPFS

[![Go Report Card](https://goreportcard.com/badge/github.com/vulcanize/ipld-eth-state-snapshot)](https://goreportcard.com/report/github.com/vulcanize/ipld-eth-state-snapshot)

## Setup

* Build the binary:

    ```bash
    make build
    ```

## Configuration

Config format:

```toml
[snapshot]
    mode         = "file"           # indicates output mode <postgres | file>
    workers      = 4                # degree of concurrency: the state trie is subdivided into sections that are traversed and processed concurrently
    blockHeight  = -1               # blockheight to perform the snapshot at (-1 indicates to use the latest blockheight found in ethdb)
    recoveryFile = "recovery_file"  # specifies a file to output recovery information on error or premature closure
    accounts = []                   # list of accounts (addresses) to take the snapshot for # SNAPSHOT_ACCOUNTS

[ethdb]
    # path to geth ethdb
    path    = "/Users/user/Library/Ethereum/geth/chaindata"         # ETHDB_PATH
    # path to geth ancient database
    ancient = "/Users/user/Library/Ethereum/geth/chaindata/ancient" # ETHDB_ANCIENT

[database]
    # when operating in 'postgres' output mode
    # db credentials
    name     = "vulcanize_public"   # DATABASE_NAME
    hostname = "localhost"          # DATABASE_HOSTNAME
    port     = 5432                 # DATABASE_PORT
    user     = "postgres"           # DATABASE_USER
    password = ""                   # DATABASE_PASSWORD

[file]
    # when operating in 'file' output mode
    # directory the CSV files are written to
    outputDir = "output_dir/"   # FILE_OUTPUT_DIR

[log]
    level = "info"      # log level (trace, debug, info, warn, error, fatal, panic) (default: info)
    file  = "log_file"  # file path for logging, leave unset to log to stdout

[prom]
    # prometheus metrics
    metrics  = true         # enable prometheus metrics         (default: false)
    http     = true         # enable prometheus http service    (default: false)
    httpAddr = "0.0.0.0"    # prometheus http host              (default: 127.0.0.1)
    httpPort = 9101         # prometheus http port              (default: 8086)
    dbStats  = true         # enable prometheus db stats        (default: false)

[ethereum]
    # node info
    clientName   = "Geth"   # ETH_CLIENT_NAME
    nodeID       = "arch1"  # ETH_NODE_ID
    networkID    = "1"      # ETH_NETWORK_ID
    chainID      = "1"      # ETH_CHAIN_ID
    genesisBlock = "0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3" # ETH_GENESIS_BLOCK
```

> **Note:** previous versions of this service used different variable names. To update, change the following:
> * `LVL_DB_PATH`, `LEVELDB_PATH` => `ETHDB_PATH`
> * `ANCIENT_DB_PATH`, `LEVELDB_ANCIENT` => `ETHDB_ANCIENT`
> * `LOGRUS_LEVEL`, `LOGRUS_FILE` => `LOG_LEVEL`, `LOG_FILE`, etc.


## Usage

* For state snapshot from EthDB:

    ```bash
    ./ipld-eth-state-snapshot stateSnapshot --config={path to toml config file}
    ```

    * Account selective snapshot: To restrict the snapshot to a list of accounts (addresses), provide the addresses in config parameter `snapshot.accounts` or env variable `SNAPSHOT_ACCOUNTS`. Only nodes related to provided addresses will be indexed.

        Example:

        ```toml
        [snapshot]
            accounts = [
                "0x825a6eec09e44Cb0fa19b84353ad0f7858d7F61a"
            ]
        ```

## Monitoring

* Enable metrics using config parameters `prom.metrics` and `prom.http`.
* `ipld-eth-state-snapshot` exposes following prometheus metrics at `/metrics` endpoint:
    * `state_node_count`: Number of state nodes processed.
    * `storage_node_count`: Number of storage nodes processed.
    * `code_node_count`: Number of code nodes processed.
    * DB stats if operating in `postgres` mode.

## Tests

* Run unit tests:

    ```bash
    # setup db
    docker-compose up -d

    # run tests after db migrations are run
    make dbtest

    # tear down db
    docker-compose down -v --remove-orphans
    ```

## Import output data in file mode into a database

* When `ipld-eth-state-snapshot stateSnapshot` is run in file mode (`database.type`), the output is in form of CSV files.

* Assuming the output files are located in host's `./output_dir` directory.

* Data post-processing:

    * Create a directory to store post-processed output:

        ```bash
        mkdir -p output_dir/processed_output
        ```

    * Combine output from multiple workers and copy to post-processed output directory:

        ```bash
        # ipld.blocks
        cat {output_dir,output_dir/*}/ipld.blocks.csv > output_dir/processed_output/combined-ipld.blocks.csv

        # eth.state_cids
        cat output_dir/*/eth.state_cids.csv > output_dir/processed_output/combined-eth.state_cids.csv

        # eth.storage_cids
        cat output_dir/*/eth.storage_cids.csv > output_dir/processed_output/combined-eth.storage_cids.csv

        # public.nodes
        cp output_dir/public.nodes.csv output_dir/processed_output/public.nodes.csv

        # eth.header_cids
        cp output_dir/eth.header_cids.csv output_dir/processed_output/eth.header_cids.csv
        ```

    * De-duplicate data:

        ```bash
        # ipld.blocks
        sort -u output_dir/processed_output/combined-ipld.blocks.csv -o output_dir/processed_output/deduped-combined-ipld.blocks.csv

        # eth.header_cids
        sort -u output_dir/processed_output/eth.header_cids.csv -o output_dir/processed_output/deduped-eth.header_cids.csv

        # eth.state_cids
        sort -u output_dir/processed_output/combined-eth.state_cids.csv -o output_dir/processed_output/deduped-combined-eth.state_cids.csv

        # eth.storage_cids
        sort -u output_dir/processed_output/combined-eth.storage_cids.csv -o output_dir/processed_output/deduped-combined-eth.storage_cids.csv
        ```

* Copy over the post-processed output files to the DB server (say in `/output_dir`).

* Start `psql` to run the import commands:

    ```bash
    psql -U <DATABASE_USER> -h <DATABASE_HOSTNAME> -p <DATABASE_PORT> <DATABASE_NAME>
    ```

* Run the following to import data:

    ```bash
    # public.nodes
    COPY public.nodes FROM '/output_dir/processed_output/public.nodes.csv' CSV;

    # ipld.blocks
    COPY ipld.blocks FROM '/output_dir/processed_output/deduped-combined-ipld.blocks.csv' CSV;

    # eth.header_cids
    COPY eth.header_cids FROM '/output_dir/processed_output/deduped-eth.header_cids.csv' CSV;

    # eth.state_cids
    COPY eth.state_cids FROM '/output_dir/processed_output/deduped-combined-eth.state_cids.csv' CSV FORCE NOT NULL state_leaf_key;

    # eth.storage_cids
    COPY eth.storage_cids FROM '/output_dir/processed_output/deduped-combined-eth.storage_cids.csv' CSV FORCE NOT NULL storage_leaf_key;
    ```

* NOTE: `COPY` command on CSVs inserts empty strings as `NULL` in the DB. Passing `FORCE_NOT_NULL <COLUMN_NAME>` forces it to insert empty strings instead. This is required to maintain compatibility of the imported snapshot data with the data generated by statediffing. Reference: https://www.postgresql.org/docs/14/sql-copy.html

### Troubleshooting

* Run the following command to find any rows (in data dumps in `file` mode) having unexpected number of columns:

    ```bash
    ./scripts/find-bad-rows.sh -i <input-file> -c <expected-columns> -o [output-file] -d true
    ```

* Run the following command to select rows (from data dumps in `file` mode) other than the ones having unexpected number of columns:

    ```bash
    ./scripts/filter-bad-rows.sh -i <input-file> -c <expected-columns> -o <output-file>
    ```

* See [scripts](./scripts) for more details.
