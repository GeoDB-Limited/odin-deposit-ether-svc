#!/bin/bash

docker run -v $PWD:$PWD -w $PWD ethereum/solc:0.7.2 @openzeppelin/=$(pwd)/node_modules/@openzeppelin/ --overwrite --abi --bin -o ./build \
contracts/Bridge.sol

./bin/abigen --abi ./build/Bridge.abi --bin ./build/Bridge.bin --type Bridge --pkg generated --out ./generated/bridge.go

