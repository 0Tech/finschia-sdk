package composable_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

func TestGenesisState(t *testing.T) {
	classIDs := createAddresses(3, "class")
	addrs := createAddresses(2, "addr")

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
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
				},
				Balances: []composable.Balance{
					{
						Owner: addrs[0],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[0],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
									sdk.NewUint(2),
								},
							},
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
								},
							},
						},
					},
					{
						Owner: addrs[1],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
								},
							},
						},
					},
				},
				Children: []composable.Children{
					{
						Subject: composable.FullID{
							ClassId: classIDs[0],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[1],
								Id:      sdk.NewUint(3),
							},
						},
					},
					{
						Subject: composable.FullID{
							ClassId: classIDs[1],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[0],
								Id:      sdk.NewUint(3),
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
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
				},
				Balances: []composable.Balance{
					{
						Owner: addrs[0],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[0],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
									sdk.NewUint(2),
								},
							},
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
								},
							},
						},
					},
					{
						Owner: addrs[1],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
								},
							},
						},
					},
				},
				Children: []composable.Children{
					{
						Subject: composable.FullID{
							ClassId: classIDs[0],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[1],
								Id:      sdk.NewUint(3),
							},
						},
					},
					{
						Subject: composable.FullID{
							ClassId: classIDs[1],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[0],
								Id:      sdk.NewUint(3),
							},
						},
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
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
				},
				Balances: []composable.Balance{
					{
						Owner: addrs[0],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[0],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
									sdk.NewUint(2),
								},
							},
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
								},
							},
						},
					},
					{
						Owner: addrs[1],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
								},
							},
						},
					},
				},
				Children: []composable.Children{
					{
						Subject: composable.FullID{
							ClassId: classIDs[0],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[1],
								Id:      sdk.NewUint(3),
							},
						},
					},
					{
						Subject: composable.FullID{
							ClassId: classIDs[1],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[0],
								Id:      sdk.NewUint(3),
							},
						},
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
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(0),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
				},
				Balances: []composable.Balance{
					{
						Owner: addrs[0],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[0],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
									sdk.NewUint(2),
								},
							},
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
								},
							},
						},
					},
					{
						Owner: addrs[1],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
								},
							},
						},
					},
				},
				Children: []composable.Children{
					{
						Subject: composable.FullID{
							ClassId: classIDs[0],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[1],
								Id:      sdk.NewUint(3),
							},
						},
					},
					{
						Subject: composable.FullID{
							ClassId: classIDs[1],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[0],
								Id:      sdk.NewUint(3),
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
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(3),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(1),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
				},
				Balances: []composable.Balance{
					{
						Owner: addrs[0],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[0],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
									sdk.NewUint(2),
								},
							},
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
								},
							},
						},
					},
					{
						Owner: addrs[1],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
								},
							},
						},
					},
				},
				Children: []composable.Children{
					{
						Subject: composable.FullID{
							ClassId: classIDs[0],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[1],
								Id:      sdk.NewUint(3),
							},
						},
					},
					{
						Subject: composable.FullID{
							ClassId: classIDs[1],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[0],
								Id:      sdk.NewUint(3),
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
						PreviousId: sdk.NewUint(2),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
				},
				Balances: []composable.Balance{
					{
						Owner: addrs[0],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[0],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
									sdk.NewUint(2),
								},
							},
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
								},
							},
						},
					},
					{
						Owner: addrs[1],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
								},
							},
						},
					},
				},
				Children: []composable.Children{
					{
						Subject: composable.FullID{
							ClassId: classIDs[0],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[1],
								Id:      sdk.NewUint(3),
							},
						},
					},
					{
						Subject: composable.FullID{
							ClassId: classIDs[1],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[0],
								Id:      sdk.NewUint(3),
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
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
				},
				Balances: []composable.Balance{
					{
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[0],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
									sdk.NewUint(2),
								},
							},
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
								},
							},
						},
					},
					{
						Owner: addrs[1],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
								},
							},
						},
					},
				},
				Children: []composable.Children{
					{
						Subject: composable.FullID{
							ClassId: classIDs[0],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[1],
								Id:      sdk.NewUint(3),
							},
						},
					},
					{
						Subject: composable.FullID{
							ClassId: classIDs[1],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[0],
								Id:      sdk.NewUint(3),
							},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		"duplicate owners": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
				},
				Balances: []composable.Balance{
					{
						Owner: addrs[0],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[0],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
									sdk.NewUint(2),
								},
							},
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
								},
							},
						},
					},
					{
						Owner: addrs[0],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
								},
							},
						},
					},
				},
				Children: []composable.Children{
					{
						Subject: composable.FullID{
							ClassId: classIDs[0],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[1],
								Id:      sdk.NewUint(3),
							},
						},
					},
					{
						Subject: composable.FullID{
							ClassId: classIDs[1],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[0],
								Id:      sdk.NewUint(3),
							},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		"class not found": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
				},
				Balances: []composable.Balance{
					{
						Owner: addrs[0],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[2],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
									sdk.NewUint(2),
								},
							},
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
								},
							},
						},
					},
					{
						Owner: addrs[1],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
								},
							},
						},
					},
				},
				Children: []composable.Children{
					{
						Subject: composable.FullID{
							ClassId: classIDs[0],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[1],
								Id:      sdk.NewUint(3),
							},
						},
					},
					{
						Subject: composable.FullID{
							ClassId: classIDs[1],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[0],
								Id:      sdk.NewUint(3),
							},
						},
					},
				},
			},
			err: composable.ErrClassNotFound,
		},
		"invalid nft id": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
				},
				Balances: []composable.Balance{
					{
						Owner: addrs[0],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[0],
								Ids: []sdk.Uint{
									sdk.NewUint(0),
									sdk.NewUint(2),
								},
							},
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
								},
							},
						},
					},
					{
						Owner: addrs[1],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
								},
							},
						},
					},
				},
				Children: []composable.Children{
					{
						Subject: composable.FullID{
							ClassId: classIDs[0],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[1],
								Id:      sdk.NewUint(3),
							},
						},
					},
					{
						Subject: composable.FullID{
							ClassId: classIDs[1],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[0],
								Id:      sdk.NewUint(3),
							},
						},
					},
				},
			},
			err: composable.ErrInvalidNFTID,
		},
		"unsorted nft id": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
				},
				Balances: []composable.Balance{
					{
						Owner: addrs[0],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[0],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
									sdk.NewUint(1),
								},
							},
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
								},
							},
						},
					},
					{
						Owner: addrs[1],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
								},
							},
						},
					},
				},
				Children: []composable.Children{
					{
						Subject: composable.FullID{
							ClassId: classIDs[0],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[1],
								Id:      sdk.NewUint(3),
							},
						},
					},
					{
						Subject: composable.FullID{
							ClassId: classIDs[1],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[0],
								Id:      sdk.NewUint(3),
							},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		"invalid subject": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
				},
				Balances: []composable.Balance{
					{
						Owner: addrs[0],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[0],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
									sdk.NewUint(2),
								},
							},
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
								},
							},
						},
					},
					{
						Owner: addrs[1],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
								},
							},
						},
					},
				},
				Children: []composable.Children{
					{
						Subject: composable.FullID{
							Id: sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[1],
								Id:      sdk.NewUint(3),
							},
						},
					},
					{
						Subject: composable.FullID{
							ClassId: classIDs[1],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[0],
								Id:      sdk.NewUint(3),
							},
						},
					},
				},
			},
			err: composable.ErrInvalidClassID,
		},
		"empty children": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
				},
				Balances: []composable.Balance{
					{
						Owner: addrs[0],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[0],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
									sdk.NewUint(2),
								},
							},
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
								},
							},
						},
					},
					{
						Owner: addrs[1],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
								},
							},
						},
					},
				},
				Children: []composable.Children{
					{
						Subject: composable.FullID{
							ClassId: classIDs[0],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{},
					},
					{
						Subject: composable.FullID{
							ClassId: classIDs[1],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[0],
								Id:      sdk.NewUint(3),
							},
						},
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		"invalid child": {
			s: composable.GenesisState{
				Params: composable.DefaultParams(),
				Nfts: []composable.ClassNFTs{
					{
						Class: composable.Class{
							Id: classIDs[0],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
					{
						Class: composable.Class{
							Id: classIDs[1],
						},
						PreviousId: sdk.NewUint(3),
						Nfts: []composable.NFT{
							{
								Id: sdk.NewUint(1),
							},
							{
								Id: sdk.NewUint(2),
							},
							{
								Id: sdk.NewUint(3),
							},
						},
					},
				},
				Balances: []composable.Balance{
					{
						Owner: addrs[0],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[0],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
									sdk.NewUint(2),
								},
							},
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(1),
								},
							},
						},
					},
					{
						Owner: addrs[1],
						Balance: []composable.ClassBalance{
							{
								ClassId: classIDs[1],
								Ids: []sdk.Uint{
									sdk.NewUint(2),
								},
							},
						},
					},
				},
				Children: []composable.Children{
					{
						Subject: composable.FullID{
							ClassId: classIDs[0],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								Id: sdk.NewUint(3),
							},
						},
					},
					{
						Subject: composable.FullID{
							ClassId: classIDs[1],
							Id:      sdk.NewUint(2),
						},
						Children: []composable.FullID{
							{
								ClassId: classIDs[0],
								Id:      sdk.NewUint(3),
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
