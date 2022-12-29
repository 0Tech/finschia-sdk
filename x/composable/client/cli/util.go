package cli

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/line/lbm-sdk/client/flags"
	"github.com/line/lbm-sdk/codec"
	sdk "github.com/line/lbm-sdk/types"
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

func ParseFullID(fullIDString string) (*composable.FullID, error) {
	const delimiter = ":"
	splitted := strings.Split(fullIDString, delimiter)
	if len(splitted) != 2 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType.Wrap("id"), "must be [class-id]:[id]")
	}

	classID, idString := splitted[0], splitted[1]

	// TODO revisit
	id, err := sdk.ParseUint(idString)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType.Wrap("id"), err.Error())
	}

	fullID := composable.FullID{
		ClassId: classID,
		Id:      id,
	}

	return &fullID, nil
}
