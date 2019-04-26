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
	s := &SegmentTree{
		nodes:    make([]int, len(input)*2),
		treeType: segmentTreeType,
	}
	j := 0
	for i := len(input); i < len(s.nodes); i++ {
		s.nodes[i] = input[j]
		j++
	}
	node := len(input) - 1
	for node > 0 {
		switch s.treeType {
		case SegmentTreeMin:
			s.nodes[node] = int(math.Min(float64(s.nodes[node*2]), float64(s.nodes[node*2+1])))
		case SegmentTreeMax:
			s.nodes[node] = int(math.Max(float64(s.nodes[node*2]), float64(s.nodes[node*2+1])))
		case SegmentTreeSum:
			s.nodes[node] = s.nodes[node*2] + s.nodes[node*2+1]
		}
		node--
	}
	return s
}

// Query queries the result from range i to j.
// Then we can get it's real position in nodes at left and right.
// Where left = n + i, and right = n + j.
// As long as i <= j, we can exam the value at i and j first and then move up to their parents level.
func (s *SegmentTree) Query(i, j int) int {
	// invalid query.
	if i > j || i < 0 || j > len(s.nodes)/2-1 {
		return -1
	}
	var res int

	switch s.treeType {
	case SegmentTreeMin:
		res = math.MaxInt32
	case SegmentTreeMax:
		res = math.MinInt32
	case SegmentTreeSum:
		res = 0
	}

	left := len(s.nodes)/2 + i
	right := len(s.nodes)/2 + j

	for left <= right {
		// if left is the right child of its parent, then we should only take it's value but not the left child's.
		// because left child's value is out of range.
		if left%2 != 0 {
			switch s.treeType {
			case SegmentTreeMin:
				res = int(math.Min(float64(s.nodes[left]), float64(res)))
			case SegmentTreeMax:
				res = int(math.Max(float64(s.nodes[left]), float64(res)))
			case SegmentTreeSum:
				res += s.nodes[left]
			}
			left++
		}

		// if right is the left child of its parent, then we should only take it's value but not the right child's.
		// because right child's value is out of range.
		if right%2 == 0 {
			switch s.treeType {
			case SegmentTreeMin:
				res = int(math.Min(float64(s.nodes[right]), float64(res)))
			case SegmentTreeMax:
				res = int(math.Max(float64(s.nodes[right]), float64(res)))
			case SegmentTreeSum:
				res += s.nodes[right]
			}
			right--
		}

		// now move up to left and right's parents.
		left /= 2
		right /= 2
	}
	return res
}

// Update updates the value of a given index i.
// The idea is that, for a given index i, the leaf node of it is at n + i
// We will then propogate the update from bottom to up.
func (s *SegmentTree) Update(i, update int) {
	// invalid query.
	if i < 0 || i > len(s.nodes)/2-1 {
		return
	}

	node := i + len(s.nodes)/2
	s.nodes[node] = update
	for node > 1 {
		// check if the current node is the parent's left or right node.
		// if it's the left, calculate the right, if it's the right, calculate the left.
		left, right := node, node
		if node%2 == 0 {
			right = right + 1
		} else {
			left = left - 1
		}

		switch s.treeType {
		case SegmentTreeMin:
			s.nodes[node/2] = int(math.Min(float64(s.nodes[left]), float64(s.nodes[right])))
		case SegmentTreeMax:
			s.nodes[node/2] = int(math.Max(float64(s.nodes[left]), float64(s.nodes[right])))
		case SegmentTreeSum:
			s.nodes[node/2] = s.nodes[left] + s.nodes[right]
		}
		node /= 2
	}

}
