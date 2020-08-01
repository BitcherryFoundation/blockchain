package simulation

import (
	"bytes"
	"fmt"

	tmkv "github.com/hashrs/blockchain/core/consensus/dpos-pbft/libs/kv"

	"github.com/hashrs/blockchain/framework/chain-app/codec"
	"github.com/hashrs/blockchain/framework/chain-app/x/supply/internal/keeper"
	"github.com/hashrs/blockchain/framework/chain-app/x/supply/internal/types"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding supply type
func DecodeStore(cdc *codec.Codec, kvA, kvB tmkv.Pair) string {
	switch {
	case bytes.Equal(kvA.Key[:1], keeper.SupplyKey):
		var supplyA, supplyB types.Supply
		cdc.MustUnmarshalBinaryLengthPrefixed(kvA.Value, &supplyA)
		cdc.MustUnmarshalBinaryLengthPrefixed(kvB.Value, &supplyB)
		return fmt.Sprintf("%v\n%v", supplyB, supplyB)
	default:
		panic(fmt.Sprintf("invalid supply key %X", kvA.Key))
	}
}
