package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage: %s logFile\n", filepath.Base(args[0]))
		os.Exit(1)
	}

	for _, fname := range args[1:] {
		f, err := os.Open(fname)
		if err != nil {
			fmt.Println("error opening file %s", err)
			os.Exit(-1)
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("error reading file %s", err)
				break
			}

			ip := findIP(line)
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}
			fmt.Println(ip)
		}
	}
}
