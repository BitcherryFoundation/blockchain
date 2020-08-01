package simulation

import (
	"math/rand"

	simappparams "github.com/hashrs/blockchain/framework/chain-app/simapp/params"
	sdk "github.com/hashrs/blockchain/framework/chain-app/types"
	"github.com/hashrs/blockchain/framework/chain-app/x/gov/types"
	"github.com/hashrs/blockchain/framework/chain-app/x/simulation"
)

// OpWeightSubmitTextProposal app params key for text proposal
const OpWeightSubmitTextProposal = "op_weight_submit_text_proposal"

// ProposalContents defines the module weighted proposals' contents
func ProposalContents() []simulation.WeightedProposalContent {
	return []simulation.WeightedProposalContent{
		{
			AppParamsKey:       OpWeightSubmitTextProposal,
			DefaultWeight:      simappparams.DefaultWeightTextProposal,
			ContentSimulatorFn: SimulateTextProposalContent,
		},
	}
}

// SimulateTextProposalContent returns a random text proposal content.
func SimulateTextProposalContent(r *rand.Rand, _ sdk.Context, _ []simulation.Account) types.Content {
	return types.NewTextProposal(
		simulation.RandStringOfLength(r, 140),
		simulation.RandStringOfLength(r, 5000),
	)
}
