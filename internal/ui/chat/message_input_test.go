package chat

import (
	"testing"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

func TestMessageInputHeightForText(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		width  int
		chrome int
		want   int
	}{
		{
			name:   "empty input keeps compact height",
			width:  20,
			chrome: 2,
			want:   3,
		},
		{
			name:   "explicit newlines grow input",
			text:   "one\ntwo\nthree",
			width:  20,
			chrome: 2,
			want:   5,
		},
		{
			name:   "long line wraps",
			text:   "123456789012345678901",
			width:  10,
			chrome: 2,
			want:   5,
		},
		{
			name:   "height is capped",
			text:   "1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n12",
			width:  20,
			chrome: 2,
			want:   12,
		},
		{
			name:   "unknown width stays compact",
			text:   "1\n2\n3",
			width:  0,
			chrome: 2,
			want:   3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := messageInputHeightForText(test.text, test.width, test.chrome); got != test.want {
				t.Fatalf("messageInputHeightForText(%q, %d, %d) = %d, want %d", test.text, test.width, test.chrome, got, test.want)
			}
		})
	}
}

func TestMessageInputPlaceholderLineUsesInputStyle(t *testing.T) {
	style := tcell.StyleDefault.Foreground(color.Red).Background(color.Blue)
	line := messageInputPlaceholderLine("Message...", style)
	if len(line) != 1 {
		t.Fatalf("placeholder line has %d segments, want 1", len(line))
	}

	if got := line[0].Style.GetForeground(); got != color.Red {
		t.Fatalf("placeholder foreground = %v, want %v", got, color.Red)
	}
	if got := line[0].Style.GetBackground(); got != color.Blue {
		t.Fatalf("placeholder background = %v, want %v", got, color.Blue)
	}
	if !line[0].Style.HasDim() {
		t.Fatal("placeholder style is not dimmed")
	}
}
