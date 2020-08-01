package types

import (
	"github.com/hashrs/blockchain/framework/chain-app/codec"
	sdk "github.com/hashrs/blockchain/framework/chain-app/types"
	authtypes "github.com/hashrs/blockchain/framework/chain-app/x/auth/types"
	stakingtypes "github.com/hashrs/blockchain/framework/chain-app/x/staking/types"
)

// ModuleCdc defines a generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

// TODO: abstract genesis transactions registration back to staking
// required for genesis transactions
func init() {
	ModuleCdc = codec.New()
	stakingtypes.RegisterCodec(ModuleCdc)
	authtypes.RegisterCodec(ModuleCdc)
	sdk.RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
