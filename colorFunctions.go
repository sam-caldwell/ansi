package ansi

import "github.com/sam-caldwell/exit/v2"

/*
 * Ansi Colors - Functions
 *
 * This file contains the functions which can be used as
 * the root of a command-chain (see colorMethods.go) or
 * as stand-alone functions.
 *
 * (c) 2023 Sam Caldwell.  MIT License
 */

// BgBlack - set color and return a new color object
func BgBlack() *Color {
	return (&Color{}).BgBlack()
}

// BgBlue - set color and return a new color object
func BgBlue() *Color {
	return (&Color{}).BgBlue()
}

// BgCyan - set color and return a new color object
func BgCyan() *Color {
	return (&Color{}).BgCyan()
}

// BgGreen - set color and return a new color object
func BgGreen() *Color {
	return (&Color{}).BgGreen()
}

// BgMagenta - set color and return a new color object
func BgMagenta() *Color {
	return (&Color{}).BgMagenta()
}

// BgRed - set color and return a new color object
func BgRed() *Color {
	return (&Color{}).BgRed()
}

// BgWhite - set color and return a new color object
func BgWhite() *Color { return (&Color{}).BgWhite() }

// BgYellow - set color and return a new color object
func BgYellow() *Color {
	return (&Color{}).BgYellow()
}

// Black - set color and return a new color object
func Black() *Color {
	return (&Color{}).Black()
}

// Blink - set format attribute and return a new color object
func Blink() *Color {
	return (&Color{}).Blink()
}

// Blue - set color and return a new color object
func Blue() *Color {
	return (&Color{}).Blue()
}

// Bold - set format attribute and return a new color object
func Bold() *Color {
	return (&Color{}).Bold()
}

// Clear - set format attribute and return a new color object
func Clear() *Color {
	return (&Color{}).Clear()
}

// Cyan - set color and return a new color object
func Cyan() *Color {
	return (&Color{}).Cyan()
}

// Dim - decrease intensity and return a new color object
func Dim() *Color {
	return (&Color{}).Dim()
}

// Down - move cursor n units and return a new color object
func Down(n int) *Color {
	return (&Color{}).Down(n)
}

// Fatal - Terminate the program and return the exit code
func Fatal(exitCode exit.Code) *Color {
	return (&Color{}).Fatal(exitCode)
}

// Flush - Flush StdOut buffer
func Flush() *Color {
	return (&Color{}).Flush()
}

// Green - set color and return a new color object
func Green() *Color {
	return (&Color{}).Green()
}

// Hidden - set format attribute and return a new color object
func Hidden() *Color {
	return (&Color{}).Hidden()
}

// Indent - indent n spaces
func Indent(n int) *Color {
	return (&Color{}).Indent(n)
}

// Left - move cursor n units and return a new color object
func Left(n int) *Color {
	return (&Color{}).Left(n)
}

// LF - print a line feed char
func LF() *Color {
	return (&Color{}).LF()
}

// Line - Print a line of 'num' 'ch' characters
func Line(ch string, num int) *Color {
	return (&Color{}).Line(ch, num)
}

// Magenta - set color and return a new color object
func Magenta() *Color {
	return (&Color{}).Magenta()
}

// Print - print text to stdout and return color object
func Print(message string) *Color {
	return (&Color{}).Print(message)
}

// Printf - print text to stdout and return color object
func Printf(format string, a ...any) *Color {
	return (&Color{}).Printf(format, a...)
}

// Println - print text to stdout and return color object
func Println(message string) *Color {
	return (&Color{}).Println(message)
}

// Red - set color and return a new color object
func Red() *Color {
	return (&Color{}).Red()
}

// Reverse - set format attribute and return a new color object
func Reverse() *Color {
	return (&Color{}).Reverse()
}

// Reset - Send Reset to stdout and return new color object
func Reset() *Color {
	return (&Color{}).Reset()
}

// Right - move cursor n units and return a new color object
func Right(n int) *Color {
	return (&Color{}).Right(n)
}

// Space - print a Space character
func Space() *Color {
	return (&Color{}).Space()
}

// Strikethrough - set format attribute and return a new color object
func Strikethrough() *Color {
	return (&Color{}).Strikethrough()
}

// Tab - print a Tab char
func Tab() *Color {
	return (&Color{}).Tab()
}

// Time - print current time
func Time() *Color {
	return (&Color{}).Time()
}

// TopLeft - Move cursor to top left and return a new color object
func TopLeft() *Color {
	return (&Color{}).TopLeft()
}

// Underline - set format attribute and return a new color object
func Underline() *Color {
	return (&Color{}).Underline()
}

// Up - move cursor n units and return a new color object
func Up(n int) *Color {
	return (&Color{}).Up(n)
}

// White - set color and return a new color object
func White() *Color {
	return (&Color{}).White()
}

// Yellow - set color and return a new color object
func Yellow() *Color {
	return (&Color{}).Yellow()
}
