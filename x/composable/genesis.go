package composable

import (
	"fmt"

	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
)

const (
	DefaultMaxDescendants = 15
)

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

func DefaultParams() Params {
	return Params{
		MaxDescendants: DefaultMaxDescendants,
	}
}

// ValidateBasic check the given genesis state has no integrity issues
func (s GenesisState) ValidateBasic() error {
	classIDs := map[string]struct{}{}
	for classIndex, genClass := range s.Classes {
		errHint := fmt.Sprintf("classes[%d]", classIndex)

		id := genClass.Id
		if err := ValidateClassID(id); err != nil {
			return sdkerrors.Wrap(err, errHint)
		}

		if _, seen := classIDs[id]; seen {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrapf("duplicate class id %s", genClass.Id), errHint)
		}
		classIDs[id] = struct{}{}

		if err := Traits(genClass.Traits).ValidateBasic(); err != nil {
			return sdkerrors.Wrap(err, errHint)
		}

		traits := map[string]struct{}{}
		for _, trait := range genClass.Traits {
			traits[trait.Id] = struct{}{}
		}

		seenID := sdk.ZeroUint()
		for nftIndex, genNFT := range genClass.Nfts {
			errHint := fmt.Sprintf("%s.nfts[%d]", errHint, nftIndex)

			id := genNFT.Id
			if err := ValidateNFTID(id); err != nil {
				return sdkerrors.Wrap(err, errHint)
			}

			if id.LTE(seenID) {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrap("unsorted nfts"), errHint)
			}
			if id.GT(genClass.LastMintedNftId) {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrapf("id %s > last minted id %s", id, genClass.LastMintedNftId), errHint)
			}
			seenID = id

			if err := Properties(genNFT.Properties).ValidateBasic(); err != nil {
				return sdkerrors.Wrap(err, errHint)
			}

			for _, property := range genNFT.Properties {
				if _, hasTrait := traits[property.Id]; !hasTrait {
					return sdkerrors.Wrap(ErrTraitNotFound.Wrap(property.Id), errHint)
				}
			}

			// xor must be true
			hasOwner := (len(genNFT.Owner) != 0)
			hasParent := (genNFT.Parent != nil)
			if hasOwner == hasParent {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrap("owner and parent mutually exclusive"), errHint)
			}

			if hasOwner {
				if err := ValidateAddress(genNFT.Owner); err != nil {
					return sdkerrors.Wrap(sdkerrors.Wrap(err, "owner"), errHint)
				}
			}

			if hasParent {
				if err := genNFT.Parent.ValidateBasic(); err != nil {
					return sdkerrors.Wrap(sdkerrors.Wrap(err, "parent"), errHint)
				}
			}
		}
	}

	return nil
}
