package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if len(slice) < index+1 || index < 0 {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if len(slice) < index+1 || index < 0 {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(s []int, values ...int) []int {
	return append(values, s...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(s []int, i int) []int {
	if len(s) < i+1 || i < 0 {
		return s
	} else if i == 0 {
		return s[1:]
	} else if len(s) == i+1 {
		return s[:i]
	} else {
		b := s[:i]
		e := s[i+1:]
		return append(b, e...)
	}
}
