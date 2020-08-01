package coretypes

import (
	"github.com/hashrs/blockchain/core/consensus/dpos-pbft/types"
	amino "github.com/hashrs/blockchain/libs/amino"
)

func RegisterAmino(cdc *amino.Codec) {
	types.RegisterEventDatas(cdc)
	types.RegisterBlockAmino(cdc)
}
