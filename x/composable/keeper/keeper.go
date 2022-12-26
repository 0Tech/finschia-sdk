package keeper

import (
	"github.com/line/lbm-sdk/codec"
	storetypes "github.com/line/lbm-sdk/store/types"
)

type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec

	authority string
}

func NewKeeper(
	storeKey storetypes.StoreKey,
	cdc codec.BinaryCodec,
	authority string,
) Keeper {
	return Keeper{
		storeKey:  storeKey,
		cdc:       cdc,
		authority: authority,
	}
}
