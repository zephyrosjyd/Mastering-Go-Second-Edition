package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Head *Node
	Size int
}

func (q *Queue) Push(v int) bool {
	if q == nil {
		return false
	}

	node := &Node{v, nil}
	node.Next = q.Head
	q.Head = node
	q.Size++

	return true
}

func (q *Queue) Pop() (int, bool) {
	if q.Size == 0 {
		return 0, false
	}

	if q.Size == 1 {
		value := q.Head.Value
		q.Head = nil
		q.Size = 0
		return value, true
	}

	node := q.Head
	for node.Next.Next != nil {
		node = node.Next
	}

	value := (node.Next).Value
	node.Next = nil

	q.Size--
	return value, true
}

func (q *Queue) traverse() {
	if q.Size == 0 {
		fmt.Println("Empty Queue!")
		return
	}

	node := q.Head
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Value)
		node = node.Next
	}
	fmt.Printf("%d -> \n", node.Value)
}

func main() {
	queue := new(Queue)
	queue.Push(10)
	fmt.Println("Size:", queue.Size)
	queue.traverse()

	v, b := queue.Pop()
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", queue.Size)

	for i := 0; i < 5; i++ {
		queue.Push(i)
	}
	queue.traverse()
	fmt.Println("Size:", queue.Size)

	v, b = queue.Pop()
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", queue.Size)

	v, b = queue.Pop()
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", queue.Size)
	queue.traverse()
}
