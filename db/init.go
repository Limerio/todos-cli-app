package db

import (
	"database/sql"

	"github.com/Limerio/todos-cli-app/utils"
)

func Init() error {
	db, err := sql.Open("sqlite3", utils.STORE_FILE)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(`
        CREATE TABLE todos (
						id TEXT PRIMARY KEY,
            name TEXT NOT NULL,
            done TEXT NOT NULL,
						date DATETIME
        );
    `)

	if err != nil {
		return err
	}

	return nil
}
