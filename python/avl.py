"""
AVL Tree Implementation in Python
Based on Java AVLTree implementation
"""

from typing import TypeVar, Generic, Optional, List

K = TypeVar('K')
V = TypeVar('V')


class AVLNode(Generic[K, V]):
    """Node class for AVL Tree"""

    def __init__(self, key: K, value: V):
        self.key: K = key
        self.value: V = value
        self.left: Optional['AVLNode[K, V]'] = None
        self.right: Optional['AVLNode[K, V]'] = None
        self.height: int = 1

    def __str__(self) -> str:
        return f"{self.key}:{self.value}"


class AVLTree(Generic[K, V]):
    """AVL Tree (Self-balancing Binary Search Tree) implementation"""

    def __init__(self):
        self._root: Optional['AVLNode[K, V]'] = None
        self._size: int = 0

    def size(self) -> int:
        """Get tree size"""
        return self._size

    def is_empty(self) -> bool:
        """Check if tree is empty"""
        return self._size == 0

    def get_height(self, node: Optional['AVLNode[K, V]']) -> int:
        """Get node height"""
        if node is None:
            return 0
        return node.height

    def get_balance_factor(self, node: Optional['AVLNode[K, V]']) -> int:
        """Get AVL balance factor (left_height - right_height)"""
        if node is None:
            return 0
        return self.get_height(node.left) - self.get_height(node.right)

    def is_bst(self) -> bool:
        """Check if tree is a valid BST"""
        keys: List[K] = []
        self._in_order_keys(self._root, keys)
        for i in range(1, len(keys)):
            if keys[i - 1] > keys[i]:
                return False
        return True

    def _in_order_keys(self, node: Optional['AVLNode[K, V]'], keys: List[K]) -> None:
        """Collect keys in-order"""
        if node is None:
            return
        self._in_order_keys(node.left, keys)
        keys.append(node.key)
        self._in_order_keys(node.right, keys)

    def is_avl(self) -> bool:
        """Check if tree is a valid AVL tree"""
        return self._is_avl(self._root)

    def _is_avl(self, node: Optional['AVLNode[K, V]']) -> bool:
        """Recursively check AVL property"""
        if node is None:
            return True
        if abs(self.get_balance_factor(node)) > 1:
            return False
        return self._is_avl(node.left) and self._is_avl(node.right)

    # Rotations
    def right_rotate(self, y: AVLNode[K, V]) -> AVLNode[K, V]:
        """Right rotation"""
        x = y.left
        t3 = x.right

        x.right = y
        y.left = t3

        y.height = max(self.get_height(y.right), self.get_height(y.left)) + 1
        x.height = max(self.get_height(x.right), self.get_height(x.left)) + 1

        return x

    def left_rotate(self, y: AVLNode[K, V]) -> AVLNode[K, V]:
        """Left rotation"""
        x = y.right
        t2 = x.left

        x.left = y
        y.right = t2

        y.height = max(self.get_height(y.left), self.get_height(y.right)) + 1
        x.height = max(self.get_height(x.left), self.get_height(x.right)) + 1

        return x

    # Core operations
    def set(self, key: K, value: V) -> None:
        """Set key-value pair"""
        self._root = self._add(self._root, key, value)

    def _add(self, node: Optional['AVLNode[K, V]'], key: K, value: V) -> AVLNode[K, V]:
        """Recursively add node and maintain balance"""
        if node is None:
            self._size += 1
            return AVLNode[K, V](key, value)

        if key < node.key:
            node.left = self._add(node.left, key, value)
        elif key > node.key:
            node.right = self._add(node.right, key, value)
        else:  # key == node.key, update value
            node.value = value
            return node

        # Update height
        node.height = 1 + max(self.get_height(node.left), self.get_height(node.right))

        # Get balance factor
        factor = self.get_balance_factor(node)

        # LL Case
        if factor > 1 and self.get_balance_factor(node.left) >= 0:
            return self.right_rotate(node)

        # RR Case
        if factor < -1 and self.get_balance_factor(node.right) <= 0:
            return self.left_rotate(node)

        # LR Case
        if factor > 1 and self.get_balance_factor(node.left) < 0:
            node.left = self.left_rotate(node.left)
            return self.right_rotate(node)

        # RL Case
        if factor < -1 and self.get_balance_factor(node.right) > 0:
            node.right = self.right_rotate(node.right)
            return self.left_rotate(node)

        return node

    def get_node(self, node: Optional['AVLNode[K, V]'], key: K) -> Optional['AVLNode[K, V]']:
        """Get node by key"""
        if node is None:
            return None
        if key == node.key:
            return node
        elif key < node.key:
            return self.get_node(node.left, key)
        else:
            return self.get_node(node.right, key)

    def contains(self, key: K) -> bool:
        """Check if key exists"""
        return self.get_node(self._root, key) is not None

    def get(self, key: K) -> Optional[V]:
        """Get value by key"""
        node = self.get_node(self._root, key)
        return node.value if node else None

    def remove(self, key: K) -> Optional[V]:
        """Remove key-value pair by key"""
        node = self.get_node(self._root, key)
        if node is not None:
            self._root = self._remove(self._root, key)
            return node.value
        return None

    def _remove(self, node: Optional['AVLNode[K, V]'], key: K) -> Optional['AVLNode[K, V]']:
        """Recursively remove node and maintain balance"""
        if node is None:
            return None

        if key < node.key:
            node.left = self._remove(node.left, key)
            ret_node = node
        elif key > node.key:
            node.right = self._remove(node.right, key)
            ret_node = node
        else:  # key == node.key
            if node.left is None:
                right_node = node.right
                node.right = None
                self._size -= 1
                ret_node = right_node
            elif node.right is None:
                left_node = node.left
                node.left = None
                self._size -= 1
                ret_node = left_node
            else:
                # Both children exist - find successor (min in right subtree)
                successor = self._minimum(node.right)
                successor.right = self._remove(node.right, successor.key)
                successor.left = node.left
                node.left = node.right = None
                ret_node = successor

        if ret_node is None:
            return None

        # Update height
        ret_node.height = 1 + max(self.get_height(ret_node.left), self.get_height(ret_node.right))

        # Get balance factor
        factor = self.get_balance_factor(ret_node)

        # LL Case
        if factor > 1 and self.get_balance_factor(ret_node.left) >= 0:
            return self.right_rotate(ret_node)

        # RR Case
        if factor < -1 and self.get_balance_factor(ret_node.right) <= 0:
            return self.left_rotate(ret_node)

        # LR Case
        if factor > 1 and self.get_balance_factor(ret_node.left) < 0:
            ret_node.left = self.left_rotate(ret_node.left)
            return self.right_rotate(ret_node)

        # RL Case
        if factor < -1 and self.get_balance_factor(ret_node.right) > 0:
            ret_node.right = self.right_rotate(ret_node.right)
            return self.left_rotate(ret_node)

        return ret_node

    def _minimum(self, node: 'AVLNode[K, V]') -> 'AVLNode[K, V]':
        """Get minimum node in subtree"""
        if node.left is None:
            return node
        return self._minimum(node.left)

    def keys(self) -> List[K]:
        """Get all keys"""
        result: List[K] = []
        self._in_order_keys(self._root, result)
        return result

    def values(self) -> List[V]:
        """Get all values"""
        result: List[V] = []

        def collect(n: Optional['AVLNode[K, V]']) -> None:
            if n is None:
                return
            collect(n.left)
            result.append(n.value)
            collect(n.right)

        collect(self._root)
        return result

    def items(self) -> List[tuple[K, V]]:
        """Get all key-value pairs"""
        result: List[tuple[K, V]] = []

        def collect(n: Optional['AVLNode[K, V]']) -> None:
            if n is None:
                return
            collect(n.left)
            result.append((n.key, n.value))
            collect(n.right)

        collect(self._root)
        return result

    def __str__(self) -> str:
        """String representation"""
        items = [f"{k}:{v}" for k, v in self.items()]
        return f"AVLTree({{{', '.join(items)}}})"

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
    # Test AVLTree
    print("Testing AVLTree:")
    avl = AVLTree[str, int]()
    avl["one"] = 1
    avl["two"] = 2
    avl["three"] = 3
    avl["four"] = 4
    avl["five"] = 5

    print(f"Tree: {avl}")
    print(f"Size: {avl.size()}")
    print(f"Get 'two': {avl.get('two')}")
    print(f"Contains 'six': {avl.contains('six')}")
    print(f"Is BST: {avl.is_bst()}")
    print(f"Is AVL: {avl.is_avl()}")

    avl["two"] = 22  # Update
    print(f"After updating 'two' to 22: {avl}")

    del avl["three"]
    print(f"After deleting 'three': {avl}")

    # Test iteration
    print("Keys:", avl.keys())
    print("Values:", avl.values())
    print("Items:", avl.items())
