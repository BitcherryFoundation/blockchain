package keeper

import (
	abci "github.com/hashrs/blockchain/core/consensus/dpos-pbft/abci/types"

	"github.com/hashrs/blockchain/framework/chain-app/codec"
	sdk "github.com/hashrs/blockchain/framework/chain-app/types"
	sdkerrors "github.com/hashrs/blockchain/framework/chain-app/types/errors"
	"github.com/hashrs/blockchain/framework/chain-app/x/auth/types"
)

// NewQuerier creates a querier for auth REST endpoints
func NewQuerier(keeper AccountKeeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryAccount:
			return queryAccount(ctx, req, keeper)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

func queryAccount(ctx sdk.Context, req abci.RequestQuery, keeper AccountKeeper) ([]byte, error) {
	var params types.QueryAccountParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	account := keeper.GetAccount(ctx, params.Address)
	if account == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "account %s does not exist", params.Address)
	}

	bz, err := codec.MarshalJSONIndent(keeper.cdc, account)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
