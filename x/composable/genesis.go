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
	for classIndex, classNfts := range s.Nfts {
		errHint := fmt.Sprintf("nfts[%d]", classIndex)

		class := classNfts.Class
		if err := class.ValidateBasic(); err != nil {
			return sdkerrors.Wrap(err, errHint)
		}

		if _, seen := classIDs[class.Id]; seen {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrapf("duplicate class id %s", class.Id), errHint)
		}
		classIDs[class.Id] = struct{}{}

		seenID := sdk.ZeroUint()
		for nftIndex, nft := range classNfts.Nfts {
			errHint := fmt.Sprintf("%s.nfts[%d]", errHint, nftIndex)

			if err := nft.ValidateBasic(); err != nil {
				return sdkerrors.Wrap(err, errHint)
			}

			if nft.Id.LTE(seenID) {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrap("unsorted nfts"), errHint)
			}
			if nft.Id.GT(classNfts.PreviousId) {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrapf("id %s > previous id %s", nft.Id, classNfts.PreviousId), errHint)
			}
			seenID = nft.Id
		}
	}

	nftOwners := map[string]struct{}{}
	for balanceIndex, balance := range s.Balances {
		errHint := fmt.Sprintf("balances[%d]", balanceIndex)

		if err := ValidateAddress(balance.Owner); err != nil {
			return sdkerrors.Wrap(err, errHint)
		}

		if _, seen := nftOwners[balance.Owner]; seen {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrapf("duplicate owner %s", balance.Owner), errHint)
		}
		nftOwners[balance.Owner] = struct{}{}

		for classBalanceIndex, classBalance := range balance.Balance {
			errHint := fmt.Sprintf("%s.balance[%d]", errHint, classBalanceIndex)

			if _, seen := classIDs[classBalance.ClassId]; !seen {
				return sdkerrors.Wrap(ErrClassNotFound.Wrap(classBalance.ClassId), errHint)
			}

			seenID := sdk.ZeroUint()
			for idIndex, id := range classBalance.Ids {
				errHint := fmt.Sprintf("%s.ids[%d]", errHint, idIndex)

				if err := ValidateNFTID(id); err != nil {
					return sdkerrors.Wrap(err, errHint)
				}

				if id.LTE(seenID) {
					return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrap("unsorted nfts"), errHint)
				}
				seenID = id
			}
		}
	}

	for nftChildrenIndex, nftChildren := range s.Children {
		errHint := fmt.Sprintf("children[%d]", nftChildrenIndex)

		subject := nftChildren.Subject
		if err := subject.ValidateBasic(); err != nil {
			return sdkerrors.Wrap(err, errHint)
		}

		if len(nftChildren.Children) == 0 {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrap("empty children"), errHint)
		}
		for childIndex, child := range nftChildren.Children {
			errHint := fmt.Sprintf("%s.children[%d]", errHint, childIndex)

			if err := child.ValidateBasic(); err != nil {
				return sdkerrors.Wrap(err, errHint)
			}
		}
	}

	return nil
}
