package keeper

import (
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/composable"
)

func (k Keeper) InitGenesis(ctx sdk.Context, gs *composable.GenesisState) error {
	return nil
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *composable.GenesisState {
	return &composable.GenesisState{}
}
