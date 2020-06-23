package browser

import (
	"errors"
	"os/exec"
	"time"
)

func appearsSuccessful(cmd *exec.Cmd, timeout time.Duration) bool {
	errc := make(chan error, 1)
	go func() {
		errc <- cmd.Wait()
	}()

	select {
	case <-time.After(timeout):
		return true
	case err := <-errc:
		return err == nil
	}
}

func openBrowser(url string) error {
	cmd := exec.Command("xdg-open", url)
	err := cmd.Start()
	if err != nil {
		return err
	}
	if !appearsSuccessful(cmd, 3*time.Second) {
		return errors.New("Unable to successfully run xdg-open")
	}
	return nil
}

func setFlags(cmd *exec.Cmd) {}
