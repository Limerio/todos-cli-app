package db

import (
	"database/sql"

	"github.com/Limerio/go-training/todos-cli-app/utils"
	_ "github.com/mattn/go-sqlite3"
)

func FindById(id string) (Todo, error) {
	db, err := sql.Open("sqlite3", utils.STORE_FILE)

	if err != nil {
		return Todo{}, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM todos WHERE id = ?", id)

	var foundTodo Todo

	err = row.Scan(&foundTodo.Id, &foundTodo.Name, &foundTodo.Done, &foundTodo.Date)
	if err != nil {
		return Todo{}, err
	}

	return foundTodo, nil
}
