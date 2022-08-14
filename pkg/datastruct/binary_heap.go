package datastruct

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type HeapItem[P Priority] interface {
	String() string
	Priority() P
}

type Priority interface {
	constraints.Ordered
}

type Heap[K HeapItem[P], P Priority] struct {
	capacity  int
	nodes     []K
	heapOrder HeapOrder
}

type HeapOrder int8

const (
	HeapMin HeapOrder = iota
	HeapMax
)

func NewHeap[K HeapItem[P], P Priority](capacity int, heapOrder HeapOrder) *Heap[K, P] {
	return &Heap[K, P]{
		capacity:  capacity,
		nodes:     make([]K, 0, capacity),
		heapOrder: heapOrder,
	}
}

func (h *Heap[K, P]) parent(index int) int {
	return (index - 1) / 2
}

func (h *Heap[K, P]) left(index int) int {
	return (index * 2) + 1
}

func (h *Heap[K, P]) right(index int) int {
	return (index * 2) + 2
}

func (h *Heap[K, P]) swap(i, j int) {
	tmp := h.nodes[i]
	h.nodes[i] = h.nodes[j]
	h.nodes[j] = tmp
}

func (h *Heap[K, P]) compare(i, j int) bool {
	if h.heapOrder == HeapMin {
		return h.nodes[i].Priority() < h.nodes[j].Priority()
	}
	return h.nodes[i].Priority() > h.nodes[j].Priority()
}

func (h *Heap[K, P]) Size() int {
	return len(h.nodes)
}

// Insert will add a new item at the end, and a comparison is made with the newly added node against its
// parent depending on HeapOrder of Min or Max. If the node is smaller or larger than the parent, their values are swapped.
// The comparison is done continuously up the heap against each parent until the comparison no longer holds true.
func (h *Heap[K, P]) Insert(data K) error {
	if len(h.nodes) == h.capacity {
		return errors.New("Heap is full")
	}

	h.nodes = append(h.nodes, data)
	i := len(h.nodes) - 1

	for i != 0 && h.compare(i, h.parent(i)) {
		parentInd := h.parent(i)
		h.swap(i, parentInd)
		i = parentInd
	}

	return nil
}

func (h *Heap[K, P]) TopValue() (K, error) {
	var val K
	if h.Size() == 0 {
		return val, errors.New("heap is empty")
	}

	return h.nodes[0], nil
}

func (h *Heap[K, P]) Pop() (K, error) {
	val, err := h.TopValue()
	if err != nil {
		return val, err
	}

	h.nodes = h.nodes[1:]
	h.heapify(0)

	return val, nil
}

func (h *Heap[K, P]) Print() string {
	var ss []string
	for _, n := range h.nodes {
		ss = append(ss, fmt.Sprintf("%s", n))
	}
	return strings.Join(ss, ",")
}

// heapify will recursively heapify a subtree at the given index. Depending of the HeapOrder,
// Min will compare and swap values with left or right that has the smallest value with the node at the given index,
// and recursively working up the heap. For HeapOrder Max, the comparison works to move the larger value up.
func (h *Heap[K, P]) heapify(index int) {
	left := h.left(index)
	right := h.right(index)

	compared := index
	if left < h.Size() && h.compare(left, compared) {
		compared = left
	}

	if right < h.Size() && h.compare(right, compared) {
		compared = right
	}

	if compared != index {
		h.swap(index, compared)
		h.heapify(compared)
	}
}
