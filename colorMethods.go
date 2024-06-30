package ansi

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/*
 * Ansi Colors - Methods
 *
 * The Color object methods in this file allow command chaining
 * such as `ansi.Blue().Printf().LF().Reset()`
 *
 * (c) 2023 Sam Caldwell.  MIT License
 */

// BgBlack - Set color
func (c *Color) BgBlack() *Color {
	_, _ = fmt.Fprint(out, CodeBgBlack)
	return c
}

// BgBlue - Set color
func (c *Color) BgBlue() *Color {
	_, _ = fmt.Fprint(out, CodeBgBlue)
	return c
}

// BgCyan - Set color
func (c *Color) BgCyan() *Color {
	_, _ = fmt.Fprint(out, CodeBgCyan)
	return c
}

// BgGreen - Set color
func (c *Color) BgGreen() *Color {
	_, _ = fmt.Fprint(out, CodeBgGreen)
	return c
}

// BgMagenta - Set color
func (c *Color) BgMagenta() *Color {
	_, _ = fmt.Fprint(out, CodeBgMagenta)
	return c
}

// BgRed - Set color
func (c *Color) BgRed() *Color {
	_, _ = fmt.Fprint(out, CodeBgRed)
	return c
}

// BgWhite - Set color
func (c *Color) BgWhite() *Color {
	_, _ = fmt.Fprint(out, CodeBgWhite)
	return c
}

// BgYellow - Set color
func (c *Color) BgYellow() *Color {
	_, _ = fmt.Fprint(out, CodeBgYellow)
	return c
}

// Black - Set color
func (c *Color) Black() *Color {
	_, _ = fmt.Fprint(out, CodeFgBlack)
	return c
}

// Blink - Set format attribute
func (c *Color) Blink() *Color {
	_, _ = fmt.Fprint(out, CodeBlink)
	return c
}

// Blue - Set color
func (c *Color) Blue() *Color {
	_, _ = fmt.Fprint(out, CodeFgBlue)
	return c
}

// Bold - Set format attribute
func (c *Color) Bold() *Color {
	_, _ = fmt.Fprint(out, CodeBold)
	return c
}

// Clear - Set format attribute
func (c *Color) Clear() *Color {
	_, _ = fmt.Fprint(out, CodeClear)
	return c
}

// Cyan - Set color
func (c *Color) Cyan() *Color {
	_, _ = fmt.Fprint(out, CodeFgCyan)
	return c
}

// Dim - decrease intensity
func (c *Color) Dim() *Color {
	_, _ = fmt.Fprint(out, CodeDim)
	return c
}

// Down - move cursor n units
func (c *Color) Down(n int) *Color {
	fmt.Printf(CodeMoveDown, n)
	return c
}

// Fatal - Terminate the program and return the exit code
func (c *Color) Fatal(exitCode int) *Color {
	Reset()
	//safety.Defer(func() { os.Exit(exitCode) })
	os.Exit(exitCode)
	return c
}

// Flush - Flush StdOut buffer
func (c *Color) Flush() *Color {
	defer func() { _ = os.Stdout.Sync() }()
	return c
}

// Green - Set color
func (c *Color) Green() *Color {
	_, _ = fmt.Fprint(out, CodeFgGreen)
	return c
}

// Hidden - Set format attribute
func (c *Color) Hidden() *Color {
	_, _ = fmt.Fprint(out, CodeHiddenText)
	return c
}

// Indent - indent n spaces
func (c *Color) Indent(n int) *Color {
	for i := 0; i < n; i++ {
		_, _ = fmt.Fprint(out, " ")
	}
	return c
}

// Left - move cursor n units
func (c *Color) Left(n int) *Color {
	_, _ = fmt.Fprintf(out, CodeMoveLeft, n)
	return c
}

// LF - print a line feed char
func (c *Color) LF() *Color {
	_, _ = fmt.Fprint(out, LineFeed)
	return c
}

// Line - Print a line of 'num' 'ch' characters
func (c *Color) Line(ch string, num int) *Color {
	_, _ = fmt.Fprintln(out, strings.Repeat(ch, num))
	return c
}

// Magenta - Set color
func (c *Color) Magenta() *Color {
	_, _ = fmt.Fprint(out, CodeFgMagenta)
	return c
}

// Print - print text to stdout
func (c *Color) Print(message string) *Color {
	_, _ = fmt.Fprint(out, message)
	return c
}

// Printf - print text to stdout
func (c *Color) Printf(format string, a ...any) *Color {
	_, _ = fmt.Fprintf(out, format, a...)
	return c
}

// Println - print text to stdout
func (c *Color) Println(message string) *Color {
	_, _ = fmt.Fprintln(out, message)
	return c
}

// Red - Set color
func (c *Color) Red() *Color {
	_, _ = fmt.Fprint(out, CodeFgRed)
	return c
}

// Reverse - Set format attribute
func (c *Color) Reverse() *Color {
	_, _ = fmt.Fprint(out, CodeReverse)
	return c
}

// Reset - Send Reset to stdout
func (c *Color) Reset() *Color {
	defer func() { _ = os.Stdout.Sync() }()
	_, _ = fmt.Fprint(out, CodeReset) // Reset color
	return c
}

// Right - move cursor n units
func (c *Color) Right(n int) *Color {
	fmt.Printf(CodeMoveRight, n)
	return c
}

// Space - print a Space character
func (c *Color) Space() *Color {
	_, _ = fmt.Fprint(out, SpaceChar)
	return c
}

// Strikethrough - Set format attribute
func (c *Color) Strikethrough() *Color {
	_, _ = fmt.Fprint(out, CodeStrikeThrough)
	return c
}

// Tab - print a Tab char
func (c *Color) Tab() *Color {
	_, _ = fmt.Fprint(out, TabChar)
	return c
}

// Time - print current time
func (c *Color) Time() *Color {
	_, _ = fmt.Fprint(out, time.Now().Format(timeFormat))
	return c
}

// TopLeft - Move cursor to top left
func (c *Color) TopLeft() *Color {
	_, _ = fmt.Fprint(out, CodeSetTopLeft)
	return c
}

// Underline - Set format attribute
func (c *Color) Underline() *Color {
	_, _ = fmt.Fprint(out, CodeUnderline)
	return c
}

// Up - move cursor n units
func (c *Color) Up(n int) *Color {
	fmt.Printf(CodeMoveUp, n)
	return c
}

// White - set color
func (c *Color) White() *Color {
	_, _ = fmt.Fprint(out, CodeFgWhite)
	return c
}

// Yellow - set color
func (c *Color) Yellow() *Color {
	_, _ = fmt.Fprint(out, CodeFgYellow)
	return c
}
