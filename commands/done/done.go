package done

import (
	"fmt"

	"github.com/Limerio/go-training/todos-cli-app/db"
	"github.com/Limerio/go-training/todos-cli-app/utils"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	return &cobra.Command{
		Use: "done",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Fprintln(cmd.ErrOrStderr(), TODO_NEEDS_ID)

				return
			}

			todo, err := db.FindById(args[0])

			if err != nil {
				panic(err)
			}

			todo.Done = utils.YES

			newTodo, err := db.Update(todo.Id, db.UpdateTodo{Name: todo.Name, Done: todo.Done})

			if err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), SOMETHING_WENT_WRONG_WHEN_UPDATE_TODO)
				return
			}

			fmt.Fprintf(cmd.OutOrStdout(), TODO_DONE, newTodo.Id)
		},
	}
}
