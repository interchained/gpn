package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tendermint/spn/x/launch/types"
)

func CmdListRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-request [chain-id]",
		Short: "list all request",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRequestRequest{
				ChainID:    args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.RequestAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-request [chain-id] [request-id]",
		Short: "shows a request",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsRequestID, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryGetRequestRequest{
				ChainID:   args[0],
				RequestID: argsRequestID,
			}

			res, err := queryClient.Request(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
