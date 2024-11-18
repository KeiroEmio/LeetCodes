package main

import "fmt"

// 定义双向链表节点
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// 定义LRU缓存结构体，有点像一个结构体
type LRUCache struct {
	capacity int
	cache    map[int]*Node // 使用哈希表存储键值对
	head     *Node         // 双向链表头节点，最近最常使用的元素
	tail     *Node         // 双向链表尾节点，最近最少使用的元素
}

// 构造函数，初始化LRU缓存
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     nil,
		tail:     nil,
	}
}

// 获取缓存中的键值对
func (lru *LRUCache) Get(key int) int {
	if node, found := lru.cache[key]; found {
		lru.moveToFront(node) // 访问后将该节点移动到头部
		return node.value
	}
	return -1 // 未找到返回-1
}

// 将键值对插入缓存
func (lru *LRUCache) Put(key int, value int) {
	if node, found := lru.cache[key]; found {
		node.value = value // 如果已经存在，更新值
		lru.moveToFront(node)
	} else {
		newNode := &Node{key: key, value: value}
		lru.cache[key] = newNode
		lru.addToFront(newNode)

		// 如果超出容量，移除尾节点
		if len(lru.cache) > lru.capacity {
			lru.removeTail()
		}
	}
}

// 将节点移到双向链表的头部
func (lru *LRUCache) moveToFront(node *Node) {
	if node == lru.head {
		return // 如果已经是头节点，直接返回
	}
	// 从链表中移除节点
	lru.removeNode(node)
	// 将节点插入头部
	lru.addToFront(node)
}

// 在链表头部添加节点
func (lru *LRUCache) addToFront(node *Node) {
	//if lru.head == nil {
	//	lru.head = node
	//	lru.tail = node
	//} else {
	//	node.next = lru.head
	//	lru.head.prev = node
	//	lru.head = node
	//}
	if lru.head == nil {
		lru.head = node
		lru.tail = node
	} else {
		//插入节点很关键
		node.next = lru.head
		lru.head.prev = node
		lru.head = node
	}
}

// 移除链表中的节点
func (lru *LRUCache) removeNode(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		lru.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		lru.tail = node.prev
	}
}

// 移除双向链表的尾节点（最近最少使用）
func (lru *LRUCache) removeTail() {
	if lru.tail == nil {
		return
	}
	delete(lru.cache, lru.tail.key)
	lru.removeNode(lru.tail)
}

// 测试LRU缓存
func main() {
	lru := Constructor(2) // 设置LRU缓存的容量为2

	lru.Put(1, 1)           // 插入(1, 1)
	lru.Put(2, 2)           // 插入(2, 2)
	fmt.Println(lru.Get(1)) // 返回1
	lru.Put(3, 3)           // 插入(3, 3)，此时缓存容量已满，移除最近最少使用的键2
	fmt.Println(lru.Get(2)) // 返回-1（未找到）
	lru.Put(4, 4)           // 插入(4, 4)，移除最近最少使用的键1
	fmt.Println(lru.Get(1)) // 返回-1（未找到）
	fmt.Println(lru.Get(3)) // 返回3
	fmt.Println(lru.Get(4)) // 返回4
}
