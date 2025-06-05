package utils

import "fmt"

type Token struct {
	Value         string
	Type          string
	Precedence    int
	Associativity string
}

func (t Token) String() string {
	return fmt.Sprintf("Token(%s, T=%s, P=%d, A=%s)",
		t.Value, t.Type, t.Precedence, t.Associativity)
}

type Queue[T any] struct {
	Values []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.Values = append(q.Values, item)
}

func (q *Queue[T]) Dequeue() T {
	if len(q.Values) == 0 {
		return *new(T)
	}
	item := q.Values[0]
	q.Values = q.Values[1:]
	return item
}

func (q Queue[T]) Len() int {
	return len(q.Values)
}

type Stack[T any] struct {
	Values []T
}

func (s *Stack[T]) Push(item T) {
	s.Values = append(s.Values, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.Values) == 0 {
		return *new(T)
	}
	item := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]
	return item
}

func (s Stack[T]) Len() int {
	return len(s.Values)
}

func (s Stack[T]) Top() T {
	if len(s.Values) == 0 {
		return *new(T)
	}
	return s.Values[len(s.Values)-1]
}
