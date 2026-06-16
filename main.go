// data-structures: Small Go implementations for interview preparation.
//
// Run:
//   go run .
//
// Each structure lives in structures/*.go with comments explaining when and why
// you would use it.

package main

import (
	"fmt"

	"data-structures/structures"
)

func main() {
	fmt.Println("=== Arrays and Slices ===")
	structures.RunArraySliceDemo()

	fmt.Println("\n=== Linked List ===")
	structures.RunLinkedListDemo()

	fmt.Println("\n=== Stack ===")
	structures.RunStackDemo()

	fmt.Println("\n=== Queue ===")
	structures.RunQueueDemo()

	fmt.Println("\n=== Hash Map and Set ===")
	structures.RunMapSetDemo()

	fmt.Println("\n=== Binary Search Tree ===")
	structures.RunTreeDemo()

	fmt.Println("\n=== Graph ===")
	structures.RunGraphDemo()

	fmt.Println("\n=== Heap / Priority Queue ===")
	structures.RunHeapDemo()

	fmt.Println("\n=== LRU Cache ===")
	structures.RunLRUCacheDemo()

	fmt.Println("\n=== LFU Cache ===")
	structures.RunLFUCacheDemo()
}
