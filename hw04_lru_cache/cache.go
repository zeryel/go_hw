package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	var isFound bool
	if _, isFound = cache.items[key]; isFound {
		cache.items[key].Value = value
		cache.queue.MoveToFront(cache.items[key])
	} else {
		item := cache.queue.PushFront(value)
		cache.items[key] = item

		if cache.queue.Len() > cache.capacity {
			itemToDel := cache.queue.Back()
			for key, item := range cache.items {
				if item.Value == itemToDel.Value {
					delete(cache.items, key)
					break
				}
			}

			cache.queue.Remove(itemToDel)
		}
	}

	return isFound
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	var item *ListItem
	var value interface{}
	var isFound bool

	if item, isFound = cache.items[key]; isFound {
		value = item.Value
		cache.queue.MoveToFront(item)
	} else {
		value = nil
	}

	return value, isFound
}

func (cache *lruCache) Clear() {
	for key := range cache.items {
		cache.queue.Remove(cache.items[key])
		delete(cache.items, key)
	}
}
