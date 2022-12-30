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

	k.setPreviousID(ctx, class.Id, sdk.ZeroUint())

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

func (k Keeper) GetPreviousID(ctx sdk.Context, classID string) sdk.Uint {
	bz, err := k.getPreviousIDBytes(ctx, classID)
	if err != nil {
		panic(err)
	}

	var id sdk.Uint
	if err := id.Unmarshal(bz); err != nil {
		panic(err)
	}

	return id
}

func (k Keeper) getPreviousIDBytes(ctx sdk.Context, classID string) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := previousIDKey(classID)

	bz := store.Get(key)
	if bz == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound.Wrap("previous id"), classID)
	}

	return bz, nil
}

func (k Keeper) setPreviousID(ctx sdk.Context, classID string, id sdk.Uint) {
	store := ctx.KVStore(k.storeKey)
	key := previousIDKey(classID)

	bz, err := id.Marshal()
	if err != nil {
		panic(err)
	}

	store.Set(key, bz)
}

func (k Keeper) iterateClasses(ctx sdk.Context, fn func(class composable.Class)) {
	prefix := classKeyPrefix

	k.iterateImpl(ctx, prefix, func(_, value []byte) {
		var class composable.Class
		k.cdc.MustUnmarshal(value, &class)

		fn(class)
	})
}

func (k Keeper) MintNFT(ctx sdk.Context, owner sdk.AccAddress, classID string, nft composable.NFT) (*sdk.Uint, error) {
	if err := k.hasClass(ctx, classID); err != nil {
		return nil, err
	}

	nft.Id = k.GetPreviousID(ctx, classID).Incr()
	k.setPreviousID(ctx, classID, nft.Id)

	fullID := composable.FullID{
		ClassId: classID,
		Id:      nft.Id,
	}

	if err := k.hasNFT(ctx, fullID); err == nil {
		panic(sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrap("nft already exists"), fullID.String()))
	}
	k.setNFT(ctx, classID, nft)

	k.setOwner(ctx, fullID, owner)

	return &nft.Id, nil
}

func (k Keeper) BurnNFT(ctx sdk.Context, owner sdk.AccAddress, id composable.FullID) error {
	if err := k.validateOwner(ctx, id, owner); err != nil {
		return err
	}
	k.deleteOwner(ctx, id)

	if err := k.hasNFT(ctx, id); err != nil {
		panic(err)
	}
	k.deleteNFT(ctx, id)

	// TODO: prune children

	return nil
}

func (k Keeper) UpdateNFT(ctx sdk.Context, classID string, nft composable.NFT) error {
	fullID := composable.FullID{
		ClassId: classID,
		Id:      nft.Id,
	}

	if err := k.hasNFT(ctx, fullID); err != nil {
		return err
	}
	k.setNFT(ctx, classID, nft)

	return nil
}

func (k Keeper) hasNFT(ctx sdk.Context, id composable.FullID) error {
	_, err := k.getNFTBytes(ctx, id)
	return err
}

func (k Keeper) GetNFT(ctx sdk.Context, id composable.FullID) (*composable.NFT, error) {
	bz, err := k.getNFTBytes(ctx, id)
	if err != nil {
		return nil, err
	}

	var nft composable.NFT
	k.cdc.MustUnmarshal(bz, &nft)

	return &nft, nil
}

func (k Keeper) getNFTBytes(ctx sdk.Context, id composable.FullID) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := nftKey(id.ClassId, id.Id)

	bz := store.Get(key)
	if bz == nil {
		return nil, composable.ErrNFTNotFound.Wrap(id.String())
	}

	return bz, nil
}

func (k Keeper) setNFT(ctx sdk.Context, classID string, nft composable.NFT) {
	store := ctx.KVStore(k.storeKey)
	key := nftKey(classID, nft.Id)

	bz := k.cdc.MustMarshal(&nft)

	store.Set(key, bz)
}

func (k Keeper) deleteNFT(ctx sdk.Context, id composable.FullID) {
	store := ctx.KVStore(k.storeKey)
	key := nftKey(id.ClassId, id.Id)

	store.Delete(key)
}

func (k Keeper) iterateNFTsOfClass(ctx sdk.Context, classID string, fn func(nft composable.NFT)) {
	prefix := nftKeyPrefixOfClass(classID)

	k.iterateImpl(ctx, prefix, func(_, value []byte) {
		var nft composable.NFT
		k.cdc.MustUnmarshal(value, &nft)

		fn(nft)
	})
}
