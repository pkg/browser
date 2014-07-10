package browser

import (
	"os/exec"
)

func openBrowser(url string) error {
	cmd := exec.Command("open", url)
	return cmd.Run()
}
