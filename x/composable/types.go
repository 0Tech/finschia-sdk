package composable

import (
	"fmt"

	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
)

func NFTIDFromString(str string) (*sdk.Uint, error) {
	id, err := sdk.ParseUint(str)
	if err != nil {
		return nil, ErrInvalidNFTID.Wrap(err.Error())
	}

	if err := ValidateNFTID(id); err != nil {
		return nil, err
	}

	return &id, nil
}

func ValidateAddress(address string) error {
	if _, err := sdk.AccAddressFromBech32(address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(address)
	}

	return nil
}

func (class Class) ValidateBasic() error {
	if err := ValidateClassID(class.Id); err != nil {
		return err
	}

	return nil
}

func (t Trait) ValidateBasic() error {
	if len(t.Id) == 0 {
		return ErrInvalidTraitID.Wrap("empty")
	}

	return nil
}

type Traits []Trait

func (t Traits) ValidateBasic() error {
	seenNames := map[string]struct{}{}
	for i, trait := range t {
		errHint := fmt.Sprintf("trait %d", i)

		if err := trait.ValidateBasic(); err != nil {
			return sdkerrors.Wrap(err, errHint)
		}

		if _, seen := seenNames[trait.Id]; seen {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrap("duplicate id"), errHint)
		}
		seenNames[trait.Id] = struct{}{}
	}

	return nil
}

func (nft NFT) ValidateBasic() error {
	if err := ValidateClassID(nft.ClassId); err != nil {
		return err
	}

	if err := ValidateNFTID(nft.Id); err != nil {
		return err
	}

	return nil
}

func (l NFT) Equal(r NFT) bool {
	if l.ClassId != r.ClassId {
		return false
	}

	return l.Id.Equal(r.Id)
}

func (p Property) ValidateBasic() error {
	if len(p.Id) == 0 {
		return ErrInvalidTraitID.Wrap("empty")
	}

	return nil
}

type Properties []Property

func (p Properties) ValidateBasic() error {
	seenNames := map[string]struct{}{}
	for i, property := range p {
		errHint := fmt.Sprintf("property %d", i)

		if err := property.ValidateBasic(); err != nil {
			return sdkerrors.Wrap(err, errHint)
		}

		if _, seen := seenNames[property.Id]; seen {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrap("duplicate id"), errHint)
		}
		seenNames[property.Id] = struct{}{}
	}

	return nil
}

func ValidateClassID(id string) error {
	if _, err := sdk.AccAddressFromBech32(id); err != nil {
		return ErrInvalidClassID.Wrap(id)
	}

	return nil
}

func ValidateNFTID(id sdk.Uint) error {
	if id.IsZero() {
		return ErrInvalidNFTID.Wrap("zero nft id")
	}

	return nil
}

func ClassOwner(id string) sdk.AccAddress {
	return sdk.MustAccAddressFromBech32(id)
}

func ClassIDFromOwner(owner sdk.AccAddress) string {
	return owner.String()
}
