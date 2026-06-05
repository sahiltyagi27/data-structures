// CONCEPT 6: Tree
//
// A tree is hierarchical: one root, children below it.
//
// Binary Search Tree:
//   Left child values are smaller than the node.
//   Right child values are larger than the node.
//
// Average search/insert: O(log n) when balanced.
// Worst case: O(n) when the tree becomes a chain.

package structures

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

func (t *BinarySearchTree) Insert(value int) {
	t.root = insertNode(t.root, value)
}

func insertNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}
	if value < node.Value {
		node.Left = insertNode(node.Left, value)
	} else if value > node.Value {
		node.Right = insertNode(node.Right, value)
	}
	return node
}

func (t *BinarySearchTree) Contains(value int) bool {
	current := t.root
	for current != nil {
		switch {
		case value == current.Value:
			return true
		case value < current.Value:
			current = current.Left
		default:
			current = current.Right
		}
	}
	return false
}

func (t *BinarySearchTree) InOrder() []int {
	var values []int
	var walk func(*TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		walk(node.Left)
		values = append(values, node.Value)
		walk(node.Right)
	}
	walk(t.root)
	return values
}

func RunTreeDemo() {
	var tree BinarySearchTree
	for _, value := range []int{8, 3, 10, 1, 6, 14, 4, 7} {
		tree.Insert(value)
	}

	fmt.Printf("inOrder(sorted)=%v contains(6)=%v contains(99)=%v\n",
		tree.InOrder(), tree.Contains(6), tree.Contains(99))
}
