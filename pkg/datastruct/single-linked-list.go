package datastruct

type SingleLinkedList[K comparable, V any] struct {
	head *singlyNode[K, V]
}

func NewSingleLinkedList[K comparable, V any]() *SingleLinkedList[K, V] {
	return &SingleLinkedList[K, V]{head: nil}
}

func (sll *SingleLinkedList[K, V]) InsertHead(key K, item V) {
	if sll.isEmpty() {
		sll.head = newSinglyNode(key, item, nil)
	} else {
		sll.head = sll.head.AddNodeBefore(key, item)
	}
}

func (sll *SingleLinkedList[K, V]) InsertAfter(keyAfter K, keyInsert K, itemInsert V) bool {
	node := sll.findNode(keyAfter)
	if node != nil {
		node.AddNodeAfter(keyInsert, itemInsert)
		return true
	}
	return false
}

func (sll *SingleLinkedList[K, V]) Delete(key K) bool {
	previous := sll.findNodeBefore(key)
	if previous != nil {
		nodeToDelete := previous.Next
		previous.Next = nodeToDelete.Next
		nodeToDelete.Next = nil
		return true
	}
	return false
}

func (sll *SingleLinkedList[K, V]) Search(key K) (bool, V) {
	var val V
	node := sll.findNode(key)
	if node == nil {
		return false, val
	}
	return true, node.Item
}

func (sll *SingleLinkedList[K, V]) AllKeys() []K {
	var keys []K
	current := sll.head
	for current != nil {
		keys = append(keys, current.Key)
		current = current.Next
	}
	return keys
}

func (sll *SingleLinkedList[K, V]) AllItems() []V {
	var vals []V
	current := sll.head
	for current != nil {
		vals = append(vals, current.Item)
		current = current.Next
	}
	return vals
}

func (sll *SingleLinkedList[K, V]) isEmpty() bool {
	return sll.head == nil
}

func (sll *SingleLinkedList[K, V]) findNode(key K) *singlyNode[K, V] {
	current := sll.head
	for current != nil {
		if current.Key == key {
			return current
		}
		current = current.Next
	}
	return nil
}

func (sll *SingleLinkedList[K, V]) findNodeBefore(key K) *singlyNode[K, V] {
	var previous *singlyNode[K, V]
	current := sll.head
	for current != nil {
		if current.Key == key {
			return previous
		}
		previous = current
		current = current.Next
	}
	return nil
}
