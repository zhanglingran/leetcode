package problems

import "fmt"

// DoubleList 定义双向链表
type DoubleList struct {
	key  int
	val  int
	pre  *DoubleList
	next *DoubleList
}

// 记录链表的表头Head, head.pre 就是链表的第一个元素; head.next 就是最后一个元素
var head *DoubleList

// LRUCache 为一个双向List+Map来实现; map的val是一个指向DoubleList的指针
type LRUCache struct {
	list     *DoubleList
	maps     map[int]*DoubleList
	size     int // 缓存容量的大小
	capacity int
}

func Constructor(capacity int) LRUCache {
	head = &DoubleList{key: -1, val: -1, pre: nil, next: nil}
	head.next = head
	head.pre = head
	return LRUCache{list: head, maps: make(map[int]*DoubleList), size: 0, capacity: capacity}
}

// Get 返回LRU中缓存的数据
// 如果缓存中不存在数据，return -1; 否则，return val
// 需要注意的是: 如果cache中存在，那么get之后，将其放到挪到list的最前边
func (cache *LRUCache) Get(key int) int {
	obj, ok := cache.maps[key]
	if !ok {
		return -1
	}
	// TODO 将该node移动到最前边
	// 将obj从DoubleList上摘下来
	obj.next.pre = obj.pre
	obj.pre.next = obj.next

	// 将该元素接到 DoubleList 的head后边
	head.pre.next = obj
	obj.pre = head.pre

	obj.next = head
	head.pre = obj

	return obj.val
}

func (cache *LRUCache) Put(key int, value int) {

	// 判断这个key是不是存在，存在的话，将其对应的val修改即可
	if cache.Get(key) != -1 {
		cache.maps[key].val = value
		return
	}

	// 缓存已满
	if cache.size == cache.capacity {
		// 将head.next删除
		p := head.next
		p.next.pre = p.pre
		p.pre.next = p.next

		// 并将maps中对应的key删除
		delete(cache.maps, p.key)

		// 然后将size--
		cache.size--
	}

	// 创建一个新的节点
	dNode := &DoubleList{
		key:  key,
		val:  value,
		pre:  nil,
		next: nil,
	}

	// 将该节点放到head前边，并用head.pre指向它
	head.pre.next = dNode
	dNode.pre = head.pre

	dNode.next = head
	head.pre = dNode

	cache.maps[key] = dNode
	cache.size++
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func (cache *LRUCache) PrintList() {
	fmt.Println("print list:")
	p := head.pre
	for p != head {
		fmt.Println(p.key, " -> ", p.val)
		p = p.pre
	}
}
