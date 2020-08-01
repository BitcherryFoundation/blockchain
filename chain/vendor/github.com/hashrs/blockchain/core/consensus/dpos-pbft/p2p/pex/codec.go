package pex

import (
	amino "github.com/hashrs/blockchain/libs/amino"
)

var cdc *amino.Codec = amino.NewCodec()

func init() {
	RegisterMessages(cdc)
}
