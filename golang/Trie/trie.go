package main

import (
	"fmt"
	"strings"
)

// TrieNode 字典树节点
type TrieNode struct {
	isWord bool
	next   map[rune]*TrieNode
}

// newTrieNode 创建一个新的字典树节点
func newTrieNode() *TrieNode {
	return &TrieNode{
		isWord: false,
		next:   make(map[rune]*TrieNode),
	}
}

// Trie 字典树（前缀树）实现
type Trie struct {
	root *TrieNode
	size int
}

// NewTrie 创建一个新的字典树
func NewTrie() *Trie {
	return &Trie{
		root: newTrieNode(),
		size: 0,
	}
}

func (t *Trie) Size() int {
	return t.size
}

// Add 添加单词到字典树
func (t *Trie) Add(word string) {
	cur := t.root
	for _, c := range word {
		if _, ok := cur.next[c]; !ok {
			cur.next[c] = newTrieNode()
		}
		cur = cur.next[c]
	}
	if !cur.isWord {
		cur.isWord = true
		t.size++
	}
}

// Contains 查询单词是否在字典树中
func (t *Trie) Contains(word string) bool {
	cur := t.root
	for _, c := range word {
		if _, ok := cur.next[c]; !ok {
			return false
		}
		cur = cur.next[c]
	}
	return cur.isWord
}

// IsPrefix 查询是否有以指定前缀开头的单词
func (t *Trie) IsPrefix(prefix string) bool {
	cur := t.root
	for _, c := range prefix {
		if _, ok := cur.next[c]; !ok {
			return false
		}
		cur = cur.next[c]
	}
	return true
}

// Remove 从字典树中删除单词
func (t *Trie) Remove(word string) bool {
	var removeHelper func(node *TrieNode, word string, index int) bool
	removeHelper = func(node *TrieNode, word string, index int) bool {
		if index == len(word) {
			if !node.isWord {
				return false
			}
			node.isWord = false
			t.size--
			return len(node.next) == 0
		}

		c := rune(word[index])
		if _, ok := node.next[c]; !ok {
			return false
		}

		shouldDelete := removeHelper(node.next[c], word, index+1)

		if shouldDelete {
			delete(node.next, c)
			return len(node.next) == 0 && !node.isWord
		}

		return false
	}

	return removeHelper(t.root, word, 0)
}

// GetWordsWithPrefix 获取所有以指定前缀开头的单词
func (t *Trie) GetWordsWithPrefix(prefix string) []string {
	var result []string

	// 找到前缀对应的节点
	cur := t.root
	for _, c := range prefix {
		if _, ok := cur.next[c]; !ok {
			return result
		}
		cur = cur.next[c]
	}

	// 从该节点开始收集所有单词
	t.collectWords(cur, prefix, &result)
	return result
}

// collectWords 从指定节点开始收集所有单词
func (t *Trie) collectWords(node *TrieNode, prefix string, result *[]string) {
	if node.isWord {
		*result = append(*result, prefix)
	}

	for c, nextNode := range node.next {
		t.collectWords(nextNode, prefix+string(c), result)
	}
}

// GetAllWords 获取字典树中的所有单词
func (t *Trie) GetAllWords() []string {
	return t.GetWordsWithPrefix("")
}

func (t *Trie) String() string {
	words := t.GetAllWords()
	return fmt.Sprintf("Trie(words=%d, %v)", t.size, words)
}

func main() {
	fmt.Println("=== Go语言字典树示例 ===")

	trie := NewTrie()

	// 添加单词
	fmt.Println("--- 添加单词 ---")
	words := []string{"apple", "app", "application", "banana", "ball", "bat"}
	for _, word := range words {
		trie.Add(word)
		fmt.Printf("添加 '%s'\n", word)
	}

	fmt.Printf("\n字典树大小: %d\n", trie.Size())

	// 查找单词
	fmt.Println("\n--- 查找单词 ---")
	fmt.Printf("Contains 'apple': %t\n", trie.Contains("apple"))
	fmt.Printf("Contains 'app': %t\n", trie.Contains("app"))
	fmt.Printf("Contains 'application': %t\n", trie.Contains("application"))
	fmt.Printf("Contains 'apply': %t\n", trie.Contains("apply"))

	// 前缀查询
	fmt.Println("\n--- 前缀查询 ---")
	fmt.Printf("IsPrefix 'app': %t\n", trie.IsPrefix("app"))
	fmt.Printf("IsPrefix 'xyz': %t\n", trie.IsPrefix("xyz"))

	// 获取指定前缀的单词
	fmt.Println("\n--- 获取指定前缀的单词 ---")
	fmt.Printf("Words with prefix 'app': %v\n", trie.GetWordsWithPrefix("app"))
	fmt.Printf("Words with prefix 'ba': %v\n", trie.GetWordsWithPrefix("ba"))

	// 获取所有单词
	fmt.Printf("\n所有单词: %v\n", trie.GetAllWords())

	// 删除单词
	fmt.Println("\n--- 删除单词 ---")
	fmt.Printf("删除 'app': %t\n", trie.Remove("app"))
	fmt.Printf("删除后 Contains 'app': %t\n", trie.Contains("app"))
	fmt.Printf("删除后 Contains 'apple': %t\n", trie.Contains("apple"))
	fmt.Printf("删除后字典树大小: %d\n", trie.Size())

	fmt.Println("\n=== 示例程序结束 ===")
}
