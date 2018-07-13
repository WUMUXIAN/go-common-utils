package trees

import (
	"errors"

	"github.com/WUMUXIAN/go-common-utils/datastructure/shared"
)

type BinarySearchTree struct {
	Root       *BinaryTreeNode
	Comparator shared.Comparator
}

func (b *BinarySearchTree) insert(node *BinaryTreeNode, data interface{}) {
	if b.Comparator(node.Data, data) > 0 {
		if node.Left == nil {
			node.Left = &BinaryTreeNode{Data: data}
		} else {
			b.insert(node.Left, data)
		}
	} else {
		if node.Right == nil {
			node.Right = &BinaryTreeNode{Data: data}
		} else {
			b.insert(node.Right, data)
		}
	}
}

func (b *BinarySearchTree) inorderTraverse(node *BinaryTreeNode) []interface{} {
	if node == nil {
		return []interface{}{}
	}

	ordered := b.inorderTraverse(node.Left)
	ordered = append(ordered, node.Data)
	ordered = append(ordered, b.inorderTraverse(node.Right)...)

	return ordered
}

func (b *BinarySearchTree) getRightMostLeaf(node *BinaryTreeNode) (leaf *BinaryTreeNode, parent *BinaryTreeNode) {
	if node == nil {
		return nil, nil
	}

	for {
		parent = node
		leaf = node.Right
		if leaf != nil {
			parent = leaf
			leaf = leaf.Right
		}
		return
	}
}

func (b *BinarySearchTree) delete(node *BinaryTreeNode, parent *BinaryTreeNode, data interface{}) error {
	if node == nil {
		return errors.New("could not find node")
	}
	compareRes := b.Comparator(node.Data, data)
	// Found the node, let's check how to remove it.
	if compareRes == 0 {
		// If this node is a leaf, we simply remove it.
		if node.Left == nil && node.Right == nil {
			if parent.Left == node {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		} else if node.Right == nil {
			// If this node has left child only, we point the parent to this child.
			if parent.Left == node {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		} else if node.Left == nil {
			// If this node has right child only, we point the parent to this child.
			if parent.Left == node {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		} else {
			// If this node has both children, we find the right most node and replace it with node.
			rightMostLeaf, parent := b.getRightMostLeaf(node)
			node.Data = rightMostLeaf.Data
			if parent.Left == rightMostLeaf {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}
		return nil
	}

	if compareRes > 0 {
		return b.delete(node.Left, node, data)
	} else {
		return b.delete(node.Right, node, data)
	}
}

// Insert inserts a data node into the binary search tree.
func (b *BinarySearchTree) Insert(data interface{}) {
	if b.Root == nil {
		b.Root = &BinaryTreeNode{Data: data}
	} else {
		b.insert(b.Root, data)
	}
}

// Clear clears the binary search tree.
func (b *BinarySearchTree) Clear() {
	b.Root = nil
}

// Delete deletes a data node from binary search tree.
func (b *BinarySearchTree) Delete(data interface{}) error {
	if b.Root == nil {
		return errors.New("could not find node")
	}
	return b.delete(b.Root, &BinaryTreeNode{Right: b.Root}, data)
}

// ToSortedSlice traverse the tree and store the data into a sorted slice
func (b *BinarySearchTree) ToSortedSlice() []interface{} {
	return b.inorderTraverse(b.Root)
}
