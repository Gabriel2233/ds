package dl

import (
	"fmt"
)

type DoublyLinkedList struct {
	size int
	head *Node
	tail *Node
}

type Node struct {
	data int
	next *Node
	prev *Node
}

func New() *DoublyLinkedList {
	return &DoublyLinkedList{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (dl *DoublyLinkedList) AddAtHead(d int) *Node {
	node := &Node{
		data: d,
	}

	if dl.head == nil && dl.tail == nil {
		node.next = nil
		node.prev = nil
		dl.tail = node
	} else {
		node.next = dl.head
		node.prev = nil
	}

	dl.head = node
	dl.size++

	return node
}

func (dl *DoublyLinkedList) AddAtTail(d int) *Node {
	node := &Node{
		data: d,
	}

	if dl.head == nil && dl.tail == nil {
		dl.AddAtHead(d)
	} else {
		node.prev = dl.tail
		node.next = nil
		dl.tail.next = node
		dl.tail = dl.tail.next
	}

	dl.size++

	return node
}

func (dl *DoublyLinkedList) AddAt(i, d int) *Node {
	if i < 0 {
		return nil
	}

	if i == 0 {
		node := dl.AddAtHead(d)
		return node
	}

	if i == dl.size-1 {
		node := dl.AddAtTail(d)
		return node
	}

	node := &Node{
		data: d,
	}

	cur := dl.head
	idx := 0

	for cur.next != nil {
		if idx == i-1 {
			// 1 <-> 2 <-> 3
			// 1 <-> 4 <-> 2 <-> 3
			// insert 4 at index 1 (2)
			// stop at 1
			// 4 next = 1's next
			// 4 prev = 1
			// 1 next = 4
			// 2 prev = 4
			node.next = cur.next
			node.prev = cur
			cur.next = node
			cur.next.next.prev = node
		}
		cur = cur.next
		idx++
	}

	return node
}

func (dl *DoublyLinkedList) RemoveAtHead() *Node {
	if dl.head == nil {
		return nil
	}

	first := dl.head
	dl.head = dl.head.next
	dl.head.prev = nil

	return first
}

func (dl *DoublyLinkedList) RemoveAtTail() *Node {
	if dl.tail == nil {
		return nil
	}

	last := dl.tail
	dl.tail = dl.tail.prev
	dl.tail.next = nil

	return last
}

func (dl *DoublyLinkedList) RemoveAt(i int) *Node {
	if i < 0 {
		return nil
	}

	if i == 0 {
		node := dl.RemoveAtHead()
		return node
	}

	if i == dl.size-1 {
		node := dl.RemoveAtTail()
		return node
	}

	cur := dl.head
	idx := 0

	var node *Node
	for cur.next != nil {
		if idx == i {
			node = cur
			cur.prev.next = cur.next
			cur.next.prev = cur.prev
		}
		idx++
		cur = cur.next
	}

	return node
}

func (dl *DoublyLinkedList) Print() {
	for cur := dl.head; cur != nil; {
		fmt.Println(cur.data)
		cur = cur.next
	}

	fmt.Println()
}
