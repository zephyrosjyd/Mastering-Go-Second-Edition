package main

import (
	"fmt"
)

var (
	a = 1
	b = "hello"
	c float32
)

func main() {
	aMap := map[string]int{}
	aMap = nil
	fmt.Println(aMap)
	_, ok := aMap["test"]
	fmt.Println(ok, len(aMap))

	fmt.Println(a, b, c)
}
