package datastructure

import (
	"hash/crc32"
)

// 使用素数扩容
// 参考 https://planetmath.org/goodhashtableprimes
var (
	capacity []int = []int{53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593,
49157, 98317, 196613, 393241, 786433, 1572869, 3145739, 6291469,
12582917, 25165843, 50331653, 100663319, 201326611, 402653189, 805306457, 1610612741}
	upperTol = 10
	lowerTol = 2
	capacityIndex = 0
)

type hashData struct {
	size int
	hashTable map[int]map[string]interface{}
	M int
}

func Hash() *hashData {
	hash := &hashData{
		size: 0,
		M: capacity[capacityIndex],
		hashTable: map[int]map[string]interface{}{},
	}
	for i := 0; i < capacity[0]; i++ {
		hash.hashTable[i] = map[string]interface{}{}
	}
	return hash
}

func (h *hashData) Add(key string, value interface{}) {
	m := h.hashTable[h.hashCode(key)]
	m[key] = value
	h.size++

	if h.size >= upperTol * h.M && capacityIndex+1 < len(capacity) {
		capacityIndex++
		h.resize(capacity[capacityIndex])
	}
}

func (h *hashData) Remove(key string) interface{} {
	tmpM := h.hashTable[h.hashCode(key)]
	if _, ok := tmpM[key]; !ok {
		return nil
	}
	h.size--
	if h.size < lowerTol * h.M && capacityIndex - 1 >= 0 {
		capacityIndex--
		h.resize(capacity[capacityIndex])
	}
	return tmpM[key]
}

func (h *hashData) Set(key string, value interface{}) {
	tmpM := h.hashTable[h.hashCode(key)]
	if _, ok := tmpM[key]; !ok {
		panic("The item " + key + " is not found.")
	}
	tmpM[key] = value
}

func (h *hashData) Get(key string) interface{} {
	tmpM := h.hashTable[h.hashCode(key)]
	if _, ok := tmpM[key]; !ok {
		return nil
	}
	return tmpM[key]
}

func (h *hashData) Contains(key string) bool {
	tmpM := h.hashTable[h.hashCode(key)]
	if _, ok := tmpM[key]; !ok {
		return false
	}
	return true
}

func (h *hashData) hashCode(data string) int {
	v := int(crc32.ChecksumIEEE([]byte(data)))
	if v < 0 {
		v = -v
	}
	return (v & 0x7fffffff) % h.M
}

func (h *hashData) resize(capacityNum int) {
	newHashTable := map[int]map[string]interface{}{}
	for i := 0; i < capacityNum; i++ {
		newHashTable[i] = map[string]interface{}{}
	}

	oldM := h.M
	h.M = capacityNum
	for i := 0; i < oldM; i++ {
		tmpM := h.hashTable[i]
		for k, v := range tmpM {
			newHashTable[h.hashCode(k)][k] = v
		}
	}
	h.hashTable = newHashTable
}
