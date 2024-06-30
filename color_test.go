package ansi

import (
	"testing"
)

func TestAnsi(t *testing.T) {

	t.Run("ANSI Color Code Validation", func(t *testing.T) {
		mockStdout := Buffer()
		defer Stdout()

		tests := []struct {
			name     string
			function func()
			expected string
		}{
			{"Reset", func() { Reset() }, CodeReset},
			{"Bold", func() { Bold() }, CodeBold},
			{"Dim", func() { Dim() }, CodeDim},
			{"Underline", func() { Underline() }, CodeUnderline},
			{"Blink", func() { Blink() }, CodeBlink},
			{"Reverse", func() { Reverse() }, CodeReverse},
			{"Hidden", func() { Hidden() }, CodeHiddenText},
			{"Strikethrough", func() { Strikethrough() }, CodeStrikeThrough},
			{"BgBlack", func() { BgBlack() }, CodeBgBlack},
			{"BgRed", func() { BgRed() }, CodeBgRed},
			{"BgGreen", func() { BgGreen() }, CodeBgGreen},
			{"BgYellow", func() { BgYellow() }, CodeBgYellow},
			{"BgBlue", func() { BgBlue() }, CodeBgBlue},
			{"BgMagenta", func() { BgMagenta() }, CodeBgMagenta},
			{"BgCyan", func() { BgCyan() }, CodeBgCyan},
			{"BgWhite", func() { BgWhite() }, CodeBgWhite},
			{"FgBlack", func() { Black() }, CodeFgBlack},
			{"FgRed", func() { Red() }, CodeFgRed},
			{"FgGreen", func() { Green() }, CodeFgGreen},
			{"FgYellow", func() { Yellow() }, CodeFgYellow},
			{"FgBlue", func() { Blue() }, CodeFgBlue},
			{"FgMagenta", func() { Magenta() }, CodeFgMagenta},
			{"FgCyan", func() { Cyan() }, CodeFgCyan},
			{"FgWhite", func() { White() }, CodeFgWhite},
			{"Clear", func() { Clear() }, CodeClear},
			{"LF", func() { LF() }, LineFeed},
			{"Space", func() { Space() }, SpaceChar},
			{"Tab", func() { Tab() }, TabChar},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tt.function()
				if got := mockStdout.String(); got != tt.expected {
					t.Errorf("got %q, want %q", got, tt.expected)
				}
				mockStdout.Reset()
			})
		}
	})

	t.Run("DebugMode", func(t *testing.T) {
		mockStdout := Buffer()
		defer Stdout()
		EnableDebug()
		Debug("debug message")

		expected := "\033[33m[DEBUG]: debug message"
		if got := mockStdout.String(); got != expected {
			t.Errorf("got %q, want %q", got, expected)
		}

		DisableDebug()
		mockStdout.Reset()
		Debug("debug message")

		if got := mockStdout.String(); got != "" {
			t.Errorf("got %q, want %q", got, "")
		}
	})
}
