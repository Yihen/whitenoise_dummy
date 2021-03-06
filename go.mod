module whitenoise

go 1.14

require (
	github.com/golang/protobuf v1.4.0
	github.com/libp2p/go-libp2p v0.13.0
	github.com/libp2p/go-libp2p-core v0.8.0
	github.com/libp2p/go-libp2p-noise v0.1.1
	github.com/libp2p/go-libp2p-pubsub v0.0.0-00010101000000-000000000000
	github.com/mr-tron/base58 v1.2.0
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/stretchr/testify v1.6.1
	google.golang.org/protobuf v1.23.0
)

replace github.com/libp2p/go-libp2p-pubsub => ../github.com/libp2p/go-libp2p-pubsub
