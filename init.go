package ansi

import (
	"os"
)

// init
//
//	(c) 2023 Sam Caldwell.  MIT License
func init() {
	out = os.Stdout
	Reset()
}
