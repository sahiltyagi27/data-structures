// CONCEPT 7: Graph
//
// A graph contains vertices/nodes connected by edges.
//
// Visual:
//
//	Nodes / vertices:
//
//	       (1)
//	      /   \
//	    (3)   (2) -- (6)
//	   /   \  /
//	 (5)   (4)
//
// Edges are the connections between nodes:
//
//	1 -- 3
//	1 -- 2
//	3 -- 5
//	3 -- 4
//	4 -- 2
//	2 -- 6
//
// Directed graph:
//
//	A -> B means A points to B.
//
// Undirected graph:
//
//	A -- B means A connects to B and B connects to A.
//
// Common uses:
//   - social networks
//   - maps and routes
//   - dependency graphs
//
// This is an adjacency-list graph: each node maps to its neighboring nodes.
//
// Example adjacency list for the visual above:
//
//	1 -> [3, 2]
//	2 -> [1, 4, 6]
//	3 -> [1, 5, 4]
//	4 -> [3, 2]
//	5 -> [3]
//	6 -> [2]

package structures

import "fmt"

type Graph struct {
	edges map[string][]string
}

// NewGraph creates the adjacency-list map.
//
// Shape:
//
//	A -> [B, C]
//	B -> [D]
//	C -> [E]
func NewGraph() *Graph {
	return &Graph{edges: make(map[string][]string)}
}

// AddEdge adds a directed edge from -> to.
//
// AddEdge("A", "B"):
//
//	A -> [B]
//
// For an undirected graph, add both directions:
//
//	A -> B
//	B -> A
func (g *Graph) AddEdge(from, to string) {
	g.edges[from] = append(g.edges[from], to)
}

// Neighbors returns a copy of the neighbors slice.
//
// Returning a copy prevents callers from accidentally modifying the graph's
// internal adjacency list.
func (g *Graph) Neighbors(node string) []string {
	return append([]string(nil), g.edges[node]...)
}

// BFS visits nodes level by level using a queue.
//
// Example:
//
//	     A
//	   /   \
//	  B     C
//	 /       \
//	D         E
//
// BFS from A:
//
//	A, B, C, D, E
//
// Queue flow:
//
//	pop node -> visit neighbors -> enqueue unseen neighbors
func (g *Graph) BFS(start string) []string {
	var order []string
	visited := NewSet[string]()
	var queue Queue[string]

	visited.Add(start)
	queue.Enqueue(start)

	for queue.Len() > 0 {
		node, _ := queue.Dequeue()
		order = append(order, node)

		for _, next := range g.edges[node] {
			if visited.Has(next) {
				continue
			}
			visited.Add(next)
			queue.Enqueue(next)
		}
	}

	return order
}

// DFS goes as deep as possible before backtracking. This version uses a stack.
//
// Example:
//
//	     A
//	   /   \
//	  B     C
//	 /       \
//	D         E
//
// DFS from A can be:
//
//	A, B, D, C, E
//
// Stack flow:
//
//	pop node -> visit -> push neighbors -> continue deep
func (g *Graph) DFS(start string) []string {
	var order []string
	visited := NewSet[string]()
	var stack Stack[string]

	stack.Push(start)
	for stack.Len() > 0 {
		node, _ := stack.Pop()
		if visited.Has(node) {
			continue
		}

		visited.Add(node)
		order = append(order, node)

		for i := len(g.edges[node]) - 1; i >= 0; i-- {
			stack.Push(g.edges[node][i])
		}
	}

	return order
}

func RunGraphDemo() {
	graph := NewGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "E")
	graph.AddEdge("E", "F")

	fmt.Printf("neighbors(A)=%v\n", graph.Neighbors("A"))
	fmt.Printf("BFS from A=%v\n", graph.BFS("A"))
	fmt.Printf("DFS from A=%v\n", graph.DFS("A"))
}
