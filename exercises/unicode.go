package main

import (
	"fmt"
	"unicode"
)

func main() {
	// const sL = "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
	// const sL = "€£³"
	const sL = "사랑ㅎㅐ요!"

	// for i := 0; i < len(sL); i++ {
	for i, c := range sL {
		fmt.Printf("%#U ", c)
		if unicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c %c\n", c, sL[i])
		} else {
			fmt.Println("Not printable!")
		}
	}
}
