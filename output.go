package ansi

import (
	"bytes"
	"io"
	"os"
)

/*
 * Ansi Colors - output
 *
 * (c) 2023 Sam Caldwell.  MIT License
 */

var out io.Writer

// Stdout - Set output (stdout)
func Stdout() *Color {
	return (&Color{}).Stdout()
}

// Stdout - Set output (stdout)
func (c *Color) Stdout() *Color {
	out = os.Stdout
	return c
}

// Stderr - Set output (stderr)
func Stderr() *Color {
	return (&Color{}).Stderr()
}

// Stderr - Set output (stderr)
func (c *Color) Stderr() *Color {
	out = os.Stderr
	return c
}

// Buffer - Set output to a bytes buffer (for streaming, testing, etc).
func Buffer() *bytes.Buffer {
	buffer := &bytes.Buffer{}
	out = buffer
	return buffer
}
