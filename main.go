package main

import (
	"fmt"

	"github.com/Gabriel2233/ds/priorityqueue"
)

func main() {
    pq := pq.New(true)
    
    pq.Add(-1)
    pq.Add(3)
    pq.Add(2)
    pq.Add(1)
    pq.Add(0)
    pq.Add(5)
    pq.Add(10)

    pq.Print()

    pq.Pool()

    pq.Print()

    pq.RemoveAt(2)

    pq.Print()

    fmt.Println(pq.IsMinHeap(0))
}
