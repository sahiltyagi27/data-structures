// CONCEPT 9: LRU Cache
//
// LRU means Least Recently Used.
// When the cache is full, evict the item that has not been used for the
// longest time.
//
// Visual:
//
//	map:
//	  key -> pointer to linked-list node
//
//	list:
//	  head <-> [most recent] <-> [...] <-> [least recent] <-> tail
//
// Example after Put(1), Put(2), Get(1):
//
//	head <-> [1] <-> [2] <-> tail
//
// key 1 is most recently used because Get(1) moved it to the front.
// If capacity is full, evict tail.prev.
//
// Data structures:
//   - map[key]*node gives O(1) lookup
//   - doubly linked list keeps recency order
//
// Most recently used item is near the front.
// Least recently used item is near the back.

package structures

import "fmt"

type lruNode struct {
	key   int
	value int
	prev  *lruNode
	next  *lruNode
}

type LRUCache struct {
	capacity int
	items    map[int]*lruNode
	head     *lruNode
	tail     *lruNode
}

// NewLRUCache creates a cache with dummy head/tail nodes.
//
// Empty list shape:
//
//	head <-> tail
//
// After Put(1), Put(2):
//
//	head <-> [2] <-> [1] <-> tail
//
// Front = most recently used.
// Back = least recently used.
func NewLRUCache(capacity int) *LRUCache {
	head := &lruNode{}
	tail := &lruNode{}
	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		items:    make(map[int]*lruNode),
		head:     head,
		tail:     tail,
	}
}

// Get returns the value and marks the key as most recently used.
//
// Before Get(1):
//
//	head <-> [2] <-> [1] <-> tail
//
// After Get(1):
//
//	head <-> [1] <-> [2] <-> tail
func (c *LRUCache) Get(key int) (int, bool) {
	node, ok := c.items[key]
	if !ok {
		return 0, false
	}

	c.moveToFront(node)
	return node.value, true
}

// Put inserts or updates a key.
//
// If key exists:
//
//	update value
//	move node to front
//
// If key is new and cache is full:
//
//	evict node before tail
//	add new node after head
func (c *LRUCache) Put(key, value int) {
	if c.capacity <= 0 {
		return
	}

	if node, ok := c.items[key]; ok {
		node.value = value
		c.moveToFront(node)
		return
	}

	if len(c.items) == c.capacity {
		c.evictLRU()
	}

	node := &lruNode{key: key, value: value}
	c.items[key] = node
	c.addToFront(node)
}

// moveToFront refreshes recency.
//
// remove from current position, then add right after head.
func (c *LRUCache) moveToFront(node *lruNode) {
	c.remove(node)
	c.addToFront(node)
}

// addToFront inserts node after dummy head.
//
// Before:
//
//	head <-> first
//
// After:
//
//	head <-> node <-> first
func (c *LRUCache) addToFront(node *lruNode) {
	first := c.head.next
	node.prev = c.head
	node.next = first
	c.head.next = node
	first.prev = node
}

// remove detaches node from the doubly linked list.
//
// Before:
//
//	prev <-> node <-> next
//
// After:
//
//	prev <---------> next
func (c *LRUCache) remove(node *lruNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// evictLRU removes the node before dummy tail.
//
// Shape:
//
//	head <-> most recent ... least recent <-> tail
//
// tail.prev is the least recently used real node.
func (c *LRUCache) evictLRU() {
	lru := c.tail.prev
	c.remove(lru)
	delete(c.items, lru.key)
}

func RunLRUCacheDemo() {
	cache := NewLRUCache(2)
	cache.Put(1, 100)
	cache.Put(2, 200)

	value, ok := cache.Get(1) // key 1 becomes most recently used
	fmt.Printf("get(1): value=%d ok=%v\n", value, ok)

	cache.Put(3, 300) // evicts key 2

	_, ok = cache.Get(2)
	fmt.Printf("get(2): ok=%v (evicted)\n", ok)

	value, ok = cache.Get(3)
	fmt.Printf("get(3): value=%d ok=%v\n", value, ok)
}
