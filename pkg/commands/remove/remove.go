package remove

import (
	"fmt"

	"github.com/Limerio/todos-cli-app/pkg/db"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	return &cobra.Command{
		Use:     "remove",
		Aliases: []string{"del", "d", "delete", "rm"},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Fprintln(cmd.ErrOrStderr(), TODO_NEEDS_ID)

				return
			}

			todo, err := db.FindById(args[0])

			if err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), TODO_CANNOT_FIND_ID)

				return
			}

			var confirm bool

			huh.NewConfirm().
				Title("Are you sure?").
				Affirmative("Yes!").
				Negative("No.").
				Value(&confirm).Run()

			if !confirm {
				fmt.Fprintln(cmd.ErrOrStderr(), TODO_DELETE_CANCEL)

				return
			}

			if err := db.Delete(todo.Id); err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), SOMETHING_WENT_WRONG_WHEN_DELETE, todo.Id)

				return
			}

			fmt.Fprintln(cmd.OutOrStdout(), TODO_DELETED)
		},
	}

}
