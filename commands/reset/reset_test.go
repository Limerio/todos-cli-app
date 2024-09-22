package reset_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Limerio/go-training/todos-cli-app/commands/reset"
	"github.com/Limerio/go-training/todos-cli-app/db"
)

func TestResetCmd_ExecuteWithSuccess(t *testing.T) {
	cmd := reset.Cmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	if err := db.Init(); err != nil {
		t.Errorf("Unexpected error when creating the store: %v", err)
	}

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error when executing the command: %v", err)
	}

	if strings.TrimRight(stdout.String(), "\n") != reset.STORE_RESET {
		t.Errorf("Expected output: %q, but got: %q", reset.STORE_RESET, stdout.String())
	}
}

func TestResetCmd_ExecuteWithError_StoreAlreadyReset(t *testing.T) {
	cmd := reset.Cmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error when executing the command: %v", err)
	}

	if strings.TrimRight(stdout.String(), "\n") != reset.STORE_ALREADY_RESET {
		t.Errorf("Expected output: %q, but got: %q", reset.STORE_ALREADY_RESET, stdout.String())
	}
}
