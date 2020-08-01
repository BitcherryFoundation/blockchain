package merkle

import (
	amino "github.com/hashrs/blockchain/libs/amino"
)

var cdc *amino.Codec

func init() {
	cdc = amino.NewCodec()
	cdc.Seal()
}
