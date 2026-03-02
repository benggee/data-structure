"""
Trie (Prefix Tree) Implementation in Python
Based on Java Trie implementation
"""

from typing import Dict, Optional


class TrieNode:
    """Node class for Trie"""

    def __init__(self, is_word: bool = False):
        self.is_word: bool = is_word
        self.next: Dict[str, 'TrieNode'] = {}

    def __repr__(self) -> str:
        return f"TrieNode(is_word={self.is_word}, children={len(self.next)})"


class Trie:
    """Trie (Prefix Tree) implementation for string storage and retrieval"""

    def __init__(self):
        self._root: TrieNode = TrieNode()
        self._size: int = 0

    def size(self) -> int:
        """Get number of words in trie"""
        return self._size

    def add(self, word: str) -> None:
        """Add word to trie"""
        cur = self._root
        for char in word:
            if char not in cur.next:
                cur.next[char] = TrieNode()
            cur = cur.next[char]
        if not cur.is_word:
            cur.is_word = True
            self._size += 1

    def contains(self, word: str) -> bool:
        """Check if word exists in trie"""
        cur = self._root
        for char in word:
            if char not in cur.next:
                return False
            cur = cur.next[char]
        return cur.is_word

    def is_prefix(self, prefix: str) -> bool:
        """Check if there is any word in the trie that has the given prefix"""
        cur = self._root
        for char in prefix:
            if char not in cur.next:
                return False
            cur = cur.next[char]
        return True

    def remove(self, word: str) -> bool:
        """Remove word from trie"""
        def _remove(node: TrieNode, word: str, index: int) -> bool:
            if index == len(word):
                if not node.is_word:
                    return False
                node.is_word = False
                self._size -= 1
                return len(node.next) == 0

            char = word[index]
            if char not in node.next:
                return False

            should_delete = _remove(node.next[char], word, index + 1)

            if should_delete:
                del node.next[char]
                return len(node.next) == 0 and not node.is_word

            return False

        return _remove(self._root, word, 0)

    def get_words_with_prefix(self, prefix: str) -> list[str]:
        """Get all words with the given prefix"""
        result: list[str] = []

        # Find the node corresponding to prefix
        cur = self._root
        for char in prefix:
            if char not in cur.next:
                return result
            cur = cur.next[char]

        # Collect all words from this node
        self._collect_words(cur, prefix, result)
        return result

    def _collect_words(self, node: TrieNode, prefix: str, result: list[str]) -> None:
        """Collect all words from a given node"""
        if node.is_word:
            result.append(prefix)

        for char, next_node in node.next.items():
            self._collect_words(next_node, prefix + char, result)

    def get_all_words(self) -> list[str]:
        """Get all words in the trie"""
        return self.get_words_with_prefix("")

    def count_words_starting_with(self, prefix: str) -> int:
        """Count how many words start with the given prefix"""
        return len(self.get_words_with_prefix(prefix))

    def __str__(self) -> str:
        return f"Trie(words={self._size})"

    def __repr__(self) -> str:
        return self.__str__()

    def __len__(self) -> int:
        return self._size

    def __contains__(self, word: str) -> bool:
        return self.contains(word)


if __name__ == "__main__":
    # Test Trie
    print("Testing Trie:")
    trie = Trie()

    words = ["apple", "app", "application", "banana", "ball", "bat"]
    for word in words:
        trie.add(word)
        print(f"Added '{word}'")

    print(f"\nTrie size: {trie.size()}")
    print(f"Contains 'apple': {trie.contains('apple')}")
    print(f"Contains 'app': {trie.contains('app')}")
    print(f"Contains 'application': {trie.contains('application')}")
    print(f"Contains 'apply': {trie.contains('apply')}")

    print(f"\nIs 'app' a prefix: {trie.is_prefix('app')}")
    print(f"Is 'xyz' a prefix: {trie.is_prefix('xyz')}")

    print(f"\nWords starting with 'app': {trie.get_words_with_prefix('app')}")
    print(f"Words starting with 'ba': {trie.get_words_with_prefix('ba')}")

    print(f"\nAll words: {trie.get_all_words()}")

    print(f"\nRemoving 'app': {trie.remove('app')}")
    print(f"Contains 'app' after removal: {trie.contains('app')}")
    print(f"Contains 'apple' after 'app' removal: {trie.contains('apple')}")

    print(f"\nTrie size after removal: {trie.size()}")
    print(f"All words after removal: {trie.get_all_words()}")
