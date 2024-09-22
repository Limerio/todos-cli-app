package remove

import (
	"fmt"

	"github.com/Limerio/go-training/todos-cli-app/db"
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

			if err := db.Delete(todo.Id); err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), SOMETHING_WENT_WRONG_WHEN_DELETE, todo.Id)

				return
			}

			fmt.Fprintln(cmd.OutOrStdout(), TODO_DELETED)
		},
	}

}
