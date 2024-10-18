package commands

import (
	"fmt"
	"os"

	"github.com/Limerio/todos-cli-app/pkg/commands/add"
	"github.com/Limerio/todos-cli-app/pkg/commands/create"
	"github.com/Limerio/todos-cli-app/pkg/commands/done"
	"github.com/Limerio/todos-cli-app/pkg/commands/export"
	"github.com/Limerio/todos-cli-app/pkg/commands/list"
	"github.com/Limerio/todos-cli-app/pkg/commands/remove"
	"github.com/Limerio/todos-cli-app/pkg/commands/reset"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todos",
	Short: "Management of your todos",
}

func Execute() {

	rootCmd.AddCommand(add.Cmd(), export.Cmd(), list.Cmd(), create.Cmd(), remove.Cmd(), done.Cmd(), reset.Cmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
