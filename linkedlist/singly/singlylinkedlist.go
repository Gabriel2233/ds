package ll

import "fmt"

type SinglyLinkedList struct {
	size int
	head *Node
	tail *Node
}

type Node struct {
	data int
	next *Node
}

func (dn *Node) Data() int   { return dn.data }
func (dn *Node) Next() *Node { return dn.next }

func New() *SinglyLinkedList {
	return &SinglyLinkedList{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (sl *SinglyLinkedList) AddAt(i, d int) *Node {
    if(sl.size == 0) {
        return sl.AddAtHead(d);
    }

    // 1 -> 4 -> 6 -> 0
    // insert 10 at pos 1 (actual 4)
    // 1 -> 10 -> 4 -> 6 -> 0

    // cur stops at one before
    // just make 10 point to prev's next and prev point to new

	node := &Node{data: d}
    idx := 0

	for cur := sl.head; cur != nil; cur = cur.next {
		if idx == i - 1 {
            node.next = cur.next
            cur.next = node
            break
		}

        idx++
	}

    return node
}

func (sl *SinglyLinkedList) SearchAt(i int) *Node {
	if sl.size == 0 {
		return nil
	}

	var node *Node
	idx := 0

	for cur := sl.head; cur != nil; cur = cur.next {
		if i == idx {
			node = cur
			break
		}
		idx++
	}

	return node
}

func (sl *SinglyLinkedList) PeekHead() *Node {
	if sl.size == 0 {
		return nil
	}
	return sl.head
}

func (sl *SinglyLinkedList) AddAtHead(d int) *Node {
	node := &Node{data: d}
	if sl.head == nil && sl.tail == nil {
		node.next = nil
		sl.head = node
		sl.tail = node
	} else {
		node.next = sl.head
		sl.head = node
	}

	sl.size++
	return node
}

func (sl *SinglyLinkedList) RemoveAtHead() *Node {
	if sl.size == 0 {
		return nil
	}

	data := sl.head
	sl.head = sl.head.next

	sl.size--

	return data
}

func (sl *SinglyLinkedList) RemoveAtTail() *Node {
	if sl.size == 0 {
		return nil
	}

	prev := sl.SearchAt(sl.size - 2)
	data := sl.tail

	sl.tail = prev
	sl.tail.next = nil

	sl.size--

	return data
}

func (sl *SinglyLinkedList) PeekTail() *Node {
	if sl.size == 0 {
		return nil
	}
	return sl.tail
}

func (sl *SinglyLinkedList) AddAtTail(d int) *Node {
	var node *Node
	if sl.head == nil && sl.tail == nil {
		node = sl.AddAtHead(d)
	} else {
		node = &Node{data: d}
		node.next = nil
		sl.tail.next = node
		sl.tail = node
	}

	sl.size++
	return node
}

func (sl *SinglyLinkedList) RemoveAt(i int) *Node {
	if i < 0 || i >= sl.size {
		return nil
	}

	if i == 0 {
		sl.size--
		return sl.RemoveAtHead()
	}

	if i == sl.size-1 {
		sl.size--
		return sl.RemoveAtTail()
	}

	idx := 0
	cur := sl.head
	target := cur.next

	var node *Node
	for target.next != nil {
		if idx == i-1 {
			node = target
			cur.next = target.next
			target = nil
			break
		}

		target = target.next
		cur = cur.next
		idx++
	}

	sl.size--

	return node
}

func (sl *SinglyLinkedList) Size() int {
	return sl.size
}

func (sl *SinglyLinkedList) Print() {
	cur := sl.head

	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}
	fmt.Println()
}
