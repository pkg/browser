package browser

import (
	"os/exec"
	"strings"
)

func hasProgram(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func openBrowser(url string) error {
	// Windows Subsystem for Linux (bash for Windows) doesn't have xdg-open available
	// but you can execute cmd.exe from there; try to identify it
	if !hasProgram("xdg-open") && hasProgram("cmd.exe") {
		r := strings.NewReplacer("&", "^&")
		return runCmd("cmd.exe", "/c", "start", r.Replace(url))
	}

	return runCmd("xdg-open", url)
}
