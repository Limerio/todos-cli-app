package create

import (
	"fmt"
	"strings"

	"github.com/Limerio/todos-cli-app/pkg/db"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	return &cobra.Command{
		Use:     "create",
		Aliases: []string{"init"},
		Short:   "Initialize the store",
		Long:    "Initialize the store",
		Run: func(cmd *cobra.Command, args []string) {
			err := db.Init()

			if err != nil {
				if strings.Contains(err.Error(), "already exists") {
					fmt.Fprintln(cmd.ErrOrStderr(), STORE_ALREADY_INITALIZE)
					return
				} else {
					fmt.Fprintln(cmd.ErrOrStderr(), SOMETHING_WENT_WRONG_CREATING_STORE)
					return
				}
			}

			fmt.Fprintln(cmd.OutOrStdout(), STORE_CREATED)
		},
	}
}
