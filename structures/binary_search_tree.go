package structures

import "golang.org/x/exp/constraints"

type treeNode[T constraints.Ordered] struct {
	Value T
	left  *treeNode[T]
	right *treeNode[T]
}

func newTreeNode[T constraints.Ordered](value T, left *treeNode[T], right *treeNode[T]) *treeNode[T] {
	return &treeNode[T]{
		Value: value,
		left:  left,
		right: right,
	}
}

/**
 * This is a classic BinarySearchTree (BST);
 * In a BST, notes to the left have a smaller value that the current node ane nodes to the right have a higher value.
 * Insertion aims to preserve that property.
 * When data is in a random order, this is usually balanced; If data is fully sorted, the tree will become un-balanced.
 * Other trees that try to self-balance: AVL Tree, RedBlack trees, etc
 */
type bst[T constraints.Ordered] struct {
	root *treeNode[T]
	len  int
}

func BinarySearchTree[T constraints.Ordered]() *bst[T] {
	return &bst[T]{
		root: nil,
		len:  0,
	}
}

func (b *bst[T]) IsEmpty() bool {
	return b.root == nil
}

func (b *bst[T]) Length() int {
	return b.len
}

func (b *bst[T]) Insert(value T) {
	if b.root == nil {
		b.root = newTreeNode(value, nil, nil)
		b.len = b.len + 1
	} else {
		b.insert(value, b.root)
	}
}

func (b *bst[T]) insert(value T, node *treeNode[T]) {
	if value < node.Value {
		if node.left == nil {
			node.left = newTreeNode(value, nil, nil)
			b.len = b.len + 1
		} else {
			b.insert(value, node.left)
		}
	} else {
		if node.right == nil {
			node.right = newTreeNode(value, nil, nil)
			b.len = b.len + 1
		} else {
			b.insert(value, node.right)
		}
	}
}

func (b *bst[T]) ToArray() []T {
	result := make([]T, b.len)
	index := 0
	b.traverse(result, b.root, &index)
	return result
}

func (b *bst[T]) traverse(result []T, node *treeNode[T], i *int) {
	if node == nil {
		return
	}
	b.traverse(result, node.left, i)
	result[*i] = node.Value
	*i = *i + 1
	b.traverse(result, node.right, i)
}
