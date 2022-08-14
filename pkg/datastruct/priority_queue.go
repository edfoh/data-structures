package datastruct

type priorityQueueItem struct {
	priority int
	payload  string
}

func (item *priorityQueueItem) Priority() int {
	return item.priority
}

func (item *priorityQueueItem) String() string {
	return item.payload
}

type PriorityQueue struct {
	heap *Heap[*priorityQueueItem, int]
}

func NewPriorityQueue(capacity int) *PriorityQueue {
	return &PriorityQueue{
		heap: NewHeap[*priorityQueueItem, int](capacity, HeapMax),
	}
}

func (pq *PriorityQueue) Enqueue(payload string, priority int) {
	item := &priorityQueueItem{
		priority: priority,
		payload:  payload,
	}
	pq.heap.Insert(item)
}

func (pq *PriorityQueue) Dequeue() (string, error) {
	item, err := pq.heap.Pop()
	if err != nil {
		return "", err
	}
	return item.String(), nil
}
