package keeper_test

import (
	"testing"

	ocproto "github.com/line/ostracon/proto/ostracon/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/line/lbm-sdk/simapp"
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/composable"
)

func TestImportExportGenesis(t *testing.T) {
	checkTx := false
	app := simapp.Setup(checkTx)

	ctx := app.BaseApp.NewContext(checkTx, ocproto.Header{})
	keeper := app.ComposableKeeper

	classIDs := createClassIDs(2, "class")
	addr := createAddresses(1, "addr")[0]

	testCases := map[string]struct {
		gs *composable.GenesisState
	}{
		"default": {
			gs: composable.DefaultGenesisState(),
		},
		"no compositions": {
			gs: &composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(100),
						NftStates: []composable.NFTState{
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(1),
								},
								Owner: addr.String(),
							},
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(10),
								},
								Owner: addr.String(),
							},
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(100),
								},
								Owner: addr.String(),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(10000),
						NftStates: []composable.NFTState{
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(100),
								},
								Owner: addr.String(),
							},
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(1000),
								},
								Owner: addr.String(),
							},
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(10000),
								},
								Owner: addr.String(),
							},
						},
					},
				},
			},
		},
		"with compositions": {
			gs: &composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(2),
						NftStates: []composable.NFTState{
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(1),
								},
								Owner: addr.String(),
							},
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(2),
								},
								Parent: &composable.FullID{
									ClassId: classIDs[1],
									Id:      sdk.NewUint(1),
								},
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(2),
						NftStates: []composable.NFTState{
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(1),
								},
								Owner: addr.String(),
							},
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(2),
								},
								Parent: &composable.FullID{
									ClassId: classIDs[0],
									Id:      sdk.NewUint(1),
								},
							},
						},
					},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, _ := ctx.CacheContext()

			err := tc.gs.ValidateBasic()
			assert.NoError(t, err)

			err = keeper.InitGenesis(ctx, tc.gs)
			require.NoError(t, err)

			exported := keeper.ExportGenesis(ctx)
			require.Equal(t, tc.gs, exported)
		})
	}
}
