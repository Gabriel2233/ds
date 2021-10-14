package pq

import "fmt"

type PQueue struct {
	isMin bool
	heap  []int
}

func New(isMin bool) *PQueue {
	var heap []int
	return &PQueue{
		isMin: isMin,
		heap:  heap,
	}
}

func (pq *PQueue) Size() int { return len(pq.heap) }

func (pq *PQueue) IsEmpty() bool { return len(pq.heap) == 0 }

func (pq *PQueue) Peek() int { return pq.heap[0] }

func (pq *PQueue) Add(d int) {
	pq.heap = append(pq.heap, d)
	lastIdx := len(pq.heap) - 1
	pq.swim(lastIdx)
}

func (pq *PQueue) Pool() {
	pq.RemoveAt(0)
}

func (pq *PQueue) RemoveAt(i int) {
	if i >= pq.Size() {
		return
	}

	pq.swap(i, pq.Size()-1)

	pq.heap = pq.heap[:pq.Size()-1]

	pq.sink(i)

    // if the invariant is not satisfied for the max heap, try to swim the index
    if !pq.isMin && !pq.IsMinHeap(0) {
        pq.swim(i)
    }

    // if the invariant is not satisfied for the min heap, try to swim the index
    if pq.isMin && pq.IsMinHeap(0) {
        pq.swim(i)
    }
}

func (pq *PQueue) Print() {
	if pq.IsEmpty() {
		return
	}

	line := 1
	elsInLine := 0

	for i := range pq.heap {
		if elsInLine < line {
			fmt.Printf("%d ", pq.heap[i])
			elsInLine++
		} else {
			fmt.Println()
			fmt.Printf("%d ", pq.heap[i])
			elsInLine = 0
			line *= 2
			elsInLine++
		}
	}

	fmt.Println()
	fmt.Println()
}

func (pq *PQueue) IsMinHeap(i int) bool {
	if i >= pq.Size() {
		return true
	}

	l := 2*i + 1
	r := 2*i + 2

    // if the next left node is bigger, the min heap invariant is not satisfied, 
    // because all the next nodes must be smaller than their parents
	if l < pq.Size() && pq.isSmaller(l, i) {
		return false
	}
    // if the next right node is bigger, the min heap invariant is not satisfied, 
    // because all the next nodes must be smaller than their parents
	if r < pq.Size() && pq.isSmaller(r, i) {
		return false
	}

    // up in the top, we return true if we can get to the bottom of the heap. If both of the conditions are not met
    // we return false. this prevents only one part of the heap being satisfied.
	return pq.IsMinHeap(l) && pq.IsMinHeap(r)
}

func (pq *PQueue) isSmaller(i, j int) bool {
	return pq.heap[i] < pq.heap[j]
}

func (pq *PQueue) swim(i int) {
	parent := (i - 1) / 2

    // if it's a min heap, we should swap the current node with it's parent if the parent is bigger
	if pq.isMin {
		for i > 0 && pq.isSmaller(i, parent) {
			pq.swap(i, parent)
			i = parent
			parent = (i - 1) / 2
		}
	} else {
		for i > 0 && !pq.isSmaller(i, parent) {
			pq.swap(i, parent)
			i = parent
			parent = (i - 1) / 2
		}
	}
}

func (pq *PQueue) sink(i int) {
	for {
		l := 2*i + 1
		r := 2*i + 2

		var nextSelectedIndex int
		// we're at a leaf node here, cannot sink anymore :D
		if l >= pq.Size() && r >= pq.Size() {
			break
		}

		if l >= pq.Size() && r < pq.Size() {
			// a case where we only got a right child after the actual node
			// next nodes in min heap must be bigger
			if pq.isMin && pq.isSmaller(i, r) {
				break
			}
			// next nodes in max heap must be smaller
			if !pq.isMin && !pq.isSmaller(i, r) {
				break
			}
			nextSelectedIndex = r
		} else if r >= pq.Size() && l < pq.Size() {
			// a case where we only got a left child after the actual node
			// next nodes in min heap must be bigger
			if pq.isMin && pq.isSmaller(i, l) {
				break
			}
			// next nodes in max heap must be smaller
			if !pq.isMin && !pq.isSmaller(i, l) {
				break
			}
			nextSelectedIndex = l
		} else {
			// a case where right and left are present
			if pq.isMin && pq.isSmaller(l, r) {
				nextSelectedIndex = l
			} else if pq.isMin && pq.isSmaller(r, l) {
				nextSelectedIndex = r
            }

            if !pq.isMin && !pq.isSmaller(l, r) {
				nextSelectedIndex = l
			} else if !pq.isMin && !pq.isSmaller(r, l) {
				nextSelectedIndex = r
			}
		}

		pq.swap(i, nextSelectedIndex)
		i = nextSelectedIndex
	}
}

func (pq *PQueue) swap(i, j int) {
	tmp := pq.heap[j]
	pq.heap[j] = pq.heap[i]
	pq.heap[i] = tmp
}
