package stack

import dl "github.com/Gabriel2233/ds/linkedlist/doubly"

type Stack struct {
	list *dl.DoublyLinkedList
}

func New() *Stack {
	return &Stack{
		list: dl.New(),
	}
}

func (s *Stack) IsEmpty() bool {
	return s.list.Size() == 0
}

func (s *Stack) Push(d int) {
	s.list.AddAtHead(d)
}

func (s *Stack) Pop() {
	s.list.RemoveAtHead()
}

func (s *Stack) Peek() int {
	node := s.list.PeekHead()
	return node.Data()
}

func (s *Stack) Print() {
	s.list.Print()
}
