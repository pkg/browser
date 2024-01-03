package browser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func openBrowser(url string) error {
	providers := []string{"xdg-open", "x-www-browser", "www-browser"}

	wsl, err := isWSL()
	if err != nil {
		return errors.New("error determining Windows Subsystem for Linux")
	}

	if wsl {
		cmd := "cmd.exe"
		args := []string{"/c", "start"}
		url := strings.ReplaceAll(url, "&", "^&")
		args = append(args, url)
		return exec.Command(cmd, args...).Start()
	}

	// There are multiple possible providers to open a browser on linux
	// One of them is xdg-open, another is x-www-browser, then there's www-browser, etc.
	// Look for one that exists and run it
	for _, provider := range providers {
		if _, err := exec.LookPath(provider); err == nil {
			return runCmd(provider, url)
		}
	}

	return &exec.Error{Name: strings.Join(providers, ","), Err: exec.ErrNotFound}
}

func isWSL() (bool, error) {
	data, err := ioutil.ReadFile("/proc/version")
	if err != nil {
		return false, fmt.Errorf("read /proc/version: %v", err)
	}
	return strings.Contains(strings.ToLower(string(data)), "microsoft"), nil
}
