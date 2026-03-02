"""
Python Data Structures Package
Collection of common data structures implemented in Python
"""

# Linear data structures
from .array import Array
from .linklist import LinkList, Node
from .stack import Stack, ArrayStack, LinkListStack
from .queue import Queue, ArrayQueue, LinkListQueue, LoopQueue

# Set and Map
from .set import Set, LinkListSet
from .map import Map, LinkListMap, MapNode

# Trees
from .bst import BinarySearchTree, BSTNode
from .avl import AVLTree, AVLNode

# Heaps
from .heap import MaxHeap, PriorityQueue

# Hash Table
from .hashtable import HashTable

# Trie
from .trie import Trie, TrieNode

# Union Find
from .unionfind import (
    UnionFind,
    UnionFindV1,
    UnionFindV2,
    UnionFindV3,
    UnionFindV4,
    UnionFindV5
)

# Segment Tree
from .segmenttree import SegmentTree

__all__ = [
    # Linear
    'Array',
    'LinkList',
    'Node',
    'Stack',
    'ArrayStack',
    'LinkListStack',
    'Queue',
    'ArrayQueue',
    'LinkListQueue',
    'LoopQueue',
    # Set and Map
    'Set',
    'LinkListSet',
    'Map',
    'LinkListMap',
    'MapNode',
    # Trees
    'BinarySearchTree',
    'BSTNode',
    'AVLTree',
    'AVLNode',
    # Heaps
    'MaxHeap',
    'PriorityQueue',
    # Hash
    'HashTable',
    # Trie
    'Trie',
    'TrieNode',
    # Union Find
    'UnionFind',
    'UnionFindV1',
    'UnionFindV2',
    'UnionFindV3',
    'UnionFindV4',
    'UnionFindV5',
    # Segment Tree
    'SegmentTree',
]

__version__ = '1.0.0'
