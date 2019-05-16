package trees

import (
	"container/list"

	"github.com/WUMUXIAN/go-common-utils/datastructure/shared"
)

// BinaryTreeNode defines a tree node of a binary search tree
type BinaryTreeNode struct {
	Key   interface{}
	Value interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// BinaryTree defines a general binary tree
type BinaryTree struct {
	Root       *BinaryTreeNode
	Comparator shared.Comparator
}

// BreadthFirstTraverse traverse the tree breadth-first way, a.k.a level by level.
// The idea is to use a QUEUE to store candidate left and right children along the way.
func (b *BinaryTree) BreadthFirstTraverse() []interface{} {
	if b.Root == nil {
		return []interface{}{}
	}
	queue := []*BinaryTreeNode{b.Root}
	result := make([]interface{}, 0)
	for len(queue) != 0 {
		// dequeue
		node := queue[0]
		queue = queue[1:]

		// enqueue the children of this node.
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		// get value of the node
		result = append(result, node.Key)
	}
	return result
}

// DepthFirstTraverse traverse the tree depth-first way.
// The idea is to use a STACK to store candidate right and left children along the way.
func (b *BinaryTree) DepthFirstTraverse() []interface{} {
	if b.Root == nil {
		return []interface{}{}
	}
	stack := list.New()
	stack.PushBack(b.Root)
	result := make([]interface{}, 0)
	for stack.Len() != 0 {
		// pop stack
		node := stack.Remove(stack.Back()).(*BinaryTreeNode)

		// push the children of this node to stack, right first, left later.
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}

		// get value of the node
		result = append(result, node.Key)
	}
	return result
}
