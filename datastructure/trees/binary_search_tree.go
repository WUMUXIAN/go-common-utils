package trees

import (
	"errors"
)

// BinarySearchTree defines a binary search tree
type BinarySearchTree struct {
	BinaryTree
	size int
}

func (b *BinarySearchTree) insert(node *BinaryTreeNode, key, value interface{}) bool {
	if b.Comparator(node.Key, key) > 0 {
		if node.Left == nil {
			node.Left = &BinaryTreeNode{Key: key, Value: value}
			return true
		}
		return b.insert(node.Left, key, value)
	} else if b.Comparator(node.Key, key) < 0 {
		if node.Right == nil {
			node.Right = &BinaryTreeNode{Key: key, Value: value}
			return true
		}
		return b.insert(node.Right, key, value)
	}
	node.Value = value
	return false
}

func (b *BinarySearchTree) search(node *BinaryTreeNode, key interface{}) (value interface{}) {
	for node != nil {
		if b.Comparator(node.Key, key) == 0 {
			value = node.Value
			return
		}
		if b.Comparator(node.Key, key) > 0 {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return
}

func (b *BinarySearchTree) inorderTraverse(node *BinaryTreeNode) []interface{} {
	if node == nil {
		return []interface{}{}
	}

	ordered := b.inorderTraverse(node.Left)
	ordered = append(ordered, node.Key)
	ordered = append(ordered, b.inorderTraverse(node.Right)...)

	return ordered
}

func (b *BinarySearchTree) findMaxNodeWithParent(node *BinaryTreeNode, parent *BinaryTreeNode) (*BinaryTreeNode, *BinaryTreeNode) {
	for node.Right != nil {
		parent = node
		node = node.Right
	}
	return node, parent
}

func (b *BinarySearchTree) delete(node *BinaryTreeNode, parent *BinaryTreeNode, key interface{}) error {
	if node == nil {
		return errors.New("could not find node")
	}
	compareRes := b.Comparator(node.Key, key)
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
			// If this node has both children, we find the right most node in the left sub tree and replace it with node.
			// The right most node also means its the biggest in the left tree.
			maxNode, maxNodeParent := b.findMaxNodeWithParent(node.Left, node)
			node.Key = maxNode.Key
			node.Value = maxNode.Value
			if maxNodeParent.Left == maxNode {
				maxNodeParent.Left = maxNode.Left
			} else {
				maxNodeParent.Right = maxNode.Left
			}
		}
		return nil
	}

	if compareRes > 0 {
		return b.delete(node.Left, node, key)
	}
	return b.delete(node.Right, node, key)
}

// GetSize gets the size of the tree.
func (b *BinarySearchTree) GetSize() int {
	return b.size
}

// Search searchs value by key.
func (b *BinarySearchTree) Search(key interface{}) (value interface{}) {
	return b.search(b.Root, key)
}

// Put puts a data node into the binary search tree, if the key exists already, update its value.
func (b *BinarySearchTree) Put(key, value interface{}) {
	if b.Root == nil {
		b.Root = &BinaryTreeNode{Key: key, Value: value}
		b.size++
		return
	}
	if b.insert(b.Root, key, value) {
		b.size++
	}
	return

}

// Clear clears the binary search tree.
func (b *BinarySearchTree) Clear() {
	b.Root = nil
	b.size = 0
}

// Delete deletes a data node from binary search tree.
func (b *BinarySearchTree) Delete(data interface{}) error {
	fakeParent := &BinaryTreeNode{Right: b.Root}
	err := b.delete(b.Root, fakeParent, data)
	if err == nil {
		b.size--
		b.Root = fakeParent.Right
	}
	return err
}

// ToSortedSlice traverse the tree and store the data into a sorted slice
func (b *BinarySearchTree) ToSortedSlice() []interface{} {
	return b.inorderTraverse(b.Root)
}

// ConvertToDoubleLinkedList converts the BST to A Double Linked List.
func (b *BinarySearchTree) ConvertToDoubleLinkedList() (head *BinaryTreeNode, tail *BinaryTreeNode) {
	if b.Root == nil {
		return nil, nil
	}
	b.convertToDoubleLinkedList(b.Root, &head, &tail)
	return
}

func (b *BinarySearchTree) convertToDoubleLinkedList(node *BinaryTreeNode, head **BinaryTreeNode, tail **BinaryTreeNode) {
	if node.Left != nil {
		b.convertToDoubleLinkedList(node.Left, head, tail)
	}

	if (*tail) == nil {
		(*tail) = node
		(*head) = node
	} else {
		(*tail).Right = node
		node.Left = (*tail)
		(*tail) = node
	}

	if node.Right != nil {
		b.convertToDoubleLinkedList(node.Right, head, tail)
	}
}

// ConvertFromDoubleLinkedList converts double linked list back to
func (b *BinarySearchTree) ConvertFromDoubleLinkedList(head *BinaryTreeNode, length int) {
	b.Root = b.convertFromDoubleLinkedList(&head, length)
}

func (b *BinarySearchTree) convertFromDoubleLinkedList(head **BinaryTreeNode, length int) *BinaryTreeNode {
	// This means we reach the most left
	if length == 0 {
		return nil
	}

	// Otherwise, we get the root for the left subtree.
	left := b.convertFromDoubleLinkedList(head, length/2)
	root := *head
	root.Left = left
	*head = (*head).Right

	root.Right = b.convertFromDoubleLinkedList(head, length-length/2-1)

	return root
}
