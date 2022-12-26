package keeper

import (
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

func (k Keeper) NewClass(ctx sdk.Context, class composable.Class) error {
	return sdkerrors.ErrNotSupported
}

func (k Keeper) UpdateClass(ctx sdk.Context, class composable.Class) error {
	return sdkerrors.ErrNotSupported
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
