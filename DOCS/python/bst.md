# Python 二分搜索树实现

## 目录
- [什么是二分搜索树](#什么是二分搜索树)
- [为什么需要二分搜索树](#为什么需要二分搜索树)
- [核心概念与原理](#核心概念与原理)
- [代码实现详解](#代码实现详解)
- [复杂度分析](#复杂度分析)
- [应用场景](#应用场景)

---

## 什么是二分搜索树

二分搜索树（Binary Search Tree，BST）是一种特殊的二叉树，它满足以下性质：

**BST 性质：**
- 对于任意节点：
  - 左子树中所有节点的值都小于该节点的值
  - 右子树中所有节点的值都大于该节点的值
  - 左右子树也分别是二分搜索树

**示例：**
```
        5
       / \
      3   7
     / \   \
    2   4   8

左子树所有值(2,3,4) < 5 < 右子树所有值(7,8)
```

---

## 为什么需要二分搜索树

### 数组的局限性

```python
# 有序数组的查找 - 二分查找 O(log n)
arr = [1, 2, 3, 4, 5, 6, 7, 8, 9]
# 但插入需要 O(n) - 需要移动元素
```

**数组的问题：**
- 查找快：O(log n) 使用二分查找
- 插入/删除慢：O(n) 需要移动元素

### BST 的优势

```python
bst = BinarySearchTree[int]()
bst.add(5)
bst.add(3)
bst.add(7)
# 查找、插入、删除都是 O(log n) 平均情况
```

**BST 的优点：**
- 查找快：O(log n)
- 插入快：O(log n)
- 删除快：O(log n)
- 中序遍历得到有序序列

---

## 核心概念与原理

### 1. 树的基本概念

```
        A        <- 根节点 (root)
       / \
      B   C     <- B、C 是 A 的子节点
     / \   \
    D   E   F   <- D、E 是 B 的子节点，F 是 C 的子节点

- 节点：包含值和指向左右子节点的指针
- 边：连接两个节点的线
- 叶子节点：没有子节点的节点（D、E、F）
- 高度：从根到叶子节点的最长路径
- 深度：从根到某节点的路径长度
```

### 2. BST 的查找过程

```
查找值 4：

        5         4 < 5，向左
       / \
      3   7      4 > 3，向右
     / \   \
    2   4   8    找到 4！

时间复杂度：O(log n) - 每次查找范围减半
```

### 3. BST 的插入过程

```
插入值 6：

        5         6 > 5，向右
       / \
      3   7      6 < 7，向左
     / \   \
    2   4   8    6 应该是 8 的左兄弟
                  找到位置，插入

结果：
        5
       / \
      3   7
     / \ / \
    2  4 6  8
```

### 4. BST 的删除过程（最复杂）

删除有三种情况：

**情况1：删除叶子节点**
```
        5              5
       / \            / \
      3   7    =>    3   7
     / \   \        /     \
    2   4   8      2       8
        /
       6          删除6（叶子）
```

**情况2：删除只有一个子节点的节点**
```
        5              5
       / \            / \
      3   7    =>    3   7
     / \   \        / \   \
    2   4   8      2   4   8
          \              /
           6            删除6，用子节点替代

删除6，直接用其唯一子节点替代
```

**情况3：删除有两个子节点的节点**
```
        5              5
       / \            / \
      3   7    =>    3   6
     / \ / \        / \ / \
    2  4 6  8      2  4 x  8

删除7，用右子树的最小值（6）替代

算法：
1. 找到节点右子树的最小节点（后继）
2. 用后继节点替代要删除的节点
3. 删除后继节点
```

---

## 代码实现详解

### 1. 节点定义

```python
class BSTNode(Generic[T]):
    """BST 节点类"""
    def __init__(self, e: T):
        self.e: T = e                    # 节点值
        self.left: Optional['BSTNode[T]'] = None   # 左子节点
        self.right: Optional['BSTNode[T]'] = None  # 右子节点
```

**设计要点：**
- 使用泛型支持不同类型
- 使用 `Optional` 表示可能为 None 的子节点
- 节点只存储值和左右指针，结构简单

### 2. 基础操作

#### 初始化

```python
class BinarySearchTree(Generic[T]):
    def __init__(self):
        self._root: Optional[BSTNode[T]] = None  # 根节点
        self._size: int = 0                       # 节点数量
```

#### 添加元素（核心算法）

```python
def add(self, e: T) -> None:
    """向BST添加元素 - 外部接口"""
    self._root = self._add(self._root, e)

def _add(self, node: Optional[BSTNode[T]], e: T) -> BSTNode[T]:
    """递归添加元素 - 内部实现

    递归逻辑：
    1. 基准情况：如果节点为空，创建新节点
    2. 如果 e < node.e，递归插入左子树
    3. 如果 e > node.e，递归插入右子树
    4. 如果 e == node.e，不插入（避免重复）
    """
    # 基准情况：到达空节点，创建新节点
    if node is None:
        self._size += 1
        return BSTNode[T](e)

    # 递归情况：根据值的大小决定插入位置
    if e < node.e:
        node.left = self._add(node.left, e)
    elif e > node.e:
        node.right = self._add(node.right, e)
    # 如果 e == node.e，什么都不做（不重复添加）

    return node
```

**递归过程图解：**

```
向空树添加 5：

_add(None, 5)
  -> node is None，创建节点5，返回
  -> root = Node(5)

添加 3：

_add(Node(5), 3)
  -> 3 < 5，递归
  -> _add(None, 3)
     -> node is None，创建节点3，返回
  -> Node(5).left = Node(3)
  -> 返回 Node(5)

        5
       /
      3

添加 7：

_add(Node(5), 7)
  -> 7 > 5，递归
  -> _add(None, 7)
     -> node is None，创建节点7，返回
  -> Node(5).right = Node(7)
  -> 返回 Node(5)

        5
       / \
      3   7
```

#### 查找元素

```python
def contains(self, e: T) -> bool:
    """查找元素是否存在"""
    return self._contains(self._root, e)

def _contains(self, node: Optional[BSTNode[T]], e: T) -> bool:
    """递归查找"""
    if node is None:
        return False

    if e == node.e:
        return True
    elif e < node.e:
        return self._contains(node.left, e)
    else:
        return self._contains(node.right, e)
```

#### 删除元素（最复杂）

```python
def remove(self, e: T) -> None:
    """删除元素"""
    self._root = self._remove(self._root, e)

def _remove(self, node: Optional[BSTNode[T]], e: T) -> Optional[BSTNode[T]]:
    """递归删除节点

    三种情况：
    1. 删除叶子节点
    2. 删除只有一个子节点的节点
    3. 删除有两个子节点的节点
    """
    if node is None:
        return None

    if e < node.e:
        # 在左子树中删除
        node.left = self._remove(node.left, e)
        return node
    elif e > node.e:
        # 在右子树中删除
        node.right = self._remove(node.right, e)
        return node
    else:  # 找到要删除的节点
        # 情况1：左子树为空，返回右子树
        if node.left is None:
            right_node = node.right
            node.right = None
            self._size -= 1
            return right_node

        # 情况2：右子树为空，返回左子树
        if node.right is None:
            left_node = node.left
            node.left = None
            self._size -= 1
            return left_node

        # 情况3：左右子树都不为空
        # 找到右子树的最小节点（后继）
        successor = self._minimum(node.right)
        # 用后继替代当前节点
        successor.right = self._remove_min(node.right)
        successor.left = node.left

        # 断开原节点连接
        node.left = node.right = None

        return successor
```

**删除过程详解：**

```
删除有两个子节点的节点：

初始树：
        5
       / \
      3   7
         / \
        6   8

删除 5：

步骤1：找到 5 的右子树最小节点（6）
        7
       /
      6

步骤2：用 6 替代 5
        6
       / \
      3   7
           \
            8

步骤3：删除原 6 节点
        6
       / \
      3   7
         / \
        x   8
```

### 3. 遍历操作

#### 中序遍历（最常用）

```python
def in_order(self, visit_func: Optional[Callable[[T], None]] = None) -> None:
    """中序遍历：左 -> 根 -> 右

    重要性质：BST 的中序遍历结果是升序序列！
    """
    self._in_order(self._root, visit_func)

def _in_order(self, node: Optional[BSTNode[T]],
              visit_func: Optional[Callable[[T], None]] = None) -> None:
    if node is None:
        return
    self._in_order(node.left, visit_func)  # 先遍历左子树
    if visit_func:
        visit_func(node.e)                # 再访问根节点
    else:
        print(node.e)
    self._in_order(node.right, visit_func) # 最后遍历右子树
```

**中序遍历图解：**

```
        5
       / \
      3   7
     / \   \
    2   4   8

中序遍历顺序：
1. _in_order(5)
2.   _in_order(3)
3.     _in_order(2)
4.       visit 2      -> 输出: 2
5.     visit 3        -> 输出: 3
6.     _in_order(4)
7.       visit 4      -> 输出: 4
8.   visit 5          -> 输出: 5
9.   _in_order(7)
10.    _in_order(None)  # 左子树为空
11.    visit 7          -> 输出: 7
12.    _in_order(8)
13.      visit 8        -> 输出: 8

最终输出：2 3 4 5 7 8（升序！）
```

#### 前序遍历

```python
def pre_order(self, visit_func: Optional[Callable[[T], None]] = None) -> None:
    """前序遍历：根 -> 左 -> 右

    用途：复制整棵树
    """
    self._pre_order(self._root, visit_func)

def _pre_order(self, node: Optional[BSTNode[T]],
               visit_func: Optional[Callable[[T], None]] = None) -> None:
    if node is None:
        return
    if visit_func:
        visit_func(node.e)                # 先访问根节点
    else:
        print(node.e)
    self._pre_order(node.left, visit_func)  # 再遍历左子树
    self._pre_order(node.right, visit_func) # 最后遍历右子树
```

#### 后序遍历

```python
def post_order(self, visit_func: Optional[Callable[[T], None]] = None) -> None:
    """后序遍历：左 -> 右 -> 根

    用途：释放内存、计算表达式
    """
    self._post_order(self._root, visit_func)

def _post_order(self, node: Optional[BSTNode[T]],
                visit_func: Optional[Callable[[T], None]] = None) -> None:
    if node is None:
        return
    self._post_order(node.left, visit_func)  # 先遍历左子树
    self._post_order(node.right, visit_func) # 再遍历右子树
    if visit_func:
        visit_func(node.e)                  # 最后访问根节点
    else:
        print(node.e)
```

#### 层序遍历（BFS）

```python
from collections import deque

def level_order(self, visit_func: Optional[Callable[[T], None]] = None) -> None:
    """层序遍历（广度优先遍历）"""
    if self._root is None:
        return

    queue = deque([self._root])
    while queue:
        cur = queue.popleft()
        if visit_func:
            visit_func(cur.e)
        else:
            print(cur.e)
        if cur.left:
            queue.append(cur.left)
        if cur.right:
            queue.append(cur.right)
```

**层序遍历图解：**

```
        5          队列: [5]
       / \
      3   7        队列: [3, 7]
     / \   \
    2   4   8      队列: [2, 4, 8]

输出顺序：5 3 7 2 4 8
```

---

## 复杂度分析

### 时间复杂度

| 操作 | 平均情况 | 最坏情况 | 说明 |
|------|---------|---------|------|
| 添加 | O(log n) | O(n) | 树平衡时为 log n |
| 删除 | O(log n) | O(n) | 需要先查找 |
| 查找 | O(log n) | O(n) | 每次比较减半范围 |
| 最小值 | O(log n) | O(n) | 一直向左走 |
| 最大值 | O(log n) | O(n) | 一直向右走 |

**为什么最坏情况是 O(n)？**

```
退化为链表的 BST：

    5
     \
      6
       \
        7
         \
          8

在这种情况下，BST 退化为链表，所有操作都是 O(n)
```

### 空间复杂度

| 操作 | 空间复杂度 | 说明 |
|------|-----------|------|
| 存储节点 | O(n) | n 个节点 |
| 递归栈 | O(h) | h 是树的高度 |
| 最坏递归 | O(n) | 树退化为链表时 |

---

## 应用场景

### 1. 适合使用 BST 的场景

```python
# 场景1：需要动态维护有序集合
bst = BinarySearchTree[int]()
bst.add(5)
bst.add(3)
bst.add(7)
print(bst.to_list())  # [3, 5, 7] - 自动有序

# 场景2：需要快速查找、插入、删除
score_bst = BinarySearchTree[int]()
score_bst.add(85)
score_bst.add(92)
if score_bst.contains(85):
    print("找到了")

# 场景3：需要范围查询
# 找到所有在 [low, high] 之间的元素
```

### 2. 不适合的场景

```python
# 场景1：数据基本不变
# 使用排序数组更高效

# 场景2：需要频繁访问第 k 大元素
# 需要维护额外信息

# 场景3：数据可能使 BST 退化
# 使用平衡 BST（AVL、红黑树）
```

---

## 完整示例

```python
# 示例1：基本操作
bst = BinarySearchTree[int]()
data = [5, 3, 7, 2, 4, 6, 8]
for x in data:
    bst.add(x)

print(f"大小: {bst.size()}")
print(f"包含 4: {bst.contains(4)}")
print(f"中序遍历: {bst.to_list()}")  # [2, 3, 4, 5, 6, 7, 8]

# 示例2：删除操作
bst.remove(3)
print(f"删除 3 后: {bst.to_list()}")  # [2, 4, 5, 6, 7, 8]

# 示例3：遍历
print("前序遍历:")
bst.pre_order()  # 5 4 2 7 6 8

print("\n中序遍历:")
bst.in_order()  # 2 4 5 6 7 8

print("\n后序遍历:")
bst.post_order()  # 2 4 6 8 7 5

print("\n层序遍历:")
bst.level_order()  # 5 4 7 2 6 8
```

---

## 总结

二分搜索树是重要的基础数据结构：

**优点：**
- ✅ 查找、插入、删除都是 O(log n) 平均情况
- ✅ 中序遍历得到有序序列
- ✅ 实现相对简单
- ✅ 支持范围查询

**缺点：**
- ❌ 最坏情况退化为 O(n)
- ❌ 不保证平衡
- ❌ 需要额外空间存储指针

**关键要点：**
1. 理解 BST 的定义和性质
2. 掌握递归实现添加、删除、查找
3. 理解三种遍历方式的区别和用途
4. 认识到 BST 可能退化，需要平衡树

**下一步学习：**
- AVL 树（自平衡 BST）
- 红黑树（工业界常用的平衡 BST）
- B 树/B+ 树（数据库索引）

---

## 代码实现链接

- [Python完整实现](../python/bst.py)
- [Golang完整实现](../golang/AVLTree/binary_search_tree.go)
- [Java完整实现](../java/BinarySearchTree)
