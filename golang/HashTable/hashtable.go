package main

import (
	"fmt"
	"strings"
)

// HashTable 哈希表实现
type HashTable[K any, V any] struct {
	buckets    []*MapNode[K, V]
	capacity   int
	size       int
	upperTol   int
	lowerTol   int
	capacitys  []int
	capIdx     int
}

// MapNode 哈希表节点
type MapNode[K any, V any] struct {
	key   K
	value V
	next  *MapNode[K, V]
}

// NewHashTable 创建一个新的哈希表
func NewHashTable[K any, V any]() *HashTable[K, V] {
	capacitys := []int{53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593,
		49157, 98317, 196613, 393241, 786433, 1572869, 3145739,
		6291469, 12582917, 25165843, 50331653, 100663319, 201326611,
		402653189, 805306457, 1610612741}

	capacity := capacitys[0]
	buckets := make([]*MapNode[K, V], capacity)

	return &HashTable[K, V]{
		buckets:   buckets,
		capacity:  capacity,
		size:      0,
		upperTol:  10,
		lowerTol:  2,
		capacitys: capacitys,
		capIdx:    0,
	}
}

// hash 计算键的哈希值
func (h *HashTable[K, V]) hash(key K) int {
	// 简单的哈希函数 - 使用字符串表示
	keyStr := fmt.Sprintf("%v", key)
	hash := 0
	for _, c := range keyStr {
		hash = (hash << 5) + int(c)
	}
	return (hash & 0x7fffffff) % h.capacity
}

func (h *HashTable[K, V]) Size() int {
	return h.size
}

func (h *HashTable[K, V]) Add(key K, value V) {
	bucket := h.hash(key)
	node := h.getNode(h.buckets[bucket], key)

	if node != nil {
		// 键已存在，更新值
		node.value = value
	} else {
		// 键不存在，添加新节点
		newNode := &MapNode[K, V]{
			key:   key,
			value: value,
			next:  h.buckets[bucket],
		}
		h.buckets[bucket] = newNode
		h.size++

		// 检查是否需要扩容
		if h.size >= h.upperTol*h.capacity && h.capIdx+1 < len(h.capacitys) {
			h.capIdx++
			h.resize(h.capacitys[h.capIdx])
		}
	}
}

func (h *HashTable[K, V]) Remove(key K) V {
	var zero V
	bucket := h.hash(key)
	prev := &MapNode[K, V]{next: h.buckets[bucket]}
	cur := prev.next

	for cur != nil {
		if fmt.Sprintf("%v", cur.key) == fmt.Sprintf("%v", key) {
			prev.next = cur.next
			h.size--

			// 检查是否需要缩容
			if h.size < h.lowerTol*h.capacity && h.capIdx-1 >= 0 {
				h.capIdx--
				h.resize(h.capacitys[h.capIdx])
			}

			return cur.value
		}
		prev = cur
		cur = cur.next
	}

	return zero
}

func (h *HashTable[K, V]) Set(key K, value V) {
	bucket := h.hash(key)
	node := h.getNode(h.buckets[bucket], key)
	if node == nil {
		panic(fmt.Sprintf("Key '%v' doesn't exist", key))
	}
	node.value = value
}

func (h *HashTable[K, V]) Contains(key K) bool {
	bucket := h.hash(key)
	return h.getNode(h.buckets[bucket], key) != nil
}

func (h *HashTable[K, V]) Get(key K) V {
	var zero V
	bucket := h.hash(key)
	node := h.getNode(h.buckets[bucket], key)
	if node == nil {
		return zero
	}
	return node.value
}

func (h *HashTable[K, V]) getNode(node *MapNode[K, V], key K) *MapNode[K, V] {
	for node != nil {
		if fmt.Sprintf("%v", node.key) == fmt.Sprintf("%v", key) {
			return node
		}
		node = node.next
	}
	return nil
}

func (h *HashTable[K, V]) resize(newCapacity int) {
	newBuckets := make([]*MapNode[K, V], newCapacity)
	oldCapacity := h.capacity
	oldBuckets := h.buckets

	h.capacity = newCapacity
	h.buckets = newBuckets

	// 重新哈希所有元素
	for i := 0; i < oldCapacity; i++ {
		node := oldBuckets[i]
		for node != nil {
			newBucket := h.hash(node.key)
			newNode := &MapNode[K, V]{
				key:   node.key,
				value: node.value,
				next:  h.buckets[newBucket],
			}
			h.buckets[newBucket] = newNode
			node = node.next
		}
	}
}

func (h *HashTable[K, V]) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("HashTable(size=%d, capacity=%d, {", h.size, h.capacity))

	first := true
	for _, bucket := range h.buckets {
		for node := bucket; node != nil; node = node.next {
			if !first {
				builder.WriteString(", ")
			}
			builder.WriteString(fmt.Sprintf("%v:%v", node.key, node.value))
			first = false
		}
	}

	builder.WriteString("})")
	return builder.String()
}

func main() {
	fmt.Println("=== Go语言哈希表示例 ===")

	ht := NewHashTable[string, int]()

	// 添加键值对
	fmt.Println("--- 添加键值对 ---")
	ht.Add("one", 1)
	ht.Add("two", 2)
	ht.Add("three", 3)
	ht.Add("four", 4)
	ht.Add("five", 5)
	fmt.Printf("哈希表: %s\n", ht.String())
	fmt.Printf("大小: %d, 容量: %d\n", ht.Size(), ht.capacity)

	// 获取值
	fmt.Println("\n--- 获取值 ---")
	fmt.Printf("Get('two'): %d\n", ht.Get("two"))
	fmt.Printf("Get('six'): %d\n", ht.Get("six"))

	// 检查键
	fmt.Println("\n--- 检查键 ---")
	fmt.Printf("Contains('one'): %t\n", ht.Contains("one"))
	fmt.Printf("Contains('six'): %t\n", ht.Contains("six"))

	// 更新值
	fmt.Println("\n--- 更新值 ---")
	ht.Set("two", 22)
	fmt.Printf("更新后: %s\n", ht.String())

	// 删除键值对
	fmt.Println("\n--- 删除键值对 ---")
	ht.Remove("three")
	fmt.Printf("删除后: %s\n", ht.String())

	fmt.Println("\n=== 示例程序结束 ===")
}
