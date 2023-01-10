package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/line/lbm-sdk/client"
	"github.com/line/lbm-sdk/client/flags"
	"github.com/line/lbm-sdk/x/composable"
)

// NewQueryCmd returns the query commands for the module
func NewQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   composable.ModuleName,
		Short: fmt.Sprintf("%s query subcommands", composable.ModuleName),
	}

	cmd.AddCommand(
		NewQueryCmdParams(),
		NewQueryCmdClass(),
		NewQueryCmdClasses(),
		NewQueryCmdNFT(),
		NewQueryCmdNFTs(),
		NewQueryCmdOwner(),
		NewQueryCmdParent(),
	)

	return cmd
}

func NewQueryCmdParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the module parameters",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			req := composable.QueryParamsRequest{}

			queryClient := composable.NewQueryClient(clientCtx)
			res, err := queryClient.Params(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryCmdClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class [class-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query a class",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			classID := args[0]
			if err := composable.ValidateClassID(classID); err != nil {
				return err
			}

			req := composable.QueryClassRequest{
				ClassId: classID,
			}

			queryClient := composable.NewQueryClient(clientCtx)
			res, err := queryClient.Class(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryCmdClasses() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "classes",
		Args:  cobra.NoArgs,
		Short: "Query the classes",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			req := composable.QueryClassesRequest{}

			queryClient := composable.NewQueryClient(clientCtx)
			res, err := queryClient.Classes(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryCmdNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nft [id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query an nft",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			nft, err := composable.NFTFromString(args[0])
			if err != nil {
				return err
			}

			req := composable.QueryNFTRequest{
				ClassId: nft.ClassId,
				Id:      nft.Id.String(),
			}

			queryClient := composable.NewQueryClient(clientCtx)
			res, err := queryClient.NFT(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryCmdNFTs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nfts [class-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query the nfts",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			classID := args[0]
			if err := composable.ValidateClassID(classID); err != nil {
				return err
			}

			req := composable.QueryNFTsRequest{
				ClassId: classID,
			}

			queryClient := composable.NewQueryClient(clientCtx)
			res, err := queryClient.NFTs(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryCmdOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "owner [id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query the owner of an nft",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			nft, err := composable.NFTFromString(args[0])
			if err != nil {
				return err
			}

			req := composable.QueryOwnerRequest{
				ClassId: nft.ClassId,
				Id:      nft.Id.String(),
			}

			queryClient := composable.NewQueryClient(clientCtx)
			res, err := queryClient.Owner(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewQueryCmdParent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parent [id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query the parent id of an nft",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			nft, err := composable.NFTFromString(args[0])
			if err != nil {
				return err
			}

			req := composable.QueryParentRequest{
				ClassId: nft.ClassId,
				Id:      nft.Id.String(),
			}

			queryClient := composable.NewQueryClient(clientCtx)
			res, err := queryClient.Parent(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
