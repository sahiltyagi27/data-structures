# Go Data Structures

Small Go implementations of common data structures for interview preparation.

## Run

```bash
go run .
go test ./...
```

## Topic Index

| Topic | File |
|---|---|
| Arrays and slices | [structures/array_slice.go](structures/array_slice.go) |
| Linked list | [structures/linked_list.go](structures/linked_list.go) |
| Stack | [structures/stack.go](structures/stack.go) |
| Queue | [structures/queue.go](structures/queue.go) |
| Hash map and set | [structures/hashmap_set.go](structures/hashmap_set.go) |
| Binary search tree | [structures/tree.go](structures/tree.go) |
| Graph, BFS, DFS | [structures/graph.go](structures/graph.go) |
| Heap / priority queue | [structures/heap.go](structures/heap.go) |
| LRU cache | [structures/lru_cache.go](structures/lru_cache.go) |
| LFU cache | [structures/lfu_cache.go](structures/lfu_cache.go) |

## Main Guide

Read the full study notes here:

- [GUIDE.md](GUIDE.md)

## Interview Highlights

- Choose data structures based on required operations.
- Slices are Go's everyday dynamic sequence type.
- Maps provide average O(1) key lookup.
- Stacks are LIFO; queues are FIFO.
- Trees model hierarchy and ordered search.
- Graphs model relationships.
- Heaps are useful for priority queues and top-K problems.
- LRU cache evicts the least recently used item.
- LFU cache evicts the least frequently used item, breaking ties by recency.

## LeetCode Mapping

| Data Structure | Example Problem |
|---|---|
| Stack | Valid Parentheses |
| Queue / BFS | Number of Islands |
| Heap | Top K Frequent Elements |
| Graph | Course Schedule |
| Map / Set | Two Sum |
| Tree | Binary Tree Level Order Traversal |
| Linked List | Reverse Linked List |
| LRU Cache | LRU Cache |
| LFU Cache | LFU Cache |
