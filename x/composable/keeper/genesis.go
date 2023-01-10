package keeper

import (
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

func (k Keeper) InitGenesis(ctx sdk.Context, gs *composable.GenesisState) error {
	k.SetParams(ctx, gs.Params)

	for _, genClass := range gs.Classes {
		class := composable.Class{
			Id: genClass.Id,
		}
		k.setClass(ctx, class)

		// TODO: trait

		k.setPreviousID(ctx, class.Id, genClass.LastMintedNftId)

		for _, genNFT := range genClass.Nfts {
			nft := composable.NFT{
				ClassId: class.Id,
				Id:      genNFT.Id,
			}
			k.setNFT(ctx, nft)

			// TODO: property

			if owner := genNFT.Owner; len(owner) != 0 {
				k.setOwner(ctx, nft, sdk.MustAccAddressFromBech32(owner))
			}

			if parent := genNFT.Parent; parent != nil {
				k.setParent(ctx, nft, *parent)

				numDescendants := k.getNumDescendants(ctx, nft)
				diff := 1 + numDescendants
				k.iterateAncestors(ctx, *parent, func(id composable.NFT) {
					old := k.getNumDescendants(ctx, id)
					new := old + diff
					k.updateNumDescendants(ctx, id, new)
				})
			}
		}
	}

	return nil
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *composable.GenesisState {
	classes := k.getClasses(ctx)

	var genClasses []composable.GenesisClass
	if len(classes) != 0 {
		genClasses = make([]composable.GenesisClass, len(classes))
	}

	for classIndex, class := range classes {
		genClasses[classIndex].Id = class.Id
		genClasses[classIndex].LastMintedNftId = k.GetPreviousID(ctx, class.Id)

		// TODO: trait

		nfts := k.getNFTsOfClass(ctx, class.Id)

		var genNFTs []composable.GenesisNFT
		if len(nfts) != 0 {
			genNFTs = make([]composable.GenesisNFT, len(nfts))
		}

		for nftIndex, nft := range nfts {
			genNFTs[nftIndex].Id = nft.Id

			// TODO: property

			if owner, err := k.getOwner(ctx, nft); err == nil {
				genNFTs[nftIndex].Owner = owner.String()
				continue
			}

			if parent, err := k.getParent(ctx, nft); err == nil {
				genNFTs[nftIndex].Parent = parent
				continue
			}

			panic(sdkerrors.Wrap(sdkerrors.ErrNotFound.Wrap("owner or parent"), nft.String()))
		}

		genClasses[classIndex].Nfts = genNFTs
	}

	return &composable.GenesisState{
		Params:  k.GetParams(ctx),
		Classes: genClasses,
	}
}

func (k Keeper) getClasses(ctx sdk.Context) (classes []composable.Class) {
	k.iterateClasses(ctx, func(class composable.Class) {
		classes = append(classes, class)
	})

	return
}

func (k Keeper) getNFTsOfClass(ctx sdk.Context, classID string) (nfts []composable.NFT) {
	k.iterateNFTsOfClass(ctx, classID, func(nft composable.NFT) {
		nfts = append(nfts, nft)
	})

	return
}
