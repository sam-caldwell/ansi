package ansi

import (
	"fmt"
)

// Debug - Print debug message
//
//	(c) 2023 Sam Caldwell.  MIT License
func Debug(msg interface{}) {
	if debugMode {
		Yellow().Print(fmt.Sprintf("[Debug]: %v", msg)).Reset()
	}
}
