// CONCEPT 2: Linked List
//
// A linked list stores values in nodes. Each node points to the next node.
//
// Visual shape:
//
//	head
//	 |
//	 v
//	[10] -> [20] -> [30] -> nil
//
// Each box is a ListNode:
//
//	[Value | Next]
//
// head is not the value itself. head is a pointer to the first node.
//
// Strength:
//   Fast insert/delete when you already have the node reference.
//
// Weakness:
//   Slow random access. To find the 100th item, you walk node by node.

package structures

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

// PushFront inserts at the beginning in O(1).
//
// Before:
//
//	head -> [20] -> [30] -> nil
//
// Operation:
//
//	new node = [10]
//	new.Next = old head
//	head = new
//
// After:
//
//	head -> [10] -> [20] -> [30] -> nil
//
// In Go, the right side is evaluated before assignment, so this one-liner
// safely stores the old head in Next before replacing l.head.
func (l *LinkedList) PushFront(value int) {
	l.head = &ListNode{Value: value, Next: l.head}
	l.size++
}

// Append inserts at the end.
//
// Empty list:
//
//	head -> nil
//
// After appending 10:
//
//	head -> [10] -> nil
//
// Non-empty list:
//
//	head -> [10] -> [20] -> nil
//	                 ^
//	                 current stops here because current.Next == nil
//
// After appending 30:
//
//	head -> [10] -> [20] -> [30] -> nil
func (l *LinkedList) Append(value int) {
	node := &ListNode{Value: value}
	if l.head == nil {
		l.head = node
		l.size++
		return
	}

	current := l.head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = node
	l.size++
}

// Contains traverses the list from head to nil.
//
// Traversal pattern:
//
//	for current := head; current != nil; current = current.Next
//
// This is how we loop over a linked list in Go. A range loop does not work
// directly on a custom linked list.
func (l *LinkedList) Contains(value int) bool {
	for current := l.head; current != nil; current = current.Next {
		if current.Value == value {
			return true
		}
	}
	return false
}

// Values converts the linked list to a slice so it can be printed or ranged over.
//
// Linked list:
//
//	head -> [5] -> [10] -> [20] -> nil
//
// Returned slice:
//
//	[]int{5, 10, 20}
func (l *LinkedList) Values() []int {
	values := make([]int, 0, l.size)
	for current := l.head; current != nil; current = current.Next {
		values = append(values, current.Value)
	}
	return values
}

func RunLinkedListDemo() {
	var list LinkedList
	list.Append(10)
	list.Append(20)
	list.PushFront(5)

	fmt.Printf("values=%v contains(20)=%v contains(99)=%v\n",
		list.Values(), list.Contains(20), list.Contains(99))
}
