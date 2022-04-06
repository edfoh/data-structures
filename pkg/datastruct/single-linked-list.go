package datastruct

type SingleLinkedList[K comparable, V any] struct {
	head *SinglyNode[K, V]
}

func NewSingleLinkedList[K comparable, V any]() *SingleLinkedList[K, V] {
	return &SingleLinkedList[K, V]{head: nil}
}

func (sll *SingleLinkedList[K, V]) InsertHead(key K, item V) {
	newHead := newSinglyNode(key, item, sll.head)
	sll.head = newHead
}

func (sll *SingleLinkedList[K, V]) InsertAfter(keyAfter K, keyInsert K, itemInsert V) bool {
	node := sll.findNode(keyAfter)
	if node != nil {
		newNode := newSinglyNode(keyInsert, itemInsert, node.Next)
		node.Next = newNode
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
	if node != nil {
		val = node.Item
		return true, val
	}
	return false, val
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

func (sll *SingleLinkedList[K, V]) findNode(key K) *SinglyNode[K, V] {
	current := sll.head
	for current != nil {
		if current.Key == key {
			return current
		}
		current = current.Next
	}
	return nil
}

func (sll *SingleLinkedList[K, V]) findNodeBefore(key K) *SinglyNode[K, V] {
	var previous *SinglyNode[K, V]
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
