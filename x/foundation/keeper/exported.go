package keeper

import (
	"cosmossdk.io/core/store"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/Finschia/finschia-sdk/x/foundation"
	"github.com/Finschia/finschia-sdk/x/foundation/keeper/internal"
)

type Keeper interface {
	GetAuthority() string
	Accept(ctx sdk.Context, grantee sdk.AccAddress, msg sdk.Msg) error

	InitGenesis(ctx sdk.Context, gs *foundation.GenesisState) error
	ExportGenesis(ctx sdk.Context) *foundation.GenesisState
}

type keeper struct {
	impl internal.Keeper
}

func NewKeeper(
	cdc codec.Codec,
	storeService store.KVStoreService,
	router baseapp.MessageRouter,
	authKeeper foundation.AuthKeeper,
	bankKeeper foundation.BankKeeper,
	feeCollectorName string,
	config foundation.Config,
	authority string,
) Keeper {
	return &keeper{
		impl: internal.NewKeeper(
			cdc,
			storeService,
			router,
			authKeeper,
			bankKeeper,
			feeCollectorName,
			config,
			authority,
		),
	}
}

// GetAuthority returns the x/foundation module's authority.
func (k keeper) GetAuthority() string {
	return k.impl.GetAuthority()
}

func (k keeper) Accept(ctx sdk.Context, grantee sdk.AccAddress, msg sdk.Msg) error {
	return k.impl.Accept(ctx, grantee, msg)
}

func (k keeper) InitGenesis(ctx sdk.Context, gs *foundation.GenesisState) error {
	return k.impl.InitGenesis(ctx, gs)
}

func (k keeper) ExportGenesis(ctx sdk.Context) *foundation.GenesisState {
	return k.impl.ExportGenesis(ctx)
}

func NewMsgServer(k Keeper) foundation.MsgServer {
	impl := k.(*keeper).impl
	return internal.NewMsgServer(impl)
}

func NewQueryServer(k Keeper) foundation.QueryServer {
	impl := k.(*keeper).impl
	return internal.NewQueryServer(impl)
}

func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	impl := k.(*keeper).impl
	internal.RegisterInvariants(ir, impl)
}

func BeginBlocker(ctx sdk.Context, k Keeper) error {
	impl := k.(*keeper).impl
	return internal.BeginBlocker(ctx, impl)
}

func EndBlocker(ctx sdk.Context, k Keeper) error {
	impl := k.(*keeper).impl
	return internal.EndBlocker(ctx, impl)
}

func NewFoundationProposalsHandler(k Keeper) govtypes.Handler {
	impl := k.(*keeper).impl
	return internal.NewFoundationProposalsHandler(impl)
}

type Migrator = internal.Migrator

func NewMigrator(k Keeper, subspace paramstypes.Subspace) Migrator {
	impl := k.(*keeper).impl
	return internal.NewMigrator(impl, subspace)
}
