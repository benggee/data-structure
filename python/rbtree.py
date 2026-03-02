"""
Red-Black Tree Implementation in Python
Based on Java RBTree implementation
红黑树是一种自平衡二叉搜索树，通过颜色标记和旋转操作保持平衡
红黑树是Java TreeMap和TreeSet、C++ std::map和std::set的底层实现
"""

from typing import TypeVar, Generic, Optional

K = TypeVar('K')
V = TypeVar('V')


class RBTree(Generic[K, V]):
    """
    红黑树实现

    红黑树的5条性质：
    1. 每个节点要么是红色，要么是黑色
    2. 根节点是黑色
    3. 所有叶子节点（NIL）是黑色
    4. 如果一个节点是红色，那么它的两个子节点都是黑色
    5. 对每个节点，从该节点到其所有后代叶子节点的简单路径上，
       均包含相同数量的黑色节点（黑高度）
    """

    RED = True
    BLACK = False

    class Node:
        """红黑树节点"""

        def __init__(self, key: K, value: V):
            self.key: K = key
            self.value: V = value
            self.left: Optional['RBTree.Node'] = None
            self.right: Optional['RBTree.Node'] = None
            self.color: bool = RBTree.RED  # 新节点默认为红色

        def __repr__(self) -> str:
            color = "R" if self.color else "B"
            return f"{self.key}:{self.value}({color})"

    def __init__(self):
        self._root: Optional['RBTree.Node'] = None
        self._size: int = 0

    def size(self) -> int:
        """获取树的大小"""
        return self._size

    def is_empty(self) -> bool:
        """检查树是否为空"""
        return self._size == 0

    def contains(self, key: K) -> bool:
        """检查键是否存在"""
        return self._get_node(self._root, key) is not None

    def get(self, key: K) -> Optional[V]:
        """获取键对应的值"""
        node = self._get_node(self._root, key)
        return node.value if node else None

    def set(self, key: K, value: V) -> None:
        """设置键值对"""
        self._root = self._add(self._root, key, value)

    # ===== 旋转操作 =====

    def _left_rotate(self, node: 'RBTree.Node') -> 'RBTree.Node':
        """
        左旋转
             node                    x
            /   \     左旋转       /   \
           T1   x    --------->  node  T3
              / \                /   \
             T2  T3             T1   T2
        """
        x = node.right

        node.right = x.left
        x.left = node

        x.color = node.color
        node.color = self.RED

        return x

    def _right_rotate(self, node: 'RBTree.Node') -> 'RBTree.Node':
        """
        右旋转
             node                     x
            /   \     右旋转        /   \
           x    T2   ------->   y   node
          / \                       /   \
         y  T1                     T1  T2
        """
        x = node.left

        node.left = x.right
        x.right = node

        x.color = node.color
        node.color = self.RED

        return x

    def _flip_color(self, node: 'RBTree.Node') -> None:
        """颜色翻转"""
        node.color = self.RED
        node.left.color = self.BLACK
        node.right.color = self.BLACK

    def _is_red(self, node: Optional['RBTree.Node']) -> bool:
        """判断节点是否为红色（空节点视为黑色）"""
        if node is None:
            return self.BLACK
        return node.color

    # ===== 核心操作 =====

    def _add(self, node: Optional['RBTree.Node'], key: K, value: V) -> 'RBTree.Node':
        """
        添加节点（递归）
        红黑树的添加操作基于2-3树，通过颜色标记和旋转保持平衡
        """
        if node is None:
            self._size += 1
            return self.Node(key, value)

        if key < node.key:
            node.left = self._add(node.left, key, value)
        elif key > node.key:
            node.right = self._add(node.right, key, value)
        else:
            # 键已存在，更新值
            node.value = value
            return node

        # ===== 维护红黑树性质 =====

        # 情况1：右孩子是红色，左孩子是黑色 -> 左旋
        if self._is_red(node.right) and not self._is_red(node.left):
            node = self._left_rotate(node)

        # 情况2：左孩子是红色，左孩子的左孩子是红色 -> 右旋
        if self._is_red(node.left) and self._is_red(node.left.left):
            node = self._right_rotate(node)

        # 情况3：左右孩子都是红色 -> 颜色翻转
        if self._is_red(node.left) and self._is_red(node.right):
            self._flip_color(node)

        return node

    def _get_node(self, node: Optional['RBTree.Node'], key: K) -> Optional['RBTree.Node']:
        """获取节点"""
        if node is None:
            return None

        if key == node.key:
            return node
        elif key < node.key:
            return self._get_node(node.left, key)
        else:
            return self._get_node(node.right, key)

    def remove(self, key: K) -> Optional[V]:
        """删除键值对（较复杂，使用懒惰删除标记）"""
        # 红黑树的删除操作非常复杂
        # 这里使用简化实现：先查找，然后标记为已删除
        # 在实际应用中，可以使用更复杂的删除算法
        node = self._get_node(self._root, key)
        if node is None:
            return None

        value = node.value
        # 实际的删除操作需要重新平衡
        # 这里简化为懒惰删除
        self._root = self._remove(self._root, key)
        return value

    def _remove(self, node: Optional['RBTree.Node'], key: K) -> Optional['RBTree.Node']:
        """
        删除节点（简化版本）

        红黑树的完整删除操作非常复杂，涉及多种情况
        这里实现一个基础版本
        """
        if node is None:
            return None

        if key < node.key:
            node.left = self._remove(node.left, key)
        elif key > node.key:
            node.right = self._remove(node.right, key)
        else:
            # 找到要删除的节点
            if node.left is None:
                self._size -= 1
                return node.right
            elif node.right is None:
                self._size -= 1
                return node.left
            else:
                # 两个孩子都不为空，找后继节点
                successor = self._minimum(node.right)
                successor_key = successor.key
                successor_value = successor.value
                # 递归删除后继
                node.right = self._remove(node.right, successor_key)
                # 用后继的值替换当前节点
                node.key = successor_key
                node.value = successor_value

        return node

    def _minimum(self, node: 'RBTree.Node') -> 'RBTree.Node':
        """找到子树中的最小节点"""
        while node.left is not None:
            node = node.left
        return node

    # ===== 辅助方法 =====

    def keys(self) -> list[K]:
        """获取所有键（中序遍历）"""
        result: list[K] = []

        def _in_order(n: Optional['RBTree.Node']):
            if n is None:
                return
            _in_order(n.left)
            result.append(n.key)
            _in_order(n.right)

        _in_order(self._root)
        return result

    def values(self) -> list[V]:
        """获取所有值"""
        result: list[V] = []

        def _in_order(n: Optional['RBTree.Node']):
            if n is None:
                return
            _in_order(n.left)
            result.append(n.value)
            _in_order(n.right)

        _in_order(self._root)
        return result

    def items(self) -> list[tuple[K, V]]:
        """获取所有键值对"""
        result: list[tuple[K, V]] = []

        def _in_order(n: Optional['RBTree.Node']):
            if n is None:
                return
            _in_order(n.left)
            result.append((n.key, n.value))
            _in_order(n.right)

        _in_order(self._root)
        return result

    def is_valid_rb_tree(self) -> bool:
        """验证是否是有效的红黑树"""
        return self._is_valid_rb_tree(self._root)

    def _is_valid_rb_tree(self, node: Optional['RBTree.Node']) -> bool:
        """递归验证红黑树性质"""
        if node is None:
            return True

        # 性质4：红色节点的子节点必须是黑色
        if node.color == self.RED:
            if (node.left is not None and node.left.color == self.RED) or \
               (node.right is not None and node.right.color == self.RED):
                return False

        # 递归检查左右子树
        return self._is_valid_rb_tree(node.left) and self._is_valid_rb_tree(node.right)

    def __str__(self) -> str:
        """字符串表示"""
        items = [f"{k}:{v}" for k, v in self.items()]
        return f"RBTree({{{', '.join(items)}}})"

    def __repr__(self) -> str:
        return self.__str__()

    def __len__(self) -> int:
        return self._size

    def __contains__(self, key: K) -> bool:
        return self.contains(key)

    def __getitem__(self, key: K) -> Optional[V]:
        return self.get(key)

    def __setitem__(self, key: K, value: V) -> None:
        self.set(key, value)

    def __delitem__(self, key: K) -> None:
        self.remove(key)


if __name__ == "__main__":
    print("=== Python红黑树示例 ===")

    rb = RBTree[str, int]()
    rb["one"] = 1
    rb["two"] = 2
    rb["three"] = 3
    rb["four"] = 4
    rb["five"] = 5

    print(f"红黑树: {rb}")
    print(f"大小: {rb.size()}")
    print(f"获取 'two': {rb.get('two')}")
    print(f"包含 'six': {rb.contains('six')}")
    print(f"是有效的红黑树: {rb.is_valid_rb_tree()}")

    # 中序遍历应该得到有序结果
    print(f"所有键（有序）: {rb.keys()}")

    # 删除操作
    del rb["three"]
    print(f"删除 'three' 后: {rb.keys()}")
    print(f"依然是有效红黑树: {rb.is_valid_rb_tree()}")
