// CONCEPT 4: Queue
//
// Queue follows FIFO: First In, First Out.
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

func (q *Queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

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
