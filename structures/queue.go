// CONCEPT 4: Queue
//
// Queue follows FIFO: First In, First Out.
//
// Visual:
//
//	front                    back
//	  |                       |
//	  v                       v
//	[A] ----> [B] ----> [C] ----> nil
//
// Enqueue adds at the back:
//
//	[A] -> [B] -> [C] -> [D]
//
// Dequeue removes from the front:
//
//	Dequeue() -> A
//	front moves to B
//
// Common uses:
//   - task processing
//   - BFS traversal
//   - buffering requests
//
// This implementation keeps a head index so Dequeue is O(1). Removing from the
// front with items = items[1:] is simple, but can keep old array memory alive.

package structures

import "fmt"

type Queue[T any] struct {
	items []T
	head  int
}

// Enqueue adds a value at the back of the queue.
//
// Before:
//
//	front          back
//	  |             |
//	  v             v
//	[A] -> [B]
//
// Enqueue("C"):
//
//	front                 back
//	  |                    |
//	  v                    v
//	[A] -> [B] -> [C]
func (q *Queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

// Dequeue removes a value from the front of the queue.
//
// Before:
//
//	front                 back
//	  |                    |
//	  v                    v
//	[A] -> [B] -> [C]
//
// Dequeue() returns A.
//
// After:
//
//	       front          back
//	         |             |
//	         v             v
//	[A] -> [B] -> [C]
//
// The old A remains in the backing slice until occasional compaction, but
// q.head moves forward so logically B is now the front.
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.head >= len(q.items) {
		var zero T
		return zero, false
	}

	value := q.items[q.head]
	q.head++

	// Compact occasionally so the old backing array can be garbage collected.
	if q.head > 32 && q.head*2 >= len(q.items) {
		q.items = append([]T(nil), q.items[q.head:]...)
		q.head = 0
	}

	return value, true
}

// Len returns the logical queue length.
//
//	items length = 5
//	head = 2
//	logical length = 5 - 2 = 3
func (q *Queue[T]) Len() int {
	return len(q.items) - q.head
}

func RunQueueDemo() {
	var queue Queue[string]
	queue.Enqueue("A")
	queue.Enqueue("B")
	queue.Enqueue("C")

	first, _ := queue.Dequeue()
	second, _ := queue.Dequeue()
	fmt.Printf("dequeued=%q,%q remainingLen=%d\n", first, second, queue.Len())
}
