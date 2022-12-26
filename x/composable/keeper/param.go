package keeper

import (
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

func (k Keeper) GetParams(ctx sdk.Context) composable.Params {
	bz, err := k.getParamsBytes(ctx)
	if err != nil {
		panic(err)
	}

	var params composable.Params
	k.cdc.MustUnmarshal(bz, &params)

	return params
}

func (k Keeper) getParamsBytes(ctx sdk.Context) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := paramsKey

	bz := store.Get(key)
	if bz == nil {
		return nil, sdkerrors.ErrNotFound.Wrap("params")
	}

	return bz, nil
}

func (k Keeper) SetParams(ctx sdk.Context, params composable.Params) {
	store := ctx.KVStore(k.storeKey)
	key := paramsKey

	bz := k.cdc.MustMarshal(&params)

	store.Set(key, bz)
}
