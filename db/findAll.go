package db

import (
	"database/sql"
	"strings"

	"github.com/Limerio/todos-cli-app/utils"
	_ "github.com/mattn/go-sqlite3"
)

func FindAll() ([]Todo, error) {
	db, err := sql.Open("sqlite3", utils.STORE_FILE)

	if err != nil {
		return []Todo{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		if strings.Contains(err.Error(), "no such table: todos") {
			return []Todo{}, ErrorDbNotInitilialized
		} else {
			return []Todo{}, err
		}
	}
	defer rows.Close()

	var foundTodos []Todo

	for rows.Next() {
		todo := Todo{}

		err = rows.Scan(&todo.Id, &todo.Name, &todo.Done, &todo.Date)
		if err != nil {
			panic(err)
		}

		foundTodos = append(foundTodos, todo)
	}

	return foundTodos, nil
}
