package composable

import (
	"github.com/line/lbm-sdk/codec"
	"github.com/line/lbm-sdk/codec/legacy"
	"github.com/line/lbm-sdk/codec/types"
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/types/msgservice"
	authzcodec "github.com/line/lbm-sdk/x/authz/codec"
	govcodec "github.com/line/lbm-sdk/x/gov/codec"
)

// RegisterLegacyAminoCodec registers concrete types on the LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	for msg, name := range map[sdk.Msg]string{
		&MsgSend{}:        "MsgSend",
		&MsgAttach{}:      "MsgAttach",
		&MsgDetach{}:      "MsgDetach",
		&MsgNewClass{}:    "MsgNewClass",
		&MsgUpdateClass{}: "MsgUpdateClass",
		&MsgMintNFT{}:     "MsgMintNFT",
		&MsgBurnNFT{}:     "MsgBurnNFT",
		&MsgUpdateNFT{}:   "MsgUpdateNFT",
	} {
		const prefix = "lbm-sdk/composable/"
		legacy.RegisterAminoMsg(cdc, msg, prefix+name)
	}
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSend{},
		&MsgAttach{},
		&MsgDetach{},
		&MsgNewClass{},
		&MsgUpdateClass{},
		&MsgMintNFT{},
		&MsgBurnNFT{},
		&MsgUpdateNFT{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	// Register all Amino interfaces and concrete types on the authz and gov Amino codec so that this can later be
	// used to properly serialize MsgGrant, MsgExec and MsgSubmitProposal instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
	RegisterLegacyAminoCodec(govcodec.Amino)
}
