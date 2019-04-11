package trees

// BinaryIndexedTree defines a binary indexes tree
type BinaryIndexedTree struct {
	c []int
}

// NewBinaryIndexedTree creates a binary indexes tree from a given input array.
func NewBinaryIndexedTree(input []int) *BinaryIndexedTree {
	bit := &BinaryIndexedTree{
		c: make([]int, len(input)+1),
	}
	for i := 0; i < len(input); i++ {
		bit.Update(i, input[i])
	}
	return bit
}

// lowbit of a given x is: from right to left, the first 1 represented.
// e.g. lowbit of 100 is: 100.
// e.g. lowbit of 101 is: 001.
// e.g. lowbit of 110 is: 010.
func lowbit(x int) int {
	return x & -x
}

// GetSum returns the sum of values from index 0 to idx.
func (b *BinaryIndexedTree) GetSum(idx int) int {
	// we need to turn the idx into 1-indexed first.
	idx = idx + 1

	// idx is out of range
	if idx > len(b.c)-1 || idx < 1 {
		return -1
	}

	sum := 0
	for idx > 0 {
		sum += b.c[idx]
		idx -= lowbit(idx)
	}
	return sum
}

// Update updates the val at idx by a delta, idx is from 0 to N-1
func (b *BinaryIndexedTree) Update(idx, delta int) {
	// we need to turn the idx into 1-indexed.
	idx = idx + 1

	for idx <= len(b.c)-1 {
		b.c[idx] += delta
		idx += lowbit(idx)
	}
}

// RangeSum gets the range sum from idx i to j. where i and j is from [0, N).
func (b *BinaryIndexedTree) RangeSum(i, j int) int {
	// i and j is invalid.
	if i > j || i < 0 || j > (len(b.c)-2) {
		return -1
	}
	if i == 0 {
		return b.GetSum(j)
	}
	return b.GetSum(j) - b.GetSum(i-1)
}
