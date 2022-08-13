package datastruct

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	testCases := []struct {
		desc     string
		capacity int
		nodes    []int
		wantRes  []int
		wantErrs []error
	}{
		{
			desc:     "insert 1",
			capacity: 10,
			nodes:    []int{50},
			wantRes:  []int{50},
			wantErrs: []error{nil},
		},
		{
			desc:     "insert 2",
			capacity: 10,
			nodes:    []int{50, 30},
			wantRes:  []int{30, 50},
			wantErrs: []error{nil, nil},
		},
		{
			desc:     "insert 3",
			capacity: 10,
			nodes:    []int{50, 30, 60},
			wantRes:  []int{30, 50, 60},
			wantErrs: []error{nil, nil, nil},
		},
		{
			desc:     "insert 5",
			capacity: 10,
			nodes:    []int{50, 30, 60, 20, 40},
			wantRes:  []int{20, 30, 40, 50, 60},
			wantErrs: []error{nil, nil, nil, nil, nil},
		},
		{
			desc:     "insert 7",
			capacity: 10,
			nodes:    []int{50, 30, 60, 20, 40, 55, 70},
			wantRes:  []int{20, 30, 40, 50, 55, 60, 70},
			wantErrs: []error{nil, nil, nil, nil, nil, nil, nil},
		},
		{
			desc:     "over capacity",
			capacity: 1,
			nodes:    []int{10, 20},
			wantRes:  []int{10},
			wantErrs: []error{nil, errors.New("tree is full")},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tree := NewBinarySearchTree[int](tC.capacity)

			for i, node := range tC.nodes {
				gotErr := tree.Insert(node)
				require.Equal(t, tC.wantErrs[i], gotErr)
			}

			gotRes := tree.InorderResults()
			assert.Equal(t, tC.wantRes, gotRes)
		})
	}
}

func TestBinarySearchTree_Search(t *testing.T) {
	items := []int{50, 30, 60, 20, 40, 55, 70}
	setup := func(t *testing.T) *BinarySearchTree[int] {
		tree := NewBinarySearchTree[int](10)

		for _, node := range items {
			err := tree.Insert(node)
			require.NoError(t, err)
		}
		return tree
	}

	t.Run("searching for existing items", func(t *testing.T) {
		tree := setup(t)

		for _, item := range items {
			gotRes := tree.Search(item)
			assert.Equal(t, item, gotRes.Data())
		}
	})

	t.Run("searching for non existent item", func(t *testing.T) {
		tree := setup(t)

		gotRes := tree.Search(100)
		assert.Nil(t, gotRes)
	})
}

func TestBinarySearchTree_Delete(t *testing.T) {
	setup := func(t *testing.T) *BinarySearchTree[int] {
		tree := NewBinarySearchTree[int](10)

		for _, node := range []int{50, 30, 60, 20, 40, 55, 70} {
			//[]int{20, 30, 40, 50, 55, 60, 70}
			err := tree.Insert(node)
			require.NoError(t, err)
		}
		return tree
	}
	testCases := []struct {
		desc        string
		deleteItems []int
		wantRes     []int
	}{
		{
			desc:        "delete leaf - 20",
			deleteItems: []int{20},
			wantRes:     []int{30, 40, 50, 55, 60, 70},
		},
		{
			desc:        "delete leaf - 40",
			deleteItems: []int{40},
			wantRes:     []int{20, 30, 50, 55, 60, 70},
		},
		{
			desc:        "delete leaf - 55",
			deleteItems: []int{55},
			wantRes:     []int{20, 30, 40, 50, 60, 70},
		},
		{
			desc:        "delete leaf - 70",
			deleteItems: []int{70},
			wantRes:     []int{20, 30, 40, 50, 55, 60},
		},
		{
			desc:        "delete 1st level on left - 30",
			deleteItems: []int{30},
			wantRes:     []int{20, 40, 50, 55, 60, 70},
		},
		{
			desc:        "delete 1st level on left with right child - 30",
			deleteItems: []int{20, 30},
			wantRes:     []int{40, 50, 55, 60, 70},
		},
		{
			desc:        "delete 1st level on left with right child - 30",
			deleteItems: []int{40, 30},
			wantRes:     []int{20, 50, 55, 60, 70},
		},
		{
			desc:        "delete 1st level on right - 60",
			deleteItems: []int{60},
			wantRes:     []int{20, 30, 40, 50, 55, 70},
		},
		{
			desc:        "delete 1st level on right with right child - 60",
			deleteItems: []int{55, 60},
			wantRes:     []int{20, 30, 40, 50, 70},
		},
		{
			desc:        "delete 1st level on right with left child - 60",
			deleteItems: []int{70, 60},
			wantRes:     []int{20, 30, 40, 50, 55},
		},
		{
			desc:        "delete root",
			deleteItems: []int{50},
			wantRes:     []int{20, 30, 40, 55, 60, 70},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tree := setup(t)

			for _, delItem := range tC.deleteItems {
				tree.Delete(delItem)
			}

			gotRes := tree.InorderResults()
			assert.Equal(t, tC.wantRes, gotRes)
		})
	}
}
