package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/exit"
)

func main() {
	code := flag.Int("code", 0, "")
	flag.Parse()
	fmt.Printf("code: %v", *code)
	c := exit.Code(*code)
	ansi.Fatal(c)
}
