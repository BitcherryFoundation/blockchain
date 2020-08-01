package core

import (
	ctypes "github.com/hashrs/blockchain/core/consensus/dpos-pbft/rpc/core/types"
	rpctypes "github.com/hashrs/blockchain/core/consensus/dpos-pbft/rpc/lib/types"
	"github.com/hashrs/blockchain/core/consensus/dpos-pbft/types"
)

// BroadcastEvidence broadcasts evidence of the misbehavior.
// More: https://hashrs.com/rpc/#/Info/broadcast_evidence
func BroadcastEvidence(ctx *rpctypes.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error) {
	err := evidencePool.AddEvidence(ev)
	if err != nil {
		return nil, err
	}
	return &ctypes.ResultBroadcastEvidence{Hash: ev.Hash()}, nil
}
