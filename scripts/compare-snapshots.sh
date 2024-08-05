#!/bin/bash
# Compare the full snapshot output from two versions of the service
#
# Usage: compare-versions.sh [-d <output-dir>] <binary-A> <binary-B>

# Configure the input data using environment vars.
(
  set -u
  : $SNAPSHOT_BLOCK_HEIGHT
  : $ETHDB_PATH
  : $ETHDB_ANCIENT
  : $ETH_GENESIS_BLOCK
)

while getopts d: opt; do
    case $opt in
      d) output_dir="$OPTARG"
    esac
done
shift $((OPTIND - 1))

binary_A=$1
binary_B=$2
shift 2

if [[ -z $output_dir ]]; then
  output_dir=$(mktemp -d)
fi

export SNAPSHOT_MODE=postgres
export SNAPSHOT_WORKERS=32
export SNAPSHOT_RECOVERY_FILE='compare-snapshots-recovery.txt'

export DATABASE_NAME="cerc_testing"
export DATABASE_HOSTNAME="localhost"
export DATABASE_PORT=8077
export DATABASE_USER="vdbm"
export DATABASE_PASSWORD="password"

export ETH_CLIENT_NAME=test-client
export ETH_NODE_ID=test-node
export ETH_NETWORK_ID=test-network
export ETH_CHAIN_ID=4242

dump_table() {
  statement="copy (select * from $1) to stdout with csv"
  docker exec -e PGPASSWORD=password test-ipld-eth-db-1 \
    psql -q cerc_testing -U vdbm -c "$statement" | sort -u > "$2/$1.csv"
}

clear_table() {
  docker exec -e PGPASSWORD=password test-ipld-eth-db-1 \
    psql -q cerc_testing -U vdbm -c "truncate $1"
}

tables=(
  eth.log_cids
  eth.receipt_cids
  eth.state_cids
  eth.storage_cids
  eth.transaction_cids
  eth.uncle_cids
  ipld.blocks
  public.nodes
)

for table in "${tables[@]}"; do
  clear_table $table
done

$binary_A stateSnapshot

mkdir -p $output_dir/A
for table in "${tables[@]}"; do
  dump_table $table $output_dir/A
  clear_table $table
done

$binary_B stateSnapshot

mkdir -p $output_dir/B
for table in "${tables[@]}"; do
  dump_table $table $output_dir/B
  clear_table $table
done

diff -rs $output_dir/A $output_dir/B
