package composable

import (
// "fmt"

// sdk "github.com/line/lbm-sdk/types"
// sdkerrors "github.com/line/lbm-sdk/types/errors"
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
	// classIDs := map[string]struct{}{}
	// for classIndex, classNfts := range s.Nfts {
	// 	errHint := fmt.Sprintf("nfts[%d]", classIndex)

	// 	class := classNfts.Class
	// 	if err := class.ValidateBasic(); err != nil {
	// 		return sdkerrors.Wrap(sdkerrors.Wrap(err, "class"), errHint)
	// 	}

	// 	if _, seen := classIDs[class.Id]; seen {
	// 		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrapf("duplicate class id %s", class.Id), errHint)
	// 	}
	// 	classIDs[class.Id] = struct{}{}

	// 	seenID := sdk.ZeroUint()
	// 	for nftIndex, nftState := range classNfts.NftStates {
	// 		errHint := fmt.Sprintf("%s.nft_states[%d]", errHint, nftIndex)

	// 		nft := nftState.Nft
	// 		if err := nft.ValidateBasic(); err != nil {
	// 			return sdkerrors.Wrap(sdkerrors.Wrap(err, "nft"), errHint)
	// 		}

	// 		if nft.Id.LTE(seenID) {
	// 			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrap("unsorted nfts"), errHint)
	// 		}
	// 		if nft.Id.GT(classNfts.PreviousId) {
	// 			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrapf("id %s > previous id %s", nft.Id, classNfts.PreviousId), errHint)
	// 		}
	// 		seenID = nft.Id

	// 		// xor must be true
	// 		hasOwner := (len(nftState.Owner) != 0)
	// 		hasParent := (nftState.Parent != nil)
	// 		if hasOwner == hasParent {
	// 			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrap("owner and parent mutually exclusive"), errHint)
	// 		}

	// 		if hasOwner {
	// 			if err := ValidateAddress(nftState.Owner); err != nil {
	// 				return sdkerrors.Wrap(sdkerrors.Wrap(err, "owner"), errHint)
	// 			}
	// 		}

	// 		if hasParent {
	// 			if err := nftState.Parent.ValidateBasic(); err != nil {
	// 				return sdkerrors.Wrap(sdkerrors.Wrap(err, "parent"), errHint)
	// 			}
	// 		}
	// 	}
	// }

	return nil
}
