package types

var (
	EventTypeIssueFT              = "issue_ft"
	EventTypeIssueNFT             = "issue_nft"
	EventTypeMintFT               = "mint_ft"
	EventTypeBurnFT               = "burn_ft"
	EventTypeMintNFT              = "mint_nft"
	EventTypeModifyTokenURI       = "modify_token_uri_token" /* #nosec */
	EventTypeGrantPermToken       = "grant_perm"
	EventTypeRevokePermToken      = "revoke_perm"
	EventTypeCreateCollection     = "create_collection"
	EventTypeAttachToken          = "attach" /* #nosec */
	EventTypeDetachToken          = "detach" /* #nosec */
	EventTypeAttachFrom           = "attach_from"
	EventTypeDetachFrom           = "detach_from"
	EventTypeTransfer             = "transfer"
	EventTypeTransferFT           = "transfer_ft"
	EventTypeTransferNFT          = "transfer_nft"
	EventTypeTransferFTFrom       = "transfer_ft_from"
	EventTypeTransferNFTFrom      = "transfer_nft_from"
	EventTypeOperationTransferNFT = "operation_transfer_nft"
	EventTypeApproveCollection    = "approve_collection"
	EventTypeDisapproveCollection = "disapprove_collection"
	EventTypeBurnNFT              = "burn_nft"
	EventTypeBurnFTFrom           = "burn_ft_from"
	EventTypeBurnNFTFrom          = "burn_nft_from"
	EventTypeOperationBurnNFT     = "operation_burn_nft"

	AttributeKeyName        = "name"
	AttributeKeySymbol      = "symbol"
	AttributeKeyTokenID     = "token_id"
	AttributeKeyOwner       = "owner"
	AttributeKeyAmount      = "amount"
	AttributeKeyDecimals    = "decimals"
	AttributeKeyTokenURI    = "token_uri"
	AttributeKeyMintable    = "mintable"
	AttributeKeyTokenType   = "token_type"
	AttributeKeyFrom        = "from"
	AttributeKeyTo          = "to"
	AttributeKeyResource    = "perm_resource"
	AttributeKeyAction      = "perm_action"
	AttributeKeyToTokenID   = "to_token_id"
	AttributeKeyFromTokenID = "from_token_id"
	AttributeKeyApprover    = "approver"
	AttributeKeyProxy       = "proxy"

	AttributeValueCategory = ModuleName
)
