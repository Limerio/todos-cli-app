package db

import "errors"

var (
	ErrorDbGenerateId       = errors.New("⚠️ Something went wrong when generating an ID")
	ErrorDbNotInitilialized = errors.New("⚠️ The store has not been initialized. Please do the command `todos create`")
)
