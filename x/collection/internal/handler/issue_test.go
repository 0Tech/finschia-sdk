package handler

import (
	"testing"

	"github.com/line/link/x/collection/internal/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	linktype "github.com/line/link/types"
)

func TestHandleMsgIssueFT(t *testing.T) {
	t.Log("implement me - ", t.Name())
}

func TestHandleMsgIssueNFT(t *testing.T) {
	t.Log("implement me - ", t.Name())
}

func TestHandlerIssueFT(t *testing.T) {
	ctx, h := cacheKeeper()

	{
		msg := types.NewMsgCreateCollection(addr1, defaultName, defaultSymbol)
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
	}

	{
		msg := types.NewMsgIssueFT(addr1, defaultName, defaultSymbol, defaultTokenURI, sdk.NewInt(defaultAmount), sdk.NewInt(defaultDecimals), true)
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
	}
	{
		msg := types.NewMsgIssueFT(addr1, defaultName, defaultSymbol, defaultTokenURI, sdk.NewInt(defaultAmount), sdk.NewInt(defaultDecimals), true)
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
	}
	{
		msg := types.NewMsgIssueFT(addr2, defaultName, defaultSymbol, defaultTokenURI, sdk.NewInt(defaultAmount), sdk.NewInt(defaultDecimals), true)
		res := h(ctx, msg)
		require.False(t, res.Code.IsOK())
		require.Equal(t, types.DefaultCodespace, res.Codespace)
		require.Equal(t, types.CodeTokenPermission, res.Code)
	}

	permission := types.Permission{
		Action:   "issue",
		Resource: defaultSymbol,
	}

	{
		msg := types.NewMsgGrantPermission(addr1, addr2, permission)
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
	}
	{
		msg := types.NewMsgIssueFT(addr2, defaultName, defaultSymbol, defaultTokenURI, sdk.NewInt(defaultAmount), sdk.NewInt(defaultDecimals), true)
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
	}
	{
		msg := types.NewMsgRevokePermission(addr1, permission)
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
	}
	{
		msg := types.NewMsgIssueFT(addr1, defaultName, defaultSymbol, defaultTokenURI, sdk.NewInt(defaultAmount), sdk.NewInt(defaultDecimals), true)
		res := h(ctx, msg)
		require.False(t, res.Code.IsOK())
		require.Equal(t, types.DefaultCodespace, res.Codespace)
		require.Equal(t, types.CodeTokenPermission, res.Code)
	}
}

func TestHandlerIssueNFT(t *testing.T) {
	ctx, h := cacheKeeper()

	{
		msg := types.NewMsgCreateCollection(addr1, defaultName, defaultSymbol)
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
	}

	{
		//Expect token type is 1001
		{
			msg := types.NewMsgIssueNFT(addr1, defaultSymbol)
			res := h(ctx, msg)
			require.True(t, res.Code.IsOK())
		}
		//Expect token type is 1002
		{
			msg := types.NewMsgIssueNFT(addr1, defaultSymbol)
			res := h(ctx, msg)
			require.True(t, res.Code.IsOK())
		}
		{
			msg := types.NewMsgMintNFT(addr1, addr1, defaultName, defaultSymbol, defaultTokenURI, defaultTokenType)
			res := h(ctx, msg)
			require.True(t, res.Code.IsOK())
		}
		{
			msg := types.NewMsgMintNFT(addr1, addr2, defaultName, defaultSymbol, defaultTokenURI, defaultTokenType)
			res := h(ctx, msg)
			require.True(t, res.Code.IsOK())
		}
		{
			mintPermission := types.Permission{
				Action:   "mint",
				Resource: defaultSymbol + defaultTokenType,
			}
			{
				msg := types.NewMsgGrantPermission(addr1, addr2, mintPermission)
				res := h(ctx, msg)
				require.True(t, res.Code.IsOK())
			}
			{
				msg := types.NewMsgMintNFT(addr2, addr2, defaultName, defaultSymbol, defaultTokenURI, defaultTokenType)
				res := h(ctx, msg)
				require.True(t, res.Code.IsOK())
			}
			{
				msg := types.NewMsgRevokePermission(addr1, mintPermission)
				res := h(ctx, msg)
				require.True(t, res.Code.IsOK())
			}
			{
				msg := types.NewMsgMintNFT(addr1, addr1, defaultName, defaultSymbol, defaultTokenURI, defaultTokenType)
				res := h(ctx, msg)
				require.False(t, res.Code.IsOK())
				require.Equal(t, types.DefaultCodespace, res.Codespace)
				require.Equal(t, types.CodeTokenPermission, res.Code)
			}
		}
	}

	permission := types.Permission{
		Action:   "issue",
		Resource: defaultSymbol,
	}

	{
		msg := types.NewMsgGrantPermission(addr1, addr2, permission)
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
	}

	//Expect token type is 1003
	{
		msg := types.NewMsgIssueNFT(addr2, defaultSymbol)
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
	}
	{
		msg := types.NewMsgMintNFT(addr2, addr2, defaultName, defaultSymbol, defaultTokenURI, "1003")
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
	}
	{
		msg := types.NewMsgRevokePermission(addr1, permission)
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
	}
	{
		msg := types.NewMsgIssueNFT(addr1, defaultSymbol)
		res := h(ctx, msg)
		require.False(t, res.Code.IsOK())
		require.Equal(t, types.DefaultCodespace, res.Codespace)
		require.Equal(t, types.CodeTokenPermission, res.Code)
	}
}

func TestEvents(t *testing.T) {
	ctx, h := cacheKeeper()

	{
		msg := types.NewMsgCreateCollection(addr1, defaultName, defaultSymbol)
		require.NoError(t, msg.ValidateBasic())
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())

		e := sdk.Events{
			sdk.NewEvent("message", sdk.NewAttribute("module", "collection")),
			sdk.NewEvent("message", sdk.NewAttribute("sender", addr1.String())),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("to", addr1.String())),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_resource", defaultSymbol)),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_action", "issue")),
			sdk.NewEvent("create_collection", sdk.NewAttribute("name", defaultName)),
			sdk.NewEvent("create_collection", sdk.NewAttribute("symbol", defaultSymbol)),
			sdk.NewEvent("create_collection", sdk.NewAttribute("owner", addr1.String())),
		}
		verifyEventFunc(t, e, res.Events)
	}

	{
		symbolWithID := defaultSymbol + defaultTokenIDFT
		msg := types.NewMsgIssueFT(addr1, defaultName, defaultSymbol, defaultTokenURI, sdk.NewInt(defaultAmount), sdk.NewInt(defaultDecimals), true)
		require.NoError(t, msg.ValidateBasic())
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())

		e := sdk.Events{
			sdk.NewEvent("message", sdk.NewAttribute("module", "collection")),
			sdk.NewEvent("message", sdk.NewAttribute("sender", addr1.String())),
			sdk.NewEvent("issue_ft", sdk.NewAttribute("name", defaultName)),
			sdk.NewEvent("issue_ft", sdk.NewAttribute("symbol", defaultSymbol)),
			sdk.NewEvent("issue_ft", sdk.NewAttribute("token_id", defaultTokenIDFT)),
			sdk.NewEvent("issue_ft", sdk.NewAttribute("owner", addr1.String())),
			sdk.NewEvent("issue_ft", sdk.NewAttribute("amount", sdk.NewInt(defaultAmount).String())),
			sdk.NewEvent("issue_ft", sdk.NewAttribute("mintable", "true")),
			sdk.NewEvent("issue_ft", sdk.NewAttribute("decimals", sdk.NewInt(defaultDecimals).String())),
			sdk.NewEvent("issue_ft", sdk.NewAttribute("token_uri", defaultTokenURI)),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("to", addr1.String())),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_resource", symbolWithID)),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_action", types.ModifyAction)),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("to", addr1.String())),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_resource", symbolWithID)),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_action", "mint")),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("to", addr1.String())),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_resource", symbolWithID)),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_action", "burn")),
		}
		verifyEventFunc(t, e, res.Events)
	}

	{
		symbolWithID := defaultSymbol + defaultTokenIDFT
		msg := types.NewMsgMintFT(addr1, addr1, linktype.NewCoinWithTokenIDs(linktype.NewCoinWithTokenID(defaultSymbol, defaultTokenIDFT, sdk.NewInt(defaultAmount))))
		require.NoError(t, msg.ValidateBasic())
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
		e := sdk.Events{
			sdk.NewEvent("message", sdk.NewAttribute("module", "collection")),
			sdk.NewEvent("message", sdk.NewAttribute("sender", addr1.String())),
			sdk.NewEvent("mint_ft", sdk.NewAttribute("from", addr1.String())),
			sdk.NewEvent("mint_ft", sdk.NewAttribute("to", addr1.String())),
			sdk.NewEvent("mint_ft", sdk.NewAttribute("amount", sdk.NewInt(defaultAmount).String()+symbolWithID)),
		}
		verifyEventFunc(t, e, res.Events)
	}

	{
		symbolWithID := defaultSymbol + defaultTokenIDFT
		msg := types.NewMsgBurnFT(addr1, linktype.NewCoinWithTokenIDs(linktype.NewCoinWithTokenID(defaultSymbol, defaultTokenIDFT, sdk.NewInt(defaultAmount))))
		require.NoError(t, msg.ValidateBasic())
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
		e := sdk.Events{
			sdk.NewEvent("message", sdk.NewAttribute("module", "collection")),
			sdk.NewEvent("message", sdk.NewAttribute("sender", addr1.String())),
			sdk.NewEvent("burn_ft", sdk.NewAttribute("from", addr1.String())),
			sdk.NewEvent("burn_ft", sdk.NewAttribute("amount", sdk.NewInt(defaultAmount).String()+symbolWithID)),
		}
		verifyEventFunc(t, e, res.Events)
	}

	{
		msg := types.NewMsgIssueNFT(addr1, defaultSymbol)
		require.NoError(t, msg.ValidateBasic())
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
		e := sdk.Events{
			sdk.NewEvent("message", sdk.NewAttribute("module", "collection")),
			sdk.NewEvent("message", sdk.NewAttribute("sender", addr1.String())),
			sdk.NewEvent("issue_nft", sdk.NewAttribute("symbol", defaultSymbol)),
			sdk.NewEvent("issue_nft", sdk.NewAttribute("token_type", defaultTokenType)),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("to", addr1.String())),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_resource", defaultSymbol+defaultTokenType)),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_action", types.MintAction)),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("to", addr1.String())),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_resource", defaultSymbol+defaultTokenType)),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_action", types.BurnAction)),
		}
		verifyEventFunc(t, e, res.Events)
	}

	{
		symbolWithID := defaultSymbol + defaultTokenID1
		msg := types.NewMsgMintNFT(addr1, addr1, defaultName, defaultSymbol, defaultTokenURI, defaultTokenType)
		require.NoError(t, msg.ValidateBasic())
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())

		e := sdk.Events{
			sdk.NewEvent("message", sdk.NewAttribute("module", "collection")),
			sdk.NewEvent("message", sdk.NewAttribute("sender", addr1.String())),
			sdk.NewEvent("mint_nft", sdk.NewAttribute("name", defaultName)),
			sdk.NewEvent("mint_nft", sdk.NewAttribute("symbol", defaultSymbol)),
			sdk.NewEvent("mint_nft", sdk.NewAttribute("token_id", defaultTokenID1)),
			sdk.NewEvent("mint_nft", sdk.NewAttribute("from", addr1.String())),
			sdk.NewEvent("mint_nft", sdk.NewAttribute("to", addr1.String())),
			sdk.NewEvent("mint_nft", sdk.NewAttribute("token_uri", defaultTokenURI)),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("to", addr1.String())),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_resource", symbolWithID)),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_action", types.ModifyAction)),
		}
		verifyEventFunc(t, e, res.Events)
	}

	permission := types.Permission{
		Action:   "issue",
		Resource: defaultSymbol,
	}

	{
		msg := types.NewMsgGrantPermission(addr1, addr2, permission)
		require.NoError(t, msg.ValidateBasic())
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())

		e := sdk.Events{
			sdk.NewEvent("message", sdk.NewAttribute("module", "collection")),
			sdk.NewEvent("message", sdk.NewAttribute("sender", addr1.String())),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("from", addr1.String())),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("to", addr2.String())),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_resource", permission.GetResource())),
			sdk.NewEvent("grant_perm", sdk.NewAttribute("perm_action", permission.GetAction())),
		}
		verifyEventFunc(t, e, res.Events)
	}
	{
		msg := types.NewMsgRevokePermission(addr1, permission)
		require.NoError(t, msg.ValidateBasic())
		res := h(ctx, msg)
		require.True(t, res.Code.IsOK())
		e := sdk.Events{
			sdk.NewEvent("message", sdk.NewAttribute("module", "collection")),
			sdk.NewEvent("message", sdk.NewAttribute("sender", addr1.String())),
			sdk.NewEvent("revoke_perm", sdk.NewAttribute("from", addr1.String())),
			sdk.NewEvent("revoke_perm", sdk.NewAttribute("perm_resource", permission.GetResource())),
			sdk.NewEvent("revoke_perm", sdk.NewAttribute("perm_action", permission.GetAction())),
		}
		verifyEventFunc(t, e, res.Events)
	}
}
