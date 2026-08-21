package models

import "errors"

type StackInt struct {
	Elements []int
	Size     int
}

func (s *StackInt) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	lastIndex := len(s.Elements) - 1
	return s.Elements[lastIndex], nil
}

// LIFO: Last in, first out
// stack.push(5) -> stack = [5]
// stack.push(4) -> stack = [5,4]
// stack.push(6) -> stack = [5,4,6]
func (s *StackInt) Push(v int) {
	s.Elements = append(s.Elements, v)
	s.Size = len(s.Elements)
}

// LIFO: Last in, first out
// stack.push(5) -> stack = [5]
// stack.push(4) -> stack = [5,4]
// stack.push(6) -> stack = [5,4,6]
func (s *StackInt) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	lastIndex := len(s.Elements) - 1
	lastElement := s.Elements[lastIndex]

	s.Elements = s.Elements[0:lastIndex]
	s.Size = len(s.Elements)
	return lastElement, nil
}

func (s *StackInt) IsEmpty() bool {
	if s.Size == 0 {
		return true
	}
	return false
}
