package client

import (
	"github.com/gorilla/mux"

	"github.com/hashrs/blockchain/framework/chain-app/client/context"
	"github.com/hashrs/blockchain/framework/chain-app/client/rpc"
)

// Register routes
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	rpc.RegisterRPCRoutes(cliCtx, r)
}
