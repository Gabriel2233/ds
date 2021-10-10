package sll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var list = New()

func TestAddAtHead(t *testing.T) {
	t.Run("it should be the first node", func(t *testing.T) {
		got := list.AddAtHead(1)

		assert.Equal(t, list.Size(), 1, "sizes should be equal")
		assert.ObjectsAreEqual(got, list.head)
		assert.ObjectsAreEqual(got, list.tail)
	})

	t.Run("it should update the head properly", func(t *testing.T) {
		got := list.AddAtHead(0)

		assert.Equal(t, list.Size(), 2, "sizes should be equal")
		assert.ObjectsAreEqual(got, list.head)
		assert.NotEqual(t, got, list.tail)
	})
}

func TestAddAtTail(t *testing.T) {
	t.Run("it should make 2 the new tail", func(t *testing.T) {
		got := list.AddAtTail(2)

		assert.Equal(t, list.Size(), 3, "sizes should be equal")
		assert.ObjectsAreEqual(got, list.tail)
	})

	t.Run("it should make 3 the new tail", func(t *testing.T) {
		got := list.AddAtTail(3)

		assert.Equal(t, list.Size(), 4, "sizes should be equal")
		assert.ObjectsAreEqual(got, list.tail)
	})
}

func TestAddAt(t *testing.T) {
	t.Run("it should insert 1 at index 1, between 0 and 1", func(t *testing.T) {
		got := list.AddAt(1, 1)

		assert.Equal(t, list.Size(), 5, "sizes should be equal")
		assert.ObjectsAreEqual(got, list.head.next)
	})

	t.Run("it should use the AddAtHead method", func(t *testing.T) {
		got := list.AddAt(0, 2)

		assert.Equal(t, list.Size(), 6, "sizes should be equal")
		assert.ObjectsAreEqual(got, list.head)
	})

	t.Run("it should use the AddAtTail method", func(t *testing.T) {
		got := list.AddAt(list.size-1, 2)

		assert.Equal(t, list.Size(), 7, "sizes should be equal")
		assert.ObjectsAreEqual(got, list.tail)
	})
}

func TestRemoveAtHead(t *testing.T) {
	t.Run("it should update the head to the removed node's next", func(t *testing.T) {
		next := list.head.next
		list.RemoveAtHead()

		assert.Equal(t, list.Size(), 6, "sizes should be equal")
		assert.ObjectsAreEqual(next, list.head)
	})
}

func TestRemoveAtTail(t *testing.T) {
	t.Run("it should update the tail to the removed node's prev", func(t *testing.T) {
        prev := list.SearchAt(list.size - 2)
		list.RemoveAtTail()

		assert.Equal(t, list.Size(), 5, "sizes should be equal")
		assert.ObjectsAreEqual(prev, list.tail)
	})
}

func TestRemoveAt(t *testing.T) {
	t.Run("it should correctly remove an element in the middle of the list", func(t *testing.T) {
        got := list.RemoveAt(3)

		assert.Equal(t, list.Size(), 4, "sizes should be equal")
        assert.Equal(t, 2, got.Data()) 
	})

	t.Run("it should use the RemoveAtHead method", func(t *testing.T) {
        got := list.RemoveAt(0)

		assert.Equal(t, list.Size(), 3, "sizes should be equal")
        assert.ObjectsAreEqual(got, list.head) 
	})

	t.Run("it should use the RemoveAtTail method", func(t *testing.T) {
        got := list.RemoveAt(list.size - 1)

		assert.Equal(t, list.Size(), 2, "sizes should be equal")
        assert.ObjectsAreEqual(got, list.tail) 
	})
}

func TestPeekHead(t *testing.T) {
    got := list.PeekHead()

    assert.ObjectsAreEqual(got, list.head)
}

func TestPeekTail(t *testing.T) {
    got := list.PeekTail()

    assert.ObjectsAreEqual(got, list.tail)
}

