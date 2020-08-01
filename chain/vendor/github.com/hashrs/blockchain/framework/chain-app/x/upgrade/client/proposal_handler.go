package client

import (
	govclient "github.com/hashrs/blockchain/framework/chain-app/x/gov/client"
	"github.com/hashrs/blockchain/framework/chain-app/x/upgrade/client/cli"
	"github.com/hashrs/blockchain/framework/chain-app/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
