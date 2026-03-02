"""
Union Find (Disjoint Set) Implementation in Python
Based on Java UnionFind implementation
"""

from typing import List


class UnionFind:
    """Union Find with path compression and union by rank (optimized version)"""

    def __init__(self, size: int):
        self._parent: List[int] = list(range(size))
        self._rank: List[int] = [1] * size

    def size(self) -> int:
        """Get number of elements"""
        return len(self._parent)

    def is_connected(self, p: int, q: int) -> bool:
        """Check if two elements are connected"""
        return self.find(p) == self.find(q)

    def union_elements(self, p: int, q: int) -> None:
        """Union two elements"""
        p_root = self.find(p)
        q_root = self.find(q)

        if p_root == q_root:
            return

        # Union by rank - attach smaller tree to larger tree
        if self._rank[p_root] > self._rank[q_root]:
            self._parent[q_root] = p_root
        elif self._rank[p_root] < self._rank[q_root]:
            self._parent[p_root] = q_root
        else:
            self._parent[p_root] = q_root
            self._rank[q_root] += 1

    def find(self, p: int) -> int:
        """Find root of element with path compression"""
        if p < 0 or p >= len(self._parent):
            raise IndexError("Index out of range")

        # Path compression
        if p != self._parent[p]:
            self._parent[p] = self.find(self._parent[p])
        return self._parent[p]

    def count_sets(self) -> int:
        """Count number of disjoint sets"""
        roots = set()
        for i in range(len(self._parent)):
            roots.add(self.find(i))
        return len(roots)

    def get_set_sizes(self) -> dict[int, int]:
        """Get size of each set"""
        set_sizes: dict[int, int] = {}
        for i in range(len(self._parent)):
            root = self.find(i)
            set_sizes[root] = set_sizes.get(root, 0) + 1
        return set_sizes

    def __str__(self) -> str:
        return f"UnionFind(elements={len(self._parent)}, sets={self.count_sets()})"

    def __repr__(self) -> str:
        return self.__str__()


class UnionFindV1:
    """Union Find Version 1: Array-based with path copying (O(n))"""

    def __init__(self, size: int):
        self._data: List[List[int]] = [[i] for i in range(size)]

    def size(self) -> int:
        return len(self._data)

    def is_connected(self, p: int, q: int) -> bool:
        if p < 0 or p >= len(self._data) or q < 0 or q >= len(self._data):
            raise IndexError("Index out of range")
        return self._find_id(p) == self._find_id(q)

    def _find_id(self, p: int) -> int:
        """Find which set element p belongs to"""
        for i, s in enumerate(self._data):
            if p in s:
                return i
        return -1

    def union_elements(self, p: int, q: int) -> None:
        """Union two elements"""
        p_id = self._find_id(p)
        q_id = self._find_id(q)

        if p_id == q_id:
            return

        # Merge q's set into p's set
        self._data[p_id].extend(self._data[q_id])
        self._data.pop(q_id)

    def find(self, p: int) -> int:
        """Find root (set identifier)"""
        return self._find_id(p)


class UnionFindV2:
    """Union Find Version 2: Quick-find with array id (O(n))"""

    def __init__(self, size: int):
        self._id: List[int] = list(range(size))

    def size(self) -> int:
        return len(self._id)

    def is_connected(self, p: int, q: int) -> bool:
        if p < 0 or p >= len(self._id) or q < 0 or q >= len(self._id):
            raise IndexError("Index out of range")
        return self._id[p] == self._id[q]

    def find(self, p: int) -> int:
        """Find the set identifier"""
        if p < 0 or p >= len(self._id):
            raise IndexError("Index out of range")
        return self._id[p]

    def union_elements(self, p: int, q: int) -> None:
        """Union two elements"""
        p_id = self.find(p)
        q_id = self.find(q)

        if p_id == q_id:
            return

        # Change all elements with id[p] to id[q]
        for i in range(len(self._id)):
            if self._id[i] == p_id:
                self._id[i] = q_id


class UnionFindV3:
    """Union Find Version 3: Quick-union with tree structure (O(n))"""

    def __init__(self, size: int):
        self._parent: List[int] = list(range(size))

    def size(self) -> int:
        return len(self._parent)

    def is_connected(self, p: int, q: int) -> bool:
        return self.find(p) == self.find(q)

    def find(self, p: int) -> int:
        """Find root of element"""
        if p < 0 or p >= len(self._parent):
            raise IndexError("Index out of range")
        while p != self._parent[p]:
            p = self._parent[p]
        return p

    def union_elements(self, p: int, q: int) -> None:
        """Union two elements"""
        p_root = self.find(p)
        q_root = self.find(q)

        if p_root == q_root:
            return

        self._parent[p_root] = q_root


class UnionFindV4:
    """Union Find Version 4: Union by rank optimization"""

    def __init__(self, size: int):
        self._parent: List[int] = list(range(size))
        self._rank: List[int] = [1] * size

    def size(self) -> int:
        return len(self._parent)

    def is_connected(self, p: int, q: int) -> bool:
        return self.find(p) == self.find(q)

    def find(self, p: int) -> int:
        """Find root of element"""
        if p < 0 or p >= len(self._parent):
            raise IndexError("Index out of range")
        while p != self._parent[p]:
            p = self._parent[p]
        return p

    def union_elements(self, p: int, q: int) -> None:
        """Union two elements with rank optimization"""
        p_root = self.find(p)
        q_root = self.find(q)

        if p_root == q_root:
            return

        # Union by rank
        if self._rank[p_root] < self._rank[q_root]:
            self._parent[p_root] = q_root
        elif self._rank[p_root] > self._rank[q_root]:
            self._parent[q_root] = p_root
        else:
            self._parent[p_root] = q_root
            self._rank[q_root] += 1


class UnionFindV5:
    """Union Find Version 5: Path compression optimization"""

    def __init__(self, size: int):
        self._parent: List[int] = list(range(size))

    def size(self) -> int:
        return len(self._parent)

    def is_connected(self, p: int, q: int) -> bool:
        return self.find(p) == self.find(q)

    def find(self, p: int) -> int:
        """Find root of element with path compression"""
        if p < 0 or p >= len(self._parent):
            raise IndexError("Index out of range")
        if p != self._parent[p]:
            self._parent[p] = self.find(self._parent[p])
        return self._parent[p]

    def union_elements(self, p: int, q: int) -> None:
        """Union two elements"""
        p_root = self.find(p)
        q_root = self.find(q)

        if p_root == q_root:
            return

        self._parent[p_root] = q_root


if __name__ == "__main__":
    # Test UnionFind (optimized version)
    print("Testing UnionFind (V6 - optimized):")
    uf = UnionFind(10)
    print(uf)

    uf.union_elements(0, 1)
    uf.union_elements(2, 3)
    uf.union_elements(1, 2)
    uf.union_elements(4, 5)
    uf.union_elements(6, 7)

    print(f"After unions:")
    print(f"Is 0 connected to 3: {uf.is_connected(0, 3)}")
    print(f"Is 0 connected to 4: {uf.is_connected(0, 4)}")
    print(f"Count sets: {uf.count_sets()}")
    print(f"Set sizes: {uf.get_set_sizes()}")

    # Test UnionFindV2
    print("\nTesting UnionFindV2 (Quick-find):")
    uf2 = UnionFindV2(10)
    uf2.union_elements(0, 1)
    uf2.union_elements(2, 3)
    uf2.union_elements(1, 2)
    print(f"Is 0 connected to 3: {uf2.is_connected(0, 3)}")

    # Test UnionFindV3
    print("\nTesting UnionFindV3 (Quick-union):")
    uf3 = UnionFindV3(10)
    uf3.union_elements(0, 1)
    uf3.union_elements(2, 3)
    uf3.union_elements(1, 2)
    print(f"Is 0 connected to 3: {uf3.is_connected(0, 3)}")
