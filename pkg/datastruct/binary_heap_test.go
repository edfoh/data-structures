package datastruct

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testHeapItem[P Priority] struct {
	val      string
	priority P
}

func (i *testHeapItem[P]) Priority() P {
	return i.priority
}

func (i *testHeapItem[P]) String() string {
	return i.val
}

func TestHeapMin_Insert(t *testing.T) {
	testCases := []struct {
		desc      string
		capacity  int
		items     []int
		wantPrint string
	}{
		{
			desc:      "one item",
			capacity:  10,
			items:     []int{1},
			wantPrint: "1",
		},
		{
			desc:      "2 items unsorted",
			capacity:  10,
			items:     []int{30, 20},
			wantPrint: "20,30",
		},
		{
			desc:      "3 items unsorted",
			capacity:  10,
			items:     []int{30, 20, 50},
			wantPrint: "20,30,50",
		},
		{
			desc:      "4 items unsorted",
			capacity:  10,
			items:     []int{30, 20, 50, 70},
			wantPrint: "20,30,50,70",
		},
		{
			desc:      "5 items unsorted",
			capacity:  10,
			items:     []int{30, 20, 50, 70, 10},
			wantPrint: "10,20,50,70,30",
		},
		{
			desc:      "6 items unsorted",
			capacity:  10,
			items:     []int{30, 20, 50, 70, 10, 40},
			wantPrint: "10,20,40,70,30,50",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			heap := NewHeap[*testHeapItem[int], int](tC.capacity, HeapMin)

			for _, item := range tC.items {
				item := &testHeapItem[int]{
					val:      strconv.Itoa(item),
					priority: item,
				}
				err := heap.Insert(item)
				require.NoError(t, err)
			}

			gotPrint := heap.Print()
			assert.Equal(t, tC.wantPrint, gotPrint)
		})
	}
}

func TestHeapMax_Insert(t *testing.T) {
	testCases := []struct {
		desc      string
		capacity  int
		items     []int
		wantPrint string
	}{
		{
			desc:      "one item",
			capacity:  10,
			items:     []int{1},
			wantPrint: "1",
		},
		{
			desc:      "2 items unsorted",
			capacity:  10,
			items:     []int{20, 30},
			wantPrint: "30,20",
		},
		{
			desc:      "3 items unsorted",
			capacity:  10,
			items:     []int{20, 30, 50},
			wantPrint: "50,20,30",
		},
		{
			desc:      "4 items unsorted",
			capacity:  10,
			items:     []int{20, 30, 50, 70},
			wantPrint: "70,50,30,20",
		},
		{
			desc:      "5 items unsorted",
			capacity:  10,
			items:     []int{20, 30, 50, 70, 10},
			wantPrint: "70,50,30,20,10",
		},
		{
			desc:      "6 items unsorted",
			capacity:  10,
			items:     []int{20, 30, 50, 70, 10, 40},
			wantPrint: "70,50,40,20,10,30",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			heap := NewHeap[*testHeapItem[int], int](tC.capacity, HeapMax)

			for _, item := range tC.items {
				item := &testHeapItem[int]{
					val:      strconv.Itoa(item),
					priority: item,
				}
				err := heap.Insert(item)
				require.NoError(t, err)
			}

			gotPrint := heap.Print()
			assert.Equal(t, tC.wantPrint, gotPrint)
		})
	}
}

func TestHeapMin_Insert_Capacity(t *testing.T) {
	t.Run("when capacity is zero", func(t *testing.T) {
		heap := NewHeap[*testHeapItem[int], int](0, HeapMin)

		item := &testHeapItem[int]{
			val:      "1",
			priority: 1,
		}

		err := heap.Insert(item)
		assert.Error(t, err)
	})

	t.Run("when capacity is non zero and over inserts", func(t *testing.T) {
		heap := NewHeap[*testHeapItem[int], int](1, HeapMin)

		item1 := &testHeapItem[int]{
			val:      "1",
			priority: 1,
		}

		err := heap.Insert(item1)
		require.NoError(t, err)

		item2 := &testHeapItem[int]{
			val:      "2",
			priority: 2,
		}

		err = heap.Insert(item2)
		assert.Error(t, err)
	})
}

func TestHeapMax_Insert_Capacity(t *testing.T) {
	t.Run("when capacity is zero", func(t *testing.T) {
		heap := NewHeap[*testHeapItem[int], int](0, HeapMax)

		item := &testHeapItem[int]{
			val:      "1",
			priority: 1,
		}

		err := heap.Insert(item)
		assert.Error(t, err)
	})

	t.Run("when capacity is non zero and over inserts", func(t *testing.T) {
		heap := NewHeap[*testHeapItem[int], int](1, HeapMax)

		item1 := &testHeapItem[int]{
			val:      "1",
			priority: 1,
		}

		err := heap.Insert(item1)
		require.NoError(t, err)

		item2 := &testHeapItem[int]{
			val:      "2",
			priority: 2,
		}

		err = heap.Insert(item2)
		assert.Error(t, err)
	})
}

func TestHeapMin_TopValue(t *testing.T) {
	item := &testHeapItem[int]{
		val:      "1",
		priority: 1,
	}
	type args struct {
		item *testHeapItem[int]
	}

	testCases := []struct {
		desc    string
		args    *args
		wantVal *testHeapItem[int]
		wantErr error
	}{
		{
			desc:    "when heap is empty",
			wantVal: nil,
			wantErr: errors.New("heap is empty"),
		},
		{
			desc: "when heap is not empty",
			args: &args{
				item: item,
			},
			wantVal: item,
			wantErr: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			heap := NewHeap[*testHeapItem[int], int](10, HeapMin)

			if tC.args != nil {
				heap.Insert(tC.args.item)
			}
			gotVal, gotErr := heap.TopValue()
			assert.Equal(t, tC.wantVal, gotVal)
			assert.Equal(t, tC.wantErr, gotErr)

		})
	}
}

func TestHeapMax_TopValue(t *testing.T) {
	item := &testHeapItem[int]{
		val:      "1",
		priority: 1,
	}
	type args struct {
		item *testHeapItem[int]
	}

	testCases := []struct {
		desc    string
		args    *args
		wantVal *testHeapItem[int]
		wantErr error
	}{
		{
			desc:    "when heap is empty",
			wantVal: nil,
			wantErr: errors.New("heap is empty"),
		},
		{
			desc: "when heap is not empty",
			args: &args{
				item: item,
			},
			wantVal: item,
			wantErr: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			heap := NewHeap[*testHeapItem[int], int](10, HeapMax)

			if tC.args != nil {
				heap.Insert(tC.args.item)
			}
			gotVal, gotErr := heap.TopValue()
			assert.Equal(t, tC.wantVal, gotVal)
			assert.Equal(t, tC.wantErr, gotErr)

		})
	}
}

func TestHeapMin_Pop(t *testing.T) {
	setup := func(t *testing.T) *Heap[*testHeapItem[int], int] {
		heap := NewHeap[*testHeapItem[int], int](10, HeapMin)

		for _, item := range []int{30, 20, 50, 70, 10, 40} {
			item := &testHeapItem[int]{
				val:      strconv.Itoa(item),
				priority: item,
			}
			err := heap.Insert(item)
			require.NoError(t, err)
		}
		return heap
	}

	testCases := []struct {
		desc              string
		wantPopPriorities []int
		wantPopErrs       []error
		wantPrint         string
	}{
		{
			desc:              "pop once",
			wantPopPriorities: []int{10},
			wantPopErrs:       []error{nil},
			wantPrint:         "20,40,70,30,50",
		},
		{
			desc:              "pop twice",
			wantPopPriorities: []int{10, 20},
			wantPopErrs:       []error{nil, nil},
			wantPrint:         "30,70,40,50",
		},
		{
			desc:              "pop thrice",
			wantPopPriorities: []int{10, 20, 30},
			wantPopErrs:       []error{nil, nil, nil},
			wantPrint:         "40,70,50",
		},
		{
			desc:              "pop 4 times",
			wantPopPriorities: []int{10, 20, 30, 40},
			wantPopErrs:       []error{nil, nil, nil, nil},
			wantPrint:         "50,70",
		},
		{
			desc:              "pop 5 times",
			wantPopPriorities: []int{10, 20, 30, 40, 50},
			wantPopErrs:       []error{nil, nil, nil, nil, nil},
			wantPrint:         "70",
		},
		{
			desc:              "pop 6 times",
			wantPopPriorities: []int{10, 20, 30, 40, 50, 70},
			wantPopErrs:       []error{nil, nil, nil, nil, nil, nil},
			wantPrint:         "",
		},
		{
			desc:              "pop 7 times",
			wantPopPriorities: []int{10, 20, 30, 40, 50, 70},
			wantPopErrs:       []error{nil, nil, nil, nil, nil, nil, errors.New("Heap is empty")},
			wantPrint:         "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			heap := setup(t)

			for index, wantPopPriority := range tC.wantPopPriorities {
				gotItem, gotErr := heap.Pop()
				require.Equal(t, tC.wantPopErrs[index], gotErr)
				assert.Equal(t, wantPopPriority, gotItem.priority)
			}

			gotPrint := heap.Print()
			assert.Equal(t, tC.wantPrint, gotPrint)
		})
	}
}

func TestHeapMax_Pop(t *testing.T) {
	setup := func(t *testing.T) *Heap[*testHeapItem[int], int] {
		heap := NewHeap[*testHeapItem[int], int](10, HeapMax)

		for _, item := range []int{20, 30, 50, 70, 10, 40} {
			item := &testHeapItem[int]{
				val:      strconv.Itoa(item),
				priority: item,
			}
			err := heap.Insert(item)
			require.NoError(t, err)
		}
		return heap
	}

	testCases := []struct {
		desc              string
		wantPopPriorities []int
		wantPopErrs       []error
		wantPrint         string
	}{
		{
			desc:              "pop once",
			wantPopPriorities: []int{70},
			wantPopErrs:       []error{nil},
			wantPrint:         "50,40,20,10,30",
		},
		{
			desc:              "pop twice",
			wantPopPriorities: []int{70, 50},
			wantPopErrs:       []error{nil, nil},
			wantPrint:         "40,20,10,30",
		},
		{
			desc:              "pop thrice",
			wantPopPriorities: []int{70, 50, 40},
			wantPopErrs:       []error{nil, nil, nil},
			wantPrint:         "30,10,20",
		},
		{
			desc:              "pop 4 times",
			wantPopPriorities: []int{70, 50, 40, 30},
			wantPopErrs:       []error{nil, nil, nil, nil},
			wantPrint:         "20,10",
		},
		{
			desc:              "pop 5 times",
			wantPopPriorities: []int{70, 50, 40, 30, 20},
			wantPopErrs:       []error{nil, nil, nil, nil, nil},
			wantPrint:         "10",
		},
		{
			desc:              "pop 6 times",
			wantPopPriorities: []int{70, 50, 40, 30, 20, 10},
			wantPopErrs:       []error{nil, nil, nil, nil, nil, nil},
			wantPrint:         "",
		},
		{
			desc:              "pop 7 times",
			wantPopPriorities: []int{70, 50, 40, 30, 20, 10},
			wantPopErrs:       []error{nil, nil, nil, nil, nil, nil, errors.New("Heap is empty")},
			wantPrint:         "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			heap := setup(t)

			for index, wantPopPriority := range tC.wantPopPriorities {
				gotItem, gotErr := heap.Pop()
				require.Equal(t, tC.wantPopErrs[index], gotErr)
				assert.Equal(t, wantPopPriority, gotItem.priority)
			}

			gotPrint := heap.Print()
			assert.Equal(t, tC.wantPrint, gotPrint)
		})
	}
}
