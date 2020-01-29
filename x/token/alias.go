package token

import (
	"github.com/line/link/x/token/internal/keeper"
	"github.com/line/link/x/token/internal/types"
)

const (
	ModuleName = types.ModuleName
	StoreKey   = types.StoreKey
	RouterKey  = types.RouterKey

	DefaultCodespace = types.DefaultCodespace
)

type (
	Token  = types.Token
	Tokens = types.Tokens

	FT  = types.FT
	NFT = types.NFT

	Collection            = types.Collection
	Collections           = types.Collections
	CollectionWithTokens  = types.CollectionWithTokens
	CollectionsWithTokens = types.CollectionsWithTokens

	MsgIssue              = types.MsgIssue
	MsgIssueCollection    = types.MsgIssueCollection
	MsgIssueNFT           = types.MsgIssueNFT
	MsgIssueNFTCollection = types.MsgIssueNFTCollection
	MsgMint               = types.MsgMint
	MsgBurn               = types.MsgBurn
	MsgGrantPermission    = types.MsgGrantPermission
	MsgRevokePermission   = types.MsgRevokePermission
	MsgModifyTokenURI     = types.MsgModifyTokenURI
	MsgTransferFT         = types.MsgTransferFT
	MsgTransferIDFT       = types.MsgTransferIDFT
	MsgTransferNFT        = types.MsgTransferNFT
	MsgTransferIDNFT      = types.MsgTransferIDNFT

	MsgAttach = types.MsgAttach
	MsgDetach = types.MsgDetach

	PermissionI = types.PermissionI
	Permissions = types.Permissions

	Keeper = keeper.Keeper
)

var (
	NewFT    = types.NewFT
	NewNFT   = types.NewNFT
	NewIDFT  = types.NewIDFT
	NewIDNFT = types.NewIDNFT

	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec

	ErrTokenExist = types.ErrTokenExist

	NewIssuePermission          = types.NewIssuePermission
	NewModifyTokenURIPermission = types.NewModifyTokenURIPermission

	NewKeeper = keeper.NewKeeper
)
