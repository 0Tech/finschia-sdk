package composable

import (
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable/codec"
)

var _ sdk.Msg = (*MsgSend)(nil)

// ValidateBasic implements Msg.
func (m MsgSend) ValidateBasic() error {
	if err := ValidateAddress(m.Sender); err != nil {
		return sdkerrors.Wrap(err, "sender")
	}

	if err := ValidateAddress(m.Recipient); err != nil {
		return sdkerrors.Wrap(err, "recipient")
	}

	id := FullID{
		ClassId: m.ClassId,
		Id:      m.Id,
	}
	if err := id.ValidateBasic(); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgSend) GetSigners() []sdk.AccAddress {
	signer := sdk.MustAccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgSend) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgSend) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgSend) GetSignBytes() []byte {
	return sdk.MustSortJSON(codec.ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgAttach)(nil)

// ValidateBasic implements Msg.
func (m MsgAttach) ValidateBasic() error {
	if err := ValidateAddress(m.Owner); err != nil {
		return sdkerrors.Wrap(err, "owner")
	}

	subject := FullID{
		ClassId: m.SubjectClassId,
		Id:      m.SubjectId,
	}
	if err := subject.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "subject")
	}

	target := FullID{
		ClassId: m.TargetClassId,
		Id:      m.TargetId,
	}
	if err := target.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "target")
	}

	if target.Equal(subject) {
		return ErrInvalidComposition.Wrap("subject == target")
	}

	return nil
}

// GetSigners implements Msg
func (m MsgAttach) GetSigners() []sdk.AccAddress {
	signer := sdk.MustAccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgAttach) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgAttach) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgAttach) GetSignBytes() []byte {
	return sdk.MustSortJSON(codec.ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgDetach)(nil)

// ValidateBasic implements Msg.
func (m MsgDetach) ValidateBasic() error {
	if err := ValidateAddress(m.Owner); err != nil {
		return sdkerrors.Wrap(err, "owner")
	}

	id := FullID{
		ClassId: m.ClassId,
		Id:      m.Id,
	}
	if err := id.ValidateBasic(); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgDetach) GetSigners() []sdk.AccAddress {
	signer := sdk.MustAccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgDetach) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgDetach) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgDetach) GetSignBytes() []byte {
	return sdk.MustSortJSON(codec.ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgNewClass)(nil)

// ValidateBasic implements Msg.
func (m MsgNewClass) ValidateBasic() error {
	if err := ValidateAddress(m.Owner); err != nil {
		return sdkerrors.Wrap(err, "owner")
	}

	if err := ValidateURIHash(m.Uri, m.UriHash); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgNewClass) GetSigners() []sdk.AccAddress {
	signer := sdk.MustAccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgNewClass) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgNewClass) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgNewClass) GetSignBytes() []byte {
	return sdk.MustSortJSON(codec.ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgUpdateClass)(nil)

// ValidateBasic implements Msg.
func (m MsgUpdateClass) ValidateBasic() error {
	class := Class{
		Id:      m.ClassId,
		Uri:     m.Uri,
		UriHash: m.UriHash,
	}
	if err := class.ValidateBasic(); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgUpdateClass) GetSigners() []sdk.AccAddress {
	signer := ClassOwner(m.ClassId)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgUpdateClass) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgUpdateClass) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgUpdateClass) GetSignBytes() []byte {
	return sdk.MustSortJSON(codec.ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgMintNFT)(nil)

// ValidateBasic implements Msg.
func (m MsgMintNFT) ValidateBasic() error {
	if err := ValidateClassID(m.ClassId); err != nil {
		return err
	}

	if err := ValidateURIHash(m.Uri, m.UriHash); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgMintNFT) GetSigners() []sdk.AccAddress {
	signer := ClassOwner(m.ClassId)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgMintNFT) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgMintNFT) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgMintNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(codec.ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgBurnNFT)(nil)

// ValidateBasic implements Msg.
func (m MsgBurnNFT) ValidateBasic() error {
	if err := ValidateAddress(m.Owner); err != nil {
		return sdkerrors.Wrap(err, "owner")
	}

	id := FullID{
		ClassId: m.ClassId,
		Id:      m.Id,
	}
	if err := id.ValidateBasic(); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgBurnNFT) GetSigners() []sdk.AccAddress {
	signer := sdk.MustAccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgBurnNFT) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgBurnNFT) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgBurnNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(codec.ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.Msg = (*MsgUpdateNFT)(nil)

// ValidateBasic implements Msg.
func (m MsgUpdateNFT) ValidateBasic() error {
	if err := ValidateClassID(m.ClassId); err != nil {
		return err
	}

	nft := NFT{
		Id:      m.Id,
		Uri:     m.Uri,
		UriHash: m.UriHash,
	}
	if err := nft.ValidateBasic(); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgUpdateNFT) GetSigners() []sdk.AccAddress {
	signer := ClassOwner(m.ClassId)
	return []sdk.AccAddress{signer}
}

// Type implements the LegacyMsg.Type method.
func (m MsgUpdateNFT) Type() string {
	return sdk.MsgTypeURL(&m)
}

// Route implements the LegacyMsg.Route method.
func (m MsgUpdateNFT) Route() string {
	return RouterKey
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method.
func (m MsgUpdateNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(codec.ModuleCdc.MustMarshalJSON(&m))
}
