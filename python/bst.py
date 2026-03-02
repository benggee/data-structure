"""
Binary Search Tree Implementation in Python
Based on Java BinarySearchTree implementation
"""

from typing import TypeVar, Generic, Optional, List, Callable
from collections import deque

T = TypeVar('T')


class BSTNode(Generic[T]):
    """Node class for Binary Search Tree"""

    def __init__(self, e: T):
        self.e: T = e
        self.left: Optional['BSTNode[T]'] = None
        self.right: Optional['BSTNode[T]'] = None

    def __str__(self) -> str:
        return str(self.e)


class BinarySearchTree(Generic[T]):
    """Binary Search Tree implementation"""

    def __init__(self):
        self._root: Optional[BSTNode[T]] = None
        self._size: int = 0

    def size(self) -> int:
        """Get tree size"""
        return self._size

    def is_empty(self) -> bool:
        """Check if tree is empty"""
        return self._size == 0

    def add(self, e: T) -> None:
        """Add element to tree"""
        self._root = self._add(self._root, e)

    def _add(self, node: Optional[BSTNode[T]], e: T) -> BSTNode[T]:
        """Recursively add element to subtree"""
        if node is None:
            self._size += 1
            return BSTNode[T](e)

        if e < node.e:
            node.left = self._add(node.left, e)
        elif e > node.e:
            node.right = self._add(node.right, e)
        # If e == node.e, don't add duplicate

        return node

    def contains(self, e: T) -> bool:
        """Check if element exists in tree"""
        return self._contains(self._root, e)

    def _contains(self, node: Optional[BSTNode[T]], e: T) -> bool:
        """Recursively check if element exists in subtree"""
        if node is None:
            return False

        if e == node.e:
            return True
        elif e < node.e:
            return self._contains(node.left, e)
        else:
            return self._contains(node.right, e)

    # Traversals
    def pre_order(self, visit_func: Optional[Callable[[T], None]] = None) -> None:
        """Pre-order traversal (Root, Left, Right)"""
        self._pre_order(self._root, visit_func)

    def _pre_order(self, node: Optional[BSTNode[T]],
                   visit_func: Optional[Callable[[T], None]] = None) -> None:
        if node is None:
            return
        if visit_func:
            visit_func(node.e)
        else:
            print(node.e)
        self._pre_order(node.left, visit_func)
        self._pre_order(node.right, visit_func)

    def pre_order_nr(self, visit_func: Optional[Callable[[T], None]] = None) -> None:
        """Non-recursive pre-order traversal"""
        if self._root is None:
            return

        stack = [self._root]
        while stack:
            cur = stack.pop()
            if visit_func:
                visit_func(cur.e)
            else:
                print(cur.e)
            if cur.right:
                stack.append(cur.right)
            if cur.left:
                stack.append(cur.left)

    def in_order(self, visit_func: Optional[Callable[[T], None]] = None) -> None:
        """In-order traversal (Left, Root, Right)"""
        self._in_order(self._root, visit_func)

    def _in_order(self, node: Optional[BSTNode[T]],
                  visit_func: Optional[Callable[[T], None]] = None) -> None:
        if node is None:
            return
        self._in_order(node.left, visit_func)
        if visit_func:
            visit_func(node.e)
        else:
            print(node.e)
        self._in_order(node.right, visit_func)

    def post_order(self, visit_func: Optional[Callable[[T], None]] = None) -> None:
        """Post-order traversal (Left, Right, Root)"""
        self._post_order(self._root, visit_func)

    def _post_order(self, node: Optional[BSTNode[T]],
                    visit_func: Optional[Callable[[T], None]] = None) -> None:
        if node is None:
            return
        self._post_order(node.left, visit_func)
        self._post_order(node.right, visit_func)
        if visit_func:
            visit_func(node.e)
        else:
            print(node.e)

    def level_order(self, visit_func: Optional[Callable[[T], None]] = None) -> None:
        """Level-order traversal (Breadth-first)"""
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

    # Min/Max operations
    def minimum(self) -> T:
        """Get minimum element"""
        if self._size == 0:
            raise ValueError("Tree is empty")
        return self._minimum(self._root).e

    def _minimum(self, node: BSTNode[T]) -> BSTNode[T]:
        """Get node with minimum element in subtree"""
        if node.left is None:
            return node
        return self._minimum(node.left)

    def maximum(self) -> T:
        """Get maximum element"""
        if self._size == 0:
            raise ValueError("Tree is empty")
        return self._maximum(self._root).e

    def _maximum(self, node: BSTNode[T]) -> BSTNode[T]:
        """Get node with maximum element in subtree"""
        if node.right is None:
            return node
        return self._maximum(node.right)

    def remove_min(self) -> T:
        """Remove minimum element"""
        if self._size == 0:
            raise ValueError("Tree is empty")
        e = self.minimum()
        self._root = self._remove_min(self._root)
        return e

    def _remove_min(self, node: BSTNode[T]) -> Optional[BSTNode[T]]:
        """Remove minimum element from subtree"""
        if node.left is None:
            right_node = node.right
            node.right = None
            self._size -= 1
            return right_node
        node.left = self._remove_min(node.left)
        return node

    def remove_max(self) -> T:
        """Remove maximum element"""
        if self._size == 0:
            raise ValueError("Tree is empty")
        e = self.maximum()
        self._root = self._remove_max(self._root)
        return e

    def _remove_max(self, node: BSTNode[T]) -> Optional[BSTNode[T]]:
        """Remove maximum element from subtree"""
        if node.right is None:
            left_node = node.left
            node.left = None
            self._size -= 1
            return left_node
        node.right = self._remove_max(node.right)
        return node

    def remove(self, e: T) -> None:
        """Remove element from tree"""
        self._root = self._remove(self._root, e)

    def _remove(self, node: Optional[BSTNode[T]], e: T) -> Optional[BSTNode[T]]:
        """Remove element from subtree"""
        if node is None:
            return None

        if e < node.e:
            node.left = self._remove(node.left, e)
            return node
        elif e > node.e:
            node.right = self._remove(node.right, e)
            return node
        else:  # e == node.e
            # Left child is empty
            if node.left is None:
                right_node = node.right
                node.right = None
                self._size -= 1
                return right_node
            # Right child is empty
            if node.right is None:
                left_node = node.left
                node.left = None
                self._size -= 1
                return left_node
            # Both children exist - find successor (minimum in right subtree)
            successor = self._minimum(node.right)
            successor.right = self._remove_min(node.right)
            successor.left = node.left
            node.left = node.right = None
            return successor

    def to_list(self) -> List[T]:
        """Convert tree to sorted list"""
        result: List[T] = []

        def _in_order_collect(n: Optional[BSTNode[T]]) -> None:
            if n is None:
                return
            _in_order_collect(n.left)
            result.append(n.e)
            _in_order_collect(n.right)

        _in_order_collect(self._root)
        return result

    def __str__(self) -> str:
        """String representation (in-order)"""
        return str(self.to_list())

    def __repr__(self) -> str:
        return f"BinarySearchTree({self.to_list()})"

    def __len__(self) -> int:
        return self._size

    def __contains__(self, item: T) -> bool:
        return self.contains(item)

    def __iter__(self):
        """Make tree iterable (in-order)"""
        return iter(self.to_list())


if __name__ == "__main__":
    # Test BinarySearchTree
    print("Testing BinarySearchTree:")
    bst = BinarySearchTree[int]()
    tree_data = [3, 4, 5, 12, 343, 8, 10, 22]
    for data in tree_data:
        bst.add(data)

    print(f"Size: {bst.size()}")
    print(f"Contains 8: {bst.contains(8)}")
    print(f"Contains 100: {bst.contains(100)}")
    print(f"Minimum: {bst.minimum()}")
    print(f"Maximum: {bst.maximum()}")

    print("\nPre-order traversal:")
    bst.pre_order()

    print("\nIn-order traversal (sorted):")
    bst.in_order()

    print("\nLevel-order traversal:")
    bst.level_order()

    print(f"\nTo list: {bst.to_list()}")

    bst.remove(5)
    print("\nAfter removing 5:")
    print(f"Tree: {bst}")

    # Test iteration
    print("\nIteration:")
    for item in bst:
        print(item, end=" ")
    print()
