
# browser

Helpers to open URLs, readers, or files in the system default web browser.

This fork adds:

- [safety](https://github.com/cli/safeexec#readme) when running inside of an untrusted directory on Windows;
- WSL compatibility;
- `OpenReader` error wrapping;
- `ErrNotFound` error wrapping on BSD;
- Go 1.13 support.

## Usage

``` go
import "github.com/cli/browser"

err = browser.OpenURL(url)
err = browser.OpenFile(path)
err = browser.OpenReader(reader)
```
