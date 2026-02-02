package main

import (
	"errors"
)

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Empty() {
		return *new(T), errors.New("stack is empty")
	}
	size := len(s.data)
	item := s.data[size-1]
	s.data = s.data[0 : size-1]
	return item, nil
}

func (s Stack[T]) Size() int {
	return len(s.data)
}

func (s Stack[T]) Empty() bool {
	if len(s.data) == 0 {
		return true
	}
	return false
}

func (s Stack[T]) ToSlice() []T {
	slc := make([]T, s.Size())
	copy(slc, s.data)
	return slc
}

func (s *Stack[T]) Clear() {
	s.data = make([]T, 0)
}
