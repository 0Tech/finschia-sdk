package composable_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

func TestGenesisState(t *testing.T) {
	classIDs := createClassIDs(2, "class")
	addr := createAddresses(1, "addr")[0]

	testCases := map[string]struct {
		s   composable.GenesisState
		err error
	}{
		"default genesis": {
			s: *composable.DefaultGenesisState(),
		},
		"all features": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
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
								Owner: addr.String(),
							},
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(3),
								},
								Parent: &composable.FullID{
									ClassId: classIDs[1],
									Id:      sdk.NewUint(2),
								},
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(3),
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
								Owner: addr.String(),
							},
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(3),
								},
								Parent: &composable.FullID{
									ClassId: classIDs[0],
									Id:      sdk.NewUint(2),
								},
							},
						},
					},
				},
			},
		},
		"invalid class": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class:      composable.Class{},
						PreviousId: sdk.NewUint(3),
					},
				},
			},
			err: composable.ErrInvalidClassID,
		},
		"duplicate class": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
					},
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		"invalid nft": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
						NftStates: []composable.NFTState{
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(0),
								},
								Owner: addr.String(),
							},
						},
					},
				},
			},
			err: composable.ErrInvalidNFTID,
		},
		"unsorted nfts": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
						NftStates: []composable.NFTState{
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(3),
								},
								Owner: addr.String(),
							},
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(2),
								},
								Owner: addr.String(),
							},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		"greater than previous id": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(0),
						NftStates: []composable.NFTState{
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(1),
								},
								Owner: addr.String(),
							},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		"has both owner and parent": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
						NftStates: []composable.NFTState{
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(2),
								},
								Owner: addr.String(),
							},
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(3),
								},
								Owner: addr.String(),
								Parent: &composable.FullID{
									ClassId: classIDs[0],
									Id:      sdk.NewUint(2),
								},
							},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		"invalid owner": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
						NftStates: []composable.NFTState{
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(1),
								},
								Owner: "invalid",
							},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		"invalid parent": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
						NftStates: []composable.NFTState{
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(2),
								},
								Owner: addr.String(),
							},
							{
								Nft: composable.NFT{
									Id: sdk.NewUint(3),
								},
								Parent: &composable.FullID{
									Id: sdk.NewUint(2),
								},
							},
						},
					},
				},
			},
			err: composable.ErrInvalidClassID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.s.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
		})
	}
}
