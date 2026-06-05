// CONCEPT 2: Linked List
//
// A linked list stores values in nodes. Each node points to the next node.
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

func (l *LinkedList) PushFront(value int) {
	l.head = &ListNode{Value: value, Next: l.head}
	l.size++
}

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

func (l *LinkedList) Contains(value int) bool {
	for current := l.head; current != nil; current = current.Next {
		if current.Value == value {
			return true
		}
	}
	return false
}

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
