module github.com/hashrs/blockchain/chain

go 1.12

require (
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/hashrs/blockchain/core/consensus/dpos-pbft v0.0.0-00010101000000-000000000000
	github.com/hashrs/blockchain/framework/chain-app v0.0.0-00010101000000-000000000000
	github.com/hashrs/blockchain/framework/chain-cli v0.0.0-00010101000000-000000000000
	github.com/hashrs/blockchain/framework/chain-starter v0.0.0-00010101000000-000000000000
	github.com/hashrs/blockchain/libs/state-db/tm-db v0.0.0-00010101000000-000000000000
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.7
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7 // indirect
	golang.org/x/text v0.3.2 // indirect
)

replace github.com/hashrs/blockchain/libs/amino => ../libs/amino

replace github.com/hashrs/blockchain/libs/log => ../libs/log

replace github.com/hashrs/blockchain/libs/crypto => ../libs/crypto

replace github.com/hashrs/blockchain/libs/go-bip39 => ../libs/go-bip39

replace github.com/hashrs/blockchain/framework/chain-app => ../framework/chain-app

replace github.com/hashrs/blockchain/libs/state-db/tm-db => ../libs/state-db/tm-db

replace github.com/hashrs/blockchain/framework/chain-iavl => ../framework/chain-iavl

replace github.com/hashrs/blockchain/framework/chain-starter => ../framework/chain-starter

replace github.com/hashrs/blockchain/framework/chain-cli => ../framework/chain-cli

replace github.com/hashrs/blockchain/core/consensus/dpos-pbft => ../core/consensus/dpos-pbft
