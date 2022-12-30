package keeper

import (
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

func (k Keeper) InitGenesis(ctx sdk.Context, gs *composable.GenesisState) error {
	k.SetParams(ctx, gs.Params)

	for _, classState := range gs.Nfts {
		class := classState.Class
		k.setClass(ctx, class)

		k.setPreviousID(ctx, class.Id, classState.PreviousId)

		for _, nftState := range classState.NftStates {
			nft := nftState.Nft
			k.setNFT(ctx, class.Id, nft)

			id := composable.FullID{
				ClassId: class.Id,
				Id:      nft.Id,
			}

			if owner := nftState.Owner; len(owner) != 0 {
				k.setOwner(ctx, id, sdk.MustAccAddressFromBech32(owner))
			}

			if parent := nftState.Parent; parent != nil {
				k.setParent(ctx, id, *parent)

				numDescendants := k.getNumDescendants(ctx, id)
				diff := 1 + numDescendants
				k.iterateAncestors(ctx, *parent, func(id composable.FullID) {
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

	var classStates []composable.ClassNFTs
	if len(classes) != 0 {
		classStates = make([]composable.ClassNFTs, len(classes))
	}

	for classIndex := range classStates {
		class := classes[classIndex]

		classStates[classIndex].Class = class
		classStates[classIndex].PreviousId = k.GetPreviousID(ctx, class.Id)

		nfts := k.getNFTsOfClass(ctx, class.Id)

		var nftStates []composable.NFTState
		if len(nfts) != 0 {
			nftStates = make([]composable.NFTState, len(nfts))
		}

		for nftIndex := range nftStates {
			nft := nfts[nftIndex]

			nftStates[nftIndex].Nft = nft

			id := composable.FullID{
				ClassId: class.Id,
				Id:      nft.Id,
			}

			if owner, err := k.getOwner(ctx, id); err == nil {
				nftStates[nftIndex].Owner = owner.String()
				continue
			}

			if parent, err := k.getParent(ctx, id); err == nil {
				nftStates[nftIndex].Parent = parent
				continue
			}

			panic(sdkerrors.Wrap(sdkerrors.ErrNotFound.Wrap("owner or parent"), id.String()))
		}

		classStates[classIndex].NftStates = nftStates
	}

	return &composable.GenesisState{
		Params: k.GetParams(ctx),
		Nfts:   classStates,
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
