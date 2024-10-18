package reset

import (
	"fmt"
	"os"

	"github.com/Limerio/todos-cli-app/pkg/utils"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Long:  "Delete the store of your todos",
		Short: "Delete the store of your todos",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := os.OpenFile(utils.STORE_FILE, os.O_RDONLY, 0644)
			if err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), STORE_ALREADY_RESET)

				return
			}

			var confirm bool

			huh.NewConfirm().
				Title("Are you sure?").
				Affirmative("Yes!").
				Negative("No.").
				Value(&confirm).Run()

			if !confirm {
				fmt.Fprintln(cmd.ErrOrStderr(), STORE_RESET_CANCEL)

				return
			}

			if err := os.Remove(utils.STORE_FILE); err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err)

				return
			}

			fmt.Fprintln(cmd.OutOrStdout(), STORE_RESET)
		},
	}
}
