package datastruct

type singlyNode[K comparable, V any] struct {
	Key  K
	Item V
	Next *singlyNode[K, V]
}

func newSinglyNode[K comparable, V any](key K, item V, next *singlyNode[K, V]) *singlyNode[K, V] {
	return &singlyNode[K, V]{Key: key, Item: item, Next: next}
}

func (s *singlyNode[K, V]) AddNodeAfter(key K, item V) *singlyNode[K, V] {
	n := newSinglyNode(key, item, s.Next)
	s.Next = n
	return n
}

func (s *singlyNode[K, V]) AddNodeBefore(key K, item V) *singlyNode[K, V] {
	n := newSinglyNode(key, item, s)
	return n
}

func (s *singlyNode[K, V]) IsTail() bool {
	return s.Next == nil
}

type doublyNode[K comparable, V any] struct {
	Key      K
	Item     V
	Next     *doublyNode[K, V]
	Previous *doublyNode[K, V]
}

func newDoublyNode[K comparable, V any](key K, item V, prev *doublyNode[K, V], next *doublyNode[K, V]) *doublyNode[K, V] {
	return &doublyNode[K, V]{Key: key, Item: item, Previous: prev, Next: next}
}

func (d *doublyNode[K, V]) IsTail() bool {
	return d.Next == nil
}

func (d *doublyNode[K, V]) IsHead() bool {
	return d.Previous == nil
}

func (d *doublyNode[K, V]) AddNodeAfter(key K, item V) *doublyNode[K, V] {
	n := newDoublyNode(key, item, d, d.Next)
	if d.Next != nil {
		d.Next.Previous = n
	}
	d.Next = n
	return n
}

func (d *doublyNode[K, V]) AddNodeBefore(key K, item V) *doublyNode[K, V] {
	n := newDoublyNode(key, item, d.Previous, d)
	if d.Previous != nil {
		d.Previous.Next = n
	}
	d.Previous = n
	return n
}
