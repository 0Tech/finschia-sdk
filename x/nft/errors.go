package nft

import (
	sdkerrors "github.com/line/lbm-sdk/types/errors"
)

// x/nft module sentinel errors
var (
	ErrClassExists    = sdkerrors.Register(ModuleName, 3, "nft class already exists")
	ErrClassNotExists = sdkerrors.Register(ModuleName, 4, "nft class does not exist")
	ErrNFTExists      = sdkerrors.Register(ModuleName, 5, "nft already exists")
	ErrNFTNotExists   = sdkerrors.Register(ModuleName, 6, "nft does not exist")
	ErrEmptyClassID   = sdkerrors.Register(ModuleName, 7, "empty class id")
	ErrEmptyNFTID     = sdkerrors.Register(ModuleName, 8, "empty nft id")
)
