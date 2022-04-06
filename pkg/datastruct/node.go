package datastruct

type SinglyNode[K comparable, V any] struct {
	Key  K
	Item V
	Next *SinglyNode[K, V]
}

func newSinglyNode[K comparable, V any](key K, item V, next *SinglyNode[K, V]) *SinglyNode[K, V] {
	return &SinglyNode[K, V]{Key: key, Item: item, Next: next}
}
