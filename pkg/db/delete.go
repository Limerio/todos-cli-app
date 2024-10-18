package db

import (
	"database/sql"

	"github.com/Limerio/todos-cli-app/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
)

func Delete(id string) error {
	db, err := sql.Open("sqlite3", utils.STORE_FILE)

	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
