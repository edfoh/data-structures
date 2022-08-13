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

func (h *Heap[K, P]) heapify(index int) {
	left := h.left(index)
	right := h.right(index)

	smallest := index
	if left < h.Size() && h.compare(left, smallest) {
		smallest = left
	}

	if right < h.Size() && h.compare(right, smallest) {
		smallest = right
	}

	if smallest != index {
		h.swap(index, smallest)
		h.heapify(smallest)
	}
}
