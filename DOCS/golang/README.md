# Golang 数据结构实现文档

## 目录
- [Golang数据结构特性](#golang数据结构特性)
- [泛型支持](#泛型支持)
- [内存管理](#内存管理)
- [实现列表](#实现列表)
- [代码示例](#代码示例)

---

## Golang数据结构特性

### 1. 使用Go 1.18+ 泛型

Go 1.18引入了泛型支持，使我们能够编写类型安全且可复用的数据结构：

```go
// 泛型动态数组
type Array[E any] struct {
    arr  []E
    size int
}

// 约束：键必须可比较
type AVLTree[K comparable, V any] struct {
    root *Node[K, V]
    size int
}
```

**类型参数说明：**
- `E any` - 任意类型
- `K comparable` - 可比较类型（用于map的key）
- `V any` - 任意值类型

### 2. 与其他语言的对比

| 特性 | Go | Java | Python | C++ |
|------|-----|------|--------|-----|
| 泛型 | Go 1.18+ | ✅ | 类型提示 | 模板 |
| 内存管理 | GC | GC | GC | 手动 |
| 指针 | 显式 | 引用 | 无 | 显式 |
| 零值 | 有 | null | None | 未定义 |

---

## 核心数据结构实现

### 1. 动态数组 (Array)

**文件：** `golang/Array/array.go`

**核心特性：**
- 自动扩容/缩容
- 泛型支持任意类型
- 扩容时容量翻倍

```go
// 创建数组
arr := NewArrayDefault[int]()
arr.AddLast(1)
arr.AddLast(2)
arr.Add(1, 99)  // 在索引1插入99

// 访问元素
value := arr.Get(1)  // 返回99

// 删除元素
arr.Remove(1)
```

**扩容机制：**
```go
func (a *Array[E]) Add(index int, e E) {
    if index < 0 || index > a.size {
        panic("Add failed. Index out of range")
    }

    // 检查是否需要扩容
    if a.size == len(a.arr) {
        a.resize(2 * len(a.arr))  // 容量翻倍
    }

    // 移动元素并插入
    for i := a.size - 1; i >= index; i-- {
        a.arr[i+1] = a.arr[i]
    }
    a.arr[index] = e
    a.size++
}
```

### 2. 链表 (LinkList)

**文件：** `golang/LinkList/linklist.go`

**核心特性：**
- 虚拟头节点简化操作
- 支持双向遍历（通过实现）
- 泛型节点

```go
type Node[E any] struct {
    e    E
    next *Node[E]
}

type LinkList[E any] struct {
    dummyHead *Node[E]  // 虚拟头节点
    size      int
}
```

**添加元素：**
```go
func (l *LinkList[E]) Add(index int, e E) {
    // 找到插入位置的前一个节点
    pre := l.dummyHead
    for i := 0; i < index; i++ {
        pre = pre.next
    }

    // 创建新节点
    node := newNode(e)
    node.next = pre.next
    pre.next = node
    l.size++
}
```

### 3. 栈 (Stack)

**文件：** `golang/Stack/stack.go`

**实现方式：**
- 基于动态数组的栈
- LIFO（后进先出）

```go
type ArrayStack[E any] struct {
    array *Array[E]
}

func (s *ArrayStack[E]) Push(e E) {
    s.array.AddLast(e)
}

func (s *ArrayStack[E]) Pop() E {
    return s.array.RemoveLast()
}

func (s *ArrayStack[E]) Peek() E {
    return s.array.Get(s.array.Size() - 1)
}
```

### 4. 队列 (Queue)

**文件：** `golang/Queue/queue.go`

**两种实现：**

**数组队列：**
```go
type ArrayQueue[E any] struct {
    array *Array[E]
}

// O(1) 入队
func (q *ArrayQueue[E]) Enqueue(e E) {
    q.array.AddLast(e)
}

// O(n) 出队（需要移动所有元素）
func (q *ArrayQueue[E]) Dequeue() E {
    return q.array.RemoveFirst()
}
```

**循环队列：**
```go
type LoopQueue[E any] struct {
    data  []E
    front int
    tail  int
    size  int
}

// O(1) 入队和出队
func (q *LoopQueue[E]) Enqueue(e E) {
    if (q.tail+1)%len(q.data) == q.front {
        q.resize(q.Capacity() * 2)
    }
    q.data[q.tail] = e
    q.tail = (q.tail + 1) % len(q.data)
    q.size++
}
```

### 5. AVL树

**文件：** `golang/AVLTree/avl_tree.go`

**核心特性：**
- 键值对存储（Map结构）
- 自动平衡
- 高度信息存储

```go
type Node[K comparable, V any] struct {
    key    K
    value  V
    left   *Node[K, V]
    right  *Node[K, V]
    height int
}
```

**旋转操作：**
```go
// 右旋转
func (t *AVLTree[K, V]) rightRotate(y *Node[K, V]) *Node[K, V] {
    x := y.left
    T3 := x.right

    // 执行旋转
    x.right = y
    y.left = T3

    // 更新高度
    y.height = t.max(t.getHeight(y.right), t.getHeight(y.left)) + 1
    x.height = t.max(t.getHeight(x.right), t.getHeight(x.left)) + 1

    return x
}
```

### 6. 哈希表

**文件：** `golang/HashTable/hashtable.go`

**实现特点：**
- 素数容量表
- 链地址法
- 动态扩容

```go
type HashTable[K any, V any] struct {
    buckets   []*MapNode[K, V]
    capacity  int
    size      int
    capacitys []int
}

// 哈希函数
func (h *HashTable[K, V]) hash(key K) int {
    keyStr := fmt.Sprintf("%v", key)
    hashValue := 0
    for _, c := range keyStr {
        hashValue = (hashValue << 5) + int(c)
    }
    return (hashValue & 0x7fffffff) % h.capacity
}
```

### 7. 字典树 (Trie)

**文件：** `golang/Trie/trie.go`

**用途：**
- 字符串搜索
- 前缀匹配
- 自动补全

```go
type TrieNode struct {
    isWord bool
    next   map[rune]*TrieNode
}

func (t *Trie) Add(word string) {
    cur := t.root
    for _, c := range word {
        if _, ok := cur.next[c]; !ok {
            cur.next[c] = newTrieNode()
        }
        cur = cur.next[c]
    }
    if !cur.isWord {
        cur.isWord = true
        t.size++
    }
}
```

### 8. 并查集 (UnionFind)

**文件：** `golang/UnionFind/unionfind.go`

**优化版本：**
- 路径压缩
- 按秩合并

```go
type UnionFind struct {
    parent []int
    rank   []int
}

// 查找根节点（带路径压缩）
func (uf *UnionFind) find(p int) int {
    if p != uf.parent[p] {
        uf.parent[p] = uf.find(uf.parent[p])
    }
    return uf.parent[p]
}

// 合并两个集合
func (uf *UnionFind) UnionElements(p, q int) {
    pRoot := uf.find(p)
    qRoot := uf.find(q)

    if pRoot == qRoot {
        return
    }

    // 按秩合并
    if uf.rank[pRoot] > uf.rank[qRoot] {
        uf.parent[qRoot] = pRoot
    } else if uf.rank[pRoot] < uf.rank[qRoot] {
        uf.parent[pRoot] = qRoot
    } else {
        uf.parent[pRoot] = qRoot
        uf.rank[qRoot]++
    }
}
```

### 9. 线段树 (SegmentTree)

**文件：** `golang/SegmentTree/segmenttree.go`

**用途：**
- 区间查询
- 区间更新
- O(log n)查询

```go
type SegmentTree[T any] struct {
    tree   []T
    data   []T
    merger Merger[T]
}

type Merger[T any] func(a, b T) T

// 查询区间 [queryL, queryR]
func (st *SegmentTree[T]) Query(queryL, queryR int) T {
    return st.query(0, 0, len(st.data)-1, queryL, queryR)
}
```

---

## 内存管理与性能

### 1. 值类型 vs 指针类型

```go
// 值类型 - 数据复制
type Array[E any] struct {
    arr  []E    // 切片本身是引用
    size int    // 值
}

// 指针类型 - 避免复制
func NewArray[E any](capacity int) *Array[E] {
    return &Array[E]{...}  // 返回指针
}
```

**最佳实践：**
- 结构较大时使用指针
- 需要修改原结构时使用指针
- 小结构可以传值

### 2. 零值利用

```go
// 利用零值简化代码
func (q *LoopQueue[E]) IsEmpty() bool {
    return q.front == q.tail  // 0 == 0
}

// 检查nil
if node == nil {
    return nil
}
```

### 3. Go特有的性能考虑

```go
// 使用切片而非链表（缓存友好）
type ArrayStack[E any] struct {
    array *Array[E]  // 连续内存
}

// 使用make预分配
data := make([]E, 0, capacity)  // 预分配容量

// 避免不必要的接口转换
type Comparable interface {
    Compare(other Comparable) int
}
```

---

## 完整示例

### 示例1：使用动态数组

```go
package main

import "fmt"

func main() {
    // 创建整型数组
    arr := NewArrayDefault[int]()

    // 添加元素
    for i := 0; i < 10; i++ {
        arr.AddLast(i)
    }

    // 在指定位置插入
    arr.Add(2, 100)

    // 获取和设置
    fmt.Printf("索引2的值: %d\n", arr.Get(2))
    arr.Set(2, 200)

    // 删除元素
    arr.Remove(2)

    // 转换为切片
    slice := arr.ToSlice()
    fmt.Printf("切片: %v\n", slice)
}
```

### 示例2：使用AVL树

```go
package main

import "fmt"

func main() {
    // 创建字符串->整数的映射
    avl := NewAVLTree[string, int]()

    // 添加键值对
    avl.Set("apple", 5)
    avl.Set("banana", 3)
    avl.Set("cherry", 8)

    // 查找
    if value, ok := avl.Get("banana"); ok {
        fmt.Printf("banana = %d\n", value)
    }

    // 检查是否存在
    fmt.Printf("包含'durian': %t\n", avl.Contains("durian"))

    // 验证平衡性
    fmt.Printf("是BST: %t\n", avl.IsBST())
    fmt.Printf("是AVL: %t\n", avl.IsAVL())

    // 删除
    avl.Remove("banana")
}
```

### 示例3：使用哈希表

```go
package main

import "fmt"

func main() {
    ht := NewHashTable[string, int]()

    // 添加
    ht.Add("one", 1)
    ht.Add("two", 2)
    ht.Add("three", 3)

    // 获取
    fmt.Printf("Get('two'): %d\n", ht.Get("two"))

    // 检查容量变化
    fmt.Printf("初始: size=%d, capacity=%d\n", ht.Size(), ht.capacity)

    // 添加更多元素触发扩容
    for i := 0; i < 100; i++ {
        ht.Add(fmt.Sprintf("key_%d", i), i)
    }
    fmt.Printf("扩容后: size=%d, capacity=%d\n", ht.Size(), ht.capacity)
}
```

---

## Go vs 其他语言实现对比

### 1. 类型系统

| 特性 | Go | Java | Python |
|------|-----|------|--------|
| 泛型约束 | `comparable`, `any` | `<T extends Comparable>` | `TypeVar`, `Protocol` |
| 类型推断 | 强 | 强 | 鸭子类型 |
| 空值 | `nil` | `null` | `None` |

### 2. 错误处理

```go
// Go: 显式错误处理
func (a *Array[E]) Get(index int) (E, error) {
    if index < 0 || index >= a.size {
        var zero E
        return zero, fmt.Errorf("index out of range")
    }
    return a.arr[index], nil
}

// 使用
value, err := arr.Get(5)
if err != nil {
    // 处理错误
}
```

### 3. 接口实现

```go
// Go: 隐式接口
type Stack interface {
    Size() int
    Push(e E)
    Pop() E
}

// ArrayStack自动实现Stack接口
type ArrayStack[E any] struct {
    array *Array[E]
}

// 只需实现方法，无需显式声明
func (s *ArrayStack[E]) Size() int {
    return s.array.Size()
}
```

---

## 总结

Go语言数据结构实现的特点：

**优点：**
- ✅ Go 1.18+ 泛型支持类型安全
- ✅ 简洁清晰的语法
- ✅ 高效的垃圾回收
- ✅ 原生支持并发（可用于并发数据结构）
- ✅ 零值机制简化代码

**注意事项：**
- ⚠️ 泛型类型约束需要理解
- ⚠️ 指针和值类型要区分清楚
- ⚠️ 错误处理是显式的
- ⚠️ 比较函数需要手动实现

**适用场景：**
- 云原生应用
- 微服务架构
- 高性能服务
- 并发系统

---

## 代码实现链接

- [Golang完整实现](../golang/)
- [Array](../golang/Array/array.go)
- [LinkList](../golang/LinkList/linklist.go)
- [AVLTree](../golang/AVLTree/avl_tree.go)
- [HashTable](../golang/HashTable/hashtable.go)
- [Trie](../golang/Trie/trie.go)
- [UnionFind](../golang/UnionFind/unionfind.go)
- [SegmentTree](../golang/SegmentTree/segmenttree.go)
