package browser

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

func openBrowser(url string) error {
	r := strings.NewReplacer("&", "^&")
	return runCmd("cmd", "/c", "start", r.Replace(url))
}

func setFlags(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}

func shell() (shell, flag string) {
	return "cmd", "/c"
}

func fmtBrowserCmd(browser, url string) string {
	return fmt.Sprintf("%s %s", browser, url)
}
