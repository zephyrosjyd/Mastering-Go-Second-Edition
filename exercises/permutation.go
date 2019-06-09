package main

import (
	"fmt"
)

func goPermutation(str string) {
	permutation(str, "")
}

func permutation(str string, prefix string) {
	if len(str) == 0 {
		fmt.Println(prefix)
	} else {
		for i := 0; i < len(str); i++ {
			rem := str[0:i] + str[i+1:]
			permutation(rem, prefix+string(str[i]))
		}
	}
}

func main() {
	goPermutation("abc")
	fmt.Println()
	goPermutation("xyx")
}
