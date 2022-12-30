package keeper

import (
	"github.com/line/lbm-sdk/codec"
	storetypes "github.com/line/lbm-sdk/store/types"
	sdk "github.com/line/lbm-sdk/types"
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

func (k Keeper) iterateImpl(ctx sdk.Context, prefix []byte, fn func(key, value []byte)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		fn(iterator.Key(), iterator.Value())
	}
}
