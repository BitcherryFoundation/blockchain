package rest

import (
	"github.com/gorilla/mux"

	"github.com/hashrs/blockchain/framework/chain-app/client/context"
)

// RegisterRoutes registers staking-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}
