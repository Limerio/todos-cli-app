package remove_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/Limerio/todos-cli-app/commands/remove"
	"github.com/Limerio/todos-cli-app/db"
	"github.com/Limerio/todos-cli-app/utils"
)

func TestRemoveCmd_ExecuteWithSuccess(t *testing.T) {
	cmd := remove.Cmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	if err := db.Init(); err != nil {
		t.Errorf("Unexpected error when creating the store: %v", err)
	}

	todo, err := db.Insert(db.CreateTodo{Name: "Test todo"})

	if err != nil {
		t.Errorf("Unexpected error when adding a test todo in the store: %v", err)
	}

	cmd.SetArgs([]string{todo.Id})

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error when executing the command: %v", err)
	}

	if strings.TrimRight(stdout.String(), "\n") != remove.TODO_DELETED {
		t.Errorf("Expected output: %q, but got: %q", remove.TODO_DELETED, stdout.String())
	}

	if err := os.Remove(utils.STORE_FILE); err != nil {
		t.Errorf("Unexpected error when deleting the store: %v", err)

		return
	}
}

func TestRemoveCmd_ExecuteWithError_NoArguments(t *testing.T) {
	cmd := remove.Cmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error when executing the command: %v", err)
	}

	if strings.TrimRight(stdout.String(), "\n") != remove.TODO_NEEDS_ID {
		t.Errorf("Expected output: %q, but got: %q", remove.TODO_NEEDS_ID, stdout.String())
	}
}

func TestRemoveCmd_ExecuteWithError_CannotFindId(t *testing.T) {
	cmd := remove.Cmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	if err := db.Init(); err != nil {
		t.Errorf("Unexpected error when creating the store: %v", err)
	}

	id, err := utils.GenerateUuid()
	if err != nil {
		t.Errorf("Unexpected error when generating an id: %v", err)
	}

	cmd.SetArgs([]string{id.String()})

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error when executing the command: %v", err)
	}

	if strings.TrimRight(stdout.String(), "\n") != remove.TODO_CANNOT_FIND_ID {
		t.Errorf("Expected output: %q, but got: %q", remove.TODO_CANNOT_FIND_ID, stdout.String())
	}

	if err := os.Remove(utils.STORE_FILE); err != nil {
		t.Errorf("Unexpected error when deleting the store: %v", err)

		return
	}
}
