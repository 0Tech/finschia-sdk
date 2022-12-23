package keeper

import (
	"context"

	// sdk "github.com/line/lbm-sdk/types"
	// sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

type msgServer struct {
	keeper Keeper
}

var _ composable.MsgServer = (*msgServer)(nil)

// NewMsgServer returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServer(keeper Keeper) composable.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

// Send defines a method to send an nft from one account to another account.
func (s msgServer) Send(c context.Context, req *composable.MsgSend) (*composable.MsgSendResponse, error) {
	d := composable.UnimplementedMsgServer{}
	return d.Send(c, req)
}

// Attach defines a method to attach a root nft to another nft.
func (s msgServer) Attach(c context.Context, req *composable.MsgAttach) (*composable.MsgAttachResponse, error) {
	d := composable.UnimplementedMsgServer{}
	return d.Attach(c, req)
}

// Detach defines a method to detach an nft from another nft.
func (s msgServer) Detach(c context.Context, req *composable.MsgDetach) (*composable.MsgDetachResponse, error) {
	d := composable.UnimplementedMsgServer{}
	return d.Detach(c, req)
}

// NewClass defines a method to create a class.
func (s msgServer) NewClass(c context.Context, req *composable.MsgNewClass) (*composable.MsgNewClassResponse, error) {
	d := composable.UnimplementedMsgServer{}
	return d.NewClass(c, req)
}

// UpdateClass defines a method to update a class.
func (s msgServer) UpdateClass(c context.Context, req *composable.MsgUpdateClass) (*composable.MsgUpdateClassResponse, error) {
	d := composable.UnimplementedMsgServer{}
	return d.UpdateClass(c, req)
}

// MintNFT defines a method to mint an nft.
func (s msgServer) MintNFT(c context.Context, req *composable.MsgMintNFT) (*composable.MsgMintNFTResponse, error) {
	d := composable.UnimplementedMsgServer{}
	return d.MintNFT(c, req)
}

// BurnNFT defines a method to burn an nft.
func (s msgServer) BurnNFT(c context.Context, req *composable.MsgBurnNFT) (*composable.MsgBurnNFTResponse, error) {
	d := composable.UnimplementedMsgServer{}
	return d.BurnNFT(c, req)
}

// UpdateNFT defines a method to update an nft.
func (s msgServer) UpdateNFT(c context.Context, req *composable.MsgUpdateNFT) (*composable.MsgUpdateNFTResponse, error) {
	d := composable.UnimplementedMsgServer{}
	return d.UpdateNFT(c, req)
}
