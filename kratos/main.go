package main

import (
	"log"

	"github.com/voltaspace/volta-kratos-cli/kratos/v1/internal/change"
	"github.com/voltaspace/volta-kratos-cli/kratos/v1/internal/project"
	"github.com/voltaspace/volta-kratos-cli/kratos/v1/internal/proto"
	"github.com/voltaspace/volta-kratos-cli/kratos/v1/internal/run"
	"github.com/voltaspace/volta-kratos-cli/kratos/v1/internal/upgrade"

	"github.com/spf13/cobra"
)
const release = "v2.1.2"
var rootCmd = &cobra.Command{
	Use:     "kratos",
	Short:   "Kratos: An elegant toolkit for Go microservices.",
	Long:    `Kratos: An elegant toolkit for Go microservices.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(change.CmdChange)
	rootCmd.AddCommand(run.CmdRun)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
