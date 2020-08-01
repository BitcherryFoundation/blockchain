package privval

import (
	"github.com/hashrs/blockchain/core/consensus/dpos-pbft/crypto"
	"github.com/hashrs/blockchain/core/consensus/dpos-pbft/types"
	amino "github.com/hashrs/blockchain/libs/amino"
)

// SignerMessage is sent between Signer Clients and Servers.
type SignerMessage interface{}

func RegisterRemoteSignerMsg(cdc *amino.Codec) {
	cdc.RegisterInterface((*SignerMessage)(nil), nil)
	cdc.RegisterConcrete(&PubKeyRequest{}, "hashrs/remotesigner/PubKeyRequest", nil)
	cdc.RegisterConcrete(&PubKeyResponse{}, "hashrs/remotesigner/PubKeyResponse", nil)
	cdc.RegisterConcrete(&SignVoteRequest{}, "hashrs/remotesigner/SignVoteRequest", nil)
	cdc.RegisterConcrete(&SignedVoteResponse{}, "hashrs/remotesigner/SignedVoteResponse", nil)
	cdc.RegisterConcrete(&SignProposalRequest{}, "hashrs/remotesigner/SignProposalRequest", nil)
	cdc.RegisterConcrete(&SignedProposalResponse{}, "hashrs/remotesigner/SignedProposalResponse", nil)

	cdc.RegisterConcrete(&PingRequest{}, "hashrs/remotesigner/PingRequest", nil)
	cdc.RegisterConcrete(&PingResponse{}, "hashrs/remotesigner/PingResponse", nil)
}

// TODO: Add ChainIDRequest

// PubKeyRequest requests the consensus public key from the remote signer.
type PubKeyRequest struct{}

// PubKeyResponse is a response message containing the public key.
type PubKeyResponse struct {
	PubKey crypto.PubKey
	Error  *RemoteSignerError
}

// SignVoteRequest is a request to sign a vote
type SignVoteRequest struct {
	Vote *types.Vote
}

// SignedVoteResponse is a response containing a signed vote or an error
type SignedVoteResponse struct {
	Vote  *types.Vote
	Error *RemoteSignerError
}

// SignProposalRequest is a request to sign a proposal
type SignProposalRequest struct {
	Proposal *types.Proposal
}

// SignedProposalResponse is response containing a signed proposal or an error
type SignedProposalResponse struct {
	Proposal *types.Proposal
	Error    *RemoteSignerError
}

// PingRequest is a request to confirm that the connection is alive.
type PingRequest struct {
}

// PingResponse is a response to confirm that the connection is alive.
type PingResponse struct {
}
