package cli

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/line/lbm-sdk/client/flags"
	"github.com/line/lbm-sdk/codec"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

func validateGenerateOnly(cmd *cobra.Command) error {
	generateOnly, err := cmd.Flags().GetBool(flags.FlagGenerateOnly)
	if err != nil {
		return err
	}

	if !generateOnly {
		return sdkerrors.ErrNotSupported.Wrapf("must use it with the flag --%s", flags.FlagGenerateOnly)
	}

	return nil
}

func parseParams(codec codec.Codec, paramsJSON string) (*composable.Params, error) {
	var params composable.Params
	if err := codec.UnmarshalJSON([]byte(paramsJSON), &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType.Wrapf("params"), err.Error())
	}

	return &params, nil
}

func ParseNFT(nftString string) (*composable.NFT, error) {
	const delimiter = ":"
	splitted := strings.Split(nftString, delimiter)
	if len(splitted) != 2 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType.Wrap("nft"), "must be [class-id]:[id]")
	}

	classID, idStr := splitted[0], splitted[1]

	id, err := composable.NFTIDFromString(idStr)
	if err != nil {
		return nil, err
	}

	nft := composable.NFT{
		ClassId: classID,
		Id:      *id,
	}
	if err := nft.ValidateBasic(); err != nil {
		return nil, err
	}

	return &nft, nil
}
