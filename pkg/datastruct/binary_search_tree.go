package datastruct

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type TreeNode[K constraints.Ordered] struct {
	data  K
	left  *TreeNode[K]
	right *TreeNode[K]
}

func NewTreeNode[K constraints.Ordered](data K) *TreeNode[K] {
	return &TreeNode[K]{
		data: data,
	}
}

func (node *TreeNode[K]) Left() *TreeNode[K] {
	return node.left
}

func (node *TreeNode[K]) Right() *TreeNode[K] {
	return node.right
}

func (node *TreeNode[K]) Data() K {
	return node.data
}

func (node *TreeNode[K]) SetData(data K) {
	node.data = data
}

func (node *TreeNode[K]) SetLeft(n *TreeNode[K]) {
	node.left = n
}

func (node *TreeNode[K]) SetRight(n *TreeNode[K]) {
	node.right = n
}

func (node *TreeNode[K]) MinValue() K {
	curr := node.Right()
	for curr.Left() != nil {
		curr = curr.Left()
	}
	return curr.Data()
}

type BinarySearchTree[K constraints.Ordered] struct {
	root      *TreeNode[K]
	capacity  int
	currNodes int
}

func NewBinarySearchTree[K constraints.Ordered](capacity int) *BinarySearchTree[K] {
	return &BinarySearchTree[K]{
		root:     nil,
		capacity: capacity,
	}
}

func (tree *BinarySearchTree[K]) Insert(data K) error {
	if tree.currNodes == tree.capacity {
		return errors.New("tree is full")
	}
	tree.root = tree.insert(tree.root, data)
	tree.currNodes++
	return nil
}

// insert will add new data as a leaf node whilst recursively walking down the tree, either as a left or right node.
// if the value is smaller than the current node, it will go left, otherwise right.
func (tree *BinarySearchTree[K]) insert(node *TreeNode[K], data K) *TreeNode[K] {
	if node == nil {
		return NewTreeNode(data)
	}

	if data < node.Data() {
		node.SetLeft(tree.insert(node.left, data))
	} else if data > node.Data() {
		node.SetRight(tree.insert(node.right, data))
	}

	return node
}

func (tree *BinarySearchTree[K]) Search(data K) *TreeNode[K] {
	val := tree.search(tree.root, data)
	return val
}

// search will recursively walk the tree, if the data is smaller than the current node, it will go left, otherwise right
func (tree *BinarySearchTree[K]) search(node *TreeNode[K], data K) *TreeNode[K] {
	if node == nil {
		return nil
	}

	if data == node.Data() {
		return node
	}
	if data < node.Data() {
		return tree.search(node.Left(), data)
	}
	if data > node.Data() {
		return tree.search(node.Right(), data)
	}
	return nil
}

func (tree *BinarySearchTree[K]) Delete(data K) {
	tree.delete(tree.root, data)
}

// delete will recursively walk the tree to find the matching node. If the matching node is a leaf (no children),
// it will just be removed. If the node has 1 child, it will be replaced with the child.
// if the node has both childrent, its value will be replaced with the smallest possible value, which obtained by
// going to the right child, then traversing left until a leaf node is reached. The value of that is replaced at the
// node and the delete is recursively called on its right with that min value to delete it.
func (tree *BinarySearchTree[K]) delete(node *TreeNode[K], data K) *TreeNode[K] {
	if node == nil {
		return node
	}

	if data < node.Data() {
		node.SetLeft(tree.delete(node.Left(), data))
	} else if data > node.Data() {
		node.SetRight(tree.delete(node.Right(), data))
	} else {
		// data == node's data

		// if either left or right is empty, return the other side to swap
		if node.Left() == nil {
			return node.Right()
		}
		if node.Right() == nil {
			return node.Left()
		}

		// node has both left and right
		// find min value from right node to replace itself
		minValue := node.MinValue()
		node.SetData(minValue)

		//remove old node from right side
		node.SetRight(tree.delete(node.Right(), minValue))
	}

	return node
}

func (tree *BinarySearchTree[K]) InorderResults() []K {
	var results []K
	tree.inorder(tree.root, &results)
	return results
}

// inorder will build an inorder traversal, starting at the leaf on the left, with the order
// of left, center, right. It will work its way up to root, then inorder traversal starts on the right leaf node.
func (tree *BinarySearchTree[K]) inorder(node *TreeNode[K], results *[]K) {
	if node != nil {
		tree.inorder(node.left, results)
		*results = append(*results, node.Data())
		tree.inorder(node.right, results)
	}
}
