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
	ErrTraitNotFound      = sdkerrors.Register(composableCodespace, 8, "trait not found")
	ErrTraitImmutable     = sdkerrors.Register(composableCodespace, 9, "trait immutable")
	ErrNFTNotFound        = sdkerrors.Register(composableCodespace, 10, "nft not found")
	ErrInsufficientNFT    = sdkerrors.Register(composableCodespace, 11, "insufficient nft")
	ErrParentNotFound     = sdkerrors.Register(composableCodespace, 12, "parent not found")
	ErrTooManyDescendants = sdkerrors.Register(composableCodespace, 13, "too many descendants")
)
