package add_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/Limerio/todos-cli-app/pkg/commands/add"
	"github.com/Limerio/todos-cli-app/pkg/db"
	"github.com/Limerio/todos-cli-app/pkg/utils"
)

func TestAddCmd_Execute(t *testing.T) {
	cmd := add.Cmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	cmd.SetArgs([]string{"Random todo"})

	if err := db.Init(); err != nil {
		t.Errorf("Unexpected error when creating the store: %v", err)
	}

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error when executing the command: %v", err)
	}

	if strings.TrimRight(stdout.String(), "\n") != add.TODO_CREATED {
		t.Errorf("Expected output: %q, but got: %q", add.TODO_CREATED, stdout.String())
	}

	if err := os.Remove(utils.STORE_FILE); err != nil {
		t.Errorf("Unexpected error when deleting the store: %v", err)

		return
	}
}

func TestAddCmd_ExecuteWithError_NoArguments(t *testing.T) {
	cmd := add.Cmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error when executing the command: %v", err)
	}

	if strings.TrimRight(stdout.String(), "\n") != add.TODO_GIVE_NAME {
		t.Errorf("Expected output: %q, but got: %q", add.TODO_GIVE_NAME, stdout.String())
	}
}

func TestAddCmd_ExecuteWithError_SomethingWentWrong(t *testing.T) {
	cmd := add.Cmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)
	cmd.SetArgs([]string{"dioesjoijed"})

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error when executing the command: %v", err)
	}

	if strings.TrimRight(stdout.String(), "\n") != add.SOMETHING_WENT_WRONG {
		t.Errorf("Expected output: %q, but got: %q", add.SOMETHING_WENT_WRONG, stdout.String())
	}

	if err := os.Remove(utils.STORE_FILE); err != nil {
		t.Errorf("Unexpected error when deleting the store: %v", err)

		return
	}
}
