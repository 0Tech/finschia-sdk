package keeper

import (
	"fmt"

	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/foundation"
)

const (
	treasuryInvariant = "treasury"
)

func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(foundation.ModuleName, treasuryInvariant, TreasuryInvariant(k))
}

func TreasuryInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		// cache, we don't want to write changes
		ctx, _ = ctx.CacheContext()

		treasuryAcc := k.authKeeper.GetModuleAccount(ctx, foundation.TreasuryName)
		balance := k.bankKeeper.GetAllBalances(ctx, treasuryAcc.GetAddress())

		treasury := k.GetTreasury(ctx)
		broken := !treasury.IsEqual(sdk.NewDecCoinsFromCoins(balance...))

		return sdk.FormatInvariant(foundation.ModuleName, treasuryInvariant, fmt.Sprintf("amount of coins in the treasury; expected %s, got %s", treasury, balance)), broken
	}
}
