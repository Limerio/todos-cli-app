package db

import (
	"time"

	"github.com/Limerio/todos-cli-app/utils"
)

type Todo struct {
	Id   string
	Name string
	Date time.Time
	Done utils.IsDone
}

type UpdateTodo struct {
	Name string
	Done utils.IsDone
}

type CreateTodo struct {
	Name string
}
