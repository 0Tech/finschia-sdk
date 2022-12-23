package keeper

import (
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

func (k Keeper) Send(ctx sdk.Context, sender, recipient sdk.AccAddress, id composable.FullID) error {
	return sdkerrors.ErrNotSupported
}

func (k Keeper) Attach(ctx sdk.Context, owner sdk.AccAddress, subjectID, targetID composable.FullID) error {
	return sdkerrors.ErrNotSupported
}

func (k Keeper) Detach(ctx sdk.Context, owner sdk.AccAddress, id composable.FullID) error {
	return sdkerrors.ErrNotSupported
}
