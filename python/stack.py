"""
Stack Implementation in Python
Based on Java Stack implementation
"""

from typing import TypeVar, Generic, Optional
from array import Array

T = TypeVar('T')


class Stack(Generic[T]):
    """Stack interface"""

    def size(self) -> int:
        """Get stack size"""
        raise NotImplementedError

    def is_empty(self) -> bool:
        """Check if stack is empty"""
        raise NotImplementedError

    def push(self, e: T) -> None:
        """Push element onto stack"""
        raise NotImplementedError

    def pop(self) -> T:
        """Pop element from stack"""
        raise NotImplementedError

    def peek(self) -> Optional[T]:
        """Get top element without removing"""
        raise NotImplementedError


class ArrayStack(Stack[T]):
    """Stack implementation using dynamic array"""

    def __init__(self):
        self._array: Array[T] = Array[T]()

    def size(self) -> int:
        return self._array.size()

    def is_empty(self) -> bool:
        return self._array.is_empty()

    def push(self, e: T) -> None:
        self._array.add_last(e)

    def pop(self) -> T:
        if self.is_empty():
            raise IndexError("Stack is empty")
        return self._array.remove_last()

    def peek(self) -> Optional[T]:
        if self.is_empty():
            return None
        return self._array.get(self._array.size() - 1)

    def __str__(self) -> str:
        return f"ArrayStack: [{', '.join(str(self._array.get(i)) for i in range(self._array.size()))}] <- top"

    def __repr__(self) -> str:
        return self.__str__()


class LinkListStack(Stack[T]):
    """Stack implementation using linked list"""

    from typing import List
    from linklist import LinkList

    def __init__(self):
        self._list: LinkList[T] = LinkList[T]()

    def size(self) -> int:
        return self._list.get_size()

    def is_empty(self) -> bool:
        return self._list.is_empty()

    def push(self, e: T) -> None:
        self._list.add_first(e)

    def pop(self) -> T:
        if self.is_empty():
            raise IndexError("Stack is empty")
        return self._list.delete_first()

    def peek(self) -> Optional[T]:
        if self.is_empty():
            return None
        return self._list.get_first()

    def __str__(self) -> str:
        return f"LinkListStack: top <- [{self._list.__str__()}]"


if __name__ == "__main__":
    # Test ArrayStack
    print("Testing ArrayStack:")
    stack = ArrayStack[int]()
    for i in range(5):
        stack.push(i)
        print(f"Push {i}: {stack}")

    print(f"Peek: {stack.peek()}")
    print(f"Pop: {stack.pop()}")
    print(f"Pop: {stack.pop()}")
    print(f"After pops: {stack}")

    # Test LinkListStack
    print("\nTesting LinkListStack:")
    ll_stack = LinkListStack[int]()
    for i in range(5):
        ll_stack.push(i)
        print(f"Push {i}: {ll_stack}")

    print(f"Peek: {ll_stack.peek()}")
    print(f"Pop: {ll_stack.pop()}")
    print(f"After pop: {ll_stack}")
