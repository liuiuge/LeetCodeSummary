package code

import (
	"container/list"
)

// https://leetcode.com/problems/lfu-cache/

type LFUCache struct {
	bykey    map[int]*list.Element
	freqs    *list.List
	capacity int
}

type FreqItem struct {
	count int
	items *list.List
}

type CacheItem struct {
	val    int
	key    int
	parent *list.Element
}

func LFUConstructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		bykey:    make(map[int]*list.Element),
		freqs:    list.New(),
	}
}

func (l *LFUCache) Get(key int) int {
	e, ok := l.bykey[key]
	if !ok {
		return -1
	}
	l.increment(e)
	return e.Value.(*CacheItem).val
}

func (l *LFUCache) Put(key int, value int) {
	if l.capacity == 0 {
		return
	}
	e, ok := l.bykey[key]
	if ok {
		cItem := e.Value.(*CacheItem)
		cItem.val = value
		l.increment(e)
		return
	}
	if len(l.bykey) == l.capacity {
		l.evict()
	}

	// add
	freq := l.freqs.Front()
	if freq == nil || freq.Value.(*FreqItem).count != 1 {
		freq = l.freqs.PushFront(&FreqItem{count: 1, items: list.New()})
	}
	fItem := freq.Value.(*FreqItem)
	cItem := &CacheItem{key: key, val: value, parent: freq}
	e = fItem.items.PushBack(cItem)
	l.bykey[key] = e
}

func (l *LFUCache) increment(e *list.Element) {
	cItem := e.Value.(*CacheItem)
	currentFreq := cItem.parent
	nextFreq := cItem.parent.Next()
	fItem := currentFreq.Value.(*FreqItem)
	if nextFreq == nil || fItem.count+1 != nextFreq.Value.(*FreqItem).count {
		newF := &FreqItem{
			count: fItem.count + 1,
			items: list.New(),
		}
		nextFreq = l.freqs.InsertAfter(newF, currentFreq)
	}
	// remove e from currentFreq, move e to new freq
	fItem.items.Remove(e)
	if fItem.items.Len() == 0 {
		// remove freq node if empty
		l.freqs.Remove(currentFreq)
	}
	cItem.parent = nextFreq
	e = nextFreq.Value.(*FreqItem).items.PushBack(cItem)
	// e now belong to a new list, so we'll need to update bykey with new pointer
	l.bykey[e.Value.(*CacheItem).key] = e
}

func (l *LFUCache) evict() {
	if len(l.bykey) == 0 {
		return
	}
	freq := l.freqs.Front()
	if freq == nil {
		return
	}
	fItem := freq.Value.(*FreqItem)
	e := fItem.items.Front()
	fItem.items.Remove(e)
	if fItem.items.Len() == 0 {
		l.freqs.Remove(freq)
	}
	delete(l.bykey, e.Value.(*CacheItem).key)
}
