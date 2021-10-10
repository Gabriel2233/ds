package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var queue = New()

func TestEnqueue(t *testing.T) {
    assert.Equal(t, 0, queue.Size())
    assert.True(t, true, queue.IsEmpty())

    assert.Equal(t, 4, queue.Enqueue(4))
    assert.Equal(t, 1, queue.Size())

    assert.Equal(t, 3, queue.Enqueue(3))
    assert.Equal(t, 2, queue.Size())
}

func TestDequeue(t *testing.T) {
    assert.Equal(t, 4, queue.Dequeue())
    assert.Equal(t, 1, queue.Size())

    assert.Equal(t, 3, queue.Dequeue())
    assert.Equal(t, 0, queue.Size())


    assert.Equal(t, -1, queue.Dequeue())
    assert.Equal(t, 0, queue.Size())
}

func TestPeek(t *testing.T) {
    queue.Enqueue(2)
    assert.Equal(t, 1, queue.Size())
    assert.Equal(t, 2, queue.Peek())

    queue.Enqueue(1)
    assert.Equal(t, 2, queue.Size())
    assert.Equal(t, 2, queue.Peek())

    queue.Dequeue()
    assert.Equal(t, 1, queue.Size())
    assert.Equal(t, 1, queue.Peek())

    queue.Dequeue()
    assert.Equal(t, 0, queue.Size())
    assert.Equal(t, -1, queue.Peek())
}
