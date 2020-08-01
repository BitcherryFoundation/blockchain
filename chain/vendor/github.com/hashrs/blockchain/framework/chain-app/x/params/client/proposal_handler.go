package client

import (
	govclient "github.com/hashrs/blockchain/framework/chain-app/x/gov/client"
	"github.com/hashrs/blockchain/framework/chain-app/x/params/client/cli"
	"github.com/hashrs/blockchain/framework/chain-app/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
