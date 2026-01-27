// Package browser provides helpers to open files, readers, and urls in a browser window.
//
// The choice of which browser is started is entirely client dependant.
package browser

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// Stdout is the io.Writer to which executed commands write standard output.
var Stdout io.Writer = os.Stdout

// Stderr is the io.Writer to which executed commands write standard error.
var Stderr io.Writer = os.Stderr

// Opener allows customizing browser opening behavior.
type Opener struct {
	// Stdout is the io.Writer to which executed commands write standard output.
	// If nil, os.Stdout is used.
	Stdout io.Writer

	// Stderr is the io.Writer to which executed commands write standard error.
	// If nil, os.Stderr is used.
	Stderr io.Writer
}

func (o *Opener) stdout() io.Writer {
	if o.Stdout != nil {
		return o.Stdout
	}
	return Stdout
}

func (o *Opener) stderr() io.Writer {
	if o.Stderr != nil {
		return o.Stderr
	}
	return Stderr
}

func (o *Opener) runCmd(prog string, args ...string) error {
	cmd := exec.Command(prog, args...)
	cmd.Stdout = o.stdout()
	cmd.Stderr = o.stderr()
	return cmd.Run()
}

// OpenFile opens new browser window for the file path.
func (o *Opener) OpenFile(path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	return o.OpenURL("file://" + path)
}

// OpenReader consumes the contents of r and presents the
// results in a new browser window.
func (o *Opener) OpenReader(r io.Reader) error {
	f, err := ioutil.TempFile("", "browser.*.html")
	if err != nil {
		return fmt.Errorf("browser: could not create temporary file: %v", err)
	}
	if _, err := io.Copy(f, r); err != nil {
		f.Close()
		return fmt.Errorf("browser: caching temporary file failed: %v", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("browser: caching temporary file failed: %v", err)
	}
	return o.OpenFile(f.Name())
}

// OpenURL opens a new browser window pointing to url.
func (o *Opener) OpenURL(url string) error {
	return o.openBrowser(url)
}

// defaultOpener returns an Opener configured with the package-level Stdout/Stderr.
// This is done as a function to always grab the latest values of Stdout/Stderr.
func defaultOpener() *Opener {
	return &Opener{
		Stdout: Stdout,
		Stderr: Stderr,
	}
}

// OpenFile opens new browser window for the file path.
func OpenFile(path string) error {
	return defaultOpener().OpenFile(path)
}

// OpenReader consumes the contents of r and presents the
// results in a new browser window.
func OpenReader(r io.Reader) error {
	return defaultOpener().OpenReader(r)
}

// OpenURL opens a new browser window pointing to url.
func OpenURL(url string) error {
	return defaultOpener().OpenURL(url)
}
