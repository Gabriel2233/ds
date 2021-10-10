package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var stack = New()

func TestPush(t *testing.T) {
    assert.Equal(t, 4, stack.Push(4))
    assert.Equal(t, 3, stack.Push(3))
}

func TestPop(t *testing.T) {
    assert.Equal(t, 3, stack.Pop())
    assert.Equal(t, 4, stack.Pop())
    assert.Equal(t, -1, stack.Pop())
}

func TestPeek(t *testing.T) {
    stack.Push(2)
    assert.Equal(t, 2, stack.Peek())


    stack.Push(1)
    assert.Equal(t, 1, stack.Peek())

    stack.Pop()
    assert.Equal(t, 2, stack.Peek())

    stack.Pop()
    assert.Equal(t, -1, stack.Peek())
}
