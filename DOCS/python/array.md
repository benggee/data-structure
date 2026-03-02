# Python 动态数组实现

## 目录
- [什么是动态数组](#什么是动态数组)
- [为什么需要动态数组](#为什么需要动态数组)
- [核心概念与原理](#核心概念与原理)
- [代码实现详解](#代码实现详解)
- [复杂度分析](#复杂度分析)
- [应用场景](#应用场景)

---

## 什么是动态数组

动态数组（Dynamic Array）是一种可以自动调整大小的数组数据结构。与普通数组不同，动态数组可以在元素数量超出当前容量时自动扩容，在元素数量减少时自动缩容。

**核心特性：**
- 支持随机访问：O(1) 时间访问任意索引位置的元素
- 自动扩容/缩容：根据元素数量动态调整内部存储空间
- 内存连续：元素在内存中连续存储，充分利用CPU缓存

---

## 为什么需要动态数组

### 传统数组的问题

```python
# 传统静态数组的局限性
arr = [1, 2, 3]  # 容量固定为3
# 如果想添加第4个元素，需要创建一个新数组，效率低下
```

**静态数组的缺点：**
1. 容量固定，必须预先知道所需大小
2. 扩容需要手动创建新数组并复制所有元素
3. 容易造成空间浪费或空间不足

### 动态数组的优势

```python
# 动态数组的优势
arr = Array[int]()
for i in range(100):
    arr.add_last(i)  # 自动扩容，无需关心容量
```

**动态数组的优点：**
1. 自动管理容量，使用方便
2. 扩容策略（通常翻倍）使均摊时间复杂度为 O(1)
3. 按需分配内存，减少浪费

---

## 核心概念与原理

### 1. 内部结构

```python
class Array[E]:
    def __init__(self, capacity: int = 10):
        self._data: List[Optional[T]] = [None] * capacity  # 内部数组
        self._size: int = 0                                 # 实际元素数量
```

**关键概念：**
- `capacity`（容量）：内部数组 `_data` 的实际大小
- `size`（大小）：实际存储的元素数量
- **容量 ≥ 大小**：总是有空间容纳元素

### 2. 扩容机制

```
初始状态：
data = [_, _, _, _, _, _, _, _, _, _]  (capacity = 10)
size = 0

添加元素后：
data = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
size = 10

触发扩容（添加第11个元素时）：
data = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, _, _, _, _, _, _, _, _, _, _]
     (capacity = 20，翻倍)
```

**扩容策略：**
1. 当 `size == capacity` 时触发扩容
2. 新容量 = 旧容量 × 2
3. 创建新数组，复制所有元素
4. 将新数组设为内部存储

### 3. 缩容机制

```
满载状态：
data = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
capacity = 10, size = 10

删除元素后：
data = [1, 2, 3, 4, 5]
capacity = 10, size = 5

触发缩容：
data = [1, 2, 3, 4, 5, _, _, _, _, _]
capacity = 5, size = 5
```

**缩容策略：**
1. 当 `size == capacity / 2` 时触发缩容
2. 新容量 = 旧容量 / 2
3. 避免频繁扩容/缩容，保持效率

---

## 代码实现详解

### 1. 基础操作

#### 初始化

```python
def __init__(self, capacity: int = 10):
    """初始化动态数组

    Args:
        capacity: 初始容量，默认10
    """
    self._data: List[Optional[T]] = [None] * capacity
    self._size: int = 0
```

**代码分析：**
- 使用 `List[Optional[T]]` 支持泛型
- 初始化时创建指定容量的数组，所有元素为 `None`
- `_size` 初始化为 0，表示还没有实际元素

#### 获取基本信息

```python
def capacity(self) -> int:
    """获取数组容量"""
    return len(self._data)

def size(self) -> int:
    """获取元素数量"""
    return self._size

def is_empty(self) -> bool:
    """检查数组是否为空"""
    return self._size == 0
```

### 2. 添加元素

#### 在指定位置插入

```python
def add(self, index: int, e: T) -> None:
    """在指定索引位置插入元素

    算法步骤：
    1. 检查索引是否合法
    2. 如果数组已满，先扩容
    3. 将 index 及之后的元素向后移动一位
    4. 在 index 位置放入新元素
    5. size + 1
    """
    # 步骤1：边界检查
    if index < 0 or index > self._size:
        raise IndexError(f"Add failed. Index must be >= 0 and <= size")

    # 步骤2：检查并扩容
    if self._size == len(self._data):
        self._resize(2 * len(self._data))

    # 步骤3：移动元素
    for i in range(self._size - 1, index - 1, -1):
        self._data[i + 1] = self._data[i]

    # 步骤4：插入新元素
    self._data[index] = e

    # 步骤5：更新大小
    self._size += 1
```

**图解插入过程：**

```
初始状态：[1, 2, 3, 4, _, _]
插入 index=2, value=99

步骤1：检查 index=2 合法 ✓
步骤2：容量充足，无需扩容
步骤3：移动元素
        [1, 2, 3, 4, _, _]  <- 原始
        [1, 2, _, 3, 4, _]  <- 4向后移
        [1, 2, _, _, 3, 4]  <- 3向后移
步骤4：插入元素
        [1, 2, 99, 3, 4, _]
步骤5：size = 5
```

**时间复杂度分析：**
- 最坏情况：O(n) - 需要移动所有元素
- 最好情况：O(1) - 在末尾插入
- 平均情况：O(n)

#### 在末尾添加（最常用）

```python
def add_last(self, e: T) -> None:
    """在数组末尾添加元素 - O(1)均摊时间复杂度"""
    self.add(self._size, e)
```

### 3. 删除元素

#### 删除指定位置元素

```python
def remove(self, index: int) -> T:
    """删除指定索引位置的元素

    算法步骤：
    1. 检查索引是否合法
    2. 保存要删除的元素
    3. 将 index+1 及之后的元素向前移动一位
    4. size - 1
    5. 检查是否需要缩容
    """
    # 步骤1：边界检查
    if index < 0 or index >= self._size:
        raise IndexError("Index out of range")

    # 步骤2：保存要删除的元素
    ret = self._data[index]

    # 步骤3：移动元素
    for i in range(index + 1, self._size):
        self._data[i - 1] = self._data[i]

    # 步骤4：更新大小
    self._size -= 1
    self._data[self._size] = None  # 清除引用

    # 步骤5：检查并缩容
    if self._size == len(self._data) // 2:
        self._resize(len(self._data) // 2)

    return ret
```

**图解删除过程：**

```
初始状态：[1, 2, 99, 3, 4, _]
删除 index=2

步骤1：检查 index=2 合法 ✓
步骤2：保存 ret = 99
步骤3：移动元素
        [1, 2, 99, 3, 4, _]  <- 原始
        [1, 2, 3, 3, 4, _]  <- 3向前移
        [1, 2, 3, 4, 4, _]  <- 4向前移
步骤4：size = 4，清除末尾
        [1, 2, 3, 4, None, _]
步骤5：size=4, capacity/2=3，不缩容
```

### 4. 扩容与缩容的实现

```python
def _resize(self, new_capacity: int) -> None:
    """调整数组容量

    过程：
    1. 创建新容量的数组
    2. 将旧数组的所有元素复制到新数组
    3. 将新数组设为内部存储
    """
    new_data = [None] * new_capacity
    for i in range(self._size):
        new_data[i] = self._data[i]
    self._data = new_data
```

**关键点：**
- 扩容时容量翻倍：`new_capacity = old_capacity * 2`
- 缩容时容量减半：`new_capacity = old_capacity / 2`
- 只复制实际使用的元素（`_size` 个），不复制整个 `_data`

---

## 复杂度分析

### 时间复杂度

| 操作 | 平均情况 | 最坏情况 | 均摊分析 |
|------|---------|---------|---------|
| `add_last(e)` | O(1) | O(n) | O(1) |
| `add_first(e)` | O(n) | O(n) | O(n) |
| `add(index, e)` | O(n) | O(n) | O(n) |
| `remove_last()` | O(1) | O(1) | O(1) |
| `remove(index)` | O(n) | O(n) | O(n) |
| `get(index)` | O(1) | O(1) | O(1) |
| `set(index, e)` | O(1) | O(1) | O(1) |
| `find(e)` | O(n) | O(n) | O(n) |

**为什么 `add_last` 的均摊复杂度是 O(1)？**

```
假设初始容量为1，连续添加n个元素：

第1次添加：容量1->2，复制1个元素
第2次添加：容量2->4，复制2个元素
第3次添加：直接添加
第4次添加：容量4->8，复制4个元素
第5-7次添加：直接添加
第8次添加：容量8->16，复制8个元素
...

总复制次数 = 1 + 2 + 4 + 8 + ... + n/2 = 2n - 2
均摊每次操作 = 2n/n = O(1)
```

### 空间复杂度

| 操作 | 空间复杂度 | 说明 |
|------|-----------|------|
| 初始化 | O(capacity) | 分配初始容量的空间 |
| 扩容 | O(n) | 需要额外的n空间临时存储 |
| 缩容 | O(n) | 需要额外的n空间临时存储 |
| 其他操作 | O(1) | 原地操作 |

---

## 应用场景

### 1. 适合使用动态数组的场景

```python
# 场景1：需要频繁随机访问
arr = Array[int]()
for i in range(1000):
    arr.add_last(i)
value = arr.get(500)  # O(1) 随机访问

# 场景2：元素数量变化较大
items = Array[str]()
while True:
    item = get_next_item()
    if item is None:
        break
    items.add_last(item)  # 自动扩容
```

### 2. 不适合的场景

```python
# 场景1：频繁在头部插入/删除
# 不推荐：动态数组
arr = Array[int]()
arr.add_first(1)  # O(n) - 需要移动所有元素

# 推荐：链表
ll = LinkList[int]()
ll.add_first(1)  # O(1)

# 场景2：频繁查找/删除特定值
arr.find(999)  # O(n) - 需要遍历
arr.remove_element(999)  # O(n) - 先查找，再移动
```

---

## 完整示例

```python
# 示例1：基本使用
arr = Array[int]()
print(f"初始: 容量={arr.capacity()}, 大小={arr.size()}")

# 添加元素
for i in range(10):
    arr.add_last(i)
print(f"添加10个元素后: {arr}")

# 插入元素
arr.add(2, 99)
print(f"在索引2插入99: {arr}")

# 删除元素
removed = arr.remove(2)
print(f"删除索引2的元素({removed}): {arr}")

# 示例2：自动扩容测试
arr = Array[int](2)
print(f"\n初始容量: {arr.capacity()}")
for i in range(10):
    arr.add_last(i)
    print(f"添加{i}: 容量={arr.capacity()}, 大小={arr.size()}")

# 示例3：不同类型
str_arr = Array[str]()
str_arr.add_last("Hello")
str_arr.add_last("World")
print(f"\n字符串数组: {str_arr}")
```

---

## 总结

动态数组是最基础也是最重要的数据结构之一：

**优点：**
- ✅ 随机访问 O(1)
- ✅ 内存连续，缓存友好
- ✅ 自动管理容量
- ✅ 尾部操作均摊 O(1)

**缺点：**
- ❌ 中间插入/删除 O(n)
- ❌ 扩容时有性能开销
- ❌ 可能浪费空间（缩容前）

**关键要点：**
1. 理解容量 vs 大小的区别
2. 掌握扩容/缩容策略（翻倍/减半）
3. 理解均摊时间复杂度的计算
4. 知道何时使用动态数组，何时使用其他结构

---

## 代码实现链接

- [Python完整实现](../python/array.py)
- [Golang完整实现](../golang/Array/array.go)
- [Java完整实现](../java/Array)
