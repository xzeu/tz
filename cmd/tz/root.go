package tz

import (
	"fmt"

	"github.com/xzeu/tz/config"
	"github.com/xzeu/tz/internal/command/wire"

	"github.com/spf13/cobra"
	"github.com/xzeu/tz/internal/command/create"
	"github.com/xzeu/tz/internal/command/new"
	"github.com/xzeu/tz/internal/command/run"
	"github.com/xzeu/tz/internal/command/upgrade"
)

var CmdRoot = &cobra.Command{
	Use:     "tz",
	Example: "tz new demo-api",
	Short:   " ██████████ ████████
	░░░░░██░░░ ░░░░░░██
	░██         ██
	░██        ██
	░██       ██
	░██      ██
	░██     ████████
	░░     ░░░░░░░░ ",
	Version: fmt.Sprintf("██████████ ████████\n\t░░░░░██░░░ ░░░░░░██\n\t░██         ██\n\t░██        ██\n\t░██       ██\n\t░██      ██\n\t░██     ████████\n\t░░     ░░░░░░░░ Tz%s - Copyright (c) 2023-2025 Tz\nReleased under the MIT License.\n\n", config.Version),
}

func init() {
	CmdRoot.AddCommand(new.CmdNew)
	CmdRoot.AddCommand(create.CmdCreate)
	CmdRoot.AddCommand(run.CmdRun)

	CmdRoot.AddCommand(upgrade.CmdUpgrade)
	create.CmdCreate.AddCommand(create.CmdCreateHandler)
	create.CmdCreate.AddCommand(create.CmdCreateService)
	create.CmdCreate.AddCommand(create.CmdCreateRepository)
	create.CmdCreate.AddCommand(create.CmdCreateModel)
	create.CmdCreate.AddCommand(create.CmdCreateAll)

	CmdRoot.AddCommand(wire.CmdWire)
	wire.CmdWire.AddCommand(wire.CmdWireAll)
}

// Execute executes the root command.
func Execute() error {
	return CmdRoot.Execute()
}
