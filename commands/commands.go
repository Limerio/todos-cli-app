package commands

import (
	"fmt"
	"os"

	"github.com/Limerio/go-training/todos-cli-app/commands/add"
	"github.com/Limerio/go-training/todos-cli-app/commands/create"
	"github.com/Limerio/go-training/todos-cli-app/commands/done"
	"github.com/Limerio/go-training/todos-cli-app/commands/export"
	"github.com/Limerio/go-training/todos-cli-app/commands/list"
	"github.com/Limerio/go-training/todos-cli-app/commands/remove"
	"github.com/Limerio/go-training/todos-cli-app/commands/reset"
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
