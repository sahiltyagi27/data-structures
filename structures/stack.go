// CONCEPT 3: Stack
//
// Stack follows LIFO: Last In, First Out.
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

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

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
