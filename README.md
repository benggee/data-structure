# 数据结构与算法 - 多语言实现

> Java、C、C++、Golang、Python 五种编程语言的完整实现

## 📚 项目简介

数据结构与算法是计算机科学的基础，但学好它并不容易。本项目提供了：
- **完整实现**：五种主流编程语言的数据结构
- **详细文档**：每个数据结构都有原理讲解和代码分析
- **对比学习**：同一数据结构在不同语言的实现对比
- **实践导向**：包含可运行的代码示例

**配套资源：**
- 公众号："程序员班吉" - 系列文章讲解
- LeetCode题单：[点击查看](DOCS/leetcode.md)

---

## 🚀 快速开始

### Java实现（参考标准）

```bash
cd java/Array/src
javac Main.java
java Main
```

### Python实现

```bash
cd python
python array.py
```

### Golang实现

```bash
cd golang/Array
go run array.go
```

### C++实现

```bash
cd c++/array
g++ main.cpp -o array
./array
```

---

## 📊 完整实现对比

### 线性数据结构

| 数据结构 | Java | C++ | C | Golang | Python |
|:--------|:----:|:--:|:-:|:------:|:------:|
| **动态数组** | `Array/` | `array/ArrayList.hpp` | - | `Array/` | `array.py` |
| **链表** | `LinkList/` | `linkList/` | `list.c` | `LinkList/` | `linklist.py` |
| **栈** | `Stack/` | `stack/` | `stack.c` | `Stack/` | `stack.py` |
| **队列** | `Queue/` | `queue/` | `queue.c` | `Queue/` | `queue.py` |
| **集合** | `Set/` | `set/` | `set.c` | `Set/` | `set.py` |
| **映射** | `Map/` | `map/` | - | `Map/` | `map.py` |
| **哈希表** | `Hash/` | `hash/` | `chtbl.c` | `HashTable/` | `hashtable.py` |

### 树形数据结构

| 数据结构 | Java | C++ | C | Golang | Python |
|:--------|:----:|:--:|:-:|:------:|:------:|
| **二分搜索树** | `BinarySearchTree/` | `tree/` | `bitree.c` | - | `bst.py` |
| **AVL树** | `AVLTree/` | `tree/avl_tree.h` | - | `AVLTree/` | `avl.py` |
| **红黑树** | `RBTree/` | `tree/rbtree.hpp` | - | `RBTree/` | `rbtree.py` ⭐ |
| **堆** | `Heap/` | `heap/` | `heap.c` | `Heap/` | `heap.py` |
| **线段树** | `SegmentTree/` | `segmentTree/` | - | `SegmentTree/` | `segmenttree.py` |
| **字典树** | `Trie/` | `trie/` | - | `Trie/` | `trie.py` |
| **并查集** | `UnionFind/` | `unionFind/` | - | `UnionFind/` | `unionfind.py` |

### 图论算法

| 数据结构 | Java | C++ | C | Golang | Python |
|:--------|:----:|:--:|:-:|:------:|:------:|
| **无权图** | `Graph/` | `graph/` | `graph.c` | - | - |
| **有权图** | `WeightGraph/` | - | - | - | - |
| **有向图** | `DirectionGraph/` | - | - | - | - |

### 排序算法

| 算法 | Java | C++ | C | Golang | Python |
|:----:|:----:|:--:|:-:|:------:|:------:|
| **选择排序** | - | `sort/selectionsort.h` | - | - | - |
| **冒泡排序** | - | `sort/bubblesort.h` | - | - | - |
| **插入排序** | - | `sort/insertsort.h` | - | - | - |
| **希尔排序** | - | `sort/shellsort.h` | - | - | - |
| **归并排序** | - | `sort/mergesort.h` | - | - | - |
| **快速排序** | - | `sort/quicksort.h` | - | - | - |

**图例说明：**
- ✅ 完整实现
- ⚠️ 部分实现
- ❌ 缺失

---

## 📁 目录结构

```
data-structure/
├── java/           # Java参考实现（最完整）
├── python/         # Python实现（14个核心结构）
├── golang/         # Golang实现（13个核心结构）
├── c++/            # C++实现（包含图论+排序）
├── c/              # C语言实现（基础结构）
├── DOCS/           # 详细文档
└── README.md       # 本文件
```

---

## 📖 各语言快速导航

### 🇵🇦 Python实现

```python
# 使用示例
from python import Array, AVLTree

# 动态数组
arr = Array[int]()
for i in range(10):
    arr.add_last(i)

# AVL树
avl = AVLTree[str, int]()
avl["apple"] = 5
avl.contains("apple")  # True
```

**文件列表：**
- [array.py](python/array.py) - 动态数组
- [linklist.py](python/linklist.py) - 链表
- [stack.py](python/stack.py) - 栈
- [queue.py](python/queue.py) - 队列
- [set.py](python/set.py) - 集合
- [map.py](python/map.py) - 映射
- [bst.py](python/bst.py) - 二分搜索树
- [avl.py](python/avl.py) - AVL树
- [rbtree.py](python/rbtree.py) - **红黑树**
- [heap.py](python/heap.py) - 堆
- [hashtable.py](python/hashtable.py) - 哈希表
- [trie.py](python/trie.py) - 字典树
- [unionfind.py](python/unionfind.py) - 并查集
- [segmenttree.py](python/segmenttree.py) - 线段树

**详细文档：** [DOCS/python/](DOCS/python/)

### 🇮🇪 Golang实现

```go
// 使用示例
package main

import "golang/Array"

func main() {
    arr := NewArrayDefault[int]()
    arr.AddLast(1)
    arr.AddLast(2)
    fmt.Println(arr.String())
}
```

**目录结构：**
- [Array/array.go](golang/Array/array.go) - 动态数组
- [LinkList/](golang/LinkList/) - 链表
- [Stack/](golang/Stack/) - 栈
- [Queue/](golang/Queue/) - 队列
- [Set/](golang/Set/) - 集合
- [Map/](golang/Map/) - 映射
- [AVLTree/](golang/AVLTree/) - AVL树
- [RBTree/](golang/RBTree/) - **红黑树**
- [Heap/](golang/Heap/) - 堆
- [HashTable/](golang/HashTable/) - 哈希表
- [Trie/](golang/Trie/) - 字典树
- [UnionFind/](golang/UnionFind/) - 并查集
- [SegmentTree/](golang/SegmentTree/) - 线段树

**详细文档：** [DOCS/golang/](DOCS/golang/)

### ☕ C++实现

```cpp
// 使用示例
#include "array/ArrayList.hpp"
#include "tree/avl_tree.h"

int main() {
    ArrayList<int> arr;
    arr.add(42);
    std::cout << arr.getSize() << std::endl;
    return 0;
}
```

**目录结构：**
- [array/ArrayList.hpp](c++/array/ArrayList.hpp) - 动态数组
- [tree/avl_tree.h](c++/tree/avl_tree.h) - AVL树
- [hash/](c++/hash/) - 哈希表
- [heap/](c++/heap/) - 堆
- [graph/](c++/graph/) - 图论
- [sort/](c++/sort/) - 8种排序算法

**详细文档：** [DOCS/cpp/](DOCS/cpp/)

### 🔧 C实现

```c
// 使用示例
#include "list.h"
#include <stdio.h>

int main() {
    List list;
    list_init(&list, free);

    int *data = malloc(sizeof(int));
    *data = 42;
    list_ins_next(&list, NULL, data);

    printf("Size: %d\n", list_size(&list));
    list_destroy(&list);
    return 0;
}
```

**文件列表：**
- [list.c](c/list.c), [list.h](c/list.h) - 链表
- [stack.c](c/stack.c), [stack.h](c/stack.h) - 栈
- [queue.c](c/queue.c), [queue.h](c/queue.h) - 队列
- [heap.c](c/heap.c), [heap.h](c/heap.h) - 堆
- [chtbl.c](c/chtbl.c), [chtbl.h](c/chtbl.h) - 哈希表
- [bistree.c](c/bistree.c), [bistree.h](c/bistree.h) - AVL树
- [graph.c](c/graph.c), [graph.h](c/graph.h) - 图

**详细文档：** [DOCS/c/](DOCS/c/)

---

## 🔗 线性数据结构

### 动态数组 (Array)

**特点：** O(1)随机访问，自动扩容/缩容

| 语言 | 实现文件 | 扩容策略 |
|------|---------|---------|
| Java | `Array/src/Array.java` | 满时翻倍 |
| Python | `python/array.py` | 满时翻倍 |
| Golang | `golang/Array/array.go` | 满时翻倍 |
| C++ | `c++/array/ArrayList.hpp` | 满时翻倍 |

```java
// Java示例
Array<Integer> arr = new Array<>();
for (int i = 0; i < 100; i++) {
    arr.addLast(i);  // 自动扩容
}
```

### 链表 (LinkList)

**特点：** O(1)插入删除，动态大小

| 语言 | 实现文件 | 特点 |
|------|---------|------|
| Java | `LinkList/src/LinkList.java` | 虚拟头节点 |
| Python | `python/linklist.py` | 类型提示 |
| Golang | `golang/LinkList/linklist.go` | 泛型支持 |
| C | `c/list.c` | void*泛型 |

### 栈 (Stack) & 队列 (Queue)

**栈 - LIFO (后进先出)：**
| 语言 | 文件 |
|------|------|
| Java | `Stack/src/ArrayStack.java` |
| Python | `python/stack.py` |
| Golang | `golang/Stack/stack.go` |
| C++ | `c++/stack/` |
| C | `c/stack.c` |

**队列 - FIFO (先进先出)：**
| 语言 | 文件 | 实现方式 |
|------|------|---------|
| Java | `Queue/src/ArrayQueue.java` | 数组队列 |
| Java | `Queue/src/LoopQueue.java` | 循环队列 |
| Python | `python/queue.py` | 3种实现 |
| Golang | `golang/Queue/queue.go` | 数组+循环 |
| C++ | `c++/queue/` | 数组+链表 |
| C | `c/queue.c` | 链表 |

---

## 🌳 树形数据结构

### 二分搜索树 (BST)

**特点：** 中序遍历有序，O(log n)查找

```java
// Java示例
BinarySearchTree<Integer> bst = new BinarySearchTree<>();
bst.add(5);
bst.add(3);
bst.contains(5);  // true
```

**实现文件：**
- Java: `BinarySearchTree/src/BinarySearchTree.java`
- Python: `python/bst.py`
- Golang: `golang/AVLTree/binary_search_tree.go`
- C++: `c++/tree/`
- C: `c/bitree.c`, `c/bistree.c`

### AVL树 vs 红黑树

| 特性 | AVL树 | 红黑树 |
|------|-------|--------|
| **平衡性** | 更严格 (|BF|≤1) | 相对宽松 |
| **查找** | 更快 | 稍慢 |
| **插入/删除** | 更多旋转 | 更少旋转 |
| **实际应用** | - | C++ std::map, Java TreeMap |

**实现文件：**
| 语言 | AVL树 | 红黑树 |
|------|-------|--------|
| Java | `AVLTree/` | `RBTree/` |
| Python | `python/avl.py` | `python/rbtree.py` |
| Golang | `golang/AVLTree/` | `golang/RBTree/` |
| C++ | `c++/tree/avl_tree.h` | `c++/tree/rbtree.hpp` |

---

## 🗺️ 高级数据结构

### 哈希表 (Hash)

**冲突解决：** 链地址法

```python
# Python示例
ht = HashTable[str, int]()
ht["one"] = 1
ht["two"] = 2
value = ht.get("one")  # O(1) 平均
```

**实现文件：**
- Java: `Hash/src/HashTable.java`
- Python: `python/hashtable.py`
- Golang: `golang/HashTable/hashtable.go`
- C++: `c++/hash/`
- C: `c/chtbl.c`

### 并查集 (UnionFind)

**6个版本的演进：**
1. v1: 数组复制 - O(n)
2. v2: Quick Find - O(n)
3. v3: Quick Union - O(n)
4. v4: 基于Size优化
5. v5: 路径压缩
6. **v6: 路径压缩+按秩合并** - O(α(n))

```python
# Python示例
uf = UnionFind(10)
uf.unionElements(0, 1)  # 合并
uf.isConnected(0, 1)  # 检查连接
```

---

## 📚 详细文档

### Python数据结构详解
- [动态数组原理与实现](DOCS/python/array.md)
- [二分搜索树详解](DOCS/python/bst.md)
- [AVL树平衡原理](DOCS/python/avl.md)
- [哈希表详解](DOCS/python/hashtable.md)

### 多语言对比学习
- [Java vs 其他语言对比](DOCS/java-comparison.md)
- [语言特性对比](DOCS/language-comparison.md)
- [各语言实现特点](DOCS/golang/README.md)
- [C语言实现详解](DOCS/c/README.md)
- [C++实现详解](DOCS/cpp/README.md)

---

## 🎯 学习路径建议

### 初学者路径
```
1️⃣ Python（理解概念）
   ↓
2️⃣ Java（面向对象）
   ↓
3️⃣ Golang（现代实践）
   ↓
4️⃣ C++（性能优化）
```

### 面试准备路径
```
1. 掌握核心数据结构原理
2. 对比不同语言的实现
3. 理解时间/空间复杂度
4. LeetCode实战练习
```

### 语言选择指南

| 场景 | 推荐语言 | 理由 |
|------|---------|------|
| 快速开发/原型 | Python | 代码简洁，库丰富 |
| 企业级应用 | Java | 生态成熟，框架完善 |
| 云原生/微服务 | Golang | 性能好，并发强 |
| 高性能/游戏 | C++ | 性能最优 |
| 嵌入式/系统 | C | 完全控制 |

---

## 💡 使用技巧

### 1. 快速查找数据结构实现

**想要看动态数组？**
```bash
grep -r "class Array" java/ python/ golang/
```

**想要看栈的实现？**
```bash
grep -r "push\|pop" java/ python/ golang/
```

### 2. 运行示例代码

**Python:**
```bash
cd python
python array.py
```

**Golang:**
```bash
cd golang/Array
go run array.go
```

**Java:**
```bash
cd java/Array/src
javac Main.java
java Main
```

**C++:**
```bash
cd c++/array
g++ main.cpp -o main
./main
```

### 3. 对比阅读同一数据结构

想深入理解动态数组？阅读：
1. [java/Array/src/Array.java](java/Array/src/Array.java) - 参考实现
2. [python/array.py](python/array.py) - 简洁版本
3. [golang/Array/array.go](golang/Array/array.go) - Go风格
4. [c++/array/ArrayList.hpp](c++/array/ArrayList.hpp) - C++模板

---

## 📊 实现完整度总结

| 类别 | Java | Python | Golang | C++ | C |
|:----:|:----:|:------:|:------:|:---:|:-:|
| **线性结构** | 8/8 | 8/8 | 8/8 | 8/8 | 7/8 |
| **树形结构** | 8/8 | 8/8 | 8/8 | 8/8 | 2/8 |
| **高级结构** | 5/5 | 5/5 | 5/5 | 5/5 | 2/5 |
| **图论算法** | 15+ | 0 | 0 | 0 | 15+ |
| **排序算法** | 0 | 0 | 0 | 8 | 0 |

**核心数据结构（不含图论和排序）：**
- Python: 14/14 ✅ (100%)
- Golang: 13/14 ✅ (93%)
- C++: 13/14 ✅ (93%)
- Java: 14/14 ✅ (100%)

---

## 🤝 贡献指南

欢迎贡献代码！建议：

1. **遵循现有代码风格**
2. **添加测试用例**
3. **更新文档**
4. **保持多语言一致性**

---

## 📞 联系方式

- **公众号**: "程序员班吉"
- **文档**: [DOCS/](DOCS/) 目录
- **源码**: 各语言目录

---

**最后更新：** 2025年
**维护者：** 程序员班吉团队
**许可证：** MIT License
