package keeper

import (
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

func (k Keeper) NewClass(ctx sdk.Context, class composable.Class) error {
	if err := k.hasClass(ctx, class.Id); err == nil {
		return composable.ErrClassAlreadyExists.Wrap(class.Id)
	}
	k.setClass(ctx, class)

	return nil
}

func (k Keeper) UpdateClass(ctx sdk.Context, class composable.Class) error {
	if err := k.hasClass(ctx, class.Id); err != nil {
		return err
	}
	k.setClass(ctx, class)

	return nil
}

func (k Keeper) hasClass(ctx sdk.Context, classID string) error {
	_, err := k.getClassBytes(ctx, classID)
	return err
}

func (k Keeper) GetClass(ctx sdk.Context, classID string) (*composable.Class, error) {
	bz, err := k.getClassBytes(ctx, classID)
	if err != nil {
		return nil, err
	}

	var class composable.Class
	k.cdc.MustUnmarshal(bz, &class)

	return &class, nil
}

func (k Keeper) getClassBytes(ctx sdk.Context, classID string) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := classKey(classID)

	bz := store.Get(key)
	if bz == nil {
		return nil, composable.ErrClassNotFound.Wrap(classID)
	}

	return bz, nil
}

func (k Keeper) setClass(ctx sdk.Context, class composable.Class) {
	store := ctx.KVStore(k.storeKey)
	key := classKey(class.Id)

	bz := k.cdc.MustMarshal(&class)

	store.Set(key, bz)
}

func (k Keeper) MintNFT(ctx sdk.Context, owner sdk.AccAddress, classID string, nft composable.NFT) (sdk.Uint, error) {
	return sdk.ZeroUint(), sdkerrors.ErrNotSupported
}

func (k Keeper) BurnNFT(ctx sdk.Context, owner sdk.AccAddress, id composable.FullID) error {
	return sdkerrors.ErrNotSupported
}

func (k Keeper) UpdateNFT(ctx sdk.Context, classID string, nft composable.NFT) error {
	return sdkerrors.ErrNotSupported
}
