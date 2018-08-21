package trees

import (
	"errors"

	"github.com/WUMUXIAN/go-common-utils/datastructure/shared"
)

// IndexedPriorityQueue defines an indexed priority queue based on a heap.
type IndexedPriorityQueue struct {
	capacity   int
	size       int
	pq         []int // This is a heap.
	qp         []int
	values     []interface{}
	heapType   HeapType
	comparator shared.Comparator
}

// NewIndexedPriorityQueue creates a new indexed priority queue with capacity, heapType and value comparator specified
func NewIndexedPriorityQueue(capacity int, heapType HeapType, comparator shared.Comparator) *IndexedPriorityQueue {
	return &IndexedPriorityQueue{
		capacity:   capacity,
		size:       0,
		pq:         make([]int, capacity+1), // Because it's used as a heap, so we go 1 indexed.
		qp:         make([]int, capacity+1), // qp is a reversed index for pq. e.g. if qp[pq[index]] = index
		values:     make([]interface{}, capacity),
		heapType:   heapType,
		comparator: comparator,
	}
}

// isIndexValid checks whether the index falls into the range of capacity
func (i *IndexedPriorityQueue) isIndexValid(index int) bool {
	return index >= 0 && index < i.capacity
}

// Contains check whether the queue contains a value with the given index.
func (i *IndexedPriorityQueue) Contains(index int) bool {
	return i.values[index] != nil
}

// Size returns the number of values stored in the queue.
func (i *IndexedPriorityQueue) Size() int {
	return i.size
}

// IsEmpty returns whether the queue is empty.
func (i *IndexedPriorityQueue) IsEmpty() bool {
	return i.size == 0
}

// Insert inters a value into the queue with a given index, if the index is not valid or is taken, error will be return.
func (i *IndexedPriorityQueue) Insert(index int, value interface{}) error {
	if !i.isIndexValid(index) {
		return errors.New("index out of range")
	}
	if i.Contains(index) {
		return errors.New("index is already used")
	}
	i.values[index] = value
	i.size++
	i.pq[i.size] = index
	i.qp[index] = i.size
	i.pqBubbleUp(i.size)
	return nil
}

// Peek peeks the priority queue, it returns the heap top value with it's index. if queue is empty, error will be returned.
func (i *IndexedPriorityQueue) Peek() (index int, value interface{}, err error) {
	if i.IsEmpty() {
		return -1, nil, errors.New("queue is empty")
	}
	return i.pq[1], i.values[i.pq[1]], nil
}

// Pop returns the heap top value of the queue, and remove it from the queue in the meantime, error will be returned if queue is empty.
func (i *IndexedPriorityQueue) Pop() (index int, value interface{}, err error) {
	if i.IsEmpty() {
		return -1, nil, errors.New("queue is empty")
	}
	index = i.pq[1]
	value = i.values[index]
	i.pq[1], i.pq[i.size] = i.pq[i.size], i.pq[1]
	i.qp[i.pq[1]] = 1
	i.qp[i.pq[i.size]] = i.size
	i.size--
	i.pqBubbleDown(1)
	i.qp[index] = -1
	i.values[index] = nil
	return
}

// GetValue gets value at a given index, if index is out of range or no value was added, error will be returned.
func (i *IndexedPriorityQueue) GetValue(index int) (interface{}, error) {
	if !i.isIndexValid(index) {
		return nil, errors.New("index out of range")
	}
	if !i.Contains(index) {
		return nil, errors.New("index does not have value")
	}
	return i.values[index], nil
}

// ChangeValue changes the value at a given index, if the index is out of range or no value was added, error will be returned.
func (i *IndexedPriorityQueue) ChangeValue(index int, value interface{}) error {
	if !i.isIndexValid(index) {
		return errors.New("index out of range")
	}
	if !i.Contains(index) {
		return errors.New("index does not have value")
	}
	i.values[index] = value
	i.pqBubbleUp(i.qp[index])
	i.pqBubbleDown(i.qp[index])
	return nil
}

// DeleteValue deletes the value at a given index, if the index is out of range or no value was added, error will be returned.
func (i *IndexedPriorityQueue) DeleteValue(index int) error {
	if !i.isIndexValid(index) {
		return errors.New("index out of range")
	}
	if !i.Contains(index) {
		return errors.New("index does not have value")
	}
	pqIndex := i.qp[index]
	i.pq[pqIndex], i.pq[i.size] = i.pq[i.size], i.pq[pqIndex]
	i.qp[i.pq[pqIndex]] = pqIndex
	i.qp[i.pq[i.size]] = i.size
	i.size--
	i.pqBubbleUp(pqIndex)
	i.pqBubbleDown(pqIndex)
	i.values[index] = nil
	i.qp[index] = -1
	return nil
}

// pqBubbleUp bubbles up the value at given index in pq.
func (i *IndexedPriorityQueue) pqBubbleUp(index int) {
	for index != 1 {
		if (i.heapType == HeapTypeMin && i.comparator(i.values[i.pq[index]], i.values[i.pq[index/2]]) <= 0) ||
			(i.heapType == HeapTypeMax && i.comparator(i.values[i.pq[index]], i.values[i.pq[index/2]]) >= 0) {
			i.pq[index], i.pq[index/2] = i.pq[index/2], i.pq[index]
			i.qp[i.pq[index]] = index
			i.qp[i.pq[index/2]] = index / 2
			index = index / 2
		} else {
			break
		}
	}
}

// pqBubbleUp bubbles down the value at given index in pq.
func (i *IndexedPriorityQueue) pqBubbleDown(index int) {
	for index*2 <= i.size {
		left := index * 2
		right := index*2 + 1
		indexToReplace := left

		if right <= i.size {
			if (i.heapType == HeapTypeMin && i.comparator(i.values[i.pq[left]], i.values[i.pq[right]]) >= 0) ||
				(i.heapType == HeapTypeMax && i.comparator(i.values[i.pq[left]], i.values[i.pq[right]]) <= 0) {
				indexToReplace = right
			}
		}

		if (i.heapType == HeapTypeMin && i.comparator(i.values[i.pq[index]], i.values[i.pq[indexToReplace]]) >= 0) ||
			(i.heapType == HeapTypeMax && i.comparator(i.values[i.pq[index]], i.values[i.pq[indexToReplace]]) <= 0) {
			i.pq[index], i.pq[indexToReplace] = i.pq[indexToReplace], i.pq[index]
			i.qp[i.pq[index]] = index
			i.qp[i.pq[indexToReplace]] = indexToReplace
			index = indexToReplace
		} else {
			break
		}
	}
}
