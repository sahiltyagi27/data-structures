# Go Data Structures Guide

> Reference for the `data structures` project.
> Run all examples with `go run .`

---

## What Are Data Structures?

A data structure is a way to organize and store data so operations like lookup, insert, delete, traversal, and update can be done efficiently.

Interview answer:

> Data structures are containers for data. We choose them based on what operations need to be fast, such as searching, inserting, deleting, ordering, or traversing.

---

## Big-O Basics

Big-O describes how cost grows as input size grows.

| Big-O | Meaning | Example |
|---|---|---|
| O(1) | Constant time | map lookup average case |
| O(log n) | Logarithmic | balanced tree search |
| O(n) | Linear | scan a slice |
| O(n log n) | Common sorting cost | efficient general sorting |
| O(n²) | Nested scan | comparing every pair |

---

## 1. Array and Slice — `structures/array_slice.go`

Array:

> Fixed-size sequence. In Go, `[3]int` and `[4]int` are different types.

Slice:

> Dynamic view over an array. Most Go code uses slices instead of arrays.

```go
arr := [3]int{1, 2, 3}
slice := []int{1, 2, 3}
slice = append(slice, 4)
```

Important slice fields:

| Field | Meaning |
|---|---|
| pointer | where the backing array starts |
| length | visible number of elements |
| capacity | how many elements can fit before reallocating |

---

## 2. Linked List — `structures/linked_list.go`

A linked list stores values in nodes. Each node points to the next node.

Best when:

- you often insert/delete near known nodes
- you do not need random index access

Tradeoff:

- access by position is O(n), because you walk node by node

---

## 3. Stack — `structures/stack.go`

Stack means LIFO: Last In, First Out.

```go
stack.Push("A")
stack.Push("B")
stack.Pop() // B
```

Common uses:

- undo history
- function calls
- parsing expressions
- DFS traversal

---

## 4. Queue — `structures/queue.go`

Queue means FIFO: First In, First Out.

```go
queue.Enqueue("A")
queue.Enqueue("B")
queue.Dequeue() // A
```

Common uses:

- task queues
- BFS traversal
- request buffering

---

## 5. Hash Map and Set — `structures/hashmap_set.go`

Hash map stores key-value pairs.

```go
ages := map[string]int{"Asha": 28}
ages["Neha"] = 26
```

Average case:

| Operation | Cost |
|---|---|
| lookup | O(1) |
| insert | O(1) |
| delete | O(1) |

Go set idiom:

```go
visited := map[string]struct{}{}
visited["Delhi"] = struct{}{}
_, ok := visited["Delhi"]
```

Use `struct{}` because it takes zero bytes.

---

## 6. Tree — `structures/tree.go`

A tree is hierarchical data: root, children, descendants.

Binary Search Tree rule:

- left side has smaller values
- right side has larger values

In-order traversal of a BST returns values in sorted order.

Average search in a balanced BST is O(log n). If the tree becomes a chain, worst case is O(n).

---

## 7. Graph — `structures/graph.go`

A graph is nodes connected by edges.

Common uses:

- social networks
- routes and maps
- dependency graphs
- recommendation systems

Representations:

| Representation | Good For |
|---|---|
| adjacency list | sparse graphs, most common |
| adjacency matrix | dense graphs, fast edge lookup |

Traversal:

- BFS uses a queue and explores level by level
- DFS uses a stack or recursion and explores deep first

---

## 8. Heap / Priority Queue — `structures/heap.go`

A heap keeps the highest or lowest priority item easy to access.

Priority queue examples:

- scheduling jobs
- Dijkstra shortest path
- top K elements
- processing urgent tasks first

Go has `container/heap`, but you implement the interface methods.

```go
heap.Push(pq, Task{Name: "urgent", Priority: 10})
task := heap.Pop(pq).(Task)
```

---

## Quick Interview Summary

| Structure | Main Idea | Typical Use |
|---|---|---|
| Array | fixed-size indexed storage | low-level fixed data |
| Slice | dynamic array view | most ordered lists in Go |
| Linked List | nodes connected by pointers | frequent insert/delete with node access |
| Stack | LIFO | undo, parsing, DFS |
| Queue | FIFO | BFS, task processing |
| Hash Map | key-value lookup | fast lookup by key |
| Set | unique values | membership checks |
| Tree | hierarchy | sorted/searchable hierarchical data |
| Graph | nodes and edges | relationships and networks |
| Heap | priority access | priority queues, top K |

