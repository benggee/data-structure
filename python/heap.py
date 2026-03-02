"""
Max Heap Implementation in Python
Based on Java MaxHeap implementation
"""

from typing import TypeVar, Generic, List, Optional
from abc import ABC, abstractmethod

T = TypeVar('T')


class MaxHeap(Generic[T]):
    """Max Heap implementation using dynamic array"""

    def __init__(self, capacity: int = 10):
        self._data: List[Optional[T]] = [None] * capacity
        self._size: int = 0

    def size(self) -> int:
        """Get heap size"""
        return self._size

    def is_empty(self) -> bool:
        """Check if heap is empty"""
        return self._size == 0

    def add(self, e: T) -> None:
        """Add element to heap"""
        if self._size == len(self._data):
            self._resize(2 * len(self._data))
        self._data[self._size] = e
        self._sift_up(self._size)
        self._size += 1

    def max(self) -> T:
        """Get maximum element (root)"""
        if self._size == 0:
            raise IndexError("Heap is empty")
        return self._data[0]  # type: ignore

    def extract_max(self) -> T:
        """Extract and return maximum element"""
        if self._size == 0:
            raise IndexError("Heap is empty")

        ret = self._data[0]
        self._swap(0, self._size - 1)
        self._data[self._size - 1] = None
        self._size -= 1
        self._sift_down(0)

        if self._size <= len(self._data) // 4 and len(self._data) // 2 != 0:
            self._resize(len(self._data) // 2)

        return ret  # type: ignore

    def replace(self, e: T) -> T:
        """Replace root element with new element and return old root"""
        if self._size == 0:
            raise IndexError("Heap is empty")
        ret = self._data[0]
        self._data[0] = e
        self._sift_down(0)
        return ret  # type: ignore

    def _sift_up(self, index: int) -> None:
        """Sift element up to maintain heap property"""
        while index > 0 and self._data[self._parent(index)] < self._data[index]:  # type: ignore
            self._swap(self._parent(index), index)
            index = self._parent(index)

    def _sift_down(self, index: int) -> None:
        """Sift element down to maintain heap property"""
        while self._left(index) < self._size:
            j = self._left(index)
            # Find the larger child
            if j + 1 < self._size and self._data[j + 1] > self._data[j]:  # type: ignore
                j = self._right(index)

            # If parent is larger or equal, stop
            if self._data[index] >= self._data[j]:  # type: ignore
                break

            self._swap(index, j)
            index = j

    def _parent(self, index: int) -> int:
        """Get parent index"""
        if index <= 0:
            raise ValueError("Index has no parent")
        return (index - 1) // 2

    def _left(self, index: int) -> int:
        """Get left child index"""
        return index * 2 + 1

    def _right(self, index: int) -> int:
        """Get right child index"""
        return index * 2 + 2

    def _swap(self, i: int, j: int) -> None:
        """Swap two elements"""
        self._data[i], self._data[j] = self._data[j], self._data[i]

    def _resize(self, new_capacity: int) -> None:
        """Resize internal array"""
        new_data: List[Optional[T]] = [None] * new_capacity
        for i in range(self._size):
            new_data[i] = self._data[i]
        self._data = new_data

    def to_list(self) -> List[T]:
        """Convert heap to list"""
        return [self._data[i] for i in range(self._size)]  # type: ignore

    def __str__(self) -> str:
        return f"MaxHeap({self.to_list()})"

    def __repr__(self) -> str:
        return self.__str__()

    def __len__(self) -> int:
        return self._size

    def __bool__(self) -> bool:
        return not self.is_empty()


# Priority Queue based on Max Heap
class PriorityQueue(Generic[T]):
    """Priority Queue implementation using Max Heap"""

    def __init__(self, capacity: int = 10):
        self._heap: MaxHeap[T] = MaxHeap[T](capacity)

    def enqueue(self, e: T) -> None:
        """Add element to queue"""
        self._heap.add(e)

    def dequeue(self) -> T:
        """Remove and return highest priority element"""
        return self._heap.extract_max()

    def front(self) -> T:
        """Get highest priority element without removing"""
        return self._heap.max()

    def size(self) -> int:
        """Get queue size"""
        return self._heap.size()

    def is_empty(self) -> bool:
        """Check if queue is empty"""
        return self._heap.is_empty()

    def __str__(self) -> str:
        return f"PriorityQueue({self._heap.to_list()})"


if __name__ == "__main__":
    # Test MaxHeap
    print("Testing MaxHeap:")
    heap = MaxHeap[int]()
    for i in range(10):
        heap.add(i)
    print(f"Heap: {heap}")
    print(f"Max: {heap.max()}")

    for _ in range(5):
        print(f"Extract max: {heap.extract_max()}")
    print(f"After 5 extracts: {heap}")

    heap.replace(100)
    print(f"After replace(100): {heap}")

    # Test PriorityQueue
    print("\nTesting PriorityQueue:")
    pq = PriorityQueue[int]()
    pq.enqueue(3)
    pq.enqueue(1)
    pq.enqueue(4)
    pq.enqueue(2)
    print(f"PriorityQueue: {pq}")
    print(f"Front: {pq.front()}")

    while not pq.is_empty():
        print(f"Dequeue: {pq.dequeue()}")
