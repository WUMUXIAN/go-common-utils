package trees

import (
	"github.com/WUMUXIAN/go-common-utils/datastructure/shared"
)

// HeapType defines the heap type
type HeapType int

// Enum values of HeapType
const (
	HeapTypeMax HeapType = iota
	HeapTypeMin
)

// Heap defines a heap
type Heap struct {
	values     []interface{}
	HeapType   HeapType
	size       int
	Comparator shared.Comparator
}

// Insert inserts a node into a heap
func (h *Heap) Insert(data interface{}) {
	if h.size == 0 {
		h.values = append([]interface{}{nil}, data)
		h.size++
	} else {
		h.values = append(h.values, data)
		h.size++
		h.bubbleUp()
	}
}

// bubbleUp try to bubble the last node of the heap up to maintain the heap
func (h *Heap) bubbleUp() {
	if h.size < 2 {
		return
	}
	// Start from the last node, for each node, compare it to its parent.
	// Loop to the top of the tree.
	// If parent's value is smaller than child's, switch value, otherwise, break.
	current := h.size
	parent := h.size / 2

	for current != 1 {
		if (h.HeapType == HeapTypeMax && h.Comparator(h.values[current], h.values[parent]) <= 0) ||
			(h.HeapType == HeapTypeMin && h.Comparator(h.values[current], h.values[parent]) >= 0) {
			break
		}

		// switch
		h.values[0] = h.values[current]
		h.values[current] = h.values[parent]
		h.values[parent] = h.values[0]

		current = parent
		parent /= 2
	}
}

func (h *Heap) bubbleDown() {
	node := 1
	// Until We've reached a node that doesn't have children.
	for node*2 <= h.size {
		left := node * 2
		right := node*2 + 1

		nodeToReplace := left
		// If we have a right child and it's better than left.
		if right <= h.size {
			if (h.HeapType == HeapTypeMax && h.Comparator(h.values[left], h.values[right]) <= 0) ||
				(h.HeapType == HeapTypeMin && h.Comparator(h.values[left], h.values[right]) >= 0) {
				nodeToReplace = right
			}
		}

		// Replace.
		if (h.HeapType == HeapTypeMax && h.Comparator(h.values[node], h.values[nodeToReplace]) <= 0) ||
			(h.HeapType == HeapTypeMin && h.Comparator(h.values[node], h.values[nodeToReplace]) >= 0) {
			h.values[0] = h.values[node]
			h.values[node] = h.values[nodeToReplace]
			h.values[nodeToReplace] = h.values[0]
			// update the node and continue
			node = nodeToReplace
		} else {
			break
		}
	}
}

// Peek returns the top value of the heap
func (h *Heap) Peek() interface{} {
	if h.size == 0 {
		return nil
	}
	return h.values[1]
}

// Pop returns the top value of the heap.
func (h *Heap) Pop() interface{} {
	if h.size == 0 {
		return nil
	}
	res := h.values[1]
	h.values[1] = h.values[h.size]
	h.values = h.values[:h.size]
	h.size--

	h.bubbleDown()

	return res
}

// GetValues gets the values of the heap
func (h *Heap) GetValues() []interface{} {
	if h.size > 0 {
		return h.values[1:]
	}
	return []interface{}{}
}

// Size gets the current size
func (h *Heap) Size() int {
	return h.size
}

// InitHeap initializes a heap using a list of values.
func (h *Heap) InitHeap(values []interface{}) {
	// Assign the values first.
	h.values = append([]interface{}{nil}, values...)
	h.size = len(values)

	// Start from the parent of the last node.
	// Loop back to the root, which is index 1.
	lastParent := h.size / 2
	for i := lastParent; i >= 1; i-- {
		// Cache the value of node to data[0]
		h.values[0] = h.values[i]
		// Loop through the subtrees to maintain.
		node := i
		// keep iterating as long as the node still has left child.
		for node*2 <= h.size {
			// get left child
			node = node * 2

			// if has left child and right child and right child is larger
			// get the right child.
			if node < h.size {
				if (h.HeapType == HeapTypeMax && h.Comparator(h.values[node], h.values[node+1]) <= 0) ||
					(h.HeapType == HeapTypeMin && h.Comparator(h.values[node], h.values[node+1]) >= 0) {
					node++
				}
			}
			// if current value is larger than the current biggest value in its children.
			if (h.HeapType == HeapTypeMax && h.Comparator(h.values[0], h.values[node]) >= 0) ||
				(h.HeapType == HeapTypeMin && h.Comparator(h.values[0], h.values[node]) <= 0) {
				break
			}
			// Otherwise, we switch the value of the parent of current child.
			h.values[node/2] = h.values[node]
			h.values[node] = h.values[0]
		}
	}
}
