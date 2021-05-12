package code

// https://leetcode.com/problems/lru-cache/

type LRUCache struct {
	keys     []int
	storage  map[int]int
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		keys:     make([]int, 0, capacity),
		storage:  make(map[int]int, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	ret, ok := this.storage[key]
	if !ok {
		return -1
	}
	new_list := make([]int, 0, this.capacity)
	new_list = append(new_list, key)
	for i := 0; i < len(this.keys); i++ {
		if this.keys[i] != key {
			new_list = append(new_list, this.keys[i])
		}
	}
	this.keys = new_list
	return ret
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.storage[key]
	if ok {
		this.storage[key] = value

		new_list := make([]int, 0, this.capacity)
		new_list = append(new_list, key)
		for i := 0; i < len(this.keys); i++ {
			if this.keys[i] != key {
				new_list = append(new_list, this.keys[i])
			}
		}
		this.keys = new_list
	} else {
		if len(this.keys) == this.capacity {
			delete(this.storage, this.keys[len(this.keys)-1])
			this.keys = this.keys[:len(this.keys)-1]
		}
		new_keys := []int{key}
		this.storage[key] = value
		this.keys = append(new_keys, this.keys...)
	}

}
