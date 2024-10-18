package db

import (
	"database/sql"

	"github.com/Limerio/todos-cli-app/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
)

func Update(id string, data UpdateTodo) (Todo, error) {
	db, err := sql.Open("sqlite3", utils.STORE_FILE)

	if err != nil {
		return Todo{}, err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE todos SET name = ?, done = ? WHERE id = ?", data.Name, data.Done, id)
	if err != nil {
		return Todo{}, err
	}

	todo, err := FindById(id)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}
