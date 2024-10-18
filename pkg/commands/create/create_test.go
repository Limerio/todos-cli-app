package create_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/Limerio/todos-cli-app/pkg/commands/create"
	"github.com/Limerio/todos-cli-app/pkg/db"
	"github.com/Limerio/todos-cli-app/pkg/utils"
)

func TestCreateCmd_ExecuteWithSuccess(t *testing.T) {
	cmd := create.Cmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error when executing the command: %v", err)
	}

	if strings.TrimRight(stdout.String(), "\n") != create.STORE_CREATED {
		t.Errorf("Expected output: %q, but got: %q", create.STORE_CREATED, stdout.String())
	}

	if err := os.Remove(utils.STORE_FILE); err != nil {
		t.Errorf("Unexpected error when deleting the store: %v", err)

		return
	}
}

func TestCreateCmd_ExecuteWithError_StoreAlreadyExist(t *testing.T) {
	cmd := create.Cmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	if err := db.Init(); err != nil {
		t.Errorf("Unexpected error when creating the store: %v", err)
	}

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error when executing the command: %v", err)
	}

	if strings.TrimRight(stdout.String(), "\n") != create.STORE_ALREADY_INITALIZE {
		t.Errorf("Expected output: %q, but got: %q", create.STORE_ALREADY_INITALIZE, stdout.String())
	}

	if err := os.Remove(utils.STORE_FILE); err != nil {
		t.Errorf("Unexpected error when deleting the store: %v", err)

		return
	}
}
