// CONCEPT: Min Stack
//
// A Min Stack supports normal stack operations plus GetMin in O(1).
//
// Visual:
//
//	Push 5, Push 3, Push 7
//
//	values stack:
//	top -> [7]
//	       [3]
//	       [5]
//
//	mins stack:
//	top -> [3]  <- min at level with 7
//	       [3]  <- min at level with 3
//	       [5]  <- min at level with 5
//
// GetMin reads the top of mins:
//
//	GetMin() -> 3
//
// The trick is to keep two stacks:
//   - values stores all pushed values
//   - mins stores the minimum value at each stack level
//
// When pushing a value, also push the smaller of:
//   - the new value
//   - the previous minimum
//
// When popping, pop from both stacks.

package structures

import "fmt"

type MinStack struct {
	values []int
	mins   []int
}

// Push stores the value and the minimum at this stack level.
//
// Example:
//
//	Push(5)
//	values: [5]
//	mins:   [5]
//
//	Push(3)
//	values: [5, 3]
//	mins:   [5, 3]
//
//	Push(7)
//	values: [5, 3, 7]
//	mins:   [5, 3, 3]
//
// GetMin is O(1) because the current min is always at the top of mins.
func (s *MinStack) Push(value int) {
	s.values = append(s.values, value)

	if len(s.mins) == 0 || value < s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, value)
		return
	}

	currentMin := s.mins[len(s.mins)-1]
	s.mins = append(s.mins, currentMin)
}

// Pop removes from both stacks.
//
// values and mins must stay the same length so each value has a matching
// "minimum at this level" entry.
func (s *MinStack) Pop() (int, bool) {
	if len(s.values) == 0 {
		return 0, false
	}

	lastIndex := len(s.values) - 1
	value := s.values[lastIndex]
	s.values = s.values[:lastIndex]
	s.mins = s.mins[:lastIndex]

	return value, true
}

// Top returns the last pushed value without removing it.
func (s *MinStack) Top() (int, bool) {
	if len(s.values) == 0 {
		return 0, false
	}
	return s.values[len(s.values)-1], true
}

// GetMin returns the current minimum in O(1).
func (s *MinStack) GetMin() (int, bool) {
	if len(s.mins) == 0 {
		return 0, false
	}
	return s.mins[len(s.mins)-1], true
}

func (s *MinStack) Len() int {
	return len(s.values)
}

func RunMinStackDemo() {
	var stack MinStack
	stack.Push(5)
	stack.Push(3)
	stack.Push(7)

	minBeforePop, _ := stack.GetMin()
	popped, _ := stack.Pop()
	minAfterPop, _ := stack.GetMin()

	fmt.Printf("minBeforePop=%d popped=%d minAfterPop=%d\n", minBeforePop, popped, minAfterPop)
}
