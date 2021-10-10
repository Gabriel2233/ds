package stack

import (
	dll "github.com/Gabriel2233/ds/linkedlist/doubly"
)

type Stack struct {
	list *dll.DoublyLinkedList
}

func New() *Stack {
	return &Stack{
		list: dll.New(),
	}
}

func (s *Stack) IsEmpty() bool {
	return s.list.Size() == 0
}

func (s *Stack) Push(d int) int {
	return s.list.AddAtHead(d).Data()
}

func (s *Stack) Pop() int {
    if s.list.Size() == 0 {
        return -1
    }
	return s.list.RemoveAtHead().Data()
}

func (s *Stack) Peek() int {
    if s.list.Size() == 0 {
        return - 1
    }
	return s.list.PeekHead().Data()
}

func (s *Stack) Print() {
	s.list.Print()
}
