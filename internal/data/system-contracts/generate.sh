solc --abi --bin -o ./build ./contracts/OdinBridge.sol --overwrite
./bin/abigen --abi ./build/OdinBridge.abi --bin ./build/OdinBridge.bin --type OdinBridge --pkg generated --out ./generated/OdinBridge.go

