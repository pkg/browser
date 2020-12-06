package browser

import (
	"os/exec"
	"strings"
	"syscall"

	"github.com/cli/safeexec"
)

func openBrowser(url string) error {
	cmdPath, err := safeexec.LookPath("cmd")
	if err != nil {
		return err
	}
	r := strings.NewReplacer("&", "^&")
	return runCmd(cmdPath, "/c", "start", r.Replace(url))
}

func setFlags(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
