"""
Singly Linked List Implementation in Python
Based on Java LinkList implementation
"""

from typing import TypeVar, Generic, Optional

T = TypeVar('T')


class Node(Generic[T]):
    """Node class for Linked List"""

    def __init__(self, e: Optional[T] = None, next_node: Optional['Node[T]'] = None):
        self.e: Optional[T] = e
        self.next: Optional['Node[T]'] = next_node

    def __str__(self) -> str:
        return str(self.e) if self.e is not None else "None"

    def __repr__(self) -> str:
        return f"Node({self.e})"


class LinkList(Generic[T]):
    """Singly Linked List with dummy head node"""

    def __init__(self):
        self._dummy_head: Node[T] = Node[T]()
        self._size: int = 0

    def get_size(self) -> int:
        """Get list size"""
        return self._size

    def is_empty(self) -> bool:
        """Check if list is empty"""
        return self._size == 0

    def add(self, index: int, e: T) -> None:
        """Add element at specified index"""
        if index < 0 or index > self._size:
            raise IndexError("Index out of range")

        prev = self._dummy_head
        for _ in range(index):
            prev = prev.next

        node = Node[T](e, prev.next)
        prev.next = node
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

        cur = self._dummy_head.next
        for _ in range(index):
            cur = cur.next
        return cur.e

    def get_first(self) -> Optional[T]:
        """Get first element"""
        return self.get(0) if self._size > 0 else None

    def get_last(self) -> Optional[T]:
        """Get last element"""
        return self.get(self._size - 1) if self._size > 0 else None

    def set(self, index: int, e: T) -> None:
        """Set element at specified index"""
        if index < 0 or index >= self._size:
            raise IndexError("Index out of range")

        cur = self._dummy_head.next
        for _ in range(index):
            cur = cur.next
        cur.e = e

    def find(self, e: T) -> bool:
        """Check if element exists in list"""
        cur = self._dummy_head.next
        while cur is not None:
            if cur.e == e:
                return True
            cur = cur.next
        return False

    def delete(self, index: int) -> T:
        """Remove element at specified index"""
        if index < 0 or index >= self._size:
            raise IndexError("Index out of range")

        prev = self._dummy_head
        for _ in range(index):
            prev = prev.next

        del_node = prev.next
        prev.next = del_node.next
        del_node.next = None
        self._size -= 1
        return del_node.e

    def delete_first(self) -> T:
        """Remove first element"""
        return self.delete(0)

    def delete_last(self) -> T:
        """Remove last element"""
        return self.delete(self._size - 1)

    def remove_element(self, e: T) -> bool:
        """Remove specified element"""
        prev = self._dummy_head
        while prev.next is not None:
            if prev.next.e == e:
                del_node = prev.next
                prev.next = del_node.next
                del_node.next = None
                self._size -= 1
                return True
            prev = prev.next
        return False

    def __str__(self) -> str:
        """String representation"""
        result = []
        cur = self._dummy_head.next
        while cur is not None:
            result.append(str(cur.e))
            cur = cur.next
        return "->".join(result) + "->NULL"

    def __repr__(self) -> str:
        return f"LinkList({self._size} elements: {self.__str__()})"

    def __len__(self) -> int:
        return self._size

    def __iter__(self):
        """Make the list iterable"""
        self._iter_current = self._dummy_head.next
        return self

    def __next__(self):
        if self._iter_current is None:
            raise StopIteration
        value = self._iter_current.e
        self._iter_current = self._iter_current.next
        return value


if __name__ == "__main__":
    # Test LinkList
    ll = LinkList[int]()
    for i in range(6):
        ll.add_first(i)
        print(ll)

    print(f"Size: {ll.get_size()}")
    print(f"Get first: {ll.get_first()}")
    print(f"Get last: {ll.get_last()}")
    print(f"Find 3: {ll.find(3)}")

    ll.add(3, 8888)
    print(f"After add(3, 8888): {ll}")

    ll.delete(3)
    print(f"After delete index 3: {ll}")

    ll.delete_first()
    print(f"After delete first: {ll}")

    # Test iteration
    print("Iteration:")
    for item in ll:
        print(item, end=" ")
    print()
