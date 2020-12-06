package browser

import (
	"errors"
	"os/exec"
)

func openBrowser(url string) error {
	browserPath, err := lookPath("xdg-open", "wslview")
	if err != nil {
		return err
	}
	return runCmd(browserPath, url)
}

func lookPath(cmdName, fallbackName string) (string, error) {
	cmdPath, err := exec.LookPath(cmdName)
	if errors.Is(err, exec.ErrNotFound) {
		if wslPath, err := exec.LookPath(fallbackName); err == nil {
			return wslPath, nil
		}
	}
	return cmdPath, err
}

func setFlags(cmd *exec.Cmd) {}
