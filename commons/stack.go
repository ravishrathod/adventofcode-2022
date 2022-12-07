package commons

import "errors"

type Stack[C any] struct {
	entries []C
}

func (s *Stack[C]) Push(value C) {
	s.entries = append(s.entries, value)
}

func (s *Stack[C]) Peek() (C, error) {
	size := len(s.entries)
	if size > 0 {
		return s.entries[size-1], nil
	}
	return s.getZeroValue(), errors.New("empty stack")
}

func (s *Stack[C]) Pop() (C, error) {
	size := len(s.entries)
	if size == 0 {
		return s.getZeroValue(), errors.New("empty stack")
	}
	index := size - 1
	value := s.entries[index]
	s.entries = s.entries[:index]
	return value, nil
}

func (s *Stack[C]) IsEmpty() bool {
	return len(s.entries) == 0
}

func (s *Stack[C]) getZeroValue() C {
	var zero C
	return zero
}
