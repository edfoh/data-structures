package datastruct_test

import (
	"testing"

	"github.com/edfoh/data-structures/pkg/datastruct"
	"github.com/stretchr/testify/assert"
)

func TestNewDoubleLinkedList(t *testing.T) {
	dll := datastruct.NewDoubleLinkedList[int, string]()

	assert.Empty(t, dll.AllKeys())
	assert.Empty(t, dll.AllKeysReverse())
	assert.Empty(t, dll.AllItems())
	assert.Empty(t, dll.AllItemsReverse())
}

func TestDoubleLinkedList_InsertHead(t *testing.T) {
	dll := datastruct.NewDoubleLinkedList[int, string]()

	t.Run("insert head works", func(t *testing.T) {
		dll.InsertHead(1, "1")

		assert.Equal(t, []int{1}, dll.AllKeys())
		assert.Equal(t, []int{1}, dll.AllKeysReverse())
		assert.Equal(t, []string{"1"}, dll.AllItems())
		assert.Equal(t, []string{"1"}, dll.AllItemsReverse())
	})

	t.Run("insert head again works", func(t *testing.T) {
		dll.InsertHead(2, "2")

		assert.Equal(t, []int{2, 1}, dll.AllKeys())
		assert.Equal(t, []int{1, 2}, dll.AllKeysReverse())
		assert.Equal(t, []string{"2", "1"}, dll.AllItems())
		assert.Equal(t, []string{"1", "2"}, dll.AllItemsReverse())
	})

	t.Run("insert head third time works", func(t *testing.T) {
		dll.InsertHead(3, "3")

		assert.Equal(t, []int{3, 2, 1}, dll.AllKeys())
		assert.Equal(t, []int{1, 2, 3}, dll.AllKeysReverse())
		assert.Equal(t, []string{"3", "2", "1"}, dll.AllItems())
		assert.Equal(t, []string{"1", "2", "3"}, dll.AllItemsReverse())
	})
}

func TestDoubleLinkedList_InsertTail(t *testing.T) {
	dll := datastruct.NewDoubleLinkedList[int, string]()

	t.Run("insert tail works", func(t *testing.T) {
		dll.InsertTail(1, "1")

		assert.Equal(t, []int{1}, dll.AllKeys())
		assert.Equal(t, []int{1}, dll.AllKeysReverse())
		assert.Equal(t, []string{"1"}, dll.AllItems())
		assert.Equal(t, []string{"1"}, dll.AllItemsReverse())
	})

	t.Run("insert tail again works", func(t *testing.T) {
		dll.InsertTail(2, "2")

		assert.Equal(t, []int{1, 2}, dll.AllKeys())
		assert.Equal(t, []int{2, 1}, dll.AllKeysReverse())
		assert.Equal(t, []string{"1", "2"}, dll.AllItems())
		assert.Equal(t, []string{"2", "1"}, dll.AllItemsReverse())
	})

	t.Run("insert tail third time works", func(t *testing.T) {
		dll.InsertTail(3, "3")

		assert.Equal(t, []int{1, 2, 3}, dll.AllKeys())
		assert.Equal(t, []int{3, 2, 1}, dll.AllKeysReverse())
		assert.Equal(t, []string{"1", "2", "3"}, dll.AllItems())
		assert.Equal(t, []string{"3", "2", "1"}, dll.AllItemsReverse())
	})
}

func TestDoubleLinkedList_InsertAfter(t *testing.T) {
	dll := datastruct.NewDoubleLinkedList[int, string]()
	dll.InsertHead(5, "5")
	dll.InsertHead(3, "3")
	dll.InsertHead(1, "1")

	t.Run("insert after an existing key works", func(t *testing.T) {
		success := dll.InsertAfter(3, 4, "4")

		assert.True(t, success)
		assert.Equal(t, []int{1, 3, 4, 5}, dll.AllKeys())
		assert.Equal(t, []int{5, 4, 3, 1}, dll.AllKeysReverse())
		assert.Equal(t, []string{"1", "3", "4", "5"}, dll.AllItems())
		assert.Equal(t, []string{"5", "4", "3", "1"}, dll.AllItemsReverse())
	})

	t.Run("insert after a non existing key does nothing", func(t *testing.T) {
		success := dll.InsertAfter(7, 5, "5")

		assert.False(t, success)
		assert.Equal(t, []int{1, 3, 4, 5}, dll.AllKeys())
		assert.Equal(t, []int{5, 4, 3, 1}, dll.AllKeysReverse())
		assert.Equal(t, []string{"1", "3", "4", "5"}, dll.AllItems())
		assert.Equal(t, []string{"5", "4", "3", "1"}, dll.AllItemsReverse())
	})

	t.Run("insert after last key works", func(t *testing.T) {
		success := dll.InsertAfter(5, 6, "6")

		assert.True(t, success)
		assert.Equal(t, []int{1, 3, 4, 5, 6}, dll.AllKeys())
		assert.Equal(t, []int{6, 5, 4, 3, 1}, dll.AllKeysReverse())
		assert.Equal(t, []string{"1", "3", "4", "5", "6"}, dll.AllItems())
		assert.Equal(t, []string{"6", "5", "4", "3", "1"}, dll.AllItemsReverse())
	})
}

func TestDoubleLinkedList_Delete(t *testing.T) {
	createListFunc := func() *datastruct.DoubleLinkedList[int, string] {
		dll := datastruct.NewDoubleLinkedList[int, string]()
		dll.InsertHead(3, "3")
		dll.InsertHead(2, "2")
		dll.InsertHead(1, "1")
		return dll
	}

	t.Run("delete head works", func(t *testing.T) {
		dll := createListFunc()
		success := dll.Delete(1)

		assert.True(t, success)
		assert.Equal(t, []int{2, 3}, dll.AllKeys())
		assert.Equal(t, []int{3, 2}, dll.AllKeysReverse())
		assert.Equal(t, []string{"2", "3"}, dll.AllItems())
		assert.Equal(t, []string{"3", "2"}, dll.AllItemsReverse())
	})

	t.Run("delete middle works", func(t *testing.T) {
		dll := createListFunc()
		success := dll.Delete(2)

		assert.True(t, success)
		assert.Equal(t, []int{1, 3}, dll.AllKeys())
		assert.Equal(t, []int{3, 1}, dll.AllKeysReverse())
		assert.Equal(t, []string{"1", "3"}, dll.AllItems())
		assert.Equal(t, []string{"3", "1"}, dll.AllItemsReverse())
	})

	t.Run("delete tail works", func(t *testing.T) {
		dll := createListFunc()
		success := dll.Delete(3)

		assert.True(t, success)
		assert.Equal(t, []int{1, 2}, dll.AllKeys())
		assert.Equal(t, []int{2, 1}, dll.AllKeysReverse())
		assert.Equal(t, []string{"1", "2"}, dll.AllItems())
		assert.Equal(t, []string{"2", "1"}, dll.AllItemsReverse())
	})

	t.Run("delete non existent key works", func(t *testing.T) {
		dll := createListFunc()
		success := dll.Delete(4)

		assert.False(t, success)
		assert.Equal(t, []int{1, 2, 3}, dll.AllKeys())
		assert.Equal(t, []int{3, 2, 1}, dll.AllKeysReverse())
		assert.Equal(t, []string{"1", "2", "3"}, dll.AllItems())
		assert.Equal(t, []string{"3", "2", "1"}, dll.AllItemsReverse())
	})
}

func TestDoubleLinkedList_DeleteHead(t *testing.T) {
	dll := datastruct.NewDoubleLinkedList[int, string]()
	dll.InsertHead(3, "3")
	dll.InsertHead(2, "2")
	dll.InsertHead(1, "1")

	t.Run("delete head with 3 items works", func(t *testing.T) {
		success := dll.DeleteHead()

		assert.True(t, success)
		assert.Equal(t, []int{2, 3}, dll.AllKeys())
		assert.Equal(t, []int{3, 2}, dll.AllKeysReverse())
		assert.Equal(t, []string{"2", "3"}, dll.AllItems())
		assert.Equal(t, []string{"3", "2"}, dll.AllItemsReverse())
	})

	t.Run("delete head with 2 items works", func(t *testing.T) {
		success := dll.DeleteHead()

		assert.True(t, success)
		assert.Equal(t, []int{3}, dll.AllKeys())
		assert.Equal(t, []int{3}, dll.AllKeysReverse())
		assert.Equal(t, []string{"3"}, dll.AllItems())
		assert.Equal(t, []string{"3"}, dll.AllItemsReverse())
	})

	t.Run("delete head with 1 item works", func(t *testing.T) {
		success := dll.DeleteHead()

		assert.True(t, success)
		assert.Empty(t, dll.AllKeys())
		assert.Empty(t, dll.AllKeysReverse())
		assert.Empty(t, dll.AllItems())
		assert.Empty(t, dll.AllItemsReverse())
	})

	t.Run("delete empty works", func(t *testing.T) {
		success := dll.DeleteHead()

		assert.False(t, success)
		assert.Empty(t, dll.AllKeys())
		assert.Empty(t, dll.AllKeysReverse())
		assert.Empty(t, dll.AllItems())
		assert.Empty(t, dll.AllItemsReverse())
	})
}

func TestDoubleLinkedList_DeleteTail(t *testing.T) {
	dll := datastruct.NewDoubleLinkedList[int, string]()
	dll.InsertHead(3, "3")
	dll.InsertHead(2, "2")
	dll.InsertHead(1, "1")

	t.Run("delete tail with 3 items works", func(t *testing.T) {
		success := dll.DeleteTail()

		assert.True(t, success)
		assert.Equal(t, []int{1, 2}, dll.AllKeys())
		assert.Equal(t, []int{2, 1}, dll.AllKeysReverse())
		assert.Equal(t, []string{"1", "2"}, dll.AllItems())
		assert.Equal(t, []string{"2", "1"}, dll.AllItemsReverse())
	})

	t.Run("delete tail with 2 items works", func(t *testing.T) {
		success := dll.DeleteTail()

		assert.True(t, success)
		assert.Equal(t, []int{1}, dll.AllKeys())
		assert.Equal(t, []int{1}, dll.AllKeysReverse())
		assert.Equal(t, []string{"1"}, dll.AllItems())
		assert.Equal(t, []string{"1"}, dll.AllItemsReverse())
	})

	t.Run("delete tail with 1 item works", func(t *testing.T) {
		success := dll.DeleteTail()

		assert.True(t, success)
		assert.Empty(t, dll.AllKeys())
		assert.Empty(t, dll.AllKeysReverse())
		assert.Empty(t, dll.AllItems())
		assert.Empty(t, dll.AllItemsReverse())
	})

	t.Run("delete empty works", func(t *testing.T) {
		success := dll.DeleteTail()

		assert.False(t, success)
		assert.Empty(t, dll.AllKeys())
		assert.Empty(t, dll.AllKeysReverse())
		assert.Empty(t, dll.AllItems())
		assert.Empty(t, dll.AllItemsReverse())
	})
}
