package multisig

import (
	"github.com/hashrs/blockchain/core/consensus/dpos-pbft/crypto"
	"github.com/hashrs/blockchain/core/consensus/dpos-pbft/crypto/ed25519"
	"github.com/hashrs/blockchain/core/consensus/dpos-pbft/crypto/secp256k1"
	"github.com/hashrs/blockchain/core/consensus/dpos-pbft/crypto/sr25519"
	amino "github.com/hashrs/blockchain/libs/amino"
)

// TODO: Figure out API for others to either add their own pubkey types, or
// to make verify / marshal accept a cdc.
const (
	PubKeyMultisigThresholdAminoRoute = "hashrs/PubKeyMultisigThreshold"
)

var cdc = amino.NewCodec()

func init() {
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(PubKeyMultisigThreshold{},
		PubKeyMultisigThresholdAminoRoute, nil)
	cdc.RegisterConcrete(ed25519.PubKeyEd25519{},
		ed25519.PubKeyAminoName, nil)
	cdc.RegisterConcrete(sr25519.PubKeySr25519{},
		sr25519.PubKeyAminoName, nil)
	cdc.RegisterConcrete(secp256k1.PubKeySecp256k1{},
		secp256k1.PubKeyAminoName, nil)
}
