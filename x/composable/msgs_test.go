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
				ClassId:   tc.classID,
				Id:        sdk.OneUint(),
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
				Owner:          tc.owner.String(),
				SubjectClassId: tc.subjectClassID,
				SubjectId:      sdk.OneUint(),
				TargetClassId:  tc.targetClassID,
				TargetId:       sdk.OneUint(),
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
				Owner:   tc.owner.String(),
				ClassId: tc.classID,
				Id:      sdk.OneUint(),
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
	uri := "https://ipfs.io/ipfs/tIBeTianfOX"
	hash := "tIBeTianfOX"

	testCases := map[string]struct {
		owner sdk.AccAddress
		uri   string
		err   error
	}{
		"valid msg": {
			owner: addr,
			uri:   uri,
		},
		"invalid owner": {
			uri: uri,
			err: sdkerrors.ErrInvalidAddress,
		},
		"invalid uri hash": {
			owner: addr,
			err:   composable.ErrInvalidUriHash,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := composable.MsgNewClass{
				Owner:   tc.owner.String(),
				Uri:     tc.uri,
				UriHash: hash,
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
	addr := createAddresses(1, "addr")[0]
	uri := "https://ipfs.io/ipfs/tIBeTianfOX"
	hash := "tIBeTianfOX"

	testCases := map[string]struct {
		classID   string
		uri       string
		recipient sdk.AccAddress
		err       error
	}{
		"valid msg": {
			classID:   classID,
			uri:       uri,
			recipient: addr,
		},
		"invalid class id": {
			uri:       uri,
			recipient: addr,
			err:       composable.ErrInvalidClassID,
		},
		"invalid uri hash": {
			classID:   classID,
			recipient: addr,
			err:       composable.ErrInvalidUriHash,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := composable.MsgMintNFT{
				ClassId:   tc.classID,
				Uri:       tc.uri,
				UriHash:   hash,
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
				Owner:   tc.owner.String(),
				ClassId: tc.classID,
				Id:      sdk.OneUint(),
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

	testCases := map[string]struct {
		classID string
		id      sdk.Uint
		err     error
	}{
		"valid msg": {
			classID: classID,
			id:      sdk.OneUint(),
		},
		"invalid class id": {
			id:  sdk.OneUint(),
			err: composable.ErrInvalidClassID,
		},
		"invalid id": {
			classID: classID,
			id:      sdk.ZeroUint(),
			err:     composable.ErrInvalidNFTID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			msg := composable.MsgUpdateNFT{
				ClassId: tc.classID,
				Id:      tc.id,
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
	hash := "tIBeTianfOX"

	testCase := []struct {
		msg legacytx.LegacyMsg
		out string
	}{
		{
			&composable.MsgSend{
				Sender:    addrs[0].String(),
				Recipient: addrs[1].String(),
				ClassId:   classIDs[0],
				Id:        id,
			},
			`{"class_id":"link1vdkxzumnxq3kswxp","id":"1","recipient":"link1v9jxgu33p9vj2k","sender":"link1v9jxgu3sunc8hy"}`,
		},
		{
			&composable.MsgAttach{
				Owner:          addrs[0].String(),
				SubjectClassId: classIDs[0],
				SubjectId:      id,
				TargetClassId:  classIDs[1],
				TargetId:       id,
			},
			`{"owner":"link1v9jxgu3sunc8hy","subject_class_id":"link1vdkxzumnxq3kswxp","subject_id":"1","target_class_id":"link1vdkxzumnxy7ujgfm","target_id":"1"}`,
		},
		{
			&composable.MsgDetach{
				Owner:   addrs[0].String(),
				ClassId: classIDs[0],
				Id:      id,
			},
			`{"class_id":"link1vdkxzumnxq3kswxp","id":"1","owner":"link1v9jxgu3sunc8hy"}`,
		},
		{
			&composable.MsgNewClass{
				Owner:   addrs[0].String(),
				Uri:     uri,
				UriHash: hash,
			},
			`{"owner":"link1v9jxgu3sunc8hy","uri":"https://ipfs.io/ipfs/tIBeTianfOX","uri_hash":"tIBeTianfOX"}`,
		},
		{
			&composable.MsgUpdateClass{
				ClassId: classIDs[0],
				Uri:     uri,
				UriHash: hash,
			},
			`{"class_id":"link1vdkxzumnxq3kswxp","uri":"https://ipfs.io/ipfs/tIBeTianfOX","uri_hash":"tIBeTianfOX"}`,
		},
		{
			&composable.MsgMintNFT{
				ClassId:   classIDs[0],
				Uri:       uri,
				UriHash:   hash,
				Recipient: addrs[0].String(),
			},
			`{"class_id":"link1vdkxzumnxq3kswxp","recipient":"link1v9jxgu3sunc8hy","uri":"https://ipfs.io/ipfs/tIBeTianfOX","uri_hash":"tIBeTianfOX"}`,
		},
		{
			&composable.MsgBurnNFT{
				Owner:   addrs[0].String(),
				ClassId: classIDs[0],
				Id:      id,
			},
			`{"class_id":"link1vdkxzumnxq3kswxp","id":"1","owner":"link1v9jxgu3sunc8hy"}`,
		},
		{
			&composable.MsgUpdateNFT{
				ClassId: classIDs[0],
				Id:      id,
				Uri:     uri,
				UriHash: hash,
			},
			`{"class_id":"link1vdkxzumnxq3kswxp","id":"1","uri":"https://ipfs.io/ipfs/tIBeTianfOX","uri_hash":"tIBeTianfOX"}`,
		},
	}

	for _, tc := range testCase {
		name := sdk.MsgTypeURL(tc.msg)
		t.Run(name, func(t *testing.T) {
			require.Equal(t, composable.RouterKey, tc.msg.Route())
			require.Equal(t, sdk.MsgTypeURL(tc.msg), tc.msg.Type())

			out := tc.msg.GetSignBytes()
			require.Equal(t, tc.out, string(out))
		})
	}
}
