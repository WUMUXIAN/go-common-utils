package trees

import (
	"math"
)

// SegmentTree represents a segment tree.
type SegmentTree struct {
	nodes    []int
	treeType SegmentTreeType
}

// SegmentTreeType defines the segment tree types.
type SegmentTreeType int

// Define 3 types of segment tree.
// These are for 3 types of query respectively, query range minimum, range max and range sum.
const (
	SegmentTreeMin SegmentTreeType = iota + 1
	SegmentTreeMax
	SegmentTreeSum
)

// NewSegmentTree creates a segment tree for the input array.
func NewSegmentTree(input []int, segmentTreeType SegmentTreeType) *SegmentTree {
	// for segment tree, the size of it will be 2*N - 1, where N is the length of input.
	// we use an array to store the nodes, and like heap, we ignore the 0-index and start from 1, so we init a 2*N size nodes array.
	segmentTree := &SegmentTree{
		nodes:    make([]int, len(input)*2),
		treeType: segmentTreeType,
	}
	segmentTree.build(1, 0, len(input)-1, input)
	return segmentTree
}

// node is the index of the tree node, range from 1 to 2N, start and end is the range this node represented.
func (s *SegmentTree) build(node, start, end int, input []int) {
	// if input is empty, nothing to do.
	if len(input) == 0 {
		return
	}
	if start == end {
		s.nodes[node] = input[start]
	} else {
		mid := (start + end) / 2
		// build the left subtree
		s.build(node*2, start, mid, input)
		// build the right subtree
		s.build(node*2+1, mid+1, end, input)
		// once the left and right tree is built, we will aggregate the value of left and right tree based on the tree type.
		switch s.treeType {
		case SegmentTreeMin:
			s.nodes[node] = int(math.Min(float64(s.nodes[node*2]), float64(s.nodes[node*2+1])))
		case SegmentTreeMax:
			s.nodes[node] = int(math.Max(float64(s.nodes[node*2]), float64(s.nodes[node*2+1])))
		case SegmentTreeSum:
			s.nodes[node] = s.nodes[node*2] + s.nodes[node*2+1]
		}

	}
}

// node is the index of the tree node, range from 1 to 2N, start and end is the range this node represented.
func (s *SegmentTree) query(node, start, end int, i, j int) int {
	// if the range that this node represents is completely out of the range of the query range from i to j, this node should not be considered.
	if i > end || j < start {
		switch s.treeType {
		case SegmentTreeMin:
			return math.MaxInt32
		case SegmentTreeMax:
			return math.MinInt32
		case SegmentTreeSum:
			return 0
		}
	}
	// if the range that this node represents is completely within the range of the query range from i to j, we can use the value of this node directly.
	if i <= start && j >= end {
		return s.nodes[node]
	}

	// if the range that this node represents is partially out of the range of the query range from i to j.
	mid := (start + end) / 2
	// left
	left := s.query(node*2, start, mid, i, j)
	// right
	right := s.query(node*2+1, mid+1, end, i, j)

	var res int
	switch s.treeType {
	case SegmentTreeMin:
		res = int(math.Min(float64(left), float64(right)))
	case SegmentTreeMax:
		res = int(math.Max(float64(left), float64(right)))
	case SegmentTreeSum:
		res = left + right
	}
	return res
}

// node is the index of the tree node, range from 1 to 2N, start and end is the range this node represented.
func (s *SegmentTree) update(node, start, end int, i, update int) {
	// we find the element.
	if start == end {
		s.nodes[node] = update
	} else {
		mid := (start + end) / 2
		if start <= i && i <= mid {
			// If idx is in the left child, recurse on the left child
			s.update(2*node, start, mid, i, update)
		} else {
			// if idx is in the right child, recurse on the right child
			s.update(2*node+1, mid+1, end, i, update)
		}
		switch s.treeType {
		case SegmentTreeMin:
			s.nodes[node] = int(math.Min(float64(s.nodes[node*2]), float64(s.nodes[node*2+1])))
		case SegmentTreeMax:
			s.nodes[node] = int(math.Max(float64(s.nodes[node*2]), float64(s.nodes[node*2+1])))
		case SegmentTreeSum:
			s.nodes[node] = s.nodes[node*2] + s.nodes[node*2+1]
		}
	}
}

// Query queries the result from range i to j.
// The result depends on what type of segment tree it is.
// If it's minimum, then the result is the minimum value from i to j.
// If it's maximum, then the result is the maximum value from i to j.
// If it's sum, then the result is the sum value from i to j.
func (s *SegmentTree) Query(i, j int) int {
	// invalid query.
	if i > j || i < 0 || j > len(s.nodes)/2-1 {
		return -1
	}
	// query starts from the root.
	return s.query(1, 0, len(s.nodes)/2-1, i, j)
}

// Update updates the value of a given index i.
// We do it recursively, if we already find it, we update the value.
// After we find it, we will update it's parents value from bottom to up.
func (s *SegmentTree) Update(i, update int) {
	// invalid update
	if i < 0 || i > len(s.nodes)/2-1 {
		return
	}

	s.update(1, 0, len(s.nodes)/2-1, i, update)
}
