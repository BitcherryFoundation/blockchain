package client

import (
	"github.com/hashrs/blockchain/core/consensus/dpos-pbft/types"
	amino "github.com/hashrs/blockchain/libs/amino"
)

var cdc = amino.NewCodec()

func init() {
	types.RegisterEvidences(cdc)
}
