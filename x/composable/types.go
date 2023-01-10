package composable

import (
	"strings"

	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
)

const didDelimiter = ":"

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
	seenIDs := map[string]struct{}{}
	for _, trait := range t {
		if err := trait.ValidateBasic(); err != nil {
			return sdkerrors.Wrap(err, trait.Id)
		}

		if _, seen := seenIDs[trait.Id]; seen {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrap("duplicate id"), trait.Id)
		}
		seenIDs[trait.Id] = struct{}{}
	}

	return nil
}

func NFTFromString(did string) (*NFT, error) {
	splitted := strings.Split(did, didDelimiter)
	if len(splitted) != 2 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType.Wrap("did"), "must be in [class-id]:[id]")
	}

	classID, idStr := splitted[0], splitted[1]

	id, err := sdk.ParseUint(idStr)
	if err != nil {
		return nil, ErrInvalidNFTID.Wrap(err.Error())
	}

	nft := NFT{
		ClassId: classID,
		Id:      id,
	}
	if err := nft.ValidateBasic(); err != nil {
		return nil, err
	}

	return &nft, nil
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

func (nft NFT) String() string {
	elems := []string{
		nft.ClassId,
		nft.Id.String(),
	}

	return strings.Join(elems, didDelimiter)
}

func (p Property) ValidateBasic() error {
	if len(p.Id) == 0 {
		return ErrInvalidTraitID.Wrap("empty")
	}

	return nil
}

type Properties []Property

func (p Properties) ValidateBasic() error {
	seenIDs := map[string]struct{}{}
	for _, property := range p {
		if err := property.ValidateBasic(); err != nil {
			return sdkerrors.Wrap(err, property.Id)
		}

		if _, seen := seenIDs[property.Id]; seen {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest.Wrap("duplicate id"), property.Id)
		}
		seenIDs[property.Id] = struct{}{}
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
