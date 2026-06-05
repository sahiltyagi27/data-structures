// CONCEPT 5: Hash Map and Set
//
// Hash Map:
//   Stores key-value pairs. Average lookup, insert, and delete are O(1).
//
// Set:
//   Stores unique values. Go has no built-in set type, so map[T]struct{} is the
//   common idiom.

package structures

import "fmt"

type Set[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]struct{})}
}

func (s *Set[T]) Add(value T) {
	s.items[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.items, value)
}

func (s *Set[T]) Has(value T) bool {
	_, ok := s.items[value]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.items)
}

func RunMapSetDemo() {
	ages := map[string]int{
		"Asha":  28,
		"Rahul": 31,
	}
	ages["Neha"] = 26
	delete(ages, "Rahul")

	age, ok := ages["Asha"]
	fmt.Printf("map lookup: Asha age=%d exists=%v map=%v\n", age, ok, ages)

	visited := NewSet[string]()
	visited.Add("Delhi")
	visited.Add("Mumbai")
	visited.Add("Delhi") // duplicate has no effect
	fmt.Printf("set: has Delhi=%v has Pune=%v len=%d\n",
		visited.Has("Delhi"), visited.Has("Pune"), visited.Len())
}
