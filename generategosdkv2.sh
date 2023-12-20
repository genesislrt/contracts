#!/bin/bash
SOLIDITY_CONTRACTS="./projects/liquid-restaking/contracts"

contracts=$(find $SOLIDITY_CONTRACTS -name "*.sol")
for contract in $contracts
do
  solc --optimize --abi --bin --base-path ./projects/liquid-restaking --include-path ./projects/liquid-restaking/node_modules/ -o $SOLIDITY_CONTRACTS/build --overwrite $contract
  pkg="${contract%.*}"
  lowercase=$(echo "$pkg" | awk '{print tolower($0)}' )

  go run ./wrap.go $SOLIDITY_CONTRACTS/build/"${lowercase##*/}".abi $SOLIDITY_CONTRACTS/build/"${lowercase##*/}".bin contract "${lowercase##*/}"
done;