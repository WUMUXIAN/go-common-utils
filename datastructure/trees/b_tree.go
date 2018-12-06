package trees

import (
	"errors"

	"github.com/TectusDreamlab/go-common-utils/datastructure/shared"
)

// BTree defines a B-Tree
type BTree struct {
	Root       *BTreeNode // root has min 2 children if it's not leaf.
	size       int
	order      int
	comparator shared.Comparator
}

// BTreeNode defines a B-Tree Node
type BTreeNode struct {
	Keys     []interface{} // The separation keys, for non-leaf node, count = count(children) - 1
	Values   []interface{}
	Children []*BTreeNode // The childrens, maximum at order
	Parent   *BTreeNode
}

// NewBTree creates a new B-Tree
func NewBTree(order int, comparator shared.Comparator) *BTree {
	if order < 3 {
		panic("the order must be at least 3")
	}
	return &BTree{nil, 0, order, comparator}
}

// Put puts an key (index) and it's value into the B-Tree, if the key exists, update the value
func (b *BTree) Put(key, value interface{}) {
	if b.Root == nil {
		b.Root = &BTreeNode{[]interface{}{key}, []interface{}{value}, []*BTreeNode{}, nil}
		b.size++
		return
	}
	if b.insert(b.Root, key, value) {
		b.size++
	}
	return
}

// Search searchs the key (index) to get the value from the B-Tree
func (b *BTree) Search(key interface{}) (node *BTreeNode, index int, err error) {
	if b.Root == nil {
		return nil, -1, errors.New("empty tree")
	}
	return b.search(b.Root, key)
}

// Delete deletes the key (index) and its value from the B-Tree
func (b *BTree) Delete(key interface{}) (value interface{}, err error) {
	if b.Root == nil {
		return nil, errors.New("empty tree")
	}
	return b.delete(b.Root, key)
}

// GetSize gets the total number of values.
func (b *BTree) GetSize() int {
	return b.size
}

// GetHeight gets the height of the B-Tree.
func (b *BTree) GetHeight() int {
	node := b.Root
	height := 0
	for ; node != nil && len(node.Children) > 0; node = node.Children[0] {
		height++
	}
	return height
}

// Clear clears the B-Tree
func (b *BTree) Clear() {
	b.Root = nil
	b.size = 0
}

func (b *BTree) insert(root *BTreeNode, key, value interface{}) (inserted bool) {
	return true
}

func (b *BTree) search(root *BTreeNode, key interface{}) (node *BTreeNode, index int, err error) {
	node = root
	for node != nil {
		// Let's binary search through the keys, to find the value or locate the sub-stree to search from if not found
		var found bool
		index, found = shared.BinarySearch(node.Keys, key, b.comparator)
		// Found
		if found {
			return
		}
		if !b.isLeaf(node) {
			node = node.Children[index+1]
		}
	}
	return nil, -1, errors.New("not found")
}

func (b *BTree) delete(root *BTreeNode, key interface{}) (value interface{}, err error) {
	return nil, nil
}

func (b *BTree) isLeaf(node *BTreeNode) bool {
	return len(node.Children) == 0
}

func (b *BTree) isFull(node *BTreeNode) bool {
	return len(node.Children) == b.order
}

func (b *BTree) minChildrenPerInternalNode() int {
	return (b.order + 1) / 2
}

func (b *BTree) maxKeysPerNode() int {
	return b.order - 1
}

func (b *BTree) minKeysPerInternalNode() int {
	return b.minChildrenPerInternalNode() - 1
}
