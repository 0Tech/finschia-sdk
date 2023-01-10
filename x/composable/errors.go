package composable

import (
	sdkerrors "github.com/line/lbm-sdk/types/errors"
)

const composableCodespace = ModuleName

// x/nft module sentinel errors
var (
	ErrInvalidClassID     = sdkerrors.Register(composableCodespace, 2, "invalid class id")
	ErrInvalidTraitID     = sdkerrors.Register(composableCodespace, 3, "invalid trait id")
	ErrInvalidNFTID       = sdkerrors.Register(composableCodespace, 4, "invalid nft id")
	ErrInvalidComposition = sdkerrors.Register(composableCodespace, 5, "invalid composition request")
	ErrClassNotFound      = sdkerrors.Register(composableCodespace, 6, "nft class not found")
	ErrClassAlreadyExists = sdkerrors.Register(composableCodespace, 7, "nft class already exists")
	ErrNFTNotFound        = sdkerrors.Register(composableCodespace, 8, "nft not found")
	ErrInsufficientNFT    = sdkerrors.Register(composableCodespace, 9, "insufficient nft")
	ErrTooManyDescendants = sdkerrors.Register(composableCodespace, 10, "too many descendants")
	ErrParentNotFound     = sdkerrors.Register(composableCodespace, 11, "parent not found")
)
