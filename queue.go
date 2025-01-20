package main

type Queue[T comparable] struct {
	items []T
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

func (q *Queue[T]) Push(v T) {
	q.items = append(q.items, v)
}

func (q *Queue[T]) Pop() T {
	v := q.items[0]
	q.items = q.items[1:]

	return v
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}
