package keeper

import (
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

func (k Keeper) Send(ctx sdk.Context, sender, recipient sdk.AccAddress, id composable.FullID) error {
	if err := k.validateOwner(ctx, id, sender); err != nil {
		return err
	}
	k.setOwner(ctx, id, recipient)

	return nil
}

func (k Keeper) Attach(ctx sdk.Context, owner sdk.AccAddress, subjectID, targetID composable.FullID) error {
	if err := k.validateOwner(ctx, subjectID, owner); err != nil {
		return sdkerrors.Wrap(err, "subject")
	}

	var ancestorIDs []composable.FullID
	k.iterateAncestors(ctx, targetID, func(id composable.FullID) {
		ancestorIDs = append(ancestorIDs, id)
	})

	rootID := ancestorIDs[len(ancestorIDs)-1]
	if err := k.validateRootOwner(ctx, rootID, owner); err != nil {
		return sdkerrors.Wrap(err, "target")
	}

	k.deleteOwner(ctx, subjectID)
	k.setParent(ctx, subjectID, targetID)

	diff := 1 + k.getNumDescendants(ctx, subjectID)
	limit := k.GetParams(ctx).MaxDescendants
	for _, id := range ancestorIDs {
		old := k.getNumDescendants(ctx, id)
		new := old + diff
		if new > limit {
			return composable.ErrTooManyDescendants
		}
		k.updateNumDescendants(ctx, id, new)
	}

	return nil
}

func (k Keeper) Detach(ctx sdk.Context, owner sdk.AccAddress, id composable.FullID) error {
	parentID, err := k.getParent(ctx, id)
	if err != nil {
		return err
	}

	var ancestorIDs []composable.FullID
	k.iterateAncestors(ctx, *parentID, func(id composable.FullID) {
		ancestorIDs = append(ancestorIDs, id)
	})

	rootID := ancestorIDs[len(ancestorIDs)-1]
	if err := k.validateRootOwner(ctx, rootID, owner); err != nil {
		return err
	}

	k.deleteParent(ctx, id)
	k.setOwner(ctx, id, owner)

	diff := -(1 + k.getNumDescendants(ctx, id))
	for _, id := range ancestorIDs {
		old := k.getNumDescendants(ctx, id)
		new := old + diff
		k.updateNumDescendants(ctx, id, new)
	}

	return nil
}

func (k Keeper) validateRootOwner(ctx sdk.Context, rootID composable.FullID, owner sdk.AccAddress) error {
	if real, err := k.GetRootOwner(ctx, rootID); err != nil || !owner.Equals(real) {
		return composable.ErrInsufficientNFT.Wrap(rootID.String())
	}

	return nil
}

func (k Keeper) validateOwner(ctx sdk.Context, id composable.FullID, owner sdk.AccAddress) error {
	if real, err := k.getOwner(ctx, id); err != nil || !owner.Equals(real) {
		return sdkerrors.Wrap(composable.ErrInsufficientNFT.Wrap("not owns root nft"), id.String())
	}

	return nil
}

func (k Keeper) GetRootOwner(ctx sdk.Context, id composable.FullID) (*sdk.AccAddress, error) {
	rootID := k.getRoot(ctx, id)
	owner, err := k.getOwner(ctx, rootID)
	if err != nil {
		return nil, err
	}

	return owner, nil
}

func (k Keeper) getRoot(ctx sdk.Context, id composable.FullID) composable.FullID {
	var res composable.FullID
	k.iterateAncestors(ctx, id, func(id composable.FullID) {
		res = id
	})

	return res
}

func (k Keeper) iterateAncestors(ctx sdk.Context, id composable.FullID, fn func(id composable.FullID)) {
	var err error
	for iter := &id; err == nil; iter, err = k.getParent(ctx, *iter) {
		fn(*iter)
	}
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

func (k Keeper) getParent(ctx sdk.Context, id composable.FullID) (*composable.FullID, error) {
	bz, err := k.getParentBytes(ctx, id)
	if err != nil {
		return nil, err
	}

	var parent composable.FullID
	k.cdc.MustUnmarshal(bz, &parent)

	return &parent, nil
}

func (k Keeper) getParentBytes(ctx sdk.Context, id composable.FullID) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := parentKey(id.ClassId, id.Id)

	bz := store.Get(key)
	if bz == nil {
		return nil, composable.ErrParentNotFound.Wrap(id.String())
	}

	return bz, nil
}

func (k Keeper) setParent(ctx sdk.Context, id composable.FullID, parent composable.FullID) {
	store := ctx.KVStore(k.storeKey)
	key := parentKey(id.ClassId, id.Id)

	bz := k.cdc.MustMarshal(&parent)

	store.Set(key, bz)
}

func (k Keeper) deleteParent(ctx sdk.Context, id composable.FullID) {
	store := ctx.KVStore(k.storeKey)
	key := parentKey(id.ClassId, id.Id)

	store.Delete(key)
}

func (k Keeper) getNumDescendants(ctx sdk.Context, id composable.FullID) uint32 {
	bz, err := k.getNumDescendantsBytes(ctx, id)
	if err != nil {
		return 0
	}

	numDescendants := uint32(bz[0])

	return numDescendants
}

func (k Keeper) getNumDescendantsBytes(ctx sdk.Context, id composable.FullID) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := numDescendantsKey(id.ClassId, id.Id)

	bz := store.Get(key)
	if bz == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound.Wrap("num descendants"), id.String())
	}

	return bz, nil
}

func (k Keeper) updateNumDescendants(ctx sdk.Context, id composable.FullID, numDescendants uint32) {
	if numDescendants == 0 {
		k.deleteNumDescendants(ctx, id)
	}
	k.setNumDescendants(ctx, id, numDescendants)
}

func (k Keeper) setNumDescendants(ctx sdk.Context, id composable.FullID, numDescendants uint32) {
	store := ctx.KVStore(k.storeKey)
	key := numDescendantsKey(id.ClassId, id.Id)

	bz := []byte{byte(numDescendants)}

	store.Set(key, bz)
}

func (k Keeper) deleteNumDescendants(ctx sdk.Context, id composable.FullID) {
	store := ctx.KVStore(k.storeKey)
	key := numDescendantsKey(id.ClassId, id.Id)

	store.Delete(key)
}
