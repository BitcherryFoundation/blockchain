package types

import (
	cryptoamino "github.com/hashrs/blockchain/core/consensus/dpos-pbft/crypto/encoding/amino"
	amino "github.com/hashrs/blockchain/libs/amino"
)

var cdc = amino.NewCodec()

func init() {
	RegisterBlockAmino(cdc)
}

func RegisterBlockAmino(cdc *amino.Codec) {
	cryptoamino.RegisterAmino(cdc)
	RegisterEvidences(cdc)
}

// GetCodec returns a codec used by the package. For testing purposes only.
func GetCodec() *amino.Codec {
	return cdc
}

// For testing purposes only
func RegisterMockEvidencesGlobal() {
	RegisterMockEvidences(cdc)
}
