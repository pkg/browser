package browser

import (
	"os/exec"
)

func openBrowser(url string) error {
	cmd := exec.Command("cmd", "/c", "start", url)
	return cmd.Run()
}
