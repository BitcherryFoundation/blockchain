package evidence

import (
	cryptoamino "github.com/hashrs/blockchain/core/consensus/dpos-pbft/crypto/encoding/amino"
	"github.com/hashrs/blockchain/core/consensus/dpos-pbft/types"
	amino "github.com/hashrs/blockchain/libs/amino"
)

var cdc = amino.NewCodec()

func init() {
	RegisterMessages(cdc)
	cryptoamino.RegisterAmino(cdc)
	types.RegisterEvidences(cdc)
}

// For testing purposes only
func RegisterMockEvidences() {
	types.RegisterMockEvidences(cdc)
}
