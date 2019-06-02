package main

import (
	"fmt"
	"sort"
)

type aStruct struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]aStruct, 0)
	mySlice = append(mySlice, aStruct{"Mihalis", 180, 90})
	mySlice = append(mySlice, aStruct{"Bill", 134, 45})
	mySlice = append(mySlice, aStruct{"Marietta", 155, 45})
	mySlice = append(mySlice, aStruct{"Epifanios", 144, 50})
	mySlice = append(mySlice, aStruct{"Athina", 134, 40})
	fmt.Println("0:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)
}
