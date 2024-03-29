#!/bin/bash

docker run -v $PWD:$PWD -w $PWD ethereum/solc:0.8.3 @openzeppelin/=$(pwd)/node_modules/@openzeppelin/ --overwrite --abi --bin -o ./build \
contracts/*.sol

# ethereum and erc20 common bridge
./bin/abigen --abi ./build/Bridge.abi --bin ./build/Bridge.bin --type Bridge --pkg generated --out ./generated/bridge.go
yarn compile-types
