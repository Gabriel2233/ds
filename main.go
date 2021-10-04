package main

import (
	ll "github.com/Gabriel2233/ds/linkedlist/singly"
)

func main() {
	l := ll.New()

	l.AddAtHead(4)
	l.AddAtHead(5)
	l.AddAtTail(6)
	l.AddAtTail(7)

	l.AddAtHead(0)
	l.AddAtTail(99)
	l.Print()

	l.RemoveAtTail()
	l.Print()

	l.RemoveAt(2)
	l.Print()

    l.AddAt(1, 10)
    l.Print()
}
