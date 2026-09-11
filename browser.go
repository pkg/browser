// Package browser provides helpers to open files, readers, and urls in a browser window.
//
// The choice of which browser is started is entirely client dependant.
package browser

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Stdout is the io.Writer to which executed commands write standard output.
var Stdout io.Writer = os.Stdout

// Stderr is the io.Writer to which executed commands write standard error.
var Stderr io.Writer = os.Stderr

// OpenFile opens new browser window for the file path.
// A URL fragment after '#' (for example "index.html#section") is preserved.
func OpenFile(path string) error {
	path, frag := splitFragment(path)
	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	return OpenURL(fileURL(abs, frag))
}

// splitFragment separates a filesystem path from an optional URL fragment.
func splitFragment(path string) (file, fragment string) {
	i := strings.LastIndex(path, "#")
	if i < 0 {
		return path, ""
	}
	return path[:i], path[i+1:]
}

// fileURL builds a file: URL for an absolute filesystem path, with optional fragment.
func fileURL(absPath, fragment string) string {
	p := filepath.ToSlash(absPath)
	if !strings.HasPrefix(p, "/") {
		// Windows drive path, e.g. C:/foo → /C:/foo
		p = "/" + p
	}
	u := url.URL{Scheme: "file", Path: p, Fragment: fragment}
	return u.String()
}

// OpenReader consumes the contents of r and presents the
// results in a new browser window.
func OpenReader(r io.Reader) error {
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
	return OpenFile(f.Name())
}

// OpenURL opens a new browser window pointing to url.
func OpenURL(url string) error {
	return openBrowser(url)
}

func runCmd(prog string, args ...string) error {
	cmd := exec.Command(prog, args...)
	cmd.Stdout = Stdout
	cmd.Stderr = Stderr
	return cmd.Run()
}
