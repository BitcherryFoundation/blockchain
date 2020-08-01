package core

import (
	ctypes "github.com/hashrs/blockchain/core/consensus/dpos-pbft/rpc/core/types"
	rpctypes "github.com/hashrs/blockchain/core/consensus/dpos-pbft/rpc/lib/types"
)

// Health gets node health. Returns empty result (200 OK) on success, no
// response - in case of an error.
// More: https://hashrs.com/rpc/#/Info/health
func Health(ctx *rpctypes.Context) (*ctypes.ResultHealth, error) {
	return &ctypes.ResultHealth{}, nil
}
