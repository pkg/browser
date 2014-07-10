package browser

import (
	"os"
	"os/exec"
)

func openBrowser(url string) error {
	// try sensible-browser first
	if err := runCmd("sensible-browser", url); err == nil {
		return nil
	}
	// sensible-browser not availble, try xdg-open
	return runCmd("xdg-open", url)
}

func runCmd(prog string, args ...string) error {
	cmd := exec.Command(prog, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
