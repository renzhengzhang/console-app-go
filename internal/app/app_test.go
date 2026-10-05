package app

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunWithArgs(t *testing.T) {
	var out bytes.Buffer
	err := Run(strings.NewReader(""), &out, []string{"Go"})
	if err != nil {
		t.Fatal(err)
	}

	if got, want := out.String(), "Hello, Go!\n"; got != want {
		t.Fatalf("got = %q, want = %q", got, want)
	}
}
