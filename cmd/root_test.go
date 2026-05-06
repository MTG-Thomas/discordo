package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/ayn2op/tview"
	"github.com/gdamore/tcell/v3"
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

func TestApplyTViewTheme(t *testing.T) {
	orig := tview.Styles
	t.Cleanup(func() { tview.Styles = orig })

	applyTViewTheme()

	tests := []struct {
		name string
		got  tcell.Color
		want tcell.Color
	}{
		{
			name: "primitive background",
			got:  tview.Styles.PrimitiveBackgroundColor,
			want: tcell.GetColor("#0f1117"),
		},
		{
			name: "contrast background",
			got:  tview.Styles.ContrastBackgroundColor,
			want: tcell.GetColor("#151820"),
		},
		{
			name: "primary text",
			got:  tview.Styles.PrimaryTextColor,
			want: tcell.GetColor("#d6d3ca"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.got != test.want {
				t.Fatalf("got = %v, want = %v", test.got, test.want)
			}
		})
	}
}
