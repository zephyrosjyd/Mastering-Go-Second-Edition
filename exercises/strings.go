package main

import (
	"fmt"
)

var (
	pl = fmt.Println
	pf = fmt.Printf
)

func main() {
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	pl(sLiteral)
	pf("x: %x\n", sLiteral)

	pf("sLiteral length: %d\n", len(sLiteral))

	for i := 0; i < len(sLiteral); i++ {
		pf("%x ", sLiteral[i])
	}
	pl()

	pf("q: %q\n", sLiteral)
	pf("+q: %+q\n", sLiteral)
	pf(" x: % x\n", sLiteral)

	pf("s: As a string: %s\n", sLiteral)

	s2 := "€£³"
	s2 = "사랑ㅎㅐ요!"
	for x, y := range s2 {
		pf("%#U starts at byte position %d\n", y, x)
	}
	pf("s2 length: %d\n", len(s2))

	const s3 = "ab12AB"
	pl("s3:", s3)
	pf("x: % x\n", s3)
	pf("s3 length: %d\n", len(s3))

	for i := 0; i < len(s3); i++ {
		pf("%x ", s3[i])
	}
	pl()
}
