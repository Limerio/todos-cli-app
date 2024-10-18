package add

import (
	"errors"
	"fmt"

	"github.com/Limerio/todos-cli-app/db"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Long:  "",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Fprintln(cmd.ErrOrStderr(), TODO_GIVE_NAME)

				return
			}

			newTodo := db.CreateTodo{
				Name: args[0],
			}

			_, err := db.Insert(newTodo)

			if err != nil {
				switch {
				case errors.Is(err, db.ErrorDbGenerateId):
					fmt.Fprintln(cmd.ErrOrStderr(), db.ErrorDbGenerateId)
				default:
					fmt.Fprintln(cmd.ErrOrStderr(), SOMETHING_WENT_WRONG)
				}
				return
			}

			fmt.Fprintln(cmd.OutOrStdout(), TODO_CREATED)
		},
	}
}
