package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunWithArgsVersion(t *testing.T) {
	var out bytes.Buffer

	if err := RunWithArgs([]string{"-version"}, &out); err != nil {
		t.Fatal(err)
	}

	got := out.String()
	for _, want := range []string{"discordo", "version=", "commit=", "date="} {
		if !strings.Contains(got, want) {
			t.Fatalf("version output %q does not contain %q", got, want)
		}
	}
}
