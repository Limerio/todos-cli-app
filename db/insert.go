package db

import (
	"database/sql"
	"time"

	"github.com/Limerio/go-training/todos-cli-app/utils"
	_ "github.com/mattn/go-sqlite3"
)

func Insert(data CreateTodo) (Todo, error) {
	db, err := sql.Open("sqlite3", utils.STORE_FILE)

	if err != nil {
		return Todo{}, err
	}
	defer db.Close()

	id, err := utils.GenerateUuid()

	if err != nil {
		return Todo{}, ErrorDbGenerateId
	}

	dateNow := time.Now()
	defaultDone := utils.NO

	_, err = db.Exec("INSERT INTO todos (id, name, date, done) VALUES (?, ?, ?, ?);", id.String(), data.Name, dateNow, defaultDone)
	if err != nil {
		return Todo{}, err
	}

	return Todo{
		Id:   id.String(),
		Name: data.Name,
		Date: dateNow,
		Done: defaultDone,
	}, nil
}
