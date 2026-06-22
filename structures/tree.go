// CONCEPT 6: Tree
//
// A tree is hierarchical: one root, children below it.
//
// Tree visual:
//
//	         root
//	          |
//	          v
//	         [8]
//	        /   \
//	      [3]   [10]
//	     /  \      \
//	   [1]  [6]    [14]
//
// Vocabulary:
//
//	root   = top node
//	child  = node below another node
//	parent = node above another node
//	leaf   = node with no children
//
// Binary Search Tree:
//   Left child values are smaller than the node.
//   Right child values are larger than the node.
//
// BST decision flow:
//
//	value < node.Value -> go left
//	value > node.Value -> go right
//	value == node.Value -> found / duplicate ignored here
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

// Insert adds a value while preserving the BST rule.
//
// BST rule:
//
//	   [8]
//	  /   \
//	[3]   [10]
//
// values smaller than 8 go left.
// values larger than 8 go right.
func (t *BinarySearchTree) Insert(value int) {
	t.root = insertNode(t.root, value)
}

// insertNode returns the subtree root after insertion.
//
// This recursive shape works because each call says:
//
//	"insert into my left or right subtree, then return me"
//
// Empty spot:
//
//	nil -> [value]
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

// Contains searches by walking left or right.
//
// Example search for 6:
//
//	   [8]
//	  /
//	[3]
//	  \
//	  [6]
//
// 6 < 8, go left.
// 6 > 3, go right.
// 6 == 6, found.
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

// InOrder returns values in sorted order for a BST.
//
// Traversal order:
//
//	left subtree -> node -> right subtree
//
// Example:
//
//	   [2]
//	  /   \
//	[1]   [3]
//
// InOrder -> [1, 2, 3]
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
