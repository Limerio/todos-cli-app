package export

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/Limerio/todos-cli-app/db"
	"github.com/Limerio/todos-cli-app/utils"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Long:  "Export your todos store in csv or json",
		Short: "Export your todos store in csv or json",
		Run: func(cmd *cobra.Command, args []string) {
			formatExport, err := cmd.Flags().GetString("format")

			if err != nil {
				panic(err)
			}

			if len(formatExport) < 1 {
				huh.NewForm(
					huh.NewGroup(
						huh.NewSelect[string]().Options(huh.NewOptions("json", "csv")...).Title("ℹ️ Choose a format").Value(&formatExport),
					),
				).Run()
			}

			if formatExport != "csv" && formatExport != "json" {
				fmt.Fprintln(cmd.ErrOrStderr(), EXPORT_FORMAT_DOESNT_EXIST)

				return
			}

			todos, err := db.FindAll()
			if err != nil {
				switch {
				case errors.Is(err, db.ErrorDbNotInitilialized):
					fmt.Fprintln(cmd.ErrOrStderr(), db.ErrorDbNotInitilialized)
				default:
					fmt.Fprintln(cmd.ErrOrStderr(), EXPORT_SOMETHING_WENT_WRONG_RETRIVE_DATA)
				}
				return
			}

			s := ""

			if formatExport == "json" {

				newString := "["

				for i, todo := range todos {
					newString += "{"
					newString += fmt.Sprintf("\"id\": \"%s\", ", todo.Id)
					newString += fmt.Sprintf("\"name\": \"%s\", ", todo.Name)
					newString += fmt.Sprintf("\"date\": \"%s\", ", todo.Date)
					newString += fmt.Sprintf("\"done\": \"%s\"", todo.Done)

					if i == len(todos)-1 {
						newString += "}"
					} else {
						newString += "},"
					}

				}
				s = newString + "]"
				fmt.Println(s)

				return
			}

			w := csv.NewWriter(os.Stdout)
			defer w.Flush()

			for _, todo := range todos {
				if err := w.Write([]string{todo.Id, todo.Name, todo.Date.Format(utils.DATE_LAYOUT), todo.Done.String()}); err != nil {
					fmt.Fprintln(cmd.ErrOrStderr(), EXPORT_SOMETHING_WENT_WRONG_WHEN_GENERATE_CSV)
				}
			}

		},
	}

	cmd.PersistentFlags().StringP("format", "f", "", "You have the ability to export your todos with various format such as json, or csv")

	return cmd
}
