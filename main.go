package main

import dl "github.com/Gabriel2233/ds/linkedlist"

func main() {
	list := dl.New()

	list.AddAtHead(5)
	list.AddAtHead(6)
	list.AddAtTail(7)

	list.Print()

	list.AddAtHead(4)
	list.AddAtTail(9)

	list.Print()

	list.RemoveAtHead()

	list.Print()

	list.RemoveAtTail()

	list.Print()

	list.AddAt(1, 80)
	list.AddAt(1, 90)

	list.AddAt(0, 100)
	list.AddAt(5, 88)

	list.RemoveAt(0)
	list.RemoveAt(6)

	list.RemoveAt(2)

	list.RemoveAt(1)

	list.Print()
}
