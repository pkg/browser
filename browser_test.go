package browser

import (
	"strings"
	"testing"
)

func TestSplitFragment(t *testing.T) {
	file, frag := splitFragment(`c:\some-file.html#basic`)
	if file != `c:\some-file.html` {
		t.Fatalf("file = %q", file)
	}
	if frag != "basic" {
		t.Fatalf("frag = %q", frag)
	}
	f2, g2 := splitFragment("index.html")
	if f2 != "index.html" || g2 != "" {
		t.Fatalf("got %q %q", f2, g2)
	}
}

func TestFileURLPreservesFragment(t *testing.T) {
	got := fileURL("/tmp/doc.html", "basic")
	want := "file:///tmp/doc.html#basic"
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
	got = fileURL("C:/some-file.html", "basic")
	if !strings.HasPrefix(got, "file:///C:/some-file.html") || !strings.HasSuffix(got, "#basic") {
		t.Fatalf("windows-style got %q", got)
	}
}
