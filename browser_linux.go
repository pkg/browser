package browser

import (
	"os/exec"
)

func openBrowser(url string) error {
	sensibleBrowser, err := exec.LookPath("sensible-browser")
	if err != nil {
		// sensible-browser not availble, try xdg-open
		return exec.Command("xdg-open", url).Run()
	}
	return exec.Command(sensibleBrowser, url).Run()
}
