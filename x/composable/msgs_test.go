package composable_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/auth/legacy/legacytx"
	"github.com/line/lbm-sdk/x/composable"
)

func TestMsgSend(t *testing.T) {
	addrs := createAddresses(2, "addr")
	classID := createClassIDs(1, "class")[0]

	testCases := map[string]struct {
		sender    sdk.AccAddress
		recipient sdk.AccAddress
		classID   string
		err       error
	}{
		"valid msg": {
			sender:    addrs[0],
			recipient: addrs[1],
			classID:   classID,
		},
		"invalid sender": {
			recipient: addrs[1],
			classID:   classID,
			err:       sdkerrors.ErrInvalidAddress,
		},
		"invalid recipient": {
			sender:  addrs[0],
			classID: classID,
			err:     sdkerrors.ErrInvalidAddress,
		},
		"invalid class id": {
			sender:    addrs[0],
			recipient: addrs[1],
			err:       composable.ErrInvalidClassID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := composable.MsgSend{
				Sender:    tc.sender.String(),
				Recipient: tc.recipient.String(),
				Nft: composable.NFT{
					ClassId: tc.classID,
					Id:      sdk.OneUint(),
				},
			}

			err := msg.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			require.Equal(t, []sdk.AccAddress{tc.sender}, msg.GetSigners())
		})
	}
}

func TestMsgAttach(t *testing.T) {
	addr := createAddresses(1, "addr")[0]
	classIDs := createClassIDs(2, "class")

	testCases := map[string]struct {
		owner          sdk.AccAddress
		subjectClassID string
		targetClassID  string
		err            error
	}{
		"valid msg": {
			owner:          addr,
			subjectClassID: classIDs[0],
			targetClassID:  classIDs[1],
		},
		"invalid owner": {
			subjectClassID: classIDs[0],
			targetClassID:  classIDs[1],
			err:            sdkerrors.ErrInvalidAddress,
		},
		"invalid subject class id": {
			owner:         addr,
			targetClassID: classIDs[1],
			err:           composable.ErrInvalidClassID,
		},
		"invalid target class id": {
			owner:          addr,
			subjectClassID: classIDs[0],
			err:            composable.ErrInvalidClassID,
		},
		"target == subject": {
			owner:          addr,
			subjectClassID: classIDs[0],
			targetClassID:  classIDs[0],
			err:            composable.ErrInvalidComposition,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := composable.MsgAttach{
				Owner: tc.owner.String(),
				Subject: composable.NFT{
					ClassId: tc.subjectClassID,
					Id:      sdk.OneUint(),
				},
				Target: composable.NFT{
					ClassId: tc.targetClassID,
					Id:      sdk.OneUint(),
				},
			}

			err := msg.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			require.Equal(t, []sdk.AccAddress{tc.owner}, msg.GetSigners())
		})
	}
}

func TestMsgDetach(t *testing.T) {
	addr := createAddresses(1, "addr")[0]
	classID := createClassIDs(1, "class")[0]

	testCases := map[string]struct {
		owner   sdk.AccAddress
		classID string
		err     error
	}{
		"valid msg": {
			owner:   addr,
			classID: classID,
		},
		"invalid owner": {
			classID: classID,
			err:     sdkerrors.ErrInvalidAddress,
		},
		"invalid class id": {
			owner: addr,
			err:   composable.ErrInvalidClassID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := composable.MsgDetach{
				Owner: tc.owner.String(),
				Nft: composable.NFT{
					ClassId: tc.classID,
					Id:      sdk.OneUint(),
				},
			}

			err := msg.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			require.Equal(t, []sdk.AccAddress{tc.owner}, msg.GetSigners())
		})
	}
}

func TestMsgNewClass(t *testing.T) {
	addr := createAddresses(1, "addr")[0]
	traitID := "uri"

	testCases := map[string]struct {
		owner   sdk.AccAddress
		traitID string
		err     error
	}{
		"valid msg": {
			owner:   addr,
			traitID: traitID,
		},
		"invalid owner": {
			traitID: traitID,
			err:     sdkerrors.ErrInvalidAddress,
		},
		"invalid trait id": {
			owner: addr,
			err:   composable.ErrInvalidTraitID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := composable.MsgNewClass{
				Owner: tc.owner.String(),
				Traits: []composable.Trait{
					{
						Id: tc.traitID,
					},
				},
			}

			err := msg.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			require.Equal(t, []sdk.AccAddress{tc.owner}, msg.GetSigners())
		})
	}
}

func TestMsgUpdateClass(t *testing.T) {
	classID := createClassIDs(1, "class")[0]

	testCases := map[string]struct {
		classID string
		err     error
	}{
		"valid msg": {
			classID: classID,
		},
		"invalid class id": {
			err: composable.ErrInvalidClassID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := composable.MsgUpdateClass{
				ClassId: tc.classID,
			}

			err := msg.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			owner := composable.ClassOwner(tc.classID)
			require.Equal(t, []sdk.AccAddress{owner}, msg.GetSigners())
		})
	}
}

func TestMsgMintNFT(t *testing.T) {
	classID := createClassIDs(1, "class")[0]
	traitID := "uri"
	addr := createAddresses(1, "addr")[0]

	testCases := map[string]struct {
		classID   string
		traitID   string
		recipient sdk.AccAddress
		err       error
	}{
		"valid msg": {
			classID:   classID,
			traitID:   traitID,
			recipient: addr,
		},
		"invalid class id": {
			recipient: addr,
			traitID:   traitID,
			err:       composable.ErrInvalidClassID,
		},
		"invalid trait id": {
			classID: classID,
			err:     composable.ErrInvalidTraitID,
		},
		"invalid recipient": {
			classID: classID,
			traitID: traitID,
			err:     sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := composable.MsgMintNFT{
				ClassId: tc.classID,
				Properties: []composable.Property{
					{
						Id: tc.traitID,
					},
				},
				Recipient: tc.recipient.String(),
			}

			err := msg.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			owner := composable.ClassOwner(tc.classID)
			require.Equal(t, []sdk.AccAddress{owner}, msg.GetSigners())
		})
	}
}

func TestMsgBurnNFT(t *testing.T) {
	addr := createAddresses(1, "addr")[0]
	classID := createClassIDs(1, "class")[0]

	testCases := map[string]struct {
		owner   sdk.AccAddress
		classID string
		err     error
	}{
		"valid msg": {
			owner:   addr,
			classID: classID,
		},
		"invalid owner": {
			classID: classID,
			err:     sdkerrors.ErrInvalidAddress,
		},
		"invalid class id": {
			owner: addr,
			err:   composable.ErrInvalidClassID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := composable.MsgBurnNFT{
				Owner: tc.owner.String(),
				Nft: composable.NFT{
					ClassId: tc.classID,
					Id:      sdk.OneUint(),
				},
			}

			err := msg.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			require.Equal(t, []sdk.AccAddress{tc.owner}, msg.GetSigners())
		})
	}
}

func TestMsgUpdateNFT(t *testing.T) {
	classID := createClassIDs(1, "class")[0]
	traitID := "uri"

	testCases := map[string]struct {
		classID  string
		traitIDs []string
		err      error
	}{
		"valid msg": {
			classID: classID,
			traitIDs: []string{
				traitID,
			},
		},
		"invalid class id": {
			traitIDs: []string{
				traitID,
			},
			err: composable.ErrInvalidClassID,
		},
		"empty properties": {
			classID: classID,
			err:     sdkerrors.ErrInvalidRequest,
		},
		"invalid trait id": {
			classID: classID,
			traitIDs: []string{
				"",
			},
			err: composable.ErrInvalidTraitID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			properties := make([]composable.Property, len(tc.traitIDs))
			for i, id := range tc.traitIDs {
				properties[i] = composable.Property{
					Id: id,
				}
			}

			msg := composable.MsgUpdateNFT{
				Nft: composable.NFT{
					ClassId: tc.classID,
					Id:      sdk.OneUint(),
				},
				Properties: properties,
			}

			err := msg.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
			if tc.err != nil {
				return
			}

			owner := composable.ClassOwner(tc.classID)
			require.Equal(t, []sdk.AccAddress{owner}, msg.GetSigners())
		})
	}
}

func TestLegacyMsg(t *testing.T) {
	addrs := createAddresses(2, "addr")
	classIDs := createClassIDs(2, "class")
	id := sdk.OneUint()
	uri := "https://ipfs.io/ipfs/tIBeTianfOX"

	testCase := []struct {
		msg legacytx.LegacyMsg
		out string
	}{
		{
			&composable.MsgSend{
				Sender:    addrs[0].String(),
				Recipient: addrs[1].String(),
				Nft: composable.NFT{
					ClassId: classIDs[0],
					Id:      id,
				},
			},
			`{"nft":{"class_id":"link1vdkxzumnxq3kswxp","id":"1"},"recipient":"link1v9jxgu33p9vj2k","sender":"link1v9jxgu3sunc8hy"}`,
		},
		{
			&composable.MsgAttach{
				Owner: addrs[0].String(),
				Subject: composable.NFT{
					ClassId: classIDs[0],
					Id:      id,
				},
				Target: composable.NFT{
					ClassId: classIDs[1],
					Id:      id,
				},
			},
			`{"owner":"link1v9jxgu3sunc8hy","subject":{"class_id":"link1vdkxzumnxq3kswxp","id":"1"},"target":{"class_id":"link1vdkxzumnxy7ujgfm","id":"1"}}`,
		},
		{
			&composable.MsgDetach{
				Owner: addrs[0].String(),
				Nft: composable.NFT{
					ClassId: classIDs[0],
					Id:      id,
				},
			},
			`{"nft":{"class_id":"link1vdkxzumnxq3kswxp","id":"1"},"owner":"link1v9jxgu3sunc8hy"}`,
		},
		{
			&composable.MsgNewClass{
				Owner: addrs[0].String(),
				Traits: []composable.Trait{
					{
						Id:      "uri",
						Mutable: true,
					},
				},
			},
			`{"owner":"link1v9jxgu3sunc8hy","traits":[{"id":"uri","mutable":true}]}`,
		},
		{
			&composable.MsgUpdateClass{
				ClassId: classIDs[0],
			},
			`{"class_id":"link1vdkxzumnxq3kswxp"}`,
		},
		{
			&composable.MsgMintNFT{
				ClassId: classIDs[0],
				Properties: []composable.Property{
					{
						Id:   "uri",
						Fact: uri,
					},
				},
				Recipient: addrs[0].String(),
			},
			`{"class_id":"link1vdkxzumnxq3kswxp","properties":[{"fact":"https://ipfs.io/ipfs/tIBeTianfOX","id":"uri"}],"recipient":"link1v9jxgu3sunc8hy"}`,
		},
		{
			&composable.MsgBurnNFT{
				Owner: addrs[0].String(),
				Nft: composable.NFT{
					ClassId: classIDs[0],
					Id:      id,
				},
			},
			`{"nft":{"class_id":"link1vdkxzumnxq3kswxp","id":"1"},"owner":"link1v9jxgu3sunc8hy"}`,
		},
		{
			&composable.MsgUpdateNFT{
				Nft: composable.NFT{
					ClassId: classIDs[0],
					Id:      id,
				},
				Properties: []composable.Property{
					{
						Id:   "uri",
						Fact: uri,
					},
				},
			},
			`{"nft":{"class_id":"link1vdkxzumnxq3kswxp","id":"1"},"properties":[{"fact":"https://ipfs.io/ipfs/tIBeTianfOX","id":"uri"}]}`,
		},
	}

	for _, tc := range testCase {
		name := sdk.MsgTypeURL(tc.msg)[1:]
		t.Run(name, func(t *testing.T) {
			require.Equal(t, composable.RouterKey, tc.msg.Route())
			require.Equal(t, sdk.MsgTypeURL(tc.msg), tc.msg.Type())

			out := tc.msg.GetSignBytes()
			require.Equal(t, tc.out, string(out))
		})
	}
}
