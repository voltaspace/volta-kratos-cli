package upgrade

import (
	"fmt"

	"github.com/voltaspace/volta-kratos-cli/kratos/v2/internal/base"

	"github.com/spf13/cobra"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the kratos tools",
	Long:  "Upgrade the kratos tools. Example: kratos upgrade",
	Run:   Run,
}

// Run upgrade the kratos tools.
func Run(cmd *cobra.Command, args []string) {
	err := base.GoInstall(
		"github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@v2.1.3",
		"github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@v2.1.3",
		"google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0",
		"github.com/favadi/protoc-go-inject-tag@v1.4.0",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.42.0",
		"github.com/envoyproxy/protoc-gen-validate@v0.1.0",
	)
	if err != nil {
		fmt.Println(err)
	}
}
