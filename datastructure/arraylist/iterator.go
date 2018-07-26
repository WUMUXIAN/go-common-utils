package arraylist

// Iterator defines an iterator for the array list
type Iterator struct {
	arrayList *ArrayList
	index     int
}

// Next moves the iterator to the next value, return false if there's no next value.
func (i *Iterator) Next() bool {
	if i.index < i.arrayList.GetSize()-1 {
		i.index++
		return true
	}
	return false
}

// Value gets the value of element the iterator is pointing at.
func (i *Iterator) Value() interface{} {
	element, err := i.arrayList.Get(i.index)
	if err != nil {
		return nil
	}
	return element
}

// Prev moves the iterator to the previous value, return false if there's no previous value.
func (i *Iterator) Prev() bool {
	if i.index >= 0 {
		i.index--
		return true
	}
	return false
}

// Begin resets the iterator to the beginning, use Next() to move to its first element.
func (i *Iterator) Begin() {
	i.index = -1
}

// End moves the iterator to the end, use Prev() to move to its last element.
func (i *Iterator) End() {
	i.index = i.arrayList.GetSize()
}
