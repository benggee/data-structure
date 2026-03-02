"""
Set Implementation in Python
Based on Java Set implementation
"""

from typing import TypeVar, Generic, Protocol, runtime_checkable
from abc import ABC, abstractmethod
from linklist import LinkList

T = TypeVar('T')


class Set(Generic[T], ABC):
    """Set interface"""

    @abstractmethod
    def add(self, e: T) -> None:
        """Add element to set"""
        pass

    @abstractmethod
    def contains(self, e: T) -> bool:
        """Check if element exists in set"""
        pass

    @abstractmethod
    def remove(self, e: T) -> None:
        """Remove element from set"""
        pass

    @abstractmethod
    def get_size(self) -> int:
        """Get set size"""
        pass

    @abstractmethod
    def is_empty(self) -> bool:
        """Check if set is empty"""
        pass


class LinkListSet(Set[T]):
    """Set implementation using linked list"""

    def __init__(self):
        self._list: LinkList[T] = LinkList[T]()

    def get_size(self) -> int:
        return self._list.get_size()

    def is_empty(self) -> bool:
        return self._list.is_empty()

    def add(self, e: T) -> None:
        """Add element to set (no duplicates)"""
        if not self._list.find(e):
            self._list.add_first(e)

    def contains(self, e: T) -> bool:
        return self._list.find(e)

    def remove(self, e: T) -> None:
        self._list.remove_element(e)

    def __str__(self) -> str:
        return f"LinkListSet({{{self._list.__str__()}}})"

    def __repr__(self) -> str:
        return self.__str__()

    def __len__(self) -> int:
        return self.get_size()

    def __contains__(self, item: T) -> bool:
        return self.contains(item)

    def __iter__(self):
        return iter(self._list)


if __name__ == "__main__":
    # Test LinkListSet
    print("Testing LinkListSet:")
    s = LinkListSet[int]()
    s.add(1)
    s.add(2)
    s.add(3)
    s.add(1)  # Duplicate, won't be added
    print(f"Set: {s}")
    print(f"Size: {s.get_size()}")
    print(f"Contains 2: {s.contains(2)}")
    print(f"Contains 5: {s.contains(5)}")

    s.remove(2)
    print(f"After remove 2: {s}")

    # Test iteration
    print("Iteration:")
    for item in s:
        print(item, end=" ")
    print()
