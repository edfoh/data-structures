package datastruct

type DoubleLinkedList[K comparable, V any] struct {
	head *doublyNode[K, V]
	tail *doublyNode[K, V]
}

func NewDoubleLinkedList[K comparable, V any]() *DoubleLinkedList[K, V] {
	return &DoubleLinkedList[K, V]{head: nil, tail: nil}
}

func (dll *DoubleLinkedList[K, V]) InsertHead(key K, item V) {
	if dll.isEmpty() {
		dll.insertFirstItem(key, item)
	} else {
		dll.head = dll.head.AddNodeBefore(key, item)
	}
}

func (dll *DoubleLinkedList[K, V]) InsertTail(key K, item V) {
	if dll.isEmpty() {
		dll.insertFirstItem(key, item)
	} else {
		dll.tail = dll.tail.AddNodeAfter(key, item)
	}
}

func (dll *DoubleLinkedList[K, V]) InsertAfter(keyAfter K, keyInsert K, itemInsert V) bool {
	node := dll.findNode(keyAfter)
	if node != nil {
		newNode := node.AddNodeAfter(keyInsert, itemInsert)
		if newNode.IsTail() {
			dll.tail = newNode
		}
		return true
	}
	return false
}

func (dll *DoubleLinkedList[K, V]) Delete(key K) bool {
	node := dll.findNode(key)
	if node != nil {
		if node.Previous != nil {
			node.Previous.Next = node.Next
		} else { // node is head
			dll.head = node.Next
		}
		if node.Next != nil {
			node.Next.Previous = node.Previous
		} else { // node is tail
			dll.tail = node.Previous
		}
		node.Previous = nil
		node.Next = nil
		return true
	}
	return false
}

func (dll *DoubleLinkedList[K, V]) DeleteHead() bool {
	if dll.isEmpty() {
		return false
	}
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
		return true
	} else if dll.head.Next == dll.tail {
		dll.head = dll.tail
		dll.head.Next = nil
		dll.tail.Previous = nil
		return true
	} else {
		dll.head = dll.head.Next
		dll.head.Previous = nil
		return true
	}
}

func (dll *DoubleLinkedList[K, V]) DeleteTail() bool {
	if dll.isEmpty() {
		return false
	}
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
		return true
	} else if dll.tail.Previous == dll.head {
		dll.tail = dll.head
		dll.head.Next = nil
		dll.tail.Previous = nil
		return true
	} else {
		dll.tail = dll.tail.Previous
		dll.tail.Next = nil
		return true
	}
}

func (dll *DoubleLinkedList[K, V]) AllKeys() []K {
	var keys []K
	current := dll.head
	for current != nil {
		keys = append(keys, current.Key)
		current = current.Next
	}
	return keys
}

func (dll *DoubleLinkedList[K, V]) AllKeysReverse() []K {
	var keys []K
	current := dll.tail
	for current != nil {
		keys = append(keys, current.Key)
		current = current.Previous
	}
	return keys
}

func (dll *DoubleLinkedList[K, V]) AllItems() []V {
	var vals []V
	current := dll.head
	for current != nil {
		vals = append(vals, current.Item)
		current = current.Next
	}
	return vals
}

func (dll *DoubleLinkedList[K, V]) AllItemsReverse() []V {
	var vals []V
	current := dll.tail
	for current != nil {
		vals = append(vals, current.Item)
		current = current.Previous
	}
	return vals
}

func (dll *DoubleLinkedList[K, V]) findNode(key K) *doublyNode[K, V] {
	current := dll.tail
	for current != nil {
		if current.Key == key {
			return current
		}
		current = current.Previous
	}
	return nil
}

func (dll *DoubleLinkedList[K, V]) insertFirstItem(key K, item V) {
	dll.tail = newDoublyNode(key, item, nil, nil)
	dll.head = dll.tail
}

func (dll *DoubleLinkedList[K, V]) isEmpty() bool {
	return dll.head == nil && dll.tail == nil
}
