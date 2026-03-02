# Python 哈希表实现

## 目录
- [什么是哈希表](#什么是哈希表)
- [为什么需要哈希表](#为什么需要哈希表)
- [核心概念与原理](#核心概念与原理)
- [代码实现详解](#代码实现详解)
- [复杂度分析](#复杂度分析)
- [应用场景](#应用场景)

---

## 什么是哈希表

哈希表（Hash Table）是一种通过**哈希函数**将键映射到存储位置的数据结构，实现了**快速的键值对存储和检索**。

**核心思想：**
```
键(key) → 哈希函数 → 哈希值(hash) → 数组索引(index) → 值(value)

示例：
"apple" → hash("apple") = 163 → index = 163 % 16 = 3 → bucket[3]
```

**示例结构：**
```
哈希表（容量16）：

Index  Bucket
  0    [ ("cat", 5) ]
  1    []
  2    [ ("dog", 8), ("frog", 2) ]  ← 冲突！两个键映射到同一位置
  3    []
  ...
 15    [ ("fish", 10) ]
```

---

## 为什么需要哈希表

### 现有数据结构的局限

```python
# 数组：查找快但需要连续索引
arr = [None] * 1000
arr[123] = "value"  # O(1)
# 但不能用字符串作为索引

# 链表：可以用任意键但查找慢
# 需要遍历整个链表：O(n)

# BST：O(log n)查找
# 但对于大多数应用，可以更快！
```

### 哈希表的优势

```python
# 哈希表：平均O(1)的查找、插入、删除
ht = HashTable[str, int]()
ht.set("apple", 5)
ht.set("banana", 3)
value = ht.get("apple")  # O(1) 平均情况！
```

**性能对比：**

| 数据结构 | 查找 | 插入 | 删除 | 有序 |
|---------|------|------|------|------|
| 数组 | O(1) | O(n) | O(n) | ✅ |
| 链表 | O(n) | O(1) | O(n) | ❌ |
| BST | O(log n) | O(log n) | O(log n) | ✅ |
| 哈希表 | O(1) | O(1) | O(1) | ❌ |

---

## 核心概念与原理

### 1. 哈希函数

**理想的哈希函数：**
- 确定性：相同输入产生相同输出
- 快速计算：O(1)时间
- 均匀分布：减少冲突
- 全域覆盖：充分利用所有索引

**常见哈希函数：**

```python
# 简单的取模哈希
def simple_hash(key: str, capacity: int) -> int:
    hash_value = 0
    for char in key:
        hash_value = (hash_value * 31 + ord(char)) % capacity
    return hash_value

# Python内置hash()
def python_hash(key: str, capacity: int) -> int:
    return (hash(key) & 0x7fffffff) % capacity
```

**哈希冲突：**

```
不同键映射到同一位置：

"cat"  → hash = 5 → index = 5 % 16 = 5
"tac"  → hash = 21 → index = 21 % 16 = 5  ← 冲突！

冲突不可避免，但可以妥善处理
```

### 2. 冲突解决方法

#### 方法1：链地址法（Separate Chaining）

```
每个bucket维护一个链表：

Index  Bucket
  0    [ ("a", 1) → ("b", 2) ]
  1    [ ("c", 3) ]
  2    []
  ...

优点：
- 简单易实现
- 可以处理任意数量的冲突
- 删除操作简单

缺点：
- 需要额外指针空间
- 性能可能退化到O(n)
```

#### 方法2：开放地址法（Open Addressing）

```
冲突时寻找下一个空位：

线性探测：
index = (hash + i) % capacity

二次探测：
index = (hash + i²) % capacity

优点：
- 无需额外指针
- 更好的缓存性能

缺点：
- 删除复杂
- 容易聚集（clustering）
```

**我们的实现使用链地址法 + AVL树：**

```
每个bucket是一个有序映射（AVL树或dict）：

Index  Bucket
  0    { "a": 1, "b": 2 }  (当冲突多时用树，少时用链表)
  1    { "c": 3 }
  2    {}
  ...
```

### 3. 动态扩容

```
扩容过程：

初始状态：
capacity = 16
size = 12
load_factor = 12/16 = 0.75

添加新元素后：
size = 13
load_factor = 13/16 = 0.8125 > upperTol (0.75)

触发扩容：
1. 创建新数组，capacity = 32 (下一个素数)
2. 重新哈希所有元素到新数组
3. 释放旧数组

扩容后：
capacity = 32
size = 13
load_factor = 13/32 = 0.406
```

**为什么使用素数作为容量？**

```
素数能更好地分散哈希值，减少冲突

示例：
容量 = 16 (2⁴)
  hash(x) % 16 只看最后4位
  hash("abc") = ...1011  → 11
  hash("xyz") = ...0011  → 11  ← 容易冲突

容量 = 17 (素数)
  hash(x) % 17 看所有位
  更均匀分布
```

---

## 代码实现详解

### 1. 基础结构

```python
class HashTable(Generic[K, V]):
    """哈希表实现（链地址法）"""

    # 素数容量表
    _CAPACITY = [53, 97, 193, 389, 769, 1543, 3079, 6151,
                 12289, 24593, 49157, 98317, 196613, 393241,
                 786433, 1572869, 3145739, 6291469, 12582917,
                 25165843, 50331653, 100663319, 201326611,
                 402653189, 805306457, 1610612741]

    def __init__(self, upper_tol: int = 10, lower_tol: int = 2):
        """
        Args:
            upper_tol: 上界容忍度（平均每个bucket的元素数）
            lower_tol: 下界容忍度
        """
        self._capacity_index: int = 0
        self._M: int = self._CAPACITY[self._capacity_index]
        self._size: int = 0
        self._upper_tol: int = upper_tol
        self._lower_tol: int = lower_tol
        # 每个位置是一个dict（相当于TreeMap）
        self._hashtable: List[dict] = [{} for _ in range(self._M)]
```

**设计要点：**
- 使用素数表作为可选容量
- 容量自动从素数表中选择
- 每个bucket用Python dict（已高度优化）

### 2. 哈希函数

```python
def _hash(self, key: K) -> int:
    """计算键的哈希值

    处理步骤：
    1. 调用Python内置hash()获取哈希值
    2. 清除符号位（确保非负）
    3. 对容量取模
    """
    return (hash(key) & 0x7fffffff) % self._M
```

**为什么 `& 0x7fffffff`？**

```
Python的hash()可能返回负数：

hash("abc") = -123456789

& 0x7fffffff 清除符号位：
-123456789 & 0x7fffffff = 123456789（正数）

确保结果始终是非负数，可以作为数组索引
```

### 3. 添加元素

```python
def add(self, key: K, value: V) -> None:
    """添加键值对"""
    map_dict = self._hashtable[self._hash(key)]

    if key in map_dict:
        # 键已存在，更新值
        map_dict[key] = value
    else:
        # 新键，添加并增加大小
        map_dict[key] = value
        self._size += 1

        # 检查是否需要扩容
        if self._size >= self._upper_tol * self._M and \
           self._capacity_index + 1 < len(self._CAPACITY):
            self._capacity_index += 1
            self._resize(self._CAPACITY[self._capacity_index])
```

**扩容判断逻辑：**

```
条件解释：
- size >= upper_tol * M
  平均每个bucket超过upper_tol个元素

- capacity_index + 1 < len(CAPACITY)
  还有更大的素数可用

示例：
M = 53, upper_tol = 10
size = 530 → 530 >= 10*53 ✓ 触发扩容
新M = 97
```

### 4. 删除元素

```python
def remove(self, key: K) -> Optional[V]:
    """删除键值对"""
    map_dict = self._hashtable[self._hash(key)]

    if key in map_dict:
        value = map_dict.pop(key)
        self._size -= 1

        # 检查是否需要缩容
        if self._size < self._lower_tol * self._M and \
           self._capacity_index - 1 >= 0:
            self._capacity_index -= 1
            self._resize(self._CAPACITY[self._capacity_index])

        return value

    return None
```

**缩容判断逻辑：**

```
条件解释：
- size < lower_tol * M
  平均每个bucket少于lower_tol个元素

- capacity_index - 1 >= 0
  还有更小的素数可用

避免频繁扩容/缩容（震荡）
```

### 5. 扩容实现

```python
def _resize(self, new_M: int) -> None:
    """调整哈希表容量

    过程：
    1. 创建新的哈希表数组
    2. 将所有元素重新哈希到新数组
    3. 替换旧数组
    """
    new_hashtable: List[dict] = [{} for _ in range(new_M)]
    old_M = self._M

    # 更新容量（必须在重新哈希前）
    self._M = new_M

    # 重新哈希所有元素
    for i in range(old_M):
        map_dict = self._hashtable[i]
        for key, value in map_dict.items():
            # 用新容量计算哈希
            new_hashtable[self._hash(key)][key] = value

    self._hashtable = new_hashtable
```

**重新哈希过程图解：**

```
原始哈希表（M=8）：
Index: 0  1  2  3  4  5  6  7
Data: a  -  b  c  -  d  e  -

扩容到M=16：

新表（空）：
Index: 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
Data: -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -

重新哈希：
hash(a) % 16 → 新位置
hash(b) % 16 → 新位置
...

所有元素重新分布！
```

---

## 复杂度分析

### 时间复杂度

| 操作 | 平均情况 | 最坏情况 | 说明 |
|------|---------|---------|------|
| 添加 | O(1) | O(n) | n个元素都在一个bucket |
| 删除 | O(1) | O(n) | 同上 |
| 查找 | O(1) | O(n) | 同上 |
| 扩容 | O(n) | O(n) | 需要重新哈希所有元素 |

**为什么平均是 O(1)？**

```
假设：
- M = 容量（bucket数量）
- N = 元素数量
- load_factor = N / M

理想情况下，每个bucket约有 load_factor 个元素
链地址法查找：遍历链表，平均 load_factor 步

如果 load_factor 是常数（如 0.75），则平均 O(1)
```

### 空间复杂度

| 项目 | 空间复杂度 | 说明 |
|------|-----------|------|
| 基本存储 | O(n) | n个键值对 |
| 额外空间 | O(M) | M个bucket |
| 扩容临时空间 | O(n) | 扩容时需要 |

**装载因子（Load Factor）：**

```
load_factor = N / M

典型值：
- Python dict: ~0.66
- Java HashMap: 0.75
- 我们的实现: 2-10（可配置）

load_factor越大，空间利用率越高，但冲突越多
load_factor越小，冲突越少，但浪费空间
```

---

## 应用场景

### 1. 适合使用哈希表的场景

```python
# 场景1：缓存实现
cache = HashTable[str, Any]()
cache.set("user:123", user_data)
user = cache.get("user:123")  # O(1) 快速访问

# 场景2：计数/统计
counter = HashTable[str, int]()
for word in text.split():
    if counter.contains(word):
        counter.set(word, counter.get(word) + 1)
    else:
        counter.set(word, 1)

# 场景3：去重
seen = HashTable[int, bool]()
unique = []
for num in nums:
    if not seen.contains(num):
        seen.set(num, True)
        unique.append(num)

# 场景4：索引/映射
index = HashTable[str, int]()
for i, word in enumerate(words):
    index.set(word, i)
```

### 2. 不适合的场景

```python
# 场景1：需要有序遍历
# 哈希表不保持顺序
# 使用 TreeMap/BST 代替

# 场景2：需要范围查询
# 不能高效查询 "所有键在 [a, b] 之间的元素"
# 使用 BST 代替

# 场景3：键不可哈希
# list、dict 不能作为键
# 需要转换为可哈希类型
```

---

## 完整示例

```python
# 示例1：基本操作
ht = HashTable[str, int]()
ht["one"] = 1
ht["two"] = 2
ht["three"] = 3
ht["four"] = 4
ht["five"] = 5

print(f"大小: {ht.size()}")
print(f"容量: {ht.capacity()}")
print(f"获取 'two': {ht.get('two')}")
print(f"包含 'six': {ht.contains('six')}")

# 示例2：自动扩容
print(f"\n初始: 大小={ht.size()}, 容量={ht.capacity()}")
for i in range(100):
    ht[f"key_{i}"] = i
print(f"添加100个后: 大小={ht.size()}, 容量={ht.capacity()}")

# 示例3：字典式操作
del ht["three"]
ht["two"] = 22  # 更新
print(f"\n操作后: {ht.keys()}")

# 示例4：遍历
print("所有键:", ht.keys())
print("所有值:", ht.values())
print("所有键值对:", ht.items())
```

---

## 哈希表 vs 其他结构

| 特性 | 哈希表 | BST/AVL | 数组 |
|------|--------|---------|------|
| 查找 | O(1) | O(log n) | O(1) |
| 插入 | O(1) | O(log n) | O(n) |
| 删除 | O(1) | O(log n) | O(n) |
| 有序 | ❌ | ✅ | ✅ |
| 范围查询 | ❌ | ✅ | ❌ |
| 空间效率 | 中等 | 较低 | 高 |

---

## 总结

哈希表是最重要的数据结构之一：

**优点：**
- ✅ 平均 O(1) 的查找、插入、删除
- ✅ 键可以是任意可哈希类型
- ✅ 灵活的动态扩容
- ✅ 实际应用中性能极佳

**缺点：**
- ❌ 最坏情况 O(n)
- ❌ 不保持顺序
- ❌ 不支持范围查询
- ❌ 需要良好的哈希函数

**关键要点：**
1. 理解哈希函数的设计要求
2. 掌握冲突解决方法（链地址法）
3. 理解装载因子对性能的影响
4. 了解动态扩容的时机和代价

**实际应用：**
- Python dict / set
- Java HashMap / HashSet
- C++ std::unordered_map
- 数据库索引
- 缓存系统
- 符号表

---

## 代码实现链接

- [Python完整实现](../python/hashtable.py)
- [Golang完整实现](../golang/HashTable/hashtable.go)
- [Java完整实现](../java/Hash)
