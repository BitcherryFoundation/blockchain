package core

import (
	abci "github.com/hashrs/blockchain/core/consensus/dpos-pbft/abci/types"
	"github.com/hashrs/blockchain/core/consensus/dpos-pbft/libs/bytes"
	"github.com/hashrs/blockchain/core/consensus/dpos-pbft/proxy"
	ctypes "github.com/hashrs/blockchain/core/consensus/dpos-pbft/rpc/core/types"
	rpctypes "github.com/hashrs/blockchain/core/consensus/dpos-pbft/rpc/lib/types"
)

// ABCIQuery queries the application for some information.
// More: https://hashrs.com/rpc/#/ABCI/abci_query
func ABCIQuery(
	ctx *rpctypes.Context,
	path string,
	data bytes.HexBytes,
	height int64,
	prove bool,
) (*ctypes.ResultABCIQuery, error) {
	resQuery, err := proxyAppQuery.QuerySync(abci.RequestQuery{
		Path:   path,
		Data:   data,
		Height: height,
		Prove:  prove,
	})
	if err != nil {
		return nil, err
	}
	logger.Info("ABCIQuery", "path", path, "data", data, "result", resQuery)
	return &ctypes.ResultABCIQuery{Response: *resQuery}, nil
}

// ABCIInfo gets some info about the application.
// More: https://hashrs.com/rpc/#/ABCI/abci_info
func ABCIInfo(ctx *rpctypes.Context) (*ctypes.ResultABCIInfo, error) {
	resInfo, err := proxyAppQuery.InfoSync(proxy.RequestInfo)
	if err != nil {
		return nil, err
	}
	return &ctypes.ResultABCIInfo{Response: *resInfo}, nil
}
