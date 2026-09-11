package browser

// Windows ShellExecute drops fragments on file: URLs. FileProtocolHandler keeps them.
func openBrowser(url string) error {
	return runCmd("rundll32", "url.dll,FileProtocolHandler", url)
}
