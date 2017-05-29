package main

import (
	"flag"
	"os"
)

var NewLine = flag.Bool("n", false, "print on newline")

const (
	SPACE   = " "
	NEWLINE = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += SPACE
		}
		s += flag.Arg(i)
	}

	if *NewLine {
		s += NEWLINE
	}
	os.Stdout.WriteString(s)
}
