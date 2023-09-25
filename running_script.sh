#!/bin/bash

KEYS[0]="validator"
KEYS[1]="airdrop"
KEYS[2]="community_pool"
KEYS[3]="strategic_reserve"
KEYS[4]="mission_control"
KEYS[5]="developer_fund"
KEYS[6]="advisors"

BALANCE[0]="1000000ollo" 	# 1M for validator
BALANCE[1]="45000000ollo"	# 45M for Airdrop
BALANCE[2]="20000000ollo"	# 20M for Community Pool
BALANCE[3]="20000000ollo"	# 20M for Strategic Reserve
BALANCE[4]="12000000ollo"	# 12M for Mission Control(Team)
BALANCE[5]="2000000ollo"	# 2M for Developer Fund
BALANCE[6]="1000000ollo"	# 1M for Advisors

total_supply=101000000		# Total 101M

CHAINID="ollo-testnet-2"
MONIKER="ollo_testnet"
# Remember to change to other types of keyring like 'file' in-case exposing to outside world,
# otherwise your balance will be wiped quickly
# The keyring test does not require private key to steal tokens from you
KEYRING="test"
LOGLEVEL="info"
# Set dedicated home directory for the ollod instance
HOMEDIR="$HOME/.ollo"
# to trace evm
#TRACE="--trace"
TRACE=""

# Path variables
CONFIG=$HOMEDIR/config/config.toml
APP_CONFIG=$HOMEDIR/config/app.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json

# validate dependencies are installed
command -v jq >/dev/null 2>&1 || {
	echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"
	exit 1
}

# used to exit on first error (any non-zero exit code)
set -e

# Reinstall daemon
make install

# User prompt if an existing local node configuration is found.
if [ -d "$HOMEDIR" ]; then
	printf "\nAn existing folder at '%s' was found. You can choose to delete this folder and start a new local node with new keys from genesis. When declined, the existing local node is started. \n" "$HOMEDIR"
	echo "Overwrite the existing configuration and start a new local node? [y/n]"
	read -r overwrite
else
	overwrite="Y"
fi

# Setup local node if overwrite is set to Yes, otherwise skip setup
if [[ $overwrite == "y" || $overwrite == "Y" ]]; then
	# Remove the previous folder
	rm -rf "$HOMEDIR"

	# Set client config
	ollod config keyring-backend $KEYRING --home "$HOMEDIR"
	ollod config chain-id $CHAINID --home "$HOMEDIR"

	# If keys exist they should be deleted
	for KEY in "${KEYS[@]}"; do
		ollod keys add $KEY --keyring-backend $KEYRING --home "$HOMEDIR" 
	done

	# Set moniker and chain-id for Evmos (Moniker can be anything, chain-id must be an integer)
	ollod init $MONIKER -o --chain-id $CHAINID --home "$HOMEDIR"

	# Change parameter token denominations to exa
	sed -i 's/stake/ollo/g' $GENESIS
	sed -i 's/utest/ollo/g' $GENESIS

	jq '.app_state["evm"]["params"]["evm_denom"]="ollo"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	jq '.app_state["market"]["params"]["commission"]="1"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["market"]["params"]["bid_close_duration"]="172800s"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["market"]["params"]["distribution"]["staking"]="0.500000000000000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["market"]["params"]["distribution"]["community_pool"]="0.500000000000000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["market"]["next_auction_number"]="1"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
    jq '.app_state["staking"]["params"]["max_validators"]="25"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# Set claims start time
	current_date=$(date -u +"%Y-%m-%dT%TZ")
	jq -r --arg current_date "$current_date" '.app_state["claims"]["params"]["airdrop_start_time"]=$current_date' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"


    # Set parameters
    sed -i 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["*"\]/g' "$CONFIG"
	sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' "$CONFIG"

    sed -i '/\[api\]/,+3 s/enable = false/enable = true/g' $APP_CONFIG
	sed -i 's/swagger = false/swagger = true/g' $APP_CONFIG
	sed -i 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g'  $APP_CONFIG
	sed -i 's/api = "eth,net,web3"/api = "eth,txpool,personal,net,debug,web3,pubsub,trace"/g' $APP_CONFIG

	# Set inflation
	jq '.app_state["mint"]["minter"]["inflation"]="1"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["mint"]["params"]["inflation_max"]="1"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["mint"]["params"]["inflation_min"]="0.0"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["mint"]["params"]["inflation_rate_change"]="0.1"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# Set token supply
	jq '.app_state["token"]["tokens"]["initial_supply"]="100000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["token"]["tokens"]["max_supply"]="5000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# Allocate genesis accounts (cosmos formatted addresses)
	for i in "${!KEYS[@]}"; do
		ollod add-genesis-account "${KEYS[$i]}" "${BALANCE[$i]}" --keyring-backend $KEYRING --home "$HOMEDIR"
	done

	# bc is required to add these big numbers
	
	jq -r --arg total_supply "$total_supply" '.app_state["bank"]["supply"][0]["denom"]="ollo"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq -r --arg total_supply "$total_supply" '.app_state["bank"]["supply"][0]["amount"]=$total_supply' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# Sign genesis transaction
	ollod gentx ${KEYS[0]} 1000000ollo --keyring-backend $KEYRING --chain-id $CHAINID --home "$HOMEDIR"
	

	# Collect genesis tx
	ollod collect-gentxs --home "$HOMEDIR"

	# Run this to ensure everything worked and that the genesis file is setup correctly
	ollod validate-genesis --home "$HOMEDIR"

	if [[ $1 == "pending" ]]; then
		echo "pending mode is on, please wait for the first block committed."
	fi
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
ollod start --pruning=nothing "$TRACE" --rpc.laddr tcp://0.0.0.0:26657 --log_level $LOGLEVEL --home "$HOMEDIR"