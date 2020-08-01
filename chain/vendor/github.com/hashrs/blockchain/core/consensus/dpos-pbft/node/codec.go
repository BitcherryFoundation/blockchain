package node

import (
	cryptoamino "github.com/hashrs/blockchain/core/consensus/dpos-pbft/crypto/encoding/amino"
	amino "github.com/hashrs/blockchain/libs/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoamino.RegisterAmino(cdc)
}
