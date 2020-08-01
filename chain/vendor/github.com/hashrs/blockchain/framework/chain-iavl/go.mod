module github.com/hashrs/blockchain/framework/chain-iavl

go 1.12

require (
	github.com/hashrs/blockchain/core/consensus/dpos-pbft v0.0.0-00010101000000-000000000000
	github.com/hashrs/blockchain/libs/amino v0.0.0-00010101000000-000000000000
	github.com/hashrs/blockchain/libs/crypto v0.0.0-20191206172530-e9b2fee46413
	github.com/hashrs/blockchain/libs/state-db/tm-db v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.5.1
)

replace github.com/hashrs/blockchain/libs/amino => ../../libs/amino

replace github.com/hashrs/blockchain/libs/crypto => ../../libs/crypto

replace github.com/hashrs/blockchain/libs/state-db/tm-db => ../../libs/state-db/tm-db

replace github.com/hashrs/blockchain/core/consensus/dpos-pbft => ../../core/consensus/dpos-pbft