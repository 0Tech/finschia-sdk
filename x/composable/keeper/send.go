package keeper

import (
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

func (k Keeper) Send(ctx sdk.Context, sender, recipient sdk.AccAddress, id composable.FullID) error {
	if owner, err := k.getOwner(ctx, id); err != nil || !sender.Equals(owner) {
		return sdkerrors.Wrap(composable.ErrInsufficientNFT.Wrap("not owns root nft"), id.String())
	}
	k.setOwner(ctx, id, recipient)

	return nil
}

func (k Keeper) Attach(ctx sdk.Context, owner sdk.AccAddress, subjectID, targetID composable.FullID) error {
	return sdkerrors.ErrNotSupported
}

func (k Keeper) Detach(ctx sdk.Context, owner sdk.AccAddress, id composable.FullID) error {
	return sdkerrors.ErrNotSupported
}

func (k Keeper) hasOwner(ctx sdk.Context, id composable.FullID) error {
	_, err := k.getOwnerBytes(ctx, id)
	return err
}

func (k Keeper) getOwner(ctx sdk.Context, id composable.FullID) (*sdk.AccAddress, error) {
	bz, err := k.getOwnerBytes(ctx, id)
	if err != nil {
		return nil, err
	}

	var owner sdk.AccAddress
	if err := owner.Unmarshal(bz); err != nil {
		panic(err)
	}

	return &owner, nil
}

func (k Keeper) getOwnerBytes(ctx sdk.Context, id composable.FullID) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := ownerKey(id.ClassId, id.Id)

	bz := store.Get(key)
	if bz == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound.Wrap("owner"), id.String())
	}

	return bz, nil
}

func (k Keeper) setOwner(ctx sdk.Context, id composable.FullID, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := ownerKey(id.ClassId, id.Id)

	bz, err := owner.Marshal()
	if err != nil {
		panic(err)
	}

	store.Set(key, bz)
}

func (k Keeper) deleteOwner(ctx sdk.Context, id composable.FullID) {
	store := ctx.KVStore(k.storeKey)
	key := ownerKey(id.ClassId, id.Id)

	store.Delete(key)
}
