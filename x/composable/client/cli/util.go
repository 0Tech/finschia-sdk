package cli

import (
	"encoding/json"

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

func ParseParams(codec codec.Codec, paramsJSON string) (*composable.Params, error) {
	var params composable.Params
	if err := codec.UnmarshalJSON([]byte(paramsJSON), &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType.Wrap("params"), err.Error())
	}

	return &params, nil
}

func ParseTraits(codec codec.Codec, traitsJSON string) ([]composable.Trait, error) {
	var traitJSONs []json.RawMessage
	if err := json.Unmarshal([]byte(traitsJSON), &traitJSONs); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType.Wrap("traits"), err.Error())
	}

	if len(traitJSONs) == 0 {
		return nil, nil
	}

	traits := make([]composable.Trait, len(traitJSONs))
	for i, traitJSON := range traitJSONs {
		if err := codec.UnmarshalJSON(traitJSON, &traits[i]); err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType.Wrapf("trait %d", i), err.Error())
		}
	}

	return traits, nil
}

func ParseProperties(codec codec.Codec, propertiesJSON string) ([]composable.Property, error) {
	var propertyJSONs []json.RawMessage
	if err := json.Unmarshal([]byte(propertiesJSON), &propertyJSONs); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType.Wrap("properties"), err.Error())
	}

	if len(propertyJSONs) == 0 {
		return nil, nil
	}

	properties := make([]composable.Property, len(propertyJSONs))
	for i, propertyJSON := range propertyJSONs {
		if err := codec.UnmarshalJSON(propertyJSON, &properties[i]); err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType.Wrapf("property %d", i), err.Error())
		}
	}

	return properties, nil
}
