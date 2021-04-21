module github.com/GeoDB-Limited/odin-deposit-ether-svc

go 1.15

require (
	github.com/GeoDB-Limited/odin-core v1.0.1-0.20210420142537-c25dc6c4d0ce
	github.com/alecthomas/kingpin v2.2.6+incompatible
	github.com/cosmos/cosmos-sdk v0.42.3
	github.com/cosmos/go-bip39 v1.0.0
	github.com/ethereum/go-ethereum v1.10.1
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.36.0
	gopkg.in/yaml.v2 v2.4.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
