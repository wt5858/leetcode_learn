package main

import "fmt"

// 剑指 Offer II 031. 最近最少使用缓存
// 运用所掌握的数据结构，设计和实现一个 LRU (Least Recently Used，最近最少使用) 缓存机制 。
//
//实现 LRUCache 类：
//
//LRUCache(int capacity) 以正整数作为容量capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value)如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
//
//示例：
//
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
//
//
//提示：
//
//1 <= capacity <= 3000
//0 <= key <= 10000
//0 <= value <= 105
//最多调用 2 * 105 次 get 和 put
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/OrIXps
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	fmt.Println(lruCache.Get(1)) // 输出 1
	lruCache.Put(3, 3)
	fmt.Println(lruCache.Get(2)) // 输出 -1，因为最近访问过的数据是 (1, 1)，2 已经被移除
	lruCache.Put(4, 4)
	fmt.Println(lruCache.Get(1)) // 输出 -1，因为最近访问过的数据是 (4, 4)，1 已经被移除
	fmt.Println(lruCache.Get(3)) // 输出 3
	fmt.Println(lruCache.Get(4)) // 输出 4
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		// 将访问的节点移动到链表头部
		c.moveToHead(node)
		return node.value
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if node, ok := c.cache[key]; ok {
		// 更新节点的值，并将其移动到链表头部
		node.value = value
		c.moveToHead(node)
	} else {
		node := &Node{
			key:   key,
			value: value,
		}
		c.cache[key] = node
		c.addToHead(node)

		if len(c.cache) > c.capacity {
			// 如果容量超过限制，移除尾部节点
			removeNode := c.removeTail()
			delete(c.cache, removeNode.key)
		}
	}
}

func (c *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) addToHead(node *Node) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) moveToHead(node *Node) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) removeTail() *Node {
	node := c.tail.prev
	c.removeNode(node)
	return node
}
