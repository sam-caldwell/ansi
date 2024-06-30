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
	fmt.Print(CodeBgBlack)
	return c
}

// BgBlue - Set color
func (c *Color) BgBlue() *Color {
	fmt.Print(CodeBgBlue)
	return c
}

// BgCyan - Set color
func (c *Color) BgCyan() *Color {
	fmt.Print(CodeBgCyan)
	return c
}

// BgGreen - Set color
func (c *Color) BgGreen() *Color {
	fmt.Print(CodeBgGreen)
	return c
}

// BgMagenta - Set color
func (c *Color) BgMagenta() *Color {
	fmt.Print(CodeBgMagenta)
	return c
}

// BgRed - Set color
func (c *Color) BgRed() *Color {
	fmt.Print(CodeBgRed)
	return c
}

// BgWhite - Set color
func (c *Color) BgWhite() *Color {
	fmt.Print(CodeBgWhite)
	return c
}

// BgYellow - Set color
func (c *Color) BgYellow() *Color {
	fmt.Print(CodeBgYellow)
	return c
}

// Black - Set color
func (c *Color) Black() *Color {
	fmt.Print(CodeFgBlack)
	return c
}

// Blink - Set format attribute
func (c *Color) Blink() *Color {
	fmt.Print(CodeBlink)
	return c
}

// Blue - Set color
func (c *Color) Blue() *Color {
	fmt.Print(CodeFgBlue)
	return c
}

// Bold - Set format attribute
func (c *Color) Bold() *Color {
	fmt.Print(CodeBold)
	return c
}

// Clear - Set format attribute
func (c *Color) Clear() *Color {
	fmt.Print(CodeClear)
	return c
}

// Cyan - Set color
func (c *Color) Cyan() *Color {
	fmt.Print(CodeFgCyan)
	return c
}

// Dim - decrease intensity
func (c *Color) Dim() *Color {
	fmt.Print(CodeDim)
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
	fmt.Print(CodeFgGreen)
	return c
}

// Hidden - Set format attribute
func (c *Color) Hidden() *Color {
	fmt.Print(CodeHiddenText)
	return c
}

// Indent - indent n spaces
func (c *Color) Indent(n int) *Color {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
	return c
}

// Left - move cursor n units
func (c *Color) Left(n int) *Color {
	fmt.Printf(CodeMoveLeft, n)
	return c
}

// LF - print a line feed char
func (c *Color) LF() *Color {
	fmt.Print(LineFeed)
	return c
}

// Line - Print a line of 'num' 'ch' characters
func (c *Color) Line(ch string, num int) *Color {
	fmt.Println(strings.Repeat(ch, num))
	return c
}

// Magenta - Set color
func (c *Color) Magenta() *Color {
	fmt.Print(CodeFgMagenta)
	return c
}

// Print - print text to stdout
func (c *Color) Print(message string) *Color {
	fmt.Print(message)
	return c
}

// Printf - print text to stdout
func (c *Color) Printf(format string, a ...any) *Color {
	fmt.Printf(format, a...)
	return c
}

// Println - print text to stdout
func (c *Color) Println(message string) *Color {
	fmt.Println(message)
	return c
}

// Red - Set color
func (c *Color) Red() *Color {
	fmt.Print(CodeFgRed)
	return c
}

// Reverse - Set format attribute
func (c *Color) Reverse() *Color {
	fmt.Print(CodeReverse)
	return c
}

// Reset - Send Reset to stdout
func (c *Color) Reset() *Color {
	defer func() { _ = os.Stdout.Sync() }()
	fmt.Print(CodeReset) // Reset color
	return c
}

// Right - move cursor n units
func (c *Color) Right(n int) *Color {
	fmt.Printf(CodeMoveRight, n)
	return c
}

// Space - print a Space character
func (c *Color) Space() *Color {
	fmt.Print(SpaceChar)
	return c
}

// Strikethrough - Set format attribute
func (c *Color) Strikethrough() *Color {
	fmt.Print(CodeStrikeThrough)
	return c
}

// Tab - print a Tab char
func (c *Color) Tab() *Color {
	fmt.Print(TabChar)
	return c
}

// Time - print current time
func (c *Color) Time() *Color {
	fmt.Print(time.Now().Format(timeFormat))
	return c
}

// TopLeft - Move cursor to top left
func (c *Color) TopLeft() *Color {
	fmt.Print(CodeSetTopLeft)
	return c
}

// Underline - Set format attribute
func (c *Color) Underline() *Color {
	fmt.Print(CodeUnderline)
	return c
}

// Up - move cursor n units
func (c *Color) Up(n int) *Color {
	fmt.Printf(CodeMoveUp, n)
	return c
}

// White - set color
func (c *Color) White() *Color {
	fmt.Print(CodeFgWhite)
	return c
}

// Yellow - set color
func (c *Color) Yellow() *Color {
	fmt.Print(CodeFgYellow)
	return c
}
