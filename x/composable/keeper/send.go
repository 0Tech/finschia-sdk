package keeper

import (
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

func (k Keeper) Send(ctx sdk.Context, sender, recipient sdk.AccAddress, nft composable.NFT) error {
	if err := k.validateOwner(ctx, nft, sender); err != nil {
		return err
	}
	k.setOwner(ctx, nft, recipient)

	return nil
}

func (k Keeper) Attach(ctx sdk.Context, owner sdk.AccAddress, subject, target composable.NFT) error {
	if err := k.validateOwner(ctx, subject, owner); err != nil {
		return sdkerrors.Wrap(err, "subject")
	}

	var ancestors []composable.NFT
	k.iterateAncestors(ctx, target, func(id composable.NFT) {
		ancestors = append(ancestors, id)
	})

	root := ancestors[len(ancestors)-1]
	if err := k.validateRootOwner(ctx, root, owner); err != nil {
		return sdkerrors.Wrap(err, "target")
	}

	k.deleteOwner(ctx, subject)
	k.setParent(ctx, subject, target)

	diff := 1 + k.getNumDescendants(ctx, subject)
	limit := k.GetParams(ctx).MaxDescendants
	for _, nft := range ancestors {
		old := k.getNumDescendants(ctx, nft)
		new := old + diff
		if new > limit {
			return composable.ErrTooManyDescendants
		}
		k.updateNumDescendants(ctx, nft, new)
	}

	return nil
}

func (k Keeper) Detach(ctx sdk.Context, owner sdk.AccAddress, nft composable.NFT) error {
	parent, err := k.getParent(ctx, nft)
	if err != nil {
		return err
	}

	var ancestors []composable.NFT
	k.iterateAncestors(ctx, *parent, func(nft composable.NFT) {
		ancestors = append(ancestors, nft)
	})

	root := ancestors[len(ancestors)-1]
	if err := k.validateRootOwner(ctx, root, owner); err != nil {
		return err
	}

	k.deleteParent(ctx, nft)
	k.setOwner(ctx, nft, owner)

	diff := -(1 + k.getNumDescendants(ctx, nft))
	for _, nft := range ancestors {
		old := k.getNumDescendants(ctx, nft)
		new := old + diff
		k.updateNumDescendants(ctx, nft, new)
	}

	return nil
}

func (k Keeper) validateRootOwner(ctx sdk.Context, root composable.NFT, owner sdk.AccAddress) error {
	if real, err := k.GetRootOwner(ctx, root); err != nil || !owner.Equals(real) {
		return composable.ErrInsufficientNFT.Wrap(root.String())
	}

	return nil
}

func (k Keeper) validateOwner(ctx sdk.Context, nft composable.NFT, owner sdk.AccAddress) error {
	if real, err := k.getOwner(ctx, nft); err != nil || !owner.Equals(real) {
		return sdkerrors.Wrap(composable.ErrInsufficientNFT.Wrap("not owns root nft"), nft.String())
	}

	return nil
}

func (k Keeper) GetRootOwner(ctx sdk.Context, nft composable.NFT) (*sdk.AccAddress, error) {
	root := k.getRoot(ctx, nft)
	owner, err := k.getOwner(ctx, root)
	if err != nil {
		return nil, composable.ErrNFTNotFound.Wrap(nft.String())
	}

	return owner, nil
}

func (k Keeper) getRoot(ctx sdk.Context, nft composable.NFT) composable.NFT {
	var res composable.NFT
	k.iterateAncestors(ctx, nft, func(nft composable.NFT) {
		res = nft
	})

	return res
}

func (k Keeper) iterateAncestors(ctx sdk.Context, nft composable.NFT, fn func(nft composable.NFT)) {
	var err error
	for iter := &nft; err == nil; iter, err = k.getParent(ctx, *iter) {
		fn(*iter)
	}
}

func (k Keeper) hasOwner(ctx sdk.Context, nft composable.NFT) error {
	_, err := k.getOwnerBytes(ctx, nft)
	return err
}

func (k Keeper) getOwner(ctx sdk.Context, nft composable.NFT) (*sdk.AccAddress, error) {
	bz, err := k.getOwnerBytes(ctx, nft)
	if err != nil {
		return nil, err
	}

	var owner sdk.AccAddress
	if err := owner.Unmarshal(bz); err != nil {
		panic(err)
	}

	return &owner, nil
}

func (k Keeper) getOwnerBytes(ctx sdk.Context, nft composable.NFT) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := ownerKey(nft.ClassId, nft.Id)

	bz := store.Get(key)
	if bz == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound.Wrap("owner"), nft.String())
	}

	return bz, nil
}

func (k Keeper) setOwner(ctx sdk.Context, nft composable.NFT, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := ownerKey(nft.ClassId, nft.Id)

	bz, err := owner.Marshal()
	if err != nil {
		panic(err)
	}

	store.Set(key, bz)
}

func (k Keeper) deleteOwner(ctx sdk.Context, nft composable.NFT) {
	store := ctx.KVStore(k.storeKey)
	key := ownerKey(nft.ClassId, nft.Id)

	store.Delete(key)
}

func (k Keeper) getParent(ctx sdk.Context, nft composable.NFT) (*composable.NFT, error) {
	bz, err := k.getParentBytes(ctx, nft)
	if err != nil {
		return nil, err
	}

	var parent composable.NFT
	k.cdc.MustUnmarshal(bz, &parent)

	return &parent, nil
}

func (k Keeper) getParentBytes(ctx sdk.Context, nft composable.NFT) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := parentKey(nft.ClassId, nft.Id)

	bz := store.Get(key)
	if bz == nil {
		return nil, composable.ErrParentNotFound.Wrap(nft.String())
	}

	return bz, nil
}

func (k Keeper) setParent(ctx sdk.Context, nft composable.NFT, parent composable.NFT) {
	store := ctx.KVStore(k.storeKey)
	key := parentKey(nft.ClassId, nft.Id)

	bz := k.cdc.MustMarshal(&parent)

	store.Set(key, bz)
}

func (k Keeper) deleteParent(ctx sdk.Context, nft composable.NFT) {
	store := ctx.KVStore(k.storeKey)
	key := parentKey(nft.ClassId, nft.Id)

	store.Delete(key)
}

func (k Keeper) getNumDescendants(ctx sdk.Context, nft composable.NFT) uint32 {
	bz, err := k.getNumDescendantsBytes(ctx, nft)
	if err != nil {
		return 0
	}

	numDescendants := uint32(bz[0])

	return numDescendants
}

func (k Keeper) getNumDescendantsBytes(ctx sdk.Context, nft composable.NFT) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := numDescendantsKey(nft.ClassId, nft.Id)

	bz := store.Get(key)
	if bz == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound.Wrap("num descendants"), nft.String())
	}

	return bz, nil
}

func (k Keeper) updateNumDescendants(ctx sdk.Context, nft composable.NFT, numDescendants uint32) {
	if numDescendants == 0 {
		k.deleteNumDescendants(ctx, nft)
	}
	k.setNumDescendants(ctx, nft, numDescendants)
}

func (k Keeper) setNumDescendants(ctx sdk.Context, nft composable.NFT, numDescendants uint32) {
	store := ctx.KVStore(k.storeKey)
	key := numDescendantsKey(nft.ClassId, nft.Id)

	bz := []byte{byte(numDescendants)}

	store.Set(key, bz)
}

func (k Keeper) deleteNumDescendants(ctx sdk.Context, nft composable.NFT) {
	store := ctx.KVStore(k.storeKey)
	key := numDescendantsKey(nft.ClassId, nft.Id)

	store.Delete(key)
}
