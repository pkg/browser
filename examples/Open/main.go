// Open is a simple example of the github.com/pkg/browser package.
//
// Usage:
//
//    # Open a file in a browser window
//    Open $FILE
//
//    # Open a URL in a browser window
//    Open $URL
//
//    # Open the contents of stdin in a browser window
//    cat $SOMEFILE | Open
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cli/browser"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n  %s {<url> | <file> | -}\n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
}

func init() {
	flag.Usage = usage
	flag.Parse()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	args := flag.Args()

	if len(args) != 1 {
		usage()
		os.Exit(1)
	}

	if args[0] == "-" {
		check(browser.OpenReader(os.Stdin))
		return
	}

	if strings.HasPrefix(args[0], "http:") || strings.HasPrefix(args[0], "https:") {
		check(browser.OpenURL(args[0]))
		return
	}

	check(browser.OpenFile(args[0]))
}
