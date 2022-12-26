package module

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	abci "github.com/line/ostracon/abci/types"
	"github.com/spf13/cobra"

	"github.com/line/lbm-sdk/client"
	"github.com/line/lbm-sdk/codec"
	codectypes "github.com/line/lbm-sdk/codec/types"
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/types/module"
	"github.com/line/lbm-sdk/x/composable"
	"github.com/line/lbm-sdk/x/composable/keeper"
)

var _ module.AppModuleBasic = AppModuleBasic{}

// AppModuleBasic defines the basic application module used by the module.
type AppModuleBasic struct{}

// Name returns the ModuleName
func (AppModuleBasic) Name() string {
	return composable.ModuleName
}

// RegisterLegacyAminoCodec registers the types on the LegacyAmino codec
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// DefaultGenesis returns default genesis state as raw bytes for the module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(composable.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var gs composable.GenesisState
	if err := cdc.UnmarshalJSON(bz, &gs); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", composable.ModuleName, err)
	}

	return gs.ValidateBasic()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := composable.RegisterQueryHandlerClient(context.Background(), mux, composable.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// GetQueryCmd returns the cli query commands for the module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	// TODO
	// return cli.NewQueryCmd()
	return nil
}

// GetTxCmd returns the transaction commands for the module
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	// TODO
	// return cli.NewTxCmd()
	return nil
}

func (b AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	composable.RegisterInterfaces(registry)
}

//____________________________________________________________________________

var _ module.AppModule = AppModule{}

// AppModule implements an application module for the module.
type AppModule struct {
	AppModuleBasic

	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		keeper: keeper,
	}
}

// RegisterInvariants does nothing, there are no invariants to enforce
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

// Route is empty, as we do not handle Messages (just proposals)
func (AppModule) Route() sdk.Route { return sdk.Route{} }

// QuerierRoute returns the route we respond to for abci queries
func (AppModule) QuerierRoute() string { return "" }

// LegacyQuerierHandler registers a query handler to respond to the module-specific queries
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers a GRPC query service to respond to the
// module-specific GRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	composable.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServer(am.keeper))
	composable.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServer(am.keeper))

	// m := keeper.NewMigrator(am.keeper)
	// migrations := map[uint64]func(sdk.Context) error{}
	// for ver, handler := range migrations {
	// 	if err := cfg.RegisterMigration(composable.ModuleName, ver, handler); err != nil {
	// 		panic(fmt.Sprintf("failed to migrate x/%s from version %d to %d: %v", composable.ModuleName, ver, ver+1, err))
	// 	}
	// }
}

// InitGenesis performs genesis initialization for the module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var gs composable.GenesisState
	cdc.MustUnmarshalJSON(data, &gs)

	if err := am.keeper.InitGenesis(ctx, &gs); err != nil {
		panic(err)
	}

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(gs)
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }
