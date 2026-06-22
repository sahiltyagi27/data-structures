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

Visual:

```text
index:   0    1    2
        +----+----+----+
value:  | 10 | 20 | 30 |
        +----+----+----+
```

Slice:

> Dynamic view over an array. Most Go code uses slices instead of arrays.

Slice header:

```text
pointer -> backing array
len     -> visible elements
cap     -> available space before reallocating
```

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

Visual shape:

```text
head
 |
 v
[10] -> [20] -> [30] -> nil
```

Each node contains:

```text
Value
Next pointer
```

Push front:

```text
Before:
head -> [20] -> [30] -> nil

PushFront(10):
new node points to old head
head moves to new node

After:
head -> [10] -> [20] -> [30] -> nil
```

The Go one-liner:

```go
l.head = &ListNode{Value: value, Next: l.head}
```

This is safe because Go evaluates the right side first. The old `l.head` is stored in `Next`, then `l.head` is updated.

Best when:

- you often insert/delete near known nodes
- you do not need random index access

Tradeoff:

- access by position is O(n), because you walk node by node

---

## 3. Stack — `structures/stack.go`

Stack means LIFO: Last In, First Out.

Visual:

```text
top
 |
 v
[C]  <- popped first
[B]
[A]
```

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

## 4. Min Stack — `structures/min_stack.go`

Min Stack is a stack that also supports `GetMin()` in O(1).

Normal stack:

```text
Push -> O(1)
Pop -> O(1)
Top -> O(1)
Find minimum -> O(n)
```

Min Stack:

```text
Push -> O(1)
Pop -> O(1)
Top -> O(1)
GetMin -> O(1)
```

Implementation idea:

```text
values stack = all values
mins stack   = current minimum at each stack level
```

When pushing:

```text
push value to values
push min(value, currentMin) to mins
```

When popping:

```text
pop from both stacks
```

Example:

```text
push 5
values: [5]
mins:   [5]

push 3
values: [5, 3]
mins:   [5, 3]

push 7
values: [5, 3, 7]
mins:   [5, 3, 3]

GetMin -> 3
```

Interview line:

> I keep one stack for values and another stack for the minimum at each level. Every push stores the current minimum, so GetMin is just the top of the min stack.

---

## 5. Queue — `structures/queue.go`

Queue means FIFO: First In, First Out.

Visual:

```text
front                    back
  |                       |
  v                       v
[A] ----> [B] ----> [C] ----> nil
```

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

## 6. Hash Map and Set — `structures/hashmap_set.go`

Hash map stores key-value pairs.

Visual:

```text
key       value
"Asha" -> 28
"Neha" -> 26
```

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

Set visual:

```text
"Delhi"  -> exists
"Mumbai" -> exists
```

Use `struct{}` because it takes zero bytes.

---

## 7. Tree — `structures/tree.go`

A tree is hierarchical data: root, children, descendants.

Visual:

```text
         root
          |
          v
         [8]
        /   \
      [3]   [10]
     /  \      \
   [1]  [6]    [14]
```

Binary Search Tree rule:

- left side has smaller values
- right side has larger values

In-order traversal of a BST returns values in sorted order.

Average search in a balanced BST is O(log n). If the tree becomes a chain, worst case is O(n).

---

## 8. Graph — `structures/graph.go`

A graph is nodes connected by edges.

Visual:

```text
Nodes / vertices:

       (1)
      /   \
    (3)   (2) -- (6)
   /   \  /
 (5)   (4)

Edges:

1 -- 3
1 -- 2
3 -- 5
3 -- 4
4 -- 2
2 -- 6
```

Vocabulary:

```text
node / vertex = item in the graph
edge          = connection between two nodes
directed      = one-way connection, A -> B
undirected    = two-way connection, A -- B
weighted      = edge has cost/distance
unweighted    = edge is just a connection
```

Adjacency list representation:

```text
1 -> [3, 2]
2 -> [1, 4, 6]
3 -> [1, 5, 4]
4 -> [3, 2]
5 -> [3]
6 -> [2]
```

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

## 9. Heap / Priority Queue — `structures/heap.go`

A heap keeps the highest or lowest priority item easy to access.

Visual tree:

```text
        [10]
       /    \
     [5]    [8]
    /  \
  [1]  [3]
```

Same heap as array:

```text
index:  0   1   2   3   4
       +---+---+---+---+---+
value: |10 | 5 | 8 | 1 | 3 |
       +---+---+---+---+---+
```

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

## 10. LRU Cache — `structures/lru_cache.go`

LRU means Least Recently Used.

When the cache is full, it evicts the item that has not been used for the longest time.

Implementation idea:

- hash map gives O(1) key lookup
- doubly linked list keeps recency order
- front = most recently used
- back = least recently used

Operations:

| Operation | Cost |
|---|---|
| Get | O(1) |
| Put | O(1) |

Interview line:

> LRU cache combines a map with a doubly linked list. The map finds nodes quickly, and the list moves recently used nodes to the front while evicting from the back.

---

## 10. LFU Cache — `structures/lfu_cache.go`

LFU means Least Frequently Used.

When the cache is full, it evicts the item with the lowest access count. If multiple items have the same frequency, evict the least recently used among them.

Implementation idea:

- hash map gives O(1) key lookup
- frequency map groups entries by access count
- each frequency bucket keeps recency order
- `minFreq` tracks the current lowest frequency

Operations:

| Operation | Cost |
|---|---|
| Get | O(1) |
| Put | O(1) |

Interview line:

> LFU cache tracks both frequency and recency. The frequency decides which bucket to evict from, and recency breaks ties inside that bucket.

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
| LRU Cache | evict least recently used | cache with recency-based eviction |
| LFU Cache | evict least frequently used | cache with frequency-based eviction |
