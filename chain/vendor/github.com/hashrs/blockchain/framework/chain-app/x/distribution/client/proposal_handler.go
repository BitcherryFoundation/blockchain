package client

import (
	"github.com/hashrs/blockchain/framework/chain-app/x/distribution/client/cli"
	"github.com/hashrs/blockchain/framework/chain-app/x/distribution/client/rest"
	govclient "github.com/hashrs/blockchain/framework/chain-app/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
