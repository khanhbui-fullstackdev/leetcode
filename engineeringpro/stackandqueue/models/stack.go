package models

import "errors"

// In Go, the most idiomatic way to build a stack is using a slice
// . Slices provide dynamic resizing, making them highly efficient for the Last-In-First-Out (LIFO) behavior required by a stack.
type Stack struct {
	elements []int
	Count    int
}

// Push adds an element to the top of the stack
func (s *Stack) Push(v int) {
	s.elements = append(s.elements, v) // append adds to the end of the slice
	s.Count = len(s.elements)
}

// Pop removes and returns the top element
func (s *Stack) Pop() (int, error) {
	if len(s.elements) == 0 {
		return 0, errors.New("stack is empty")
	}

	// Get the last element
	index := len(s.elements) - 1
	element := s.elements[index]

	// remove it by re-slicing
	s.elements = s.elements[:index] // arr := []int{0, 1, 2, 3, 4} -> arr[:len(arr)-1] = [0, 1, 2, 3].
	s.Count = len(s.elements)

	// // Tối ưu bộ nhớ: Xóa tham chiếu tới phần tử cũ
	// s.elements[index] = 0
	return element, nil
}

// Peek returns the top element without removing it
func (s *Stack) Peek() (int, error) {
	if len(s.elements) == 0 {
		return 0, errors.New("stack is empty")
	}
	s.Count = len(s.elements)
	return s.elements[len(s.elements)-1], nil
}
