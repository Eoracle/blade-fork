package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/0xPolygon/polygon-edge/command/backup"
	"github.com/0xPolygon/polygon-edge/command/bridge"
	"github.com/0xPolygon/polygon-edge/command/genesis"
	"github.com/0xPolygon/polygon-edge/command/helper"
	"github.com/0xPolygon/polygon-edge/command/mint"
	"github.com/0xPolygon/polygon-edge/command/monitor"
	"github.com/0xPolygon/polygon-edge/command/peers"
	"github.com/0xPolygon/polygon-edge/command/regenesis"
	"github.com/0xPolygon/polygon-edge/command/secrets"
	polybftsecrets "github.com/0xPolygon/polygon-edge/command/secrets/init"
	"github.com/0xPolygon/polygon-edge/command/server"
	"github.com/0xPolygon/polygon-edge/command/status"
	"github.com/0xPolygon/polygon-edge/command/txpool"
	"github.com/0xPolygon/polygon-edge/command/validator"
	"github.com/0xPolygon/polygon-edge/command/version"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Polygon Edge is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		polybftsecrets.GetCommand(),
		bridge.GetCommand(),
		regenesis.GetCommand(),
		mint.GetCommand(),
		validator.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
