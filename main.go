package main

import "github.com/data-structures/list"

func main() {
	l := list.NewInt64List()
	l.PushBack(1)
	l.PushBack(5)
	l.PushBack(6)

	l.PrintList()
	l.RPrintList()

	l.PushFront(8)
	l.PushFront(-6)
	l.PushFront(0)

	l.PrintList()
	l.RPrintList()
}
