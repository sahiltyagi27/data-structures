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
// Representation rule:
//
//	If nodes are 0 to n-1, use:
//
//	    graph := make([][]int, n)
//	    visited := make([]bool, n)
//
//	If nodes are random integer IDs, use:
//
//	    graph := map[int][]int{}
//	    visited := map[int]bool{}
//
//	If nodes are strings, use:
//
//	    graph := map[string][]string{}
//	    visited := map[string]bool{}
//
// Interview default:
//
//	Most LeetCode graph problems use nodes 0..n-1, so start with [][]int
//	and []bool unless the input uses strings or sparse/random IDs.
//
// Course Schedule:
//
//	graph := make([][]int, numCourses)
//	indegree := make([]int, numCourses)
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
//
// Interview note:
//
//	Use a slice as a queue so you do not need to implement a separate Queue
//	type during a coding round.
//
//	queue := []string{start}
//	node := queue[0]
//	queue = queue[1:]
func (g *Graph) BFS(start string) []string {
	var order []string
	visited := map[string]bool{}
	queue := []string{start}

	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)

		for _, next := range g.edges[node] {
			if visited[next] {
				continue
			}
			visited[next] = true
			queue = append(queue, next)
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
//
// Interview note:
//
//	Use a slice as a stack so you do not need to implement a separate Stack
//	type during a coding round.
//
//	stack := []string{start}
//	node := stack[len(stack)-1]
//	stack = stack[:len(stack)-1]
func (g *Graph) DFS(start string) []string {
	var order []string
	visited := map[string]bool{}
	stack := []string{start}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[node] {
			continue
		}

		visited[node] = true
		order = append(order, node)

		for i := len(g.edges[node]) - 1; i >= 0; i-- {
			stack = append(stack, g.edges[node][i])
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
