package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyNode_AddNodeAfter(t *testing.T) {
	tail := newSinglyNode(1, "1", nil)

	t.Run("adding after a node that is not a tail", func(t *testing.T) {
		node := newSinglyNode(2, "2", tail)
		newNode := node.AddNodeAfter(3, "3")

		assert.Equal(t, newNode.Key, 3)
		assert.Equal(t, newNode.Item, "3")
		assert.Equal(t, newNode.Next, tail)
		assert.Equal(t, node.Next, newNode)
	})

	t.Run("adding after a node that is a tail", func(t *testing.T) {
		newNode := tail.AddNodeAfter(3, "3")

		assert.Equal(t, newNode.Key, 3)
		assert.Equal(t, newNode.Item, "3")
		assert.Equal(t, tail.Next, newNode)
		assert.True(t, newNode.IsTail())
	})
}

func TestSinglyNode_AddNodeBefore(t *testing.T) {
	head := newSinglyNode(1, "1", nil)

	t.Run("adding before a node that is not a head", func(t *testing.T) {
		node := newSinglyNode(2, "2", nil)
		head.Next = node
		newNode := node.AddNodeBefore(3, "3")

		assert.Equal(t, newNode.Key, 3)
		assert.Equal(t, newNode.Item, "3")
		assert.Equal(t, newNode.Next, node)
	})

	t.Run("adding before a node that is a head", func(t *testing.T) {
		newNode := head.AddNodeBefore(3, "3")

		assert.Equal(t, newNode.Key, 3)
		assert.Equal(t, newNode.Item, "3")
		assert.Equal(t, newNode.Next, head)
	})
}

func TestSinglyNode_IsTail(t *testing.T) {
	tail := newSinglyNode(1, "1", nil)

	t.Run("node with no pointer to next is a tail", func(t *testing.T) {
		assert.True(t, tail.IsTail())
	})

	t.Run("node with a pointer to next is not a tail", func(t *testing.T) {
		node := newSinglyNode(1, "1", tail)

		assert.False(t, node.IsTail())
	})
}

func TestDoublyNode_IsTail(t *testing.T) {
	tail := newDoublyNode(1, "1", nil, nil)

	t.Run("node with no pointer to next is a tail", func(t *testing.T) {
		assert.True(t, tail.IsTail())
	})

	t.Run("node with a pointer to next is not a tail", func(t *testing.T) {
		node := newDoublyNode(1, "1", nil, tail)

		assert.False(t, node.IsTail())
	})
}

func TestDoublyNode_IsHead(t *testing.T) {
	head := newDoublyNode(1, "1", nil, nil)

	t.Run("node with no pointer to prev is a head", func(t *testing.T) {
		assert.True(t, head.IsHead())
	})

	t.Run("node with a pointer to prev is not a head", func(t *testing.T) {
		node := newDoublyNode(1, "1", head, nil)

		assert.False(t, node.IsHead())
	})
}

func TestDoublyNode_AddNodeAfter(t *testing.T) {
	tail := newDoublyNode(1, "1", nil, nil)

	t.Run("adding after a node that is not a tail", func(t *testing.T) {
		node := newDoublyNode(2, "2", nil, tail)
		newNode := node.AddNodeAfter(3, "3")

		assert.Equal(t, newNode.Key, 3)
		assert.Equal(t, newNode.Item, "3")
		assert.Equal(t, newNode.Previous, node)
		assert.Equal(t, newNode.Next, tail)
		assert.Equal(t, node.Next, newNode)
		assert.Equal(t, tail.Previous, newNode)
	})

	t.Run("adding after a node that is a tail", func(t *testing.T) {
		newNode := tail.AddNodeAfter(3, "3")

		assert.Equal(t, newNode.Key, 3)
		assert.Equal(t, newNode.Item, "3")
		assert.Equal(t, newNode.Previous, tail)
		assert.True(t, newNode.IsTail(), nil)
		assert.Equal(t, tail.Next, newNode)
	})
}

func TestDoublyNode_AddNodeBefore(t *testing.T) {
	head := newDoublyNode(1, "1", nil, nil)

	t.Run("adding before a node that is not a head", func(t *testing.T) {
		node := newDoublyNode(2, "2", head, nil)
		newNode := node.AddNodeBefore(3, "3")

		assert.Equal(t, newNode.Key, 3)
		assert.Equal(t, newNode.Item, "3")
		assert.Equal(t, newNode.Next, node)
		assert.Equal(t, newNode.Previous, head)
		assert.Equal(t, node.Previous, newNode)
		assert.Equal(t, head.Next, newNode)
	})

	t.Run("adding before a node that is a head", func(t *testing.T) {
		newNode := head.AddNodeBefore(3, "3")

		assert.Equal(t, newNode.Key, 3)
		assert.Equal(t, newNode.Item, "3")
		assert.Equal(t, newNode.Next, head)
		assert.True(t, newNode.IsHead(), nil)
		assert.Equal(t, head.Previous, newNode)
	})
}
