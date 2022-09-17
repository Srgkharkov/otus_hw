package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache    // Remove me after realization.
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	lc.mu.Lock()
	item, found := lc.items[key]
	lc.mu.Unlock()
	if found {
		lc.queue.MoveToFront(item)
		ci := item.Value.(cacheItem)
		return ci.value, true
	}
	return nil, false
}

func (lc *lruCache) DeleteOld() bool {
	isDelete := false
	for lc.queue.Len() > lc.capacity {
		key := lc.queue.Back().Value.(cacheItem).key
		lc.mu.Lock()
		delete(lc.items, key)
		lc.mu.Unlock()
		lc.queue.Remove(lc.queue.Back())
		isDelete = true
	}
	return isDelete
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	lc.mu.Lock()
	item, found := lc.items[key]
	lc.mu.Unlock()
	if found {
		lc.queue.MoveToFront(item)
		item.Value = cacheItem{key: key, value: value}
		return true
	} else {
		li := lc.queue.PushFront(cacheItem{key: key, value: value})
		lc.mu.Lock()
		lc.items[key] = li
		lc.mu.Unlock()
		lc.DeleteOld()
		return false
	}
}

func (lc *lruCache) Clear() {
	for lc.queue.Len() > 0 {
		key := lc.queue.Back().Value.(cacheItem).key
		lc.mu.Lock()
		delete(lc.items, key)
		lc.mu.Unlock()
		lc.queue.Remove(lc.queue.Back())
	}
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
