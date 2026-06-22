// CONCEPT 5: Hash Map and Set
//
// Hash Map:
//   Stores key-value pairs. Average lookup, insert, and delete are O(1).
//
// Hash map visual:
//
//	key       value
//	"Asha" -> 28
//	"Neha" -> 26
//
// Lookup:
//
//	ages["Asha"] -> 28
//
// Set:
//   Stores unique values. Go has no built-in set type, so map[T]struct{} is the
//   common idiom.
//
// Set visual:
//
//	"Delhi"  -> exists
//	"Mumbai" -> exists
//
// Duplicate Add("Delhi") has no effect because map keys are unique.

package structures

import "fmt"

type Set[T comparable] struct {
	items map[T]struct{}
}

// NewSet creates the internal map.
//
// Without make(), writing to the map would panic because a nil map can be read
// but cannot be written to.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]struct{})}
}

// Add stores the value as a key.
//
// Set shape:
//
//	"Delhi"  -> struct{}{}
//	"Mumbai" -> struct{}{}
//
// The value is struct{}{} because it takes no space and we only care whether
// the key exists.
func (s *Set[T]) Add(value T) {
	s.items[value] = struct{}{}
}

// Remove deletes the key. Deleting a missing key is safe in Go.
func (s *Set[T]) Remove(value T) {
	delete(s.items, value)
}

// Has checks whether the key exists.
//
// Go map lookup with comma-ok:
//
//	_, ok := map[key]
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
