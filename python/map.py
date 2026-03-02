"""
Map Implementation in Python
Based on Java Map implementation
"""

from typing import TypeVar, Generic, Optional, ABC, abstractmethod
from linklist import LinkList

K = TypeVar('K')
V = TypeVar('V')


class MapNode(Generic[K, V]):
    """Node class for Map"""

    def __init__(self, key: Optional[K] = None, value: Optional[V] = None,
                 next_node: Optional['MapNode[K, V]'] = None):
        self.key: Optional[K] = key
        self.value: Optional[V] = value
        self.next: Optional['MapNode[K, V]'] = next_node

    def __str__(self) -> str:
        return f"{self.key}:{self.value}"


class Map(Generic[K, V], ABC):
    """Map interface"""

    @abstractmethod
    def remove(self, key: K) -> Optional[V]:
        """Remove key-value pair by key"""
        pass

    @abstractmethod
    def contains(self, key: K) -> bool:
        """Check if key exists"""
        pass

    @abstractmethod
    def get(self, key: K) -> Optional[V]:
        """Get value by key"""
        pass

    @abstractmethod
    def set(self, key: K, value: V) -> None:
        """Set key-value pair"""
        pass

    @abstractmethod
    def size(self) -> int:
        """Get map size"""
        pass

    @abstractmethod
    def is_empty(self) -> bool:
        """Check if map is empty"""
        pass


class LinkListMap(Map[K, V]):
    """Map implementation using linked list"""

    def __init__(self):
        self._dummy_head: MapNode[K, V] = MapNode[K, V]()
        self._size: int = 0

    def _get_node(self, key: K) -> Optional[MapNode[K, V]]:
        """Get node by key"""
        cur = self._dummy_head.next
        while cur is not None:
            if cur.key == key:
                return cur
            cur = cur.next
        return None

    def size(self) -> int:
        return self._size

    def is_empty(self) -> bool:
        return self._size == 0

    def contains(self, key: K) -> bool:
        return self._get_node(key) is not None

    def get(self, key: K) -> Optional[V]:
        node = self._get_node(key)
        return node.value if node is not None else None

    def set(self, key: K, value: V) -> None:
        """Set key-value pair. If key exists, update value; otherwise add new pair"""
        node = self._get_node(key)
        if node is not None:
            node.value = value
        else:
            new_node = MapNode[K, V](key, value, self._dummy_head.next)
            self._dummy_head.next = new_node
            self._size += 1

    def remove(self, key: K) -> Optional[V]:
        """Remove key-value pair by key"""
        prev = self._dummy_head
        while prev.next is not None:
            if prev.next.key == key:
                del_node = prev.next
                prev.next = del_node.next
                del_node.next = None
                self._size -= 1
                return del_node.value
            prev = prev.next
        return None

    def __str__(self) -> str:
        result = []
        cur = self._dummy_head.next
        while cur is not None:
            result.append(str(cur))
            cur = cur.next
        return f"LinkListMap({{{', '.join(result)}}})"

    def __repr__(self) -> str:
        return self.__str__()

    def __len__(self) -> int:
        return self.size()

    def __contains__(self, key: K) -> bool:
        return self.contains(key)

    def __getitem__(self, key: K) -> Optional[V]:
        return self.get(key)

    def __setitem__(self, key: K, value: V) -> None:
        self.set(key, value)

    def __delitem__(self, key: K) -> None:
        self.remove(key)

    def keys(self):
        """Get all keys"""
        keys = []
        cur = self._dummy_head.next
        while cur is not None:
            if cur.key is not None:
                keys.append(cur.key)
            cur = cur.next
        return keys

    def values(self):
        """Get all values"""
        values = []
        cur = self._dummy_head.next
        while cur is not None:
            if cur.value is not None:
                values.append(cur.value)
            cur = cur.next
        return values

    def items(self):
        """Get all key-value pairs"""
        items = []
        cur = self._dummy_head.next
        while cur is not None:
            if cur.key is not None:
                items.append((cur.key, cur.value))
            cur = cur.next
        return items


if __name__ == "__main__":
    # Test LinkListMap
    print("Testing LinkListMap:")
    m = LinkListMap[str, int]()
    m.set("one", 1)
    m.set("two", 2)
    m.set("three", 3)
    print(f"Map: {m}")
    print(f"Size: {m.size()}")
    print(f"Get 'two': {m.get('two')}")
    print(f"Contains 'four': {m.contains('four')}")

    m.set("two", 22)  # Update value
    print(f"After updating 'two' to 22: {m}")

    m.remove("one")
    print(f"After removing 'one': {m}")

    # Test dictionary-like access
    m["five"] = 5
    print(f"After m['five'] = 5: {m}")
    print(f"m['three']: {m['three']}")
