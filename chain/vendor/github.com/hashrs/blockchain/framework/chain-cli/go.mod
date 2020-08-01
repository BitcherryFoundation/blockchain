module github.com/hashrs/blockchain/framework/chain-cli

go 1.12

require (
	github.com/hashrs/blockchain/core/consensus/dpos-pbft v0.0.0-00010101000000-000000000000
	github.com/hashrs/blockchain/libs/log v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.7
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
)

replace github.com/hashrs/blockchain/core/consensus/dpos-pbft => ../../core/consensus/dpos-pbft

replace github.com/hashrs/blockchain/libs/amino => ../../libs/amino

replace github.com/hashrs/blockchain/libs/crypto => ../../libs/crypto

replace github.com/hashrs/blockchain/libs/log => ../../libs/log

replace github.com/hashrs/blockchain/libs/state-db/tm-db => ../../libs/state-db/tm-db
