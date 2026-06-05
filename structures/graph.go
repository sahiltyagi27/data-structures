// CONCEPT 7: Graph
//
// A graph contains vertices/nodes connected by edges.
//
// Common uses:
//   - social networks
//   - maps and routes
//   - dependency graphs
//
// This is an adjacency-list graph: each node maps to its neighboring nodes.

package structures

import "fmt"

type Graph struct {
	edges map[string][]string
}

func NewGraph() *Graph {
	return &Graph{edges: make(map[string][]string)}
}

func (g *Graph) AddEdge(from, to string) {
	g.edges[from] = append(g.edges[from], to)
}

func (g *Graph) Neighbors(node string) []string {
	return append([]string(nil), g.edges[node]...)
}

// BFS visits nodes level by level using a queue.
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
