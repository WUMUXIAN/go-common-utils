// Package arraylist contains an array list implementation following the java's implementation of its ArrayList.
// It bascially starts with an slice of capcity of 10, and dynamically extend the size of this slice as the array list grow.
package arraylist

import (
	"errors"

	"github.com/WUMUXIAN/go-common-utils/datastructure/shared"
)

type ArrayList struct {
	elements   []interface{}
	size       int
	capacity   int
	Comparator shared.Comparator
}

// New creates a new array list with given comparator
func New(comparator shared.Comparator) *ArrayList {
	// We give a capacity of 10 initially.
	return &ArrayList{
		make([]interface{}, 10),
		0,
		10,
		comparator,
	}
}

func (a *ArrayList) ensureCapacity(newCapcity int) {
	if a.capacity > newCapcity {
		return
	}
	elements := make([]interface{}, newCapcity)
	copy(elements, a.elements)
	a.elements = elements
	a.capacity = newCapcity
}

// Append appends an element to the end of the array list
func (a *ArrayList) Append(element interface{}) {
	a.Add(a.size, element)
}

// Add adds an element in front of the given index in the array list
func (a *ArrayList) Add(index int, element interface{}) error {
	if index < 0 || index > a.size {
		return errors.New("index out of range")
	}
	if a.size == len(a.elements) {
		a.ensureCapacity(a.size*2 + 1)
	}
	for i := a.size - 1; i >= index; i-- {
		a.elements[i+1] = a.elements[i]
	}
	a.elements[index] = element
	a.size++
	return nil
}

// Remove removes an element from the array list
func (a *ArrayList) Remove(element interface{}) error {
	index, err := a.GetIndexOf(element)
	if err != nil {
		return errors.New("element not found")
	}
	return a.RemoveAt(index)
}

// RemoveAt removes an element at given index
func (a *ArrayList) RemoveAt(index int) error {
	if index < 0 || index >= a.size {
		return errors.New("index out of range")
	}
	for i := index; i < a.size-1; i++ {
		a.elements[i] = a.elements[i+1]
	}
	a.size--
	return nil
}

// GetIndexOf gets index of an element
func (a *ArrayList) GetIndexOf(element interface{}) (int, error) {
	for i := 0; i < a.size; i++ {
		if a.Comparator(a.elements[i], element) == 0 {
			return i, nil
		}
	}
	return -1, errors.New("element not found")
}

// Get gets element at given index
func (a *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= a.size {
		return nil, errors.New("index out of range")
	}
	return a.elements[index], nil
}

// Set sets element at given index
func (a *ArrayList) Set(index int, element interface{}) error {
	if index < 0 || index >= a.size {
		return errors.New("index out of range")
	}
	a.elements[index] = element
	return nil
}

// Contains checks whether the array list contains an element
func (a *ArrayList) Contains(element interface{}) bool {
	_, err := a.GetIndexOf(element)
	return err == nil
}

// GetSize gets the size of the array list
func (a *ArrayList) GetSize() int {
	return a.size
}

// Clear clears the array list
func (a *ArrayList) Clear() {
	a.elements = make([]interface{}, 10)
	a.size = 0
	a.capacity = 10
}

// ToSlice returns the array list as a slice
func (a *ArrayList) ToSlice() []interface{} {
	return a.elements[:a.size]
}
