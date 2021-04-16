package browser

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func ExampleOpenFile() {
	OpenFile("index.html")
}

func ExampleOpenReader() {
	// https://github.com/rust-lang/rust/issues/13871
	const quote = `There was a night when winds from unknown spaces
whirled us irresistibly into limitless vacum beyond all thought and entity.
Perceptions of the most maddeningly untransmissible sort thronged upon us;
perceptions of infinity which at the time convulsed us with joy, yet which
are now partly lost to my memory and partly incapable of presentation to others.`
	r := strings.NewReader(quote)
	OpenReader(r)
}

func ExampleOpenURL() {
	const url = "http://golang.org/"
	OpenURL(url)
}

func ExampleOpenURLWithOptions() {
	out := &bytes.Buffer{}
	const url = "http://golang.org/"
	OpenURL(url, func(cmd *exec.Cmd) {
		cmd.Stdout = out
		cmd.Stderr = out
	})
	fmt.Printf("browser open output: %q\n", out.String())
}
