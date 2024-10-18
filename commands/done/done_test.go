package done_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Limerio/todos-cli-app/commands/done"
	"github.com/Limerio/todos-cli-app/db"
	"github.com/Limerio/todos-cli-app/utils"
)

func TestDoneCmd_ExecuteWithSuccess(t *testing.T) {
	cmd := done.Cmd()

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

	if strings.TrimRight(stdout.String(), "\n") != fmt.Sprintf(done.TODO_DONE, todo.Id) {
		t.Errorf("Expected output: %q, but got: %q", done.TODO_DONE, stdout.String())
	}

	if err := os.Remove(utils.STORE_FILE); err != nil {
		t.Errorf("Unexpected error when deleting the store: %v", err)

		return
	}
}

func TestDoneCmd_ExecuteWithError_NoArguments(t *testing.T) {
	cmd := done.Cmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error when executing the command: %v", err)
	}

	if strings.TrimRight(stdout.String(), "\n") != done.TODO_NEEDS_ID {
		t.Errorf("Expected output: %q, but got: %q", done.TODO_NEEDS_ID, stdout.String())
	}
}
