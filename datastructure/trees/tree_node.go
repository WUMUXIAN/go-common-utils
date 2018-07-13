package trees

// BinaryTreeNode defines a tree node of a binary search tree
type BinaryTreeNode struct {
	Data  interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}
