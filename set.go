package main

type Set[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]struct{})}
}

func (s *Set[T]) Add(v T) {
	s.items[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.items, v)
}

func (s *Set[T]) Contains(v T) bool {
	_, exists := s.items[v]
	return exists
}

func (s *Set[T]) Len() int {
	return len(s.items)
}
