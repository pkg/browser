package browser

import (
	"fmt"
	"syscall/js"
)

func openBrowser(url string) error {
	window := js.Global().Get("window").Call("open", url, "_blank")
	if !window.Truthy() {
		return fmt.Errorf("Unable to open tab, check browser logs")
	}
	return nil
}
