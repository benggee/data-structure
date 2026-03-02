# Python AVL树（自平衡二分搜索树）实现

## 目录
- [什么是AVL树](#什么是avl树)
- [为什么需要AVL树](#为什么需要avl树)
- [核心概念与原理](#核心概念与原理)
- [代码实现详解](#代码实现详解)
- [复杂度分析](#复杂度分析)
- [应用场景](#应用场景)

---

## 什么是AVL树

AVL树（Adelson-Velsky and Landis Tree）是一种**自平衡二分搜索树**，它是第一个发明的平衡树。

**核心特性：**
1. 任何节点的两个子树的高度差的绝对值不超过1
2. 是二分搜索树，满足BST的所有性质
3. 通过旋转操作保持平衡

**示例：**
```
平衡的AVL树：
        5        高度差：|2-2| = 0 ≤ 1 ✓
       / \
      3   7
     / \   \
    2   4   8

不平衡的BST：
    5          高度差：|0-3| = 3 > 1 ✗
     \
      7
       \
        8
         \
          9
```

---

## 为什么需要AVL树

### BST的问题

```
插入有序序列到BST：

插入 1, 2, 3, 4, 5：

1:
    1

2:
    1
     \
      2

3:
    1
     \
      2
       \
        3

最终退化为链表！查找、插入、删除都变成 O(n)
```

**问题总结：**
- BST最坏情况下退化为链表
- 所有操作复杂度从 O(log n) 退化为 O(n)
- 数据的插入顺序严重影响树的结构

### AVL树的解决方案

```
AVL树自动平衡：

插入 1, 2, 3, 4, 5 到 AVL树：

1:
    1

2:
    1
     \
      2
    高度差 = 1，还OK

3:
    1
     \
      2
       \
        3
    高度差 = 2，触发旋转！

旋转后：
      2
     / \
    1   3
    高度差 = 1，平衡 ✓

继续插入 4, 5，自动保持平衡
```

**AVL树保证：**
- ✅ 树的高度始终是 O(log n)
- ✅ 所有操作都是 O(log n)
- ✅ 不会退化为链表

---

## 核心概念与原理

### 1. 平衡因子（Balance Factor）

```
平衡因子 = 左子树高度 - 右子树高度

对于任意节点：
- 平衡因子为 0：左右子树高度相等
- 平衡因子为 1：左子树比右子树高1
- 平衡因子为 -1：右子树比左子树高1
- 平衡因子的绝对值 > 1：不平衡！
```

**示例：**
```
        5        平衡因子 = 2 - 2 = 0
       / \
      3   7     3的BF = 1 - 1 = 0
     / \   \    7的BF = 0 - 1 = -1
    2   4   8   2,4,8的BF = 0
```

### 2. 四种不平衡情况

#### LL情况（左左）

```
        y          x
       / \        / \
      x   T4  => z   y
     / \        / \ / \
    z   T3     T1 T2 T3 T4
   / \
  T1 T2

原因：在y的左子树的左子树插入节点导致
解决：右旋转
```

#### RR情况（右右）

```
    y               x
   / \             / \
  T1  x      =>   y   z
     / \         / \ / \
    T2  z       T1 T2 T3 T4
       / \
      T3 T4

原因：在y的右子树的右子树插入节点导致
解决：左旋转
```

#### LR情况（左右）

```
      y          y          x
     / \        / \        / \
    x   T4  => z   T4  => z   y
   / \        / \        / \ / \
  T1  z      x   T3     T1 T2 T3 T4
     / \    / \
    T2 T3  T1 T2

原因：在y的左子树的右子树插入节点导致
解决：先左旋转左子节点，再右旋转
```

#### RL情况（右左）

```
  y            y                x
 / \          / \             / \
T1  x    =>  T1  z      =>   y   z
   / \            / \       / \ / \
  z  T4          T2  x     T1 T2 T3 T4
 / \                / \
T2  T3             T3  T4

原因：在y的右子树的左子树插入节点导致
解决：先右旋转右子节点，再左旋转
```

### 3. 旋转操作

#### 右旋转

```python
def right_rotate(self, y: AVLNode) -> AVLNode:
    """对节点y进行右旋转

          y                   x
        /   \               /   \
       x     T4    ==>     z     y
      / \                  / \   / \
     z   T3              T1 T2 T3 T4
    / \
   T1 T2
    """
    x = y.left
    T3 = x.right

    # 旋转
    x.right = y
    y.left = T3

    # 更新高度（先更新y，再更新x）
    y.height = 1 + max(self.get_height(y.left), self.get_height(y.right))
    x.height = 1 + max(self.get_height(x.left), self.get_height(x.right))

    return x  # 新的根节点
```

#### 左旋转

```python
def left_rotate(self, y: AVLNode) -> AVLNode:
    """对节点y进行左旋转

      y                        x
    /   \                   /   \
   T1    x      ==>        y     z
        / \               / \   / \
       T2  z            T1 T2 T3  T4
          / \
         T3 T4
    """
    x = y.right
    T2 = x.left

    # 旋转
    x.left = y
    y.right = T2

    # 更新高度
    y.height = 1 + max(self.get_height(y.left), self.get_height(y.right))
    x.height = 1 + max(self.get_height(x.left), self.get_height(x.right))

    return x  # 新的根节点
```

---

## 代码实现详解

### 1. 节点定义

```python
class AVLNode(Generic[K, V]):
    """AVL树节点"""
    def __init__(self, key: K, value: V):
        self.key: K = key
        self.value: V = value
        self.left: Optional['AVLNode[K, V]'] = None
        self.right: Optional['AVLNode[K, V]'] = None
        self.height: int = 1  # 新节点高度为1
```

**关键点：**
- AVL树通常实现为Map结构（key-value对）
- 每个节点额外存储`height`信息
- 用于快速计算平衡因子

### 2. 添加元素（核心）

```python
def set(self, key: K, value: V) -> None:
    """添加/更新键值对"""
    self._root = self._add(self._root, key, value)

def _add(self, node: Optional['AVLNode[K, V]'],
         key: K, value: V) -> 'AVLNode[K, V]':
    """递归添加节点并保持平衡"""
    # 基准情况：创建新节点
    if node is None:
        self._size += 1
        return AVLNode[K, V](key, value)

    # 递归插入
    if key < node.key:
        node.left = self._add(node.left, key, value)
    elif key > node.key:
        node.right = self._add(node.right, key, value)
    else:  # key已存在，更新value
        node.value = value
        return node

    # ===== 平衡维护 =====
    # 更新高度
    node.height = 1 + max(self.get_height(node.left),
                          self.get_height(node.right))

    # 计算平衡因子
    balance_factor = self.get_balance_factor(node)

    # LL情况：左子树更高，且左子节点的平衡因子 >= 0
    if balance_factor > 1 and self.get_balance_factor(node.left) >= 0:
        return self.right_rotate(node)

    # RR情况：右子树更高，且右子节点的平衡因子 <= 0
    if balance_factor < -1 and self.get_balance_factor(node.right) <= 0:
        return self.left_rotate(node)

    # LR情况：左子树更高，但左子节点的平衡因子 < 0
    if balance_factor > 1 and self.get_balance_factor(node.left) < 0:
        node.left = self.left_rotate(node.left)  # 先左旋左子节点
        return self.right_rotate(node)              # 再右旋

    # RL情况：右子树更高，但右子节点的平衡因子 > 0
    if balance_factor < -1 and self.get_balance_factor(node.right) > 0:
        node.right = self.right_rotate(node.right)  # 先右旋右子节点
        return self.left_rotate(node)                # 再左旋

    return node
```

**添加过程示例：**

```
向空树插入 10, 20, 30：

插入10：
    10

插入20：
    10
     \
      20
    BF = -1，还OK

插入30：
    10
     \
      20
       \
        30
    BF(20) = -1, BF(10) = -2，不平衡！

    这是RR情况，左旋转10：
      20
     /  \
   10    30
```

### 3. 删除元素

```python
def remove(self, key: K) -> Optional[V]:
    """删除键值对"""
    node = self.get_node(self._root, key)
    if node is None:
        return None
    old_value = node.value
    self._root = self._remove(self._root, key)
    return old_value

def _remove(self, node: Optional['AVLNode[K, V]'],
            key: K) -> Optional['AVLNode[K, V]']:
    """递归删除节点并保持平衡"""
    if node is None:
        return None

    # 查找要删除的节点
    if key < node.key:
        node.left = self._remove(node.left, key)
        ret_node = node
    elif key > node.key:
        node.right = self._remove(node.right, key)
        ret_node = node
    else:  # 找到要删除的节点
        # 情况1 & 2：只有一个子节点或没有子节点
        if node.left is None:
            self._size -= 1
            return node.right
        if node.right is None:
            self._size -= 1
            return node.left

        # 情况3：有两个子节点
        # 找到后继节点（右子树最小）
        successor = self._minimum(node.right)
        successor.right = self._remove_min(node.right)
        successor.left = node.left
        node.left = node.right = None
        ret_node = successor

    if ret_node is None:
        return None

    # ===== 平衡维护（与添加相同）=====
    # 更新高度
    ret_node.height = 1 + max(self.get_height(ret_node.left),
                              self.get_height(ret_node.right))

    # 计算平衡因子
    balance_factor = self.get_balance_factor(ret_node)

    # LL情况
    if balance_factor > 1 and self.get_balance_factor(ret_node.left) >= 0:
        return self.right_rotate(ret_node)

    # RR情况
    if balance_factor < -1 and self.get_balance_factor(ret_node.right) <= 0:
        return self.left_rotate(ret_node)

    # LR情况
    if balance_factor > 1 and self.get_balance_factor(ret_node.left) < 0:
        ret_node.left = self.left_rotate(ret_node.left)
        return self.right_rotate(ret_node)

    # RL情况
    if balance_factor < -1 and self.get_balance_factor(ret_node.right) > 0:
        ret_node.right = self.right_rotate(ret_node.right)
        return self.left_rotate(ret_node)

    return ret_node
```

### 4. 辅助方法

```python
def get_height(self, node: Optional['AVLNode[K, V]']) -> int:
    """获取节点高度"""
    if node is None:
        return 0
    return node.height

def get_balance_factor(self, node: Optional['AVLNode[K, V]']) -> int:
    """计算平衡因子"""
    if node is None:
        return 0
    return self.get_height(node.left) - self.get_height(node.right)
```

---

## 复杂度分析

### 时间复杂度

| 操作 | 时间复杂度 | 说明 |
|------|-----------|------|
| 添加 | O(log n) | 需要O(log n)查找 + O(1)旋转 |
| 删除 | O(log n) | 需要O(log n)查找 + O(1)旋转 |
| 查找 | O(log n) | 树的高度始终是log n |
| 更新 | O(log n) | 查找 + 更新 |
| 获取最小/最大 | O(log n) | 沿着最左/最右路径 |

**为什么是 O(log n)？**

```
AVL树的重要性质：
- 高度 h <= 1.44 * log₂(n+2)
- 对于n个节点的AVL树，高度大约是 log₂n

因此所有操作都是 O(log n)
```

### 空间复杂度

| 项目 | 空间复杂度 | 说明 |
|------|-----------|------|
| 节点存储 | O(n) | n个节点 |
| 递归栈 | O(log n) | 递归深度等于树高 |
| 额外信息 | O(n) | 每个节点存储height |

---

## 应用场景

### 1. 适合使用AVL树的场景

```python
# 场景1：需要有序存储且频繁查询
avl = AVLTree[str, int]()
avl.set("apple", 5)
avl.set("banana", 3)
avl.set("cherry", 8)
# O(log n) 查找

# 场景2：数据可能使普通BST退化
# 插入有序数据时，AVL树自动平衡
avl = AVLTree[int, int]()
for i in range(1000):
    avl.set(i, i)
# 依然保持平衡，所有操作 O(log n)

# 场景3：需要范围查询
# 中序遍历得到有序序列
```

### 2. AVL树 vs 红黑树

| 特性 | AVL树 | 红黑树 |
|------|-------|--------|
| 平衡性 | 更严格（|BF|≤1） | 相对宽松 |
| 查找 | 更快 | 稍慢 |
| 插入/删除 | 更多旋转 | 较少旋转 |
| 应用场景 | 读多写少 | 通用（STL map/set） |

---

## 完整示例

```python
# 示例1：基本操作
avl = AVLTree[str, int]()
avl["one"] = 1
avl["two"] = 2
avl["three"] = 3
avl["four"] = 4
avl["five"] = 5

print(f"大小: {avl.size()}")
print(f"获取 'two': {avl.get('two')}")
print(f"包含 'six': {avl.contains('six')}")
print(f"是否是BST: {avl.is_bst()}")
print(f"是否是AVL: {avl.is_avl()}")

# 示例2：自动平衡
avl = AVLTree[int, int]()
# 插入有序序列
for i in range(1, 11):
    avl.set(i, i * 10)
print(f"插入1-10后是AVL: {avl.is_avl()}")

# 示例3：删除
del avl["three"]
print(f"删除后: {avl.keys()}")
print(f"删除后依然是AVL: {avl.is_avl()}")

# 示例4：遍历
print("所有键:", avl.keys())
print("所有值:", avl.values())
print("所有键值对:", avl.items())
```

---

## 总结

AVL树是第一个自平衡二分搜索树：

**优点：**
- ✅ 保证 O(log n) 的所有操作
- ✅ 不会退化
- ✅ 查找效率高（比红黑树更平衡）
- ✅ 实现相对清晰

**缺点：**
- ❌ 插入/删除可能需要多次旋转
- ❌ 需要额外存储高度信息
- ❌ 实现复杂度高于普通BST

**关键要点：**
1. 理解平衡因子的概念
2. 掌握四种不平衡情况（LL/RR/LR/RL）
3. 熟练掌握左旋和右旋操作
4. 理解AVL树如何保证 O(log n) 的高度

**下一步学习：**
- 红黑树（更常用的平衡树）
- B树/B+树（多路平衡树，用于数据库）
- Skip List（跳跃表，概率平衡结构）

---

## 代码实现链接

- [Python完整实现](../python/avl.py)
- [Golang完整实现](../golang/AVLTree/avl_tree.go)
- [Java完整实现](../java/AVLTree)
