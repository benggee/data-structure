package main

import (
	"fmt"
)

// UnionFind 并查集接口（最优版本：路径压缩 + 按秩合并）
type UnionFind struct {
	parent []int
	rank   []int
}

// NewUnionFind 创建一个新的并查集
func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFind{
		parent: parent,
		rank:   rank,
	}
}

func (uf *UnionFind) Size() int {
	return len(uf.parent)
}

// IsConnected 检查两个元素是否连接
func (uf *UnionFind) IsConnected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

// UnionElements 合并两个元素
func (uf *UnionFind) UnionElements(p, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)

	if pRoot == qRoot {
		return
	}

	// 按秩合并 - 尽量保持树的深度一致
	if uf.rank[pRoot] > uf.rank[qRoot] {
		uf.parent[qRoot] = pRoot
	} else if uf.rank[pRoot] < uf.rank[qRoot] {
		uf.parent[pRoot] = qRoot
	} else {
		uf.parent[pRoot] = qRoot
		uf.rank[qRoot]++
	}
}

// find 查找根节点，带路径压缩
func (uf *UnionFind) find(p int) int {
	if p < 0 || p >= len(uf.parent) {
		panic("Index out of range")
	}

	// 路径压缩
	if p != uf.parent[p] {
		uf.parent[p] = uf.find(uf.parent[p])
	}
	return uf.parent[p]
}

// CountSets 统计集合数量
func (uf *UnionFind) CountSets() int {
	roots := make(map[int]bool)
	for i := 0; i < len(uf.parent); i++ {
		roots[uf.find(i)] = true
	}
	return len(roots)
}

func (uf *UnionFind) String() string {
	return fmt.Sprintf("UnionFind(elements=%d, sets=%d)", len(uf.parent), uf.CountSets())
}

// UnionFindV1 第一版并查集：基于数组，时间复杂度O(n)
type UnionFindV1 struct {
	data [][]int
}

// NewUnionFindV1 创建第一版并查集
func NewUnionFindV1(size int) *UnionFindV1 {
	data := make([][]int, size)
	for i := 0; i < size; i++ {
		data[i] = []int{i}
	}
	return &UnionFindV1{data: data}
}

func (uf *UnionFindV1) IsConnected(p, q int) bool {
	return uf.findId(p) == uf.findId(q)
}

func (uf *UnionFindV1) findId(p int) int {
	for i, s := range uf.data {
		for _, v := range s {
			if v == p {
				return i
			}
		}
	}
	return -1
}

func (uf *UnionFindV1) UnionElements(p, q int) {
	pId := uf.findId(p)
	qId := uf.findId(q)

	if pId == qId {
		return
	}

	// 将q的集合合并到p的集合
	uf.data[pId] = append(uf.data[pId], uf.data[qId]...)
	// 删除q的集合
	uf.data = append(uf.data[:qId], uf.data[qId+1:]...)
}

// UnionFindV2 第二版并查集：Quick Find
type UnionFindV2 struct {
	id []int
}

// NewUnionFindV2 创建第二版并查集
func NewUnionFindV2(size int) *UnionFindV2 {
	id := make([]int, size)
	for i := 0; i < size; i++ {
		id[i] = i
	}
	return &UnionFindV2{id: id}
}

func (uf *UnionFindV2) IsConnected(p, q int) bool {
	return uf.id[p] == uf.id[q]
}

func (uf *UnionFindV2) find(p int) int {
	if p < 0 || p >= len(uf.id) {
		panic("Index out of range")
	}
	return uf.id[p]
}

func (uf *UnionFindV2) UnionElements(p, q int) {
	pId := uf.find(p)
	qId := uf.find(q)

	if pId == qId {
		return
	}

	// 将所有id[p]的元素改为id[q]
	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pId {
			uf.id[i] = qId
		}
	}
}

// UnionFindV3 第三版并查集：Quick Union
type UnionFindV3 struct {
	parent []int
}

// NewUnionFindV3 创建第三版并查集
func NewUnionFindV3(size int) *UnionFindV3 {
	parent := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	return &UnionFindV3{parent: parent}
}

func (uf *UnionFindV3) IsConnected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UnionFindV3) find(p int) int {
	if p < 0 || p >= len(uf.parent) {
		panic("Index out of range")
	}
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p
}

func (uf *UnionFindV3) UnionElements(p, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)

	if pRoot == qRoot {
		return
	}

	uf.parent[pRoot] = qRoot
}

// UnionFindV4 第四版并查集：基于size的优化
type UnionFindV4 struct {
	parent []int
	size   []int
}

// NewUnionFindV4 创建第四版并查集
func NewUnionFindV4(size int) *UnionFindV4 {
	parent := make([]int, size)
	sz := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		sz[i] = 1
	}
	return &UnionFindV4{
		parent: parent,
		size:   sz,
	}
}

func (uf *UnionFindV4) IsConnected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UnionFindV4) find(p int) int {
	if p < 0 || p >= len(uf.parent) {
		panic("Index out of range")
	}
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p
}

func (uf *UnionFindV4) UnionElements(p, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)

	if pRoot == qRoot {
		return
	}

	// 基于size的优化，将小树连接到大树
	if uf.size[pRoot] < uf.size[qRoot] {
		uf.parent[pRoot] = qRoot
		uf.size[qRoot] += uf.size[pRoot]
	} else {
		uf.parent[qRoot] = pRoot
		uf.size[pRoot] += uf.size[qRoot]
	}
}

// UnionFindV5 第五版并查集：路径压缩优化
type UnionFindV5 struct {
	parent []int
	rank   []int
}

// NewUnionFindV5 创建第五版并查集
func NewUnionFindV5(size int) *UnionFindV5 {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFindV5{
		parent: parent,
		rank:   rank,
	}
}

func (uf *UnionFindV5) IsConnected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UnionFindV5) find(p int) int {
	if p < 0 || p >= len(uf.parent) {
		panic("Index out of range")
	}

	// 路径压缩
	if p != uf.parent[p] {
		uf.parent[p] = uf.find(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf *UnionFindV5) UnionElements(p, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)

	if pRoot == qRoot {
		return
	}

	if uf.rank[pRoot] < uf.rank[qRoot] {
		uf.parent[pRoot] = qRoot
	} else if uf.rank[pRoot] > uf.rank[qRoot] {
		uf.parent[qRoot] = pRoot
	} else {
		uf.parent[pRoot] = qRoot
		uf.rank[qRoot]++
	}
}

func main() {
	fmt.Println("=== Go语言并查集示例 ===")

	// 测试最优版本
	fmt.Println("\n--- UnionFind V6 (路径压缩 + 按秩合并) ---")
	uf := NewUnionFind(10)
	fmt.Println(uf)

	uf.UnionElements(0, 1)
	uf.UnionElements(2, 3)
	uf.UnionElements(1, 2)
	uf.UnionElements(4, 5)
	uf.UnionElements(6, 7)

	fmt.Printf("合并后:\n")
	fmt.Printf("IsConnected(0, 3): %t\n", uf.IsConnected(0, 3))
	fmt.Printf("IsConnected(0, 4): %t\n", uf.IsConnected(0, 4))
	fmt.Printf("CountSets: %d\n", uf.CountSets())

	// 测试V2版本
	fmt.Println("\n--- UnionFind V2 (Quick Find) ---")
	uf2 := NewUnionFindV2(10)
	uf2.UnionElements(0, 1)
	uf2.UnionElements(2, 3)
	uf2.UnionElements(1, 2)
	fmt.Printf("IsConnected(0, 3): %t\n", uf2.IsConnected(0, 3))

	// 测试V3版本
	fmt.Println("\n--- UnionFind V3 (Quick Union) ---")
	uf3 := NewUnionFindV3(10)
	uf3.UnionElements(0, 1)
	uf3.UnionElements(2, 3)
	uf3.UnionElements(1, 2)
	fmt.Printf("IsConnected(0, 3): %t\n", uf3.IsConnected(0, 3))

	fmt.Println("\n=== 示例程序结束 ===")
}
