package datastruct_test

import (
	"testing"

	"github.com/edfoh/data-structures/pkg/datastruct"
	"github.com/stretchr/testify/assert"
)

func TestNewSingleLinkList(t *testing.T) {
	sll := datastruct.NewSingleLinkedList[int, string]()

	assert.Empty(t, sll.AllKeys())
	assert.Empty(t, sll.AllItems())
}

func TestSingleLinkedList_InsertHead(t *testing.T) {
	sll := datastruct.NewSingleLinkedList[int, string]()

	t.Run("insert head works", func(t *testing.T) {
		sll.InsertHead(1, "1")

		assert.Equal(t, []int{1}, sll.AllKeys())
		assert.Equal(t, []string{"1"}, sll.AllItems())
	})

	t.Run("insert head again works", func(t *testing.T) {
		sll.InsertHead(2, "2")

		assert.Equal(t, []int{2, 1}, sll.AllKeys())
		assert.Equal(t, []string{"2", "1"}, sll.AllItems())
	})

	t.Run("insert head 3rd time works", func(t *testing.T) {
		sll.InsertHead(3, "3")

		assert.Equal(t, []int{3, 2, 1}, sll.AllKeys())
		assert.Equal(t, []string{"3", "2", "1"}, sll.AllItems())
	})
}

func TestSingleLinkedList_InsertAfter(t *testing.T) {
	sll := datastruct.NewSingleLinkedList[int, string]()
	sll.InsertHead(1, "1")
	sll.InsertHead(2, "2")
	sll.InsertHead(3, "3")

	t.Run("insert after an existing key works", func(t *testing.T) {
		success := sll.InsertAfter(2, 4, "4")

		assert.True(t, success)
		assert.Equal(t, []int{3, 2, 4, 1}, sll.AllKeys())
		assert.Equal(t, []string{"3", "2", "4", "1"}, sll.AllItems())
	})

	t.Run("insert after a non existing key does nothing", func(t *testing.T) {
		success := sll.InsertAfter(7, 5, "5")

		assert.False(t, success)
		assert.Equal(t, []int{3, 2, 4, 1}, sll.AllKeys())
		assert.Equal(t, []string{"3", "2", "4", "1"}, sll.AllItems())
	})

	t.Run("insert after last key works", func(t *testing.T) {
		success := sll.InsertAfter(1, 5, "5")

		assert.True(t, success)
		assert.Equal(t, []int{3, 2, 4, 1, 5}, sll.AllKeys())
		assert.Equal(t, []string{"3", "2", "4", "1", "5"}, sll.AllItems())
	})
}

func TestSingleLinkedList_Delete(t *testing.T) {
	sll := datastruct.NewSingleLinkedList[int, string]()
	sll.InsertHead(1, "1")
	sll.InsertHead(2, "2")
	sll.InsertHead(3, "3")

	t.Run("delete an existing key works", func(t *testing.T) {
		success := sll.Delete(1)

		assert.True(t, success)
		assert.Equal(t, []int{3, 2}, sll.AllKeys())
		assert.Equal(t, []string{"3", "2"}, sll.AllItems())
	})

	t.Run("delete a non existing key does nothing", func(t *testing.T) {
		success := sll.Delete(5)

		assert.False(t, success)
		assert.Equal(t, []int{3, 2}, sll.AllKeys())
		assert.Equal(t, []string{"3", "2"}, sll.AllItems())
	})
}

func TestSingleLinkedList_Search(t *testing.T) {
	sll := datastruct.NewSingleLinkedList[int, string]()
	sll.InsertHead(1, "1")
	sll.InsertHead(2, "2")
	sll.InsertHead(3, "3")

	t.Run("search for an existing key works", func(t *testing.T) {
		ok, val := sll.Search(2)

		assert.True(t, ok)
		assert.Equal(t, "2", val)
	})

	t.Run("search for a non existing key does nothing", func(t *testing.T) {
		ok, val := sll.Search(5)

		assert.False(t, ok)
		assert.Equal(t, "", val)
	})
}
