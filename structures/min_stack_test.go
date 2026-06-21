package structures

import "testing"

func TestMinStack(t *testing.T) {
	var stack MinStack

	stack.Push(5)
	stack.Push(3)
	stack.Push(7)

	if got, _ := stack.GetMin(); got != 3 {
		t.Fatalf("GetMin() = %d, want 3", got)
	}

	if got, _ := stack.Top(); got != 7 {
		t.Fatalf("Top() = %d, want 7", got)
	}

	if got, _ := stack.Pop(); got != 7 {
		t.Fatalf("Pop() = %d, want 7", got)
	}

	if got, _ := stack.GetMin(); got != 3 {
		t.Fatalf("GetMin() after popping 7 = %d, want 3", got)
	}

	if got, _ := stack.Pop(); got != 3 {
		t.Fatalf("Pop() = %d, want 3", got)
	}

	if got, _ := stack.GetMin(); got != 5 {
		t.Fatalf("GetMin() after popping 3 = %d, want 5", got)
	}
}

func TestMinStackEmpty(t *testing.T) {
	var stack MinStack

	if _, ok := stack.Pop(); ok {
		t.Fatal("Pop() on empty stack should return ok=false")
	}

	if _, ok := stack.Top(); ok {
		t.Fatal("Top() on empty stack should return ok=false")
	}

	if _, ok := stack.GetMin(); ok {
		t.Fatal("GetMin() on empty stack should return ok=false")
	}
}
