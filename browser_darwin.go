package browser

func (o *Opener) openBrowser(url string) error {
	return o.runCmd("open", url)
}
