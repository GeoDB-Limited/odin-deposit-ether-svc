module github.com/GeoDB-Limited/odin-deposit-ether-svc

go 1.15

require (
	github.com/alecthomas/kingpin v2.2.6+incompatible
	github.com/cosmos/cosmos-sdk v0.42.1
	github.com/ethereum/go-ethereum v1.10.1
	github.com/gogo/protobuf v1.3.3 // indirect
	gitlab.com/distributed_lab/figure v2.1.0+incompatible
	gitlab.com/distributed_lab/kit v1.8.5
	gitlab.com/distributed_lab/logan v3.8.0+incompatible
	google.golang.org/grpc v1.36.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
