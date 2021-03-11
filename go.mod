module github.com/GeoDB-Limited/odin-deposit-ether-svc

go 1.15

require (
	github.com/alecthomas/kingpin v2.2.6+incompatible
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d // indirect
	github.com/cosmos/cosmos-sdk v0.42.1 // indirect
	github.com/ethereum/go-ethereum v1.10.1
	github.com/tendermint/iavl v0.14.1 // indirect
	gitlab.com/distributed_lab/figure v2.1.0+incompatible
	gitlab.com/distributed_lab/kit v1.8.5
	gitlab.com/distributed_lab/logan v3.8.0+incompatible
	gitlab.com/distributed_lab/running v0.0.0-20200706131153-4af0e83eb96c
	google.golang.org/grpc v1.36.0
)

replace github.com/gogo/protobuf v1.3.3 => github.com/gogo/protobuf v1.3.2
