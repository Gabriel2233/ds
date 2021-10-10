package dll

import (
	"fmt"
)

type DoublyLinkedList struct {
	size int
	head *DlNode
	tail *DlNode
}

type DlNode struct {
	data int
	next *DlNode
	prev *DlNode
}

func (dn *DlNode) Data() int     { return dn.data }
func (dn *DlNode) Next() *DlNode { return dn.next }
func (dn *DlNode) Prev() *DlNode { return dn.prev }

func New() *DoublyLinkedList {
	return &DoublyLinkedList{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (dl *DoublyLinkedList) Size() int {
	return dl.size
}

func (dl *DoublyLinkedList) PeekHead() *DlNode {
	if dl.size == 0 {
		return nil
	}

	return dl.head
}

func (dl *DoublyLinkedList) AddAtHead(d int) *DlNode {
	node := &DlNode{
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

func (dl *DoublyLinkedList) PeekTail() *DlNode {
	if dl.size == 0 {
		return nil
	}

	return dl.tail
}

func (dl *DoublyLinkedList) AddAtTail(d int) *DlNode {
	node := &DlNode{
		data: d,
	}

	if dl.head == nil && dl.tail == nil {
		dl.AddAtHead(d)
	} else {
		node.prev = dl.tail
		node.next = nil
		dl.tail.next = node
		dl.tail = dl.tail.next
	    dl.size++
	}

	return node
}

func (dl *DoublyLinkedList) AddAt(i, d int) *DlNode {
	if i < 0 || i >= dl.size {
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

	node := &DlNode{
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

	dl.size++

	return node
}

func (dl *DoublyLinkedList) RemoveAtHead() *DlNode {
	if dl.head == nil {
		return nil
	}

    if dl.size == 1 {
        first := dl.head
        dl.head = nil
        dl.tail = nil
        dl.size--
        return first
    }

	first := dl.head
	dl.head = dl.head.next
	dl.head.prev = nil
	dl.size--

	return first
}

func (dl *DoublyLinkedList) RemoveAtTail() *DlNode {
	if dl.tail == nil {
		return nil
	}

    if dl.size == 1 {
        return dl.RemoveAtHead()
    }

	last := dl.tail
	dl.tail = dl.tail.prev
	dl.tail.next = nil
	dl.size--

	return last
}

func (dl *DoublyLinkedList) RemoveAt(i int) *DlNode {
	if i < 0 || i >= dl.size {
		return nil
	}

	if i == 0 {
		return dl.RemoveAtHead()
	}

	if i == dl.size-1 {
		return dl.RemoveAtTail()
	}

	cur := dl.head
	idx := 0

	var node *DlNode
	for cur.next != nil {
		if idx == i {
			node = cur
			cur.prev.next = cur.next
			cur.next.prev = cur.prev
			break
		}
		idx++
		cur = cur.next
	}

	dl.size--
	return node
}

func (dl *DoublyLinkedList) Print() {
	for cur := dl.head; cur != nil; {
		fmt.Println(cur.data)
		cur = cur.next
	}

	fmt.Println()
}
