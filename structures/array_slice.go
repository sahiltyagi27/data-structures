// CONCEPT 1: Array and Slice
//
// Array:
//   Fixed-size sequence. The length is part of the type: [3]int and [4]int are
//   different types.
//
// Slice:
//   Dynamic view over an underlying array. Slices are used much more often than
//   arrays in day-to-day Go.
//
// Interview line:
//   In Go, arrays are fixed-size values; slices are flexible descriptors over
//   arrays with length and capacity.

package structures

import "fmt"

func RunArraySliceDemo() {
	numbers := [3]int{10, 20, 30}
	fmt.Printf("array: values=%v len=%d\n", numbers, len(numbers))

	scores := []int{70, 80, 90}
	scores = append(scores, 100)
	fmt.Printf("slice after append: values=%v len=%d cap=%d\n", scores, len(scores), cap(scores))

	// Slicing does not copy values. It creates a new view over the same array.
	firstTwo := scores[:2]
	firstTwo[0] = 75
	fmt.Printf("slice shares backing array: scores=%v firstTwo=%v\n", scores, firstTwo)
}
