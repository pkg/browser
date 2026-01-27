package browser

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestErrorRedirect(t *testing.T) {
	stderr := new(bytes.Buffer)
	o := &Opener{
		Stderr: stderr,
	}

	_ = o.OpenFile(filepath.Join(t.TempDir(), "nonexistentfile.html"))
	if stderr.Len() == 0 {
		t.Errorf("expected stderr to contain error message, got empty")
	}
}
