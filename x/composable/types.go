package composable

import (
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

	if err := ValidateURIHash(class.Uri, class.UriHash); err != nil {
		return err
	}

	return nil
}

func (nft NFT) ValidateBasic() error {
	if err := ValidateNFTID(nft.Id); err != nil {
		return err
	}

	if err := ValidateURIHash(nft.Uri, nft.UriHash); err != nil {
		return err
	}

	return nil
}

func (id FullID) ValidateBasic() error {
	if err := ValidateClassID(id.ClassId); err != nil {
		return err
	}

	if err := ValidateNFTID(id.Id); err != nil {
		return err
	}

	return nil
}

func (l FullID) Equal(r FullID) bool {
	if l.ClassId != r.ClassId {
		return false
	}

	return l.Id.Equal(r.Id)
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

func ValidateURIHash(uri, hash string) error {
	if len(uri) != 0 && len(hash) == 0 {
		return ErrInvalidUriHash.Wrap("empty hash for non-empty uri")
	}

	if len(uri) == 0 && len(hash) != 0 {
		return ErrInvalidUriHash.Wrap("non-empty hash for empty uri")
	}

	return nil
}

func ClassOwner(id string) sdk.AccAddress {
	return sdk.MustAccAddressFromBech32(id)
}

func ClassIDFromOwner(owner sdk.AccAddress) string {
	return owner.String()
}
