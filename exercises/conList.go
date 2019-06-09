package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func printList(l *list.List) {
	for t := l.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()

	for t := l.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()
}

func main() {
	values := list.New()
	fmt.Printf("After New(): %v\n", values)

	e1 := values.PushBack("one")
	e2 := values.PushBack("two")
	fmt.Printf("e1: %v, e2: %v\n", e1, e2)
	values.PushFront("three")
	values.InsertBefore("four", e1)
	values.InsertAfter("five", e2)
	values.Remove(e2)
	values.Remove(e2)
	values.InsertAfter("fivefive", e2)
	values.PushBackList(values)

	printList(values)

	values.Init()
	fmt.Printf("After Init(): %v\n", values)

	for i := 0; i < 20; i++ {
		values.PushFront(strconv.Itoa(i))
	}

	printList(values)
}
