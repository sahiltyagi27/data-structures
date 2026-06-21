// CONCEPT: Min Stack
//
// A Min Stack supports normal stack operations plus GetMin in O(1).
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

func (s *MinStack) Push(value int) {
	s.values = append(s.values, value)

	if len(s.mins) == 0 || value < s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, value)
		return
	}

	currentMin := s.mins[len(s.mins)-1]
	s.mins = append(s.mins, currentMin)
}

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

func (s *MinStack) Top() (int, bool) {
	if len(s.values) == 0 {
		return 0, false
	}
	return s.values[len(s.values)-1], true
}

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
