// nolint:dupl
package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/line/link/x/token/internal/keeper"
	"github.com/line/link/x/token/internal/types"
)

func handleMsgMint(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgMint) (*sdk.Result, error) {
	err := keeper.MintToken(ctx, msg.Amount, msg.From, msg.To)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgBurn(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgBurn) (*sdk.Result, error) {
	err := keeper.BurnToken(ctx, msg.Amount, msg.From)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
