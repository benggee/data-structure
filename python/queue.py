"""
Queue Implementation in Python
Based on Java Queue implementation
"""

from typing import TypeVar, Generic, Optional
from array import Array
from linklist import LinkList

T = TypeVar('T')


class Queue(Generic[T]):
    """Queue interface"""

    def size(self) -> int:
        """Get queue size"""
        raise NotImplementedError

    def is_empty(self) -> bool:
        """Check if queue is empty"""
        raise NotImplementedError

    def enqueue(self, e: T) -> None:
        """Add element to queue"""
        raise NotImplementedError

    def dequeue(self) -> T:
        """Remove element from queue"""
        raise NotImplementedError

    def front(self) -> Optional[T]:
        """Get front element without removing"""
        raise NotImplementedError


class ArrayQueue(Queue[T]):
    """Queue implementation using dynamic array"""

    def __init__(self, capacity: int = 10):
        self._array: Array[T] = Array[T](capacity)

    def size(self) -> int:
        return self._array.size()

    def is_empty(self) -> bool:
        return self._array.is_empty()

    def enqueue(self, e: T) -> None:
        self._array.add_last(e)

    def dequeue(self) -> T:
        if self.is_empty():
            raise IndexError("Queue is empty")
        return self._array.remove_first()

    def front(self) -> Optional[T]:
        if self.is_empty():
            return None
        return self._array.get(0)

    def __str__(self) -> str:
        items = ", ".join(str(self._array.get(i)) for i in range(self._array.size()))
        return f"front:[{items}]tail"


class LinkListQueue(Queue[T]):
    """Queue implementation using linked list"""

    def __init__(self):
        self._list: LinkList[T] = LinkList[T]()

    def size(self) -> int:
        return self._list.get_size()

    def is_empty(self) -> bool:
        return self._list.is_empty()

    def enqueue(self, e: T) -> None:
        self._list.add_last(e)

    def dequeue(self) -> T:
        if self.is_empty():
            raise IndexError("Queue is empty")
        return self._list.delete_first()

    def front(self) -> Optional[T]:
        if self.is_empty():
            return None
        return self._list.get_first()

    def __str__(self) -> str:
        return f"front:[{self._list.__str__()}]tail"


class LoopQueue(Queue[T]):
    """Loop queue implementation using circular array"""

    def __init__(self, capacity: int = 10):
        self._data: list[Optional[T]] = [None] * (capacity + 1)  # +1 for distinguishing full/empty
        self._front: int = 0
        self._tail: int = 0
        self._size: int = 0

    def capacity(self) -> int:
        return len(self._data) - 1

    def size(self) -> int:
        return self._size

    def is_empty(self) -> bool:
        return self._front == self._tail

    def enqueue(self, e: T) -> None:
        if (self._tail + 1) % len(self._data) == self._front:
            self._resize(self.capacity() * 2)

        self._data[self._tail] = e
        self._tail = (self._tail + 1) % len(self._data)
        self._size += 1

    def dequeue(self) -> T:
        if self.is_empty():
            raise IndexError("Queue is empty")

        ret = self._data[self._front]
        self._data[self._front] = None
        self._front = (self._front + 1) % len(self._data)
        self._size -= 1

        if self._size <= self.capacity() // 4 and self.capacity() // 2 != 0:
            self._resize(self.capacity() // 2)

        return ret

    def front(self) -> Optional[T]:
        if self.is_empty():
            return None
        return self._data[self._front]

    def _resize(self, new_capacity: int) -> None:
        new_data: list[Optional[T]] = [None] * (new_capacity + 1)
        for i in range(self._size):
            new_data[i] = self._data[(self._front + i) % len(self._data)]
        self._data = new_data
        self._front = 0
        self._tail = self._size

    def __str__(self) -> str:
        if self._tail >= self._front:
            items = [str(self._data[i]) for i in range(self._front, self._tail)]
        else:
            items = [str(self._data[i]) for i in range(self._front, len(self._data))]
            items += [str(self._data[i]) for i in range(0, self._tail)]
        return f"front:[{', '.join(filter(None, items))}]tail"


if __name__ == "__main__":
    # Test ArrayQueue
    print("Testing ArrayQueue:")
    queue = ArrayQueue[int]()
    for i in range(5):
        queue.enqueue(i)
        print(f"Enqueue {i}: {queue}")

    print(f"Front: {queue.front()}")
    print(f"Dequeue: {queue.dequeue()}")
    print(f"After dequeue: {queue}")

    # Test LoopQueue
    print("\nTesting LoopQueue:")
    loop_queue = LoopQueue[int]()
    for i in range(10):
        loop_queue.enqueue(i)
    print(f"After 10 enqueues: {loop_queue}")

    for _ in range(5):
        loop_queue.dequeue()
    print(f"After 5 dequeues: {loop_queue}")

    # Test LinkListQueue
    print("\nTesting LinkListQueue:")
    ll_queue = LinkListQueue[int]()
    for i in range(5):
        ll_queue.enqueue(i)
    print(f"After 5 enqueues: {ll_queue}")

    print(f"Dequeue: {ll_queue.dequeue()}")
    print(f"After dequeue: {ll_queue}")
