// CONCEPT 8: Heap / Priority Queue
//
// Heap:
//   A tree-shaped structure usually stored in an array/slice.
//
// Min-heap:
//   The smallest value is always at the root.
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

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// Higher priority number comes first. Use < for a min-heap.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Task))
}

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
