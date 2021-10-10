package queue

import (
	"fmt"

	dll "github.com/Gabriel2233/ds/linkedlist/doubly"
)

type Queue struct {
	list *dll.DoublyLinkedList
}

func New() *Queue {
	list := dll.New()
	return &Queue{
		list: list,
	}
}

func (q *Queue) Size() int { return q.list.Size() }

func (q *Queue) IsEmpty() bool { return q.list.Size() == 0 }

func (q *Queue) Peek() int {
	if q.Size() == 0 {
		return -1
	}
	return q.list.PeekHead().Data()
}

func (q *Queue) Enqueue(d int) int { return q.list.AddAtTail(d).Data() }

func (q *Queue) Dequeue() int {
	if q.Size() == 0 {
		fmt.Println("Cannot dequeue from an empty queue")
		return -1
	}

	return q.list.RemoveAtHead().Data()
}

func (q *Queue) Print() { q.list.Print() }
