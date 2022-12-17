// provides a generic implementation of a simple stack using
// an array as its underlying data structure
package stackutil

import (
	"errors"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Pop() (T, error) {
	var res T
	if len(s.items) <= 0 {
		return res, errors.New("Stack is empty")
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, nil
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Peek() (T, error) {
	var res T
	if len(s.items) <= 0 {
		return res, errors.New("Stack empty")
	}
	items := s.items
	return items[len(items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) <= 0
}
