package main

import "fmt"

var p = fmt.Println

func main() {
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	p("a6:", a6)
	p("a4:", a4)

	copy(a6, a4)
	p("a6:", a6)
	p("a4:", a4)
	p()

	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}
	p("b6:", b6)
	p("b4:", b4)
	copy(b4, b6)
	p("b6:", b6)
	p("b4:", b4)
	p()

	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, 1, -1, -1, 5, -5}
	copy(s6, array4[0:])
	p("array4:", array4[0:])
	p("s6:", s6)

	p()

	array5 := [5]int{5, -5, 5, -5, 5}
	s7 := []int{7, 7, -7, -7, 7, -7, 7}
	copy(array5[:], s7)
	p("array5:", array5)
	p("s7:", s7)
}
