package keeper

import (
	"context"

	sdk "github.com/line/lbm-sdk/types"
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
	ctx := sdk.UnwrapSDKContext(c)

	sender := sdk.MustAccAddressFromBech32(req.Sender)
	recipient := sdk.MustAccAddressFromBech32(req.Recipient)

	if err := s.keeper.Send(ctx, sender, recipient, req.Nft); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&composable.EventSend{
		Sender:   req.Sender,
		Receiver: req.Recipient,
		Nft:      req.Nft,
	}); err != nil {
		panic(err)
	}

	return &composable.MsgSendResponse{}, nil
}

// Attach defines a method to attach a root nft to another nft.
func (s msgServer) Attach(c context.Context, req *composable.MsgAttach) (*composable.MsgAttachResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	owner := sdk.MustAccAddressFromBech32(req.Owner)

	if err := s.keeper.Attach(ctx, owner, req.Subject, req.Target); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&composable.EventAttach{
		Owner:   req.Owner,
		Subject: req.Subject,
		Target:  req.Target,
	}); err != nil {
		panic(err)
	}

	return &composable.MsgAttachResponse{}, nil
}

// Detach defines a method to detach an nft from another nft.
func (s msgServer) Detach(c context.Context, req *composable.MsgDetach) (*composable.MsgDetachResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	owner := sdk.MustAccAddressFromBech32(req.Owner)

	if err := s.keeper.Detach(ctx, owner, req.Nft); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&composable.EventDetach{
		Owner: req.Owner,
		Nft:   req.Nft,
	}); err != nil {
		panic(err)
	}

	return &composable.MsgDetachResponse{}, nil
}

// NewClass defines a method to create a class.
func (s msgServer) NewClass(c context.Context, req *composable.MsgNewClass) (*composable.MsgNewClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	owner := sdk.MustAccAddressFromBech32(req.Owner)
	id := composable.ClassIDFromOwner(owner)
	class := composable.Class{
		Id: id,
	}

	// TODO: traits

	if err := s.keeper.NewClass(ctx, class); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&composable.EventNewClass{
		Class: class,
		Data:  req.Data,
	}); err != nil {
		panic(err)
	}

	return &composable.MsgNewClassResponse{}, nil
}

// UpdateClass defines a method to update a class.
func (s msgServer) UpdateClass(c context.Context, req *composable.MsgUpdateClass) (*composable.MsgUpdateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	class := composable.Class{
		Id: req.ClassId,
	}

	// TODO: data

	if err := s.keeper.UpdateClass(ctx, class); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&composable.EventUpdateClass{
		Class: class,
		Data:  req.Data,
	}); err != nil {
		panic(err)
	}

	return &composable.MsgUpdateClassResponse{}, nil
}

// MintNFT defines a method to mint an nft.
func (s msgServer) MintNFT(c context.Context, req *composable.MsgMintNFT) (*composable.MsgMintNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	recipient := sdk.MustAccAddressFromBech32(req.Recipient)

	// TODO: traits

	_, err := s.keeper.MintNFT(ctx, recipient, req.ClassId)
	if err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&composable.EventMintNFT{
		Recipient: req.Recipient,
	}); err != nil {
		panic(err)
	}

	return &composable.MsgMintNFTResponse{}, nil
}

// BurnNFT defines a method to burn an nft.
func (s msgServer) BurnNFT(c context.Context, req *composable.MsgBurnNFT) (*composable.MsgBurnNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	owner := sdk.MustAccAddressFromBech32(req.Owner)

	if err := s.keeper.BurnNFT(ctx, owner, req.Nft); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&composable.EventBurnNFT{
		Owner: req.Owner,
		Nft:   req.Nft,
	}); err != nil {
		panic(err)
	}

	return &composable.MsgBurnNFTResponse{}, nil
}

// UpdateNFT defines a method to update an nft.
func (s msgServer) UpdateNFT(c context.Context, req *composable.MsgUpdateNFT) (*composable.MsgUpdateNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// TODO: property

	if err := s.keeper.UpdateNFT(ctx, req.Nft); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(&composable.EventUpdateNFT{
		Nft: req.Nft,
	}); err != nil {
		panic(err)
	}

	return &composable.MsgUpdateNFTResponse{}, nil
}
