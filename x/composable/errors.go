package composable

import (
	sdkerrors "github.com/line/lbm-sdk/types/errors"
)

const composableCodespace = ModuleName

// x/nft module sentinel errors
var (
	ErrInvalidClassID     = sdkerrors.Register(composableCodespace, 2, "invalid class id")
	ErrInvalidNFTID       = sdkerrors.Register(composableCodespace, 3, "invalid nft id")
	ErrInvalidUriHash     = sdkerrors.Register(composableCodespace, 4, "invalid uri hash")
	ErrInvalidComposition = sdkerrors.Register(composableCodespace, 5, "invalid composition request")
	ErrClassNotFound      = sdkerrors.Register(composableCodespace, 6, "nft class not found")
	ErrClassAlreadyExists = sdkerrors.Register(composableCodespace, 7, "nft class already exists")
	ErrNFTNotFound        = sdkerrors.Register(composableCodespace, 8, "nft not found")
)
