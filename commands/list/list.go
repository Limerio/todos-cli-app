package list

import (
	"errors"
	"fmt"
	"os"

	"github.com/Limerio/go-training/todos-cli-app/db"
	"github.com/Limerio/go-training/todos-cli-app/utils"
	"github.com/mergestat/timediff"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Give the list of your todos",
		Long:    "Give the list of your todos",
		Run: func(cmd *cobra.Command, args []string) {

			allView, err := cmd.Flags().GetBool("all")

			if err != nil {
				panic(err)
			}

			todos, err := db.FindAll()

			if err != nil {
				switch {
				case errors.Is(err, db.ErrorDbNotInitilialized):
					fmt.Fprintln(cmd.ErrOrStderr(), db.ErrorDbNotInitilialized)
				default:
					fmt.Fprintln(cmd.ErrOrStderr(), err)
				}
			}

			table := tablewriter.NewWriter(os.Stdout)
			if allView {
				table.SetHeader([]string{"Id", "Name", "Date", "Done"})

				for _, todo := range todos {
					table.Append([]string{todo.Id, todo.Name, timediff.TimeDiff(todo.Date), todo.Done.String()})
				}
			} else {
				table.SetHeader([]string{"Id", "Name", "Date"})

				for _, todo := range todos {
					if todo.Done == utils.NO {
						table.Append([]string{todo.Id, todo.Name, timediff.TimeDiff(todo.Date)})
					}
				}
			}
			table.Render()
		},
	}

	cmd.PersistentFlags().BoolP("all", "a", false, "View all of todos")

	return cmd
}
