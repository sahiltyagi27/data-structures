// CONCEPT 8: Heap / Priority Queue
//
// Heap:
//   A tree-shaped structure usually stored in an array/slice.
//
// Heap visual:
//
//	         [10]
//	        /    \
//	      [5]    [8]
//	     /  \
//	   [1]  [3]
//
// Same heap as array:
//
//	index:  0   1   2   3   4
//	       +---+---+---+---+---+
//	value: |10 | 5 | 8 | 1 | 3 |
//	       +---+---+---+---+---+
//
// Parent/child indexes:
//
//	left child  = 2*i + 1
//	right child = 2*i + 2
//	parent      = (i - 1) / 2
//
// Min-heap:
//   The smallest value is always at the root.
//
// Max-heap:
//   The largest value is always at the root.
//
// Priority Queue:
//   Uses a heap so the highest-priority item can be removed efficiently.
//
// Go provides container/heap. You implement Len, Less, Swap, Push, and Pop.

package structures

import (
	"container/heap"
	"fmt"
)

type Task struct {
	Name     string
	Priority int
}

type PriorityQueue []Task

// Len is required by heap.Interface.
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less decides heap priority.
//
// This implementation is a max-priority queue:
//
//	higher Priority number = comes out first
//
// Use < instead of > for a min-priority queue.
func (pq PriorityQueue) Less(i, j int) bool {
	// Higher priority number comes first. Use < for a min-heap.
	return pq[i].Priority > pq[j].Priority
}

// Swap is required so the heap package can reorder items.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push adds an item to the heap's backing slice.
//
// container/heap calls this, then restores heap order.
func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Task))
}

// Pop removes the last item from the backing slice.
//
// container/heap first swaps the highest-priority item to the end, then calls
// this method. That is why this method removes from the end.
func (pq *PriorityQueue) Pop() any {
	old := *pq
	lastIndex := len(old) - 1
	item := old[lastIndex]
	*pq = old[:lastIndex]
	return item
}

func RunHeapDemo() {
	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, Task{Name: "write tests", Priority: 2})
	heap.Push(pq, Task{Name: "fix production bug", Priority: 10})
	heap.Push(pq, Task{Name: "refactor small helper", Priority: 1})

	for pq.Len() > 0 {
		task := heap.Pop(pq).(Task)
		fmt.Printf("next=%q priority=%d\n", task.Name, task.Priority)
	}
}
