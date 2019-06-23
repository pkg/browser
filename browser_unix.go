// +build !windows

package browser

import (
	"fmt"
	"os"
)

func shell() (shell, flag string) {
	shell = os.Getenv("SHELL")

	if shell == "" {
		shell = "/bin/sh"
	}

	return shell, "-c"
}

func fmtBrowserCmd(browser, url string) string {
	return fmt.Sprintf("%s '%s'", browser, url)
}
