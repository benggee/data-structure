"""
Dynamic Array Implementation in Python
Based on Java Array implementation
"""

from typing import TypeVar, Generic, List
import copy

T = TypeVar('T')


class Array(Generic[T]):
    """Dynamic array with automatic resizing"""

    def __init__(self, capacity: int = 10):
        self._data: List[T] = [None] * capacity
        self._size: int = 0

    def capacity(self) -> int:
        """Get array capacity"""
        return len(self._data)

    def size(self) -> int:
        """Get number of elements"""
        return self._size

    def is_empty(self) -> bool:
        """Check if array is empty"""
        return self._size == 0

    def add(self, index: int, e: T) -> None:
        """Insert element at specified index"""
        if index < 0 or index > self._size:
            raise IndexError(f"Add failed. Index must be >= 0 and <= size")

        if self._size == len(self._data):
            self._resize(2 * len(self._data))

        for i in range(self._size - 1, index - 1, -1):
            self._data[i + 1] = self._data[i]

        self._data[index] = e
        self._size += 1

    def add_first(self, e: T) -> None:
        """Add element at the beginning"""
        self.add(0, e)

    def add_last(self, e: T) -> None:
        """Add element at the end"""
        self.add(self._size, e)

    def get(self, index: int) -> T:
        """Get element at specified index"""
        if index < 0 or index >= self._size:
            raise IndexError("Index out of range")
        return self._data[index]

    def set(self, index: int, e: T) -> None:
        """Set element at specified index"""
        if index < 0 or index >= self._size:
            raise IndexError("Index out of range")
        self._data[index] = e

    def swap(self, i: int, j: int) -> None:
        """Swap two elements"""
        if i < 0 or i >= self._size or j < 0 or j >= self._size:
            raise IndexError("Index out of range")
        self._data[i], self._data[j] = self._data[j], self._data[i]

    def find(self, e: T) -> int:
        """Find index of element"""
        for i in range(self._size):
            if self._data[i] == e:
                return i
        return -1

    def contains(self, e: T) -> bool:
        """Check if element exists in array"""
        return self.find(e) != -1

    def remove(self, index: int) -> T:
        """Remove element at specified index"""
        if index < 0 or index >= self._size:
            raise IndexError("Index out of range")

        ret = self._data[index]
        for i in range(index + 1, self._size):
            self._data[i - 1] = self._data[i]

        self._size -= 1
        self._data[self._size] = None

        if self._size == len(self._data) // 2:
            self._resize(len(self._data) // 2)

        return ret

    def remove_last(self) -> T:
        """Remove last element"""
        return self.remove(self._size - 1)

    def remove_first(self) -> T:
        """Remove first element"""
        return self.remove(0)

    def remove_element(self, e: T) -> None:
        """Remove specified element"""
        idx = self.find(e)
        if idx != -1:
            self.remove(idx)

    def _resize(self, new_capacity: int) -> None:
        """Resize internal array"""
        new_data = [None] * new_capacity
        for i in range(self._size):
            new_data[i] = self._data[i]
        self._data = new_data

    def __str__(self) -> str:
        """String representation"""
        return "[" + ", ".join(str(self._data[i]) for i in range(self._size)) + "]"

    def __repr__(self) -> str:
        return f"Array({self.__str__()})"

    def __len__(self) -> int:
        return self._size

    def __getitem__(self, index: int) -> T:
        return self.get(index)

    def __setitem__(self, index: int, value: T) -> None:
        self.set(index, value)


if __name__ == "__main__":
    # Test Array
    arr = Array[int]()
    for i in range(10):
        arr.add_last(i)
    print(f"Array: {arr}")
    print(f"Size: {arr.size()}, Capacity: {arr.capacity()}")

    arr.add(1, 100)
    print(f"After add(1, 100): {arr}")

    arr.remove_element(8)
    print(f"After remove 8: {arr}")

    arr.remove_first()
    arr.remove_last()
    print(f"After remove first and last: {arr}")
