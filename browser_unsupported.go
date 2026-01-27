//go:build !linux && !windows && !darwin && !openbsd && !freebsd && !netbsd
// +build !linux,!windows,!darwin,!openbsd,!freebsd,!netbsd

package browser

import (
	"fmt"
	"runtime"
)

func (o *Opener) openBrowser(url string) error {
	return fmt.Errorf("openBrowser: unsupported operating system: %v", runtime.GOOS)
}
