// CONCEPT 3: Stack
//
// Stack follows LIFO: Last In, First Out.
//
// Visual:
//
//	top
//	 |
//	 v
//	[C]  <- last pushed, first popped
//	[B]
//	[A]
//
// Push adds to the top:
//
//	Push(D)
//
//	top
//	 |
//	 v
//	[D]
//	[C]
//	[B]
//	[A]
//
// Pop removes from the top:
//
//	Pop() -> D
//
// Common uses:
//   - undo history
//   - function call stack
//   - expression parsing
//   - DFS traversal

package structures

import "fmt"

type Stack[T any] struct {
	items []T
}

// Push adds a value at the top of the stack.
//
// Before:
//
//	top
//	 |
//	 v
//	[B, A]
//
// Push("C"):
//
//	top
//	 |
//	 v
//	[C, B, A]
//
// Slice view:
//
//	items = append(items, value)
func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

// Pop removes and returns the top value.
//
// Before:
//
//	top
//	 |
//	 v
//	[C, B, A]
//
// Pop() returns C.
//
// After:
//
//	top
//	 |
//	 v
//	[B, A]
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return value, true
}

// Peek returns the top value without removing it.
//
// Stack:
//
//	top
//	 |
//	 v
//	[C, B, A]
//
// Peek() returns C, stack stays the same.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func RunStackDemo() {
	var stack Stack[string]
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")

	value, _ := stack.Pop()
	top, _ := stack.Peek()
	fmt.Printf("popped=%q currentTop=%q len=%d\n", value, top, stack.Len())
}
