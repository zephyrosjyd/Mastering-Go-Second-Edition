package main

import "io"
import "os"

func main() {
	str := ""
	args := os.Args
	var fp *os.File
	if len(args) == 1 {
		str = "Please give me one argument!"
		fp = os.Stderr
	} else {
		str = args[1]
		fp = os.Stdout
	}

	io.WriteString(fp, str)
	io.WriteString(fp, "\n")
}
