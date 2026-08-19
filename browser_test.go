package browser

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	var stdout, stderr bytes.Buffer
	Stdout = &stdout
	Stderr = &stderr
	defer func() {
		Stdout = os.Stdout
		Stderr = os.Stderr
	}()

	err := runCmd("echo", "hello", "world")
	if err != nil {
		t.Fatalf("runCmd failed: %v", err)
	}
	if !strings.Contains(stdout.String(), "hello world") {
		t.Errorf("stdout = %q, expected to contain %q", stdout.String(), "hello world")
	}
}

func TestRunCmdError(t *testing.T) {
	err := runCmd("nonexistent-command-xyz-12345")
	if err == nil {
		t.Error("expected error for nonexistent command, got nil")
	}
}
