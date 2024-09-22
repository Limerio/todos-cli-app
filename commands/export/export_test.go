package export_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/Limerio/go-training/todos-cli-app/commands/export"
	"github.com/Limerio/go-training/todos-cli-app/db"
	"github.com/Limerio/go-training/todos-cli-app/utils"
)

func TestExportCmd_ExecuteWithSuccess_Csv(t *testing.T) {
	cmd := export.Cmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	cmd.SetArgs([]string{"-f csv"})

	if err := db.Init(); err != nil {
		t.Errorf("Unexpected error when creating the store: %v", err)
	}

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error when executing the command: %v", err)
	}

	if strings.TrimRight(stdout.String(), "\n") != "" {
		t.Errorf("Expected output: %q, but got: %q", "", stdout.String())
	}

	if err := os.Remove(utils.STORE_FILE); err != nil {
		t.Errorf("Unexpected error when deleting the store: %v", err)

		return
	}
}

func TestExportCmd_ExecuteWithSuccess_Json(t *testing.T) {
	cmd := export.Cmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	cmd.SetArgs([]string{"-f json"})

	if err := db.Init(); err != nil {
		t.Errorf("Unexpected error when creating the store: %v", err)
	}

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error when executing the command: %v", err)
	}

	if strings.TrimRight(stdout.String(), "\n") != "" {
		t.Errorf("Expected output: %q, but got: %q", "", stdout.String())
	}

	if err := os.Remove(utils.STORE_FILE); err != nil {
		t.Errorf("Unexpected error when deleting the store: %v", err)

		return
	}
}

func TestExportCmd_ExecuteWithError_FormatDoesntExist(t *testing.T) {
	cmd := export.Cmd()

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	cmd.SetArgs([]string{"-f kdioesjdse"})

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error when executing the command: %v", err)
	}

	if strings.TrimRight(stdout.String(), "\n") != export.EXPORT_FORMAT_DOESNT_EXIST {
		t.Errorf("Expected output: %q, but got: %q", export.EXPORT_FORMAT_DOESNT_EXIST, stdout.String())
	}
}
