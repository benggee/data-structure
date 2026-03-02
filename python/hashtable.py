"""
Hash Table Implementation in Python
Based on Java HashTable implementation
"""

from typing import TypeVar, Generic, Optional, List
from abc import ABC, abstractmethod

K = TypeVar('K')
V = TypeVar('V')


class HashTable(Generic[K, V]):
    """Hash Table implementation with separate chaining using AVL trees"""

    # Prime numbers for capacity (good hash table primes)
    _CAPACITY = [53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593,
                 49157, 98317, 196613, 393241, 786433, 1572869, 3145739,
                 6291469, 12582917, 25165843, 50331653, 100663319, 201326611,
                 402653189, 805306457, 1610612741]

    def __init__(self, upper_tol: int = 10, lower_tol: int = 2):
        self._capacity_index: int = 0
        self._M: int = self._CAPACITY[self._capacity_index]
        self._size: int = 0
        self._upper_tol: int = upper_tol
        self._lower_tol: int = lower_tol
        # Using Python dict as the underlying map (equivalent to TreeMap in Java)
        self._hashtable: List[dict] = [{} for _ in range(self._M)]

    def _hash(self, key: K) -> int:
        """Compute hash value for key"""
        return (hash(key) & 0x7fffffff) % self._M

    def size(self) -> int:
        """Get number of key-value pairs"""
        return self._size

    def add(self, key: K, value: V) -> None:
        """Add key-value pair"""
        map_dict = self._hashtable[self._hash(key)]
        if key in map_dict:
            map_dict[key] = value  # Update existing key
        else:
            map_dict[key] = value
            self._size += 1

            # Resize if needed (upper tolerance)
            if self._size >= self._upper_tol * self._M and \
               self._capacity_index + 1 < len(self._CAPACITY):
                self._capacity_index += 1
                self._resize(self._CAPACITY[self._capacity_index])

    def remove(self, key: K) -> Optional[V]:
        """Remove key-value pair by key"""
        map_dict = self._hashtable[self._hash(key)]
        if key in map_dict:
            value = map_dict.pop(key)
            self._size -= 1

            # Resize if needed (lower tolerance)
            if self._size < self._lower_tol * self._M and self._capacity_index - 1 >= 0:
                self._capacity_index -= 1
                self._resize(self._CAPACITY[self._capacity_index])

            return value
        return None

    def set(self, key: K, value: V) -> None:
        """Set value for existing key"""
        map_dict = self._hashtable[self._hash(key)]
        if key not in map_dict:
            raise KeyError(f"Key '{key}' doesn't exist")
        map_dict[key] = value

    def contains(self, key: K) -> bool:
        """Check if key exists"""
        return key in self._hashtable[self._hash(key)]

    def get(self, key: K) -> Optional[V]:
        """Get value by key"""
        return self._hashtable[self._hash(key)].get(key)

    def _resize(self, new_M: int) -> None:
        """Resize the hash table"""
        new_hashtable: List[dict] = [{} for _ in range(new_M)]
        old_M = self._M
        self._M = new_M

        # Rehash all entries
        for i in range(old_M):
            map_dict = self._hashtable[i]
            for key, value in map_dict.items():
                new_hashtable[self._hash(key)][key] = value

        self._hashtable = new_hashtable

    def capacity(self) -> int:
        """Get current capacity"""
        return self._M

    def keys(self) -> List[K]:
        """Get all keys"""
        result: List[K] = []
        for map_dict in self._hashtable:
            result.extend(map_dict.keys())
        return result

    def values(self) -> List[V]:
        """Get all values"""
        result: List[V] = []
        for map_dict in self._hashtable:
            result.extend(map_dict.values())
        return result

    def items(self) -> List[tuple[K, V]]:
        """Get all key-value pairs"""
        result: List[tuple[K, V]] = []
        for map_dict in self._hashtable:
            result.extend(map_dict.items())
        return result

    def __str__(self) -> str:
        items = [f"{k}:{v}" for k, v in self.items()]
        return f"HashTable(size={self._size}, capacity={self._M}, {{{', '.join(items)}}})"

    def __repr__(self) -> str:
        return self.__str__()

    def __len__(self) -> int:
        return self._size

    def __contains__(self, key: K) -> bool:
        return self.contains(key)

    def __getitem__(self, key: K) -> Optional[V]:
        return self.get(key)

    def __setitem__(self, key: K, value: V) -> None:
        self.add(key, value)

    def __delitem__(self, key: K) -> None:
        value = self.remove(key)
        if value is None:
            raise KeyError(f"Key '{key}' doesn't exist")


if __name__ == "__main__":
    # Test HashTable
    print("Testing HashTable:")
    ht = HashTable[str, int]()

    # Add some key-value pairs
    ht["one"] = 1
    ht["two"] = 2
    ht["three"] = 3
    ht["four"] = 4
    ht["five"] = 5

    print(f"HashTable: {ht}")
    print(f"Size: {ht.size()}")
    print(f"Capacity: {ht.capacity()}")
    print(f"Get 'two': {ht.get('two')}")
    print(f"Contains 'six': {ht.contains('six')}")

    # Update existing key
    ht["two"] = 22
    print(f"After updating 'two' to 22: {ht}")

    # Remove a key
    del ht["three"]
    print(f"After deleting 'three': {ht}")

    # Test dictionary-like operations
    print("\nKeys:", ht.keys())
    print("Values:", ht.values())
    print("Items:", ht.items())

    # Add more items to trigger resize
    print("\nAdding more items...")
    for i in range(50):
        ht[f"key_{i}"] = i

    print(f"After adding 50 more items:")
    print(f"Size: {ht.size()}")
    print(f"Capacity: {ht.capacity()}")
