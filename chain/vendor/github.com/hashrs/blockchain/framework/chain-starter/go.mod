module github.com/hashrs/blockchain/framework/chain-starter

go 1.12

require (
	github.com/gorilla/mux v1.7.4
	github.com/hashrs/blockchain/core/consensus/dpos-pbft v0.0.0-00010101000000-000000000000
	github.com/hashrs/blockchain/framework/chain-app v0.0.0-00010101000000-000000000000
	github.com/hashrs/blockchain/libs/amino v0.0.0-00010101000000-000000000000
	github.com/hashrs/blockchain/libs/state-db/tm-db v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v0.0.7
	github.com/spf13/viper v1.6.3
)

replace github.com/hashrs/blockchain/framework/chain-iavl => ../chain-iavl

replace github.com/hashrs/blockchain/framework/chain-app => ../chain-app

replace github.com/hashrs/blockchain/libs/crypto => ../../libs/crypto

replace github.com/hashrs/blockchain/libs/amino => ../../libs/amino

replace github.com/hashrs/blockchain/libs/go-bip39 => ../../libs/go-bip39

replace github.com/hashrs/blockchain/libs/state-db/tm-db => ../../libs/state-db/tm-db

replace github.com/hashrs/blockchain/core/consensus/dpos-pbft => ../../core/consensus/dpos-pbft
