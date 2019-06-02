package main

import "C"

import (
	"fmt"
)

//export PrintMessage
func PrintMessage() {
	fmt.Println("A Go function!")
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

//export Add
func Add(a, b int) int {
	return a + b
}

//export Subtract
func Subtract(a, b int) int {
	return a - b
}

//export divide
func divide(a, b int) float32 {
	return float32(a) / float32(b)
}

func main() {}
