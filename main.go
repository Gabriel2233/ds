package main

import (

	"github.com/Gabriel2233/ds/queue"
)

func main() {
    q := queue.New()

    q.Enqueue(2)
    println("queue has only 2, head is: ", q.Peek(), " size is: ", q.Size())
    q.Enqueue(3)
    println("enqueued 3, head is still: ", q.Peek(), " size is: ", q.Size())
    q.Dequeue()
    println("removed 2, head is: ", q.Peek(), " size is: ", q.Size())
    q.Dequeue()
    println("removed 3, head is: ", q.Peek(), " size is: ", q.Size())
}
