package utils

import "fmt"

// Token represents a token in an expression with its value, type, precedence, and associativity.
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

// Queue is a generic queue implementation.
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

// Stack is a generic stack implementation.
type Stack[T any] struct {
	Values []T
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.Values = append(s.Values, item)
}

// Pop removes the top item from the stack and returns it. If the stack is empty
// Pop returns the zero value for the type T.
func (s *Stack[T]) Pop() T {
	if len(s.Values) == 0 {
		return *new(T)
	}
	item := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]
	return item
}

// Len returns the number of items in the stack.
func (s Stack[T]) Len() int {
	return len(s.Values)
}

// Top returns the item at the top of the stack without removing it. If the stack
// is empty Top returns the zero value for the type T.
func (s Stack[T]) Top() T {
	if len(s.Values) == 0 {
		return *new(T)
	}
	return s.Values[len(s.Values)-1]
}
