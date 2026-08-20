package browser

import (
	"testing"
)

func TestStdoutStderr(t *testing.T) {
	if Stdout == nil {
		t.Error("Stdout should not be nil")
	}
	if Stderr == nil {
		t.Error("Stderr should not be nil")
	}
}
