// CONCEPT 10: LFU Cache
//
// LFU means Least Frequently Used.
// When the cache is full, evict the item with the lowest access frequency.
// If multiple items have the same frequency, evict the least recently used
// item among that frequency.
//
// Data structures:
//   - map[key]*node gives O(1) lookup
//   - map[frequency]*list stores nodes grouped by frequency
//   - each frequency list is ordered by recency

package structures

import (
	"container/list"
	"fmt"
)

type lfuEntry struct {
	key   int
	value int
	freq  int
}

type LFUCache struct {
	capacity int
	minFreq  int
	items    map[int]*list.Element
	freqs    map[int]*list.List
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		items:    make(map[int]*list.Element),
		freqs:    make(map[int]*list.List),
	}
}

func (c *LFUCache) Get(key int) (int, bool) {
	element, ok := c.items[key]
	if !ok {
		return 0, false
	}

	entry := element.Value.(*lfuEntry)
	c.incrementFrequency(element)
	return entry.value, true
}

func (c *LFUCache) Put(key, value int) {
	if c.capacity <= 0 {
		return
	}

	if element, ok := c.items[key]; ok {
		entry := element.Value.(*lfuEntry)
		entry.value = value
		c.incrementFrequency(element)
		return
	}

	if len(c.items) == c.capacity {
		c.evictLFU()
	}

	entry := &lfuEntry{key: key, value: value, freq: 1}
	c.minFreq = 1
	c.ensureFreqList(1)
	c.items[key] = c.freqs[1].PushFront(entry)
}

func (c *LFUCache) incrementFrequency(element *list.Element) {
	entry := element.Value.(*lfuEntry)
	oldFreq := entry.freq
	oldList := c.freqs[oldFreq]

	oldList.Remove(element)
	if oldList.Len() == 0 {
		delete(c.freqs, oldFreq)
		if c.minFreq == oldFreq {
			c.minFreq++
		}
	}

	entry.freq++
	c.ensureFreqList(entry.freq)
	c.items[entry.key] = c.freqs[entry.freq].PushFront(entry)
}

func (c *LFUCache) evictLFU() {
	leastFreqList := c.freqs[c.minFreq]
	victim := leastFreqList.Back() // least recently used within min frequency
	entry := victim.Value.(*lfuEntry)

	leastFreqList.Remove(victim)
	if leastFreqList.Len() == 0 {
		delete(c.freqs, c.minFreq)
	}
	delete(c.items, entry.key)
}

func (c *LFUCache) ensureFreqList(freq int) {
	if c.freqs[freq] == nil {
		c.freqs[freq] = list.New()
	}
}

func RunLFUCacheDemo() {
	cache := NewLFUCache(2)
	cache.Put(1, 100)
	cache.Put(2, 200)

	cache.Get(1)      // key 1 frequency becomes 2
	cache.Put(3, 300) // evicts key 2 because key 2 frequency is lower

	_, ok := cache.Get(2)
	fmt.Printf("get(2): ok=%v (evicted)\n", ok)

	value, ok := cache.Get(1)
	fmt.Printf("get(1): value=%d ok=%v\n", value, ok)

	value, ok = cache.Get(3)
	fmt.Printf("get(3): value=%d ok=%v\n", value, ok)
}
