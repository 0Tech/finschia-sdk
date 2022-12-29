package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/line/lbm-sdk/client"
	"github.com/line/lbm-sdk/client/flags"
	"github.com/line/lbm-sdk/client/tx"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

const (
	FlagUri     = "uri"
	FlagUriHash = "uri-hash"
)

// NewTxCmd returns the transaction commands for the module
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        composable.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", composable.ModuleName),
		Long:                       "manipulate composable nfts",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		// NewTxCmdUpdateParams(),
		NewTxCmdSend(),
		NewTxCmdAttach(),
		NewTxCmdDetach(),
		NewTxCmdNewClass(),
		NewTxCmdUpdateClass(),
		NewTxCmdMintNFT(),
		NewTxCmdBurnNFT(),
		NewTxCmdUpdateNFT(),
	)

	return txCmd
}

// func NewTxCmdUpdateParams() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "update-params [authority] [params-json]",
// 		Args:  cobra.ExactArgs(2),
// 		Short: "Update the module parameters",
// 		Example: `
// Example of the content of params-json:

// {
//   "max_descendants": 15
// }
// `,
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			if err := validateGenerateOnly(cmd); err != nil {
// 				return err
// 			}

// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			params, err := parseParams(clientCtx.Codec, args[1])
// 			if err != nil {
// 				return err
// 			}

// 			msg := composable.MsgUpdateParams{
// 				Authority: args[0],
// 				Params:    *params,
// 			}
// 			if err := msg.ValidateBasic(); err != nil {
// 				return err
// 			}

// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
// 		},
// 	}

// 	flags.AddTxFlagsToCmd(cmd)

// 	return cmd
// }

func NewTxCmdSend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [sender] [recipient] [id]",
		Args:  cobra.ExactArgs(3),
		Short: "Send an nft from one account to another account",
		Example: `
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			sender := args[0]
			if err := cmd.Flags().Set(flags.FlagFrom, sender); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := ParseFullID(args[2])
			if err != nil {
				return err
			}

			msg := composable.MsgSend{
				Sender:    sender,
				Recipient: args[1],
				ClassId:   id.ClassId,
				Id:        id.Id,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdAttach() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "attach [owner] [subject-id] [target-id]",
		Args:    cobra.ExactArgs(3),
		Short:   "Attach a root nft to another nft",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			owner := args[0]
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subjectID, err := ParseFullID(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "subject")
			}

			targetID, err := ParseFullID(args[2])
			if err != nil {
				return sdkerrors.Wrap(err, "target")
			}

			msg := composable.MsgAttach{
				Owner:          owner,
				SubjectClassId: subjectID.ClassId,
				SubjectId:      subjectID.Id,
				TargetClassId:  targetID.ClassId,
				TargetId:       targetID.Id,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdDetach() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "detach [owner] [id]",
		Args:    cobra.ExactArgs(2),
		Short:   "Detach an nft from another nft",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			owner := args[0]
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := ParseFullID(args[1])
			if err != nil {
				return err
			}

			msg := composable.MsgDetach{
				Owner:   owner,
				ClassId: id.ClassId,
				Id:      id.Id,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdNewClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "new-class [owner] [--uri] [--uri-hash]",
		Args:    cobra.ExactArgs(1),
		Short:   "create a class",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			owner := args[0]
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			uri, err := cmd.Flags().GetString(FlagUri)
			if err != nil {
				return err
			}

			uriHash, err := cmd.Flags().GetString(FlagUriHash)
			if err != nil {
				return err
			}

			msg := composable.MsgNewClass{
				Owner:   owner,
				Uri:     uri,
				UriHash: uriHash,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(FlagUri, "", "set uri")
	cmd.Flags().String(FlagUriHash, "", "set uri-hash")

	return cmd
}

func NewTxCmdUpdateClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-class [class-id] [--uri] [--uri-hash]",
		Args:    cobra.ExactArgs(1),
		Short:   "update a class",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			classID := args[0]
			if err := composable.ValidateClassID(classID); err != nil {
				return err
			}

			owner := composable.ClassOwner(classID).String()
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			uri, err := cmd.Flags().GetString(FlagUri)
			if err != nil {
				return err
			}

			uriHash, err := cmd.Flags().GetString(FlagUriHash)
			if err != nil {
				return err
			}

			msg := composable.MsgUpdateClass{
				ClassId: classID,
				Uri:     uri,
				UriHash: uriHash,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(FlagUri, "", "set uri")
	cmd.Flags().String(FlagUriHash, "", "set uri-hash")

	return cmd
}

func NewTxCmdMintNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "mint-nft [class-id] [recipient] [--uri] [--uri-hash]",
		Args:    cobra.ExactArgs(2),
		Short:   "mint an nft",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			classID := args[0]
			if err := composable.ValidateClassID(classID); err != nil {
				return err
			}

			owner := composable.ClassOwner(classID).String()
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			uri, err := cmd.Flags().GetString(FlagUri)
			if err != nil {
				return err
			}

			uriHash, err := cmd.Flags().GetString(FlagUriHash)
			if err != nil {
				return err
			}

			msg := composable.MsgMintNFT{
				ClassId:   classID,
				Uri:       uri,
				UriHash:   uriHash,
				Recipient: args[1],
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(FlagUri, "", "set uri")
	cmd.Flags().String(FlagUriHash, "", "set uri-hash")

	return cmd
}

func NewTxCmdBurnNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-nft [owner] [id]",
		Args:  cobra.ExactArgs(2),
		Short: "burn an nft",
		Example: `
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			owner := args[0]
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := ParseFullID(args[1])
			if err != nil {
				return err
			}

			msg := composable.MsgBurnNFT{
				Owner:   owner,
				ClassId: id.ClassId,
				Id:      id.Id,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewTxCmdUpdateNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-nft [id] [--uri] [--uri-hash]",
		Args:    cobra.ExactArgs(1),
		Short:   "update an nft",
		Example: ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := ParseFullID(args[0])
			if err != nil {
				return err
			}

			classID := id.ClassId
			if err := composable.ValidateClassID(classID); err != nil {
				return err
			}

			owner := composable.ClassOwner(classID).String()
			if err := cmd.Flags().Set(flags.FlagFrom, owner); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			uri, err := cmd.Flags().GetString(FlagUri)
			if err != nil {
				return err
			}

			uriHash, err := cmd.Flags().GetString(FlagUriHash)
			if err != nil {
				return err
			}

			msg := composable.MsgUpdateNFT{
				ClassId: id.ClassId,
				Id:      id.Id,
				Uri:     uri,
				UriHash: uriHash,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(FlagUri, "", "set uri")
	cmd.Flags().String(FlagUriHash, "", "set uri-hash")

	return cmd
}
