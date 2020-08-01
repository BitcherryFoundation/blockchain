package mempool

import (
	amino "github.com/hashrs/blockchain/libs/amino"
)

var cdc = amino.NewCodec()

func init() {
	RegisterMessages(cdc)
}
