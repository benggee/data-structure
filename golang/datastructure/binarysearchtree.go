package datastructure

import (
	"fmt"
)

type binarySearchTreeNode struct {
	data interface{}
	score int64
	left *binarySearchTreeNode
	right *binarySearchTreeNode
}

type binarySearchTree struct {
	root *binarySearchTreeNode
	size int
}

func BinarySearchTree() *binarySearchTree {
	return &binarySearchTree{root:nil, size:0}
}

func (b *binarySearchTree) Add(data interface{}, score int64) {
	b.root = b.add(b.root, data, score)
}

func (b *binarySearchTree) add(node *binarySearchTreeNode, e interface{}, score int64) *binarySearchTreeNode {
	if (node == nil) {
		b.size++
		return &binarySearchTreeNode{data:e, score:score, left: nil, right: nil}
	}
	if score < node.score {
		node.left = b.add(node.left, e, score)
	}
	if score > node.score {
		node.right = b.add(node.right, e, score)
	}
	return node
}

// 前序遍历
func (b *binarySearchTree) PreOrder() {
	b.preOrder(b.root)
}

func (b *binarySearchTree) preOrder(node *binarySearchTreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	b.preOrder(node.left)
	b.preOrder(node.right)
}

// 百递归前序遍历
func (b *binarySearchTree) PreOrderNR() {
	b.preOrderNR(b.root)
}

func (b *binarySearchTree) preOrderNR(node *binarySearchTreeNode) {
	if b.root == nil {
		return
	}
	s := Stack()
	s.Push(b.root)

	for s.Size() > 0 {
		cur, _ := s.Pop().(*binarySearchTreeNode)
		fmt.Println(cur.data)

		if cur.left != nil {
			s.Push(cur.left)
		}
		if cur.right != nil {
			s.Push(cur.right)
		}
	}
}

// 中序遍历
func (b *binarySearchTree) InOrder() {
	b.inOrder(b.root)
}

func (b *binarySearchTree) inOrder(node *binarySearchTreeNode) {
	if node == nil {
		return
	}
	b.inOrder(node.left)
	fmt.Println(node.data)
	b.inOrder(node.right)
}

// 后序遍历
func (b *binarySearchTree) PostOrder() {
	b.postOrder(b.root)
}

func (b *binarySearchTree) postOrder(node *binarySearchTreeNode) {
	if node == nil {
		return
	}
	b.postOrder(node.left)
	b.postOrder(node.right)
	fmt.Println(node.data)
}

// 层序遍历
func (b *binarySearchTree) LevelOrder() {
	if b.root == nil {
		return
	}
	q := Queue()
	q.Push(*b.root)
	for q.Size() > 0 {
		cur, _ := q.Pop().(binarySearchTreeNode)
		fmt.Println(cur.data)

		if cur.left != nil {
			q.Push(*cur.left)
		}
		if cur.right != nil {
			q.Push(*cur.right)
		}
	}
}

// 获取最小元素
func (b *binarySearchTree) Min() interface{} {
	return b.min(b.root).data
}

func (b *binarySearchTree) min(node *binarySearchTreeNode) *binarySearchTreeNode {
	if node.left == nil {
		return node
	}
	return b.min(node.left)
}

// 获取最大值
func (b *binarySearchTree) Max() interface{} {
	return b.max(b.root)
}

func (b *binarySearchTree) max(node *binarySearchTreeNode) *binarySearchTreeNode {
	if node.right == nil {
		return node
	}
	return b.max(node.right)
}

// 删除最小元素
func (b *binarySearchTree) DelMin() interface{} {
	min := b.Min()
	b.root = b.delMin(b.root)
	return min
}

func (b *binarySearchTree) delMin(node *binarySearchTreeNode) *binarySearchTreeNode {
	if node.left == nil {
		tmpNode := node.right
		node.right = nil
		b.size--
		return tmpNode
	}
	node.left = b.delMin(node.left)
	return node
}

// 删除最大元素
func (b *binarySearchTree) DelMax() interface{} {
	max := b.Max()
	b.root = b.delMax(b.root)
	return max
}

func (b *binarySearchTree) delMax(node *binarySearchTreeNode) *binarySearchTreeNode {
	if node.right == nil {
		tmp := node.left
		node.right = nil
		b.size--
		return tmp
	}
	node.right = b.delMax(node.right)
	return node
}

// 删除任意元素（通过score）
func (b *binarySearchTree) Del(score int64) {
	b.root = b.del(b.root, score)
}

func (b *binarySearchTree) del(node *binarySearchTreeNode, score int64) *binarySearchTreeNode {
	if node == nil {
		return nil
	}
	if node.score > score {
		node.left = b.del(node.left, score)
		return node
	}
	if node.score < score {
		node.right = b.del(node.right, score)
		return node
	}
	// 找到元素
	if node.left == nil {
		tmp := node.right
		node.right = nil
		b.size--
		return tmp
	}
	if node.right == nil {
		tmp := node.left
		node.left = nil
		b.size--
		return tmp
	}

	// 到这里说明左右子树都不为nil
	// 将右子树最小元素放到当前要删除的元素位置
	minLeftNode := b.min(node.right)
	minLeftNode.right = b.delMin(node.right)
	minLeftNode.left = node.left

	node.left = nil
	node.right = nil

	return minLeftNode
}