"""
Segment Tree Implementation in Python
Based on Java SegmentTree implementation
"""

from typing import TypeVar, Generic, List, Callable, Optional

T = TypeVar('T')


class Merger(Generic[T]):
    """Functional interface for merging two values"""

    def merge(self, a: T, b: T) -> T:
        """Merge two values"""
        raise NotImplementedError


class SegmentTree(Generic[T]):
    """Segment Tree for range queries and updates"""

    def __init__(self, arr: List[T], merger: Callable[[T, T], T]):
        """
        Initialize segment tree

        Args:
            arr: Input array
            merger: Function to merge two values (e.g., sum, max, min)
        """
        self._data: List[T] = arr.copy()
        self._merger: Callable[[T, T], T] = merger
        self._tree: List[Optional[T]] = [None] * (4 * len(arr))
        self._create_segment_tree(0, 0, len(self._data) - 1)

    def _create_segment_tree(self, tree_index: int, l: int, r: int) -> None:
        """Build segment tree recursively"""
        if l == r:
            self._tree[tree_index] = self._data[l]
            return

        left_tree_index = self._left_child(tree_index)
        right_tree_index = self._right_child(tree_index)

        mid = l + (r - l) // 2
        self._create_segment_tree(left_tree_index, l, mid)
        self._create_segment_tree(right_tree_index, mid + 1, r)

        self._tree[tree_index] = self._merger(
            self._tree[left_tree_index],  # type: ignore
            self._tree[right_tree_index]  # type: ignore
        )

    def query(self, query_l: int, query_r: int) -> T:
        """
        Query range [query_l, query_r]

        Args:
            query_l: Left bound of query (inclusive)
            query_r: Right bound of query (inclusive)

        Returns:
            Result of the query
        """
        if query_l < 0 or query_l >= len(self._data) or \
           query_r < 0 or query_r >= len(self._data) or query_l > query_r:
            raise ValueError("Query range is invalid")

        return self._query(0, 0, len(self._data) - 1, query_l, query_r)

    def _query(self, tree_index: int, l: int, r: int,
               query_l: int, query_r: int) -> T:
        """Recursively query range"""
        if l == query_l and r == query_r:
            return self._tree[tree_index]  # type: ignore

        mid = l + (r - l) // 2
        left_tree_index = self._left_child(tree_index)
        right_tree_index = self._right_child(tree_index)

        if query_l > mid:
            # Query is entirely in right subtree
            return self._query(right_tree_index, mid + 1, r, query_l, query_r)
        elif query_r <= mid:
            # Query is entirely in left subtree
            return self._query(left_tree_index, l, mid, query_l, query_r)
        else:
            # Query spans both subtrees
            left_result = self._query(left_tree_index, l, mid, query_l, mid)
            right_result = self._query(right_tree_index, mid + 1, r, mid + 1, query_r)
            return self._merger(left_result, right_result)

    def set(self, index: int, value: T) -> None:
        """
        Set element at index to value

        Args:
            index: Index to update
            value: New value
        """
        if index < 0 or index >= len(self._data):
            raise IndexError("Index out of range")

        self._data[index] = value
        self._set(0, 0, len(self._data) - 1, index, value)

    def _set(self, tree_index: int, l: int, r: int, index: int, value: T) -> None:
        """Recursively update element"""
        if l == r:
            self._tree[tree_index] = value
            return

        mid = l + (r - l) // 2
        left_tree_index = self._left_child(tree_index)
        right_tree_index = self._right_child(tree_index)

        if index <= mid:
            self._set(left_tree_index, l, mid, index, value)
        else:
            self._set(right_tree_index, mid + 1, r, index, value)

        # Recalculate parent node after update
        self._tree[tree_index] = self._merger(
            self._tree[left_tree_index],  # type: ignore
            self._tree[right_tree_index]  # type: ignore
        )

    def get_size(self) -> int:
        """Get size of data"""
        return len(self._data)

    def get(self, index: int) -> T:
        """Get element at index"""
        if index < 0 or index >= len(self._data):
            raise IndexError("Index out of range")
        return self._data[index]

    def _left_child(self, index: int) -> int:
        """Get left child index"""
        return 2 * index + 1

    def _right_child(self, index: int) -> int:
        """Get right child index"""
        return 2 * index + 2

    def __str__(self) -> str:
        """String representation of tree array"""
        return "[" + ", ".join(
            str(v) if v is not None else "null" for v in self._tree
        ) + "]"

    def __repr__(self) -> str:
        return f"SegmentTree(size={len(self._data)})"

    def __len__(self) -> int:
        return len(self._data)

    def __getitem__(self, index: int) -> T:
        return self.get(index)

    def __setitem__(self, index: int, value: T) -> None:
        self.set(index, value)


if __name__ == "__main__":
    # Test SegmentTree with sum operation
    print("Testing SegmentTree (sum):")
    arr = [1, 3, 5, 7, 9, 11, 13, 15]
    seg_tree = SegmentTree[int](arr, lambda a, b: a + b)

    print(f"Original array: {arr}")
    print(f"Tree: {seg_tree}")

    print(f"\nSum of [0, 3]: {seg_tree.query(0, 3)}")  # 1+3+5+7 = 16
    print(f"Sum of [2, 5]: {seg_tree.query(2, 5)}")  # 5+7+9+11 = 32
    print(f"Sum of [0, 7]: {seg_tree.query(0, 7)}")  # Total

    # Test update
    seg_tree.set(3, 10)
    print(f"\nAfter setting index 3 to 10:")
    print(f"Sum of [0, 3]: {seg_tree.query(0, 3)}")  # 1+3+5+10 = 19

    # Test SegmentTree with max operation
    print("\n\nTesting SegmentTree (max):")
    arr2 = [1, 3, 5, 7, 9, 11, 13, 15]
    seg_tree_max = SegmentTree[int](arr2, max)

    print(f"Original array: {arr2}")
    print(f"Max of [0, 7]: {seg_tree_max.query(0, 7)}")  # 15
    print(f"Max of [2, 5]: {seg_tree_max.query(2, 5)}")  # 11

    seg_tree_max.set(3, 20)
    print(f"\nAfter setting index 3 to 20:")
    print(f"Max of [0, 7]: {seg_tree_max.query(0, 7)}")  # 20

    # Test SegmentTree with min operation
    print("\n\nTesting SegmentTree (min):")
    arr3 = [5, 3, 8, 1, 9, 2, 7, 4]
    seg_tree_min = SegmentTree[int](arr3, min)

    print(f"Original array: {arr3}")
    print(f"Min of [0, 7]: {seg_tree_min.query(0, 7)}")  # 1
    print(f"Min of [2, 5]: {seg_tree_min.query(2, 5)}")  # 2
