package evidence

import (
	"fmt"

	sdk "github.com/hashrs/blockchain/framework/chain-app/types"

	abci "github.com/hashrs/blockchain/core/consensus/dpos-pbft/abci/types"
	tmtypes "github.com/hashrs/blockchain/core/consensus/dpos-pbft/types"
)

// BeginBlocker iterates through and handles any newly discovered evidence of
// misbehavior submitted by Tendermint. Currently, only equivocation is handled.
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k Keeper) {
	for _, tmEvidence := range req.ByzantineValidators {
		switch tmEvidence.Type {
		case tmtypes.ABCIEvidenceTypeDuplicateVote:
			evidence := ConvertDuplicateVoteEvidence(tmEvidence)
			k.HandleDoubleSign(ctx, evidence.(Equivocation))

		default:
			k.Logger(ctx).Error(fmt.Sprintf("ignored unknown evidence type: %s", tmEvidence.Type))
		}
	}
}
