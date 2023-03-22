package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/fanfury-sports/fury/x/sportevent/types"
	"github.com/spf13/cobra"
)

// CmdResolveEvent registers the resolve-event command
func CmdResolveEvent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolve-sport-event [ticket]",
		Short: "set resolution of an event",
		Long:  "Resolve a sport-event with ticket.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgResolveEvent(
				clientCtx.GetFromAddress().String(),
				args[0],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
