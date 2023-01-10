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
	traitID := "uri"
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
				Classes: []composable.GenesisClass{
					{
						Id: classIDs[0],
						Traits: []composable.Trait{
							{
								Id: traitID,
							},
						},
						LastMintedNftId: sdk.NewUint(3),
						Nfts: []composable.GenesisNFT{
							{
								Id:    sdk.NewUint(1),
								Owner: addr.String(),
							},
							{
								Id:    sdk.NewUint(2),
								Owner: addr.String(),
							},
							{
								Id: sdk.NewUint(3),
								Parent: &composable.NFT{
									ClassId: classIDs[1],
									Id:      sdk.NewUint(2),
								},
							},
						},
					},
					{
						Id:              classIDs[1],
						LastMintedNftId: sdk.NewUint(3),
						Nfts: []composable.GenesisNFT{
							{
								Id:    sdk.NewUint(1),
								Owner: addr.String(),
							},
							{
								Id:    sdk.NewUint(2),
								Owner: addr.String(),
							},
							{
								Id: sdk.NewUint(3),
								Parent: &composable.NFT{
									ClassId: classIDs[0],
									Id:      sdk.NewUint(2),
								},
							},
						},
					},
				},
			},
		},
		"invalid class id": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Classes: []composable.GenesisClass{
					{
						LastMintedNftId: sdk.NewUint(3),
					},
				},
			},
			err: composable.ErrInvalidClassID,
		},
		"invalid trait id": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Classes: []composable.GenesisClass{
					{
						Id: classIDs[0],
						Traits: []composable.Trait{
							{},
						},
						LastMintedNftId: sdk.NewUint(3),
					},
				},
			},
			err: composable.ErrInvalidTraitID,
		},
		"duplicate class": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Classes: []composable.GenesisClass{
					{
						Id:              classIDs[0],
						LastMintedNftId: sdk.NewUint(3),
					},
					{
						Id:              classIDs[0],
						LastMintedNftId: sdk.NewUint(3),
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		"invalid nft id": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Classes: []composable.GenesisClass{
					{
						Id:              classIDs[0],
						LastMintedNftId: sdk.NewUint(3),
						Nfts: []composable.GenesisNFT{
							{
								Id:    sdk.NewUint(0),
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
				Classes: []composable.GenesisClass{
					{
						Id:              classIDs[0],
						LastMintedNftId: sdk.NewUint(3),
						Nfts: []composable.GenesisNFT{
							{
								Id:    sdk.NewUint(3),
								Owner: addr.String(),
							},
							{
								Id:    sdk.NewUint(2),
								Owner: addr.String(),
							},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		"greater than last minted nft id": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Classes: []composable.GenesisClass{
					{
						Id:              classIDs[0],
						LastMintedNftId: sdk.NewUint(0),
						Nfts: []composable.GenesisNFT{
							{
								Id:    sdk.NewUint(1),
								Owner: addr.String(),
							},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		"invalid property id": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Classes: []composable.GenesisClass{
					{
						Id: classIDs[0],
						Traits: []composable.Trait{
							{
								Id: traitID,
							},
						},
						LastMintedNftId: sdk.NewUint(3),
						Nfts: []composable.GenesisNFT{
							{
								Id: sdk.NewUint(1),
								Properties: []composable.Property{
									{},
								},
								Owner: addr.String(),
							},
						},
					},
				},
			},
			err: composable.ErrInvalidTraitID,
		},
		"no corresponding trait": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Classes: []composable.GenesisClass{
					{
						Id: classIDs[0],
						Traits: []composable.Trait{
							{
								Id: traitID,
							},
						},
						LastMintedNftId: sdk.NewUint(3),
						Nfts: []composable.GenesisNFT{
							{
								Id: sdk.NewUint(1),
								Properties: []composable.Property{
									{
										Id: "nosuchid",
									},
								},
								Owner: addr.String(),
							},
						},
					},
				},
			},
			err: composable.ErrTraitNotFound,
		},
		"both owner and parent": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Classes: []composable.GenesisClass{
					{
						Id:              classIDs[0],
						LastMintedNftId: sdk.NewUint(3),
						Nfts: []composable.GenesisNFT{
							{
								Id:    sdk.NewUint(2),
								Owner: addr.String(),
							},
							{
								Id:    sdk.NewUint(3),
								Owner: addr.String(),
								Parent: &composable.NFT{
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
				Classes: []composable.GenesisClass{
					{
						Id:              classIDs[0],
						LastMintedNftId: sdk.NewUint(3),
						Nfts: []composable.GenesisNFT{
							{
								Id:    sdk.NewUint(1),
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
				Classes: []composable.GenesisClass{
					{
						Id:              classIDs[0],
						LastMintedNftId: sdk.NewUint(3),
						Nfts: []composable.GenesisNFT{
							{
								Id:    sdk.NewUint(2),
								Owner: addr.String(),
							},
							{
								Id: sdk.NewUint(3),
								Parent: &composable.NFT{
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
