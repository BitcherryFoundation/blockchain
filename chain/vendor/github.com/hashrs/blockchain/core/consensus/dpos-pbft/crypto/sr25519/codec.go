package sr25519

import (
	"github.com/hashrs/blockchain/core/consensus/dpos-pbft/crypto"
	amino "github.com/hashrs/blockchain/libs/amino"
)

var _ crypto.PrivKey = PrivKeySr25519{}

const (
	PrivKeyAminoName = "hashrs/PrivKeySr25519"
	PubKeyAminoName  = "hashrs/PubKeySr25519"

	// SignatureSize is the size of an Edwards25519 signature. Namely the size of a compressed
	// Sr25519 point, and a field element. Both of which are 32 bytes.
	SignatureSize = 64
)

var cdc = amino.NewCodec()

func init() {
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(PubKeySr25519{},
		PubKeyAminoName, nil)

	cdc.RegisterInterface((*crypto.PrivKey)(nil), nil)
	cdc.RegisterConcrete(PrivKeySr25519{},
		PrivKeyAminoName, nil)
}
