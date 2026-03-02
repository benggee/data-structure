package avltree

import (
	"math"
)

// AVLTree AVL平衡二叉搜索树
// AVL树是一种自平衡的二叉搜索树，通过旋转操作保持树的平衡
// 平衡因子定义为：左子树高度 - 右子树高度
// 当平衡因子的绝对值大于1时，需要进行旋转操作来恢复平衡
type AVLTree[K comparable, V any] struct {
	root *Node[K, V] // 根节点
	size int         // 树中节点的数量
}

// Node AVL树的节点结构
type Node[K comparable, V any] struct {
	key    K           // 节点的键
	value  V           // 节点的值
	left   *Node[K, V] // 左子节点
	right  *Node[K, V] // 右子节点
	height int         // 节点的高度
}

// NewAVLTree 创建一个新的AVL树
func NewAVLTree[K comparable, V any]() *AVLTree[K, V] {
	return &AVLTree[K, V]{
		root: nil,
		size: 0,
	}
}

// newNode 创建一个新的节点
func newNode[K comparable, V any](key K, value V) *Node[K, V] {
	return &Node[K, V]{
		key:    key,
		value:  value,
		left:   nil,
		right:  nil,
		height: 1, // 新节点的高度为1
	}
}

// Remove 从AVL树中删除指定键的节点
// 返回被删除节点的值，如果键不存在则返回零值
func (t *AVLTree[K, V]) Remove(key K) (V, bool) {
	var zero V
	node := t.getNode(t.root, key)
	if node == nil {
		return zero, false
	}

	oldValue := node.value
	t.root = t.remove(t.root, key)
	return oldValue, true
}

// Contains 检查AVL树中是否包含指定的键
func (t *AVLTree[K, V]) Contains(key K) bool {
	return t.getNode(t.root, key) != nil
}

// Get 获取指定键对应的值
// 如果键不存在，返回零值和false
func (t *AVLTree[K, V]) Get(key K) (V, bool) {
	var zero V
	node := t.getNode(t.root, key)
	if node == nil {
		return zero, false
	}
	return node.value, true
}

// Set 设置指定键的值，如果键已存在则更新值
func (t *AVLTree[K, V]) Set(key K, value V) {
	t.root = t.add(t.root, key, value)
}

// Size 返回AVL树中节点的数量
func (t *AVLTree[K, V]) Size() int {
	return t.size
}

// Empty 检查AVL树是否为空
func (t *AVLTree[K, V]) Empty() bool {
	return t.size == 0
}

// IsBST 判断是否是二叉搜索树
// 利用二叉搜索树的特点：中序遍历得到的是一个从小到大的有序列表
func (t *AVLTree[K, V]) IsBST() bool {
	var list []K
	t.inorderTraversal(t.root, &list)

	// 检查中序遍历结果是否有序
	for i := 1; i < len(list); i++ {
		if !t.less(list[i-1], list[i]) {
			return false
		}
	}
	return true
}

// inorderTraversal 中序遍历，将节点键值按顺序添加到列表中
func (t *AVLTree[K, V]) inorderTraversal(node *Node[K, V], list *[]K) {
	if node == nil {
		return
	}
	t.inorderTraversal(node.left, list)
	*list = append(*list, node.key)
	t.inorderTraversal(node.right, list)
}

// IsAVL 判断是否是AVL树
// 利用平衡因子不能大于1的特点
func (t *AVLTree[K, V]) IsAVL() bool {
	return t.isAVL(t.root)
}

// isAVL 递归检查每个节点的平衡因子
func (t *AVLTree[K, V]) isAVL(node *Node[K, V]) bool {
	if node == nil {
		return true
	}

	// 检查当前节点的平衡因子
	if math.Abs(float64(t.getBalanceFactor(node))) > 1 {
		return false
	}

	// 递归检查左右子树
	return t.isAVL(node.left) && t.isAVL(node.right)
}

// getHeight 获取节点的高度
// 空节点的高度为0
func (t *AVLTree[K, V]) getHeight(node *Node[K, V]) int {
	if node == nil {
		return 0
	}
	return node.height
}

// getBalanceFactor 计算节点的平衡因子
// 平衡因子 = 左子树高度 - 右子树高度
func (t *AVLTree[K, V]) getBalanceFactor(node *Node[K, V]) int {
	if node == nil {
		return 0
	}
	return t.getHeight(node.left) - t.getHeight(node.right)
}

// rightRotate 右旋转操作
// 用于修复LL情况（左左情况）
//
//	     y                              x
//	    / \                           /   \
//	   x   T4     向右旋转 (y)        z     y
//	  / \       - - - - - - - ->    / \   / \
//	 z   T3                       T1  T2 T3 T4
//	/ \
//
// T1   T2
func (t *AVLTree[K, V]) rightRotate(y *Node[K, V]) *Node[K, V] {
	x := y.left
	T3 := x.right

	// 执行旋转
	x.right = y
	y.left = T3

	// 更新高度
	y.height = t.max(t.getHeight(y.right), t.getHeight(y.left)) + 1
	x.height = t.max(t.getHeight(x.right), t.getHeight(x.left)) + 1

	return x
}

// leftRotate 左旋转操作
// 用于修复RR情况（右右情况）
//
//	  y                             x
//	/  \                          /   \
//
// T1   x      向左旋转 (y)       y     z
//
//	  / \   - - - - - - - ->   / \   / \
//	T2  z                     T1 T2 T3 T4
//	   / \
//	  T3 T4
func (t *AVLTree[K, V]) leftRotate(y *Node[K, V]) *Node[K, V] {
	x := y.right
	T2 := x.left

	// 执行旋转
	x.left = y
	y.right = T2

	// 更新高度
	y.height = t.max(t.getHeight(y.left), t.getHeight(y.right)) + 1
	x.height = t.max(t.getHeight(x.left), t.getHeight(x.right)) + 1

	return x
}

// add 向AVL树中添加节点
// 这是一个递归函数，在添加节点后会自动进行平衡操作
func (t *AVLTree[K, V]) add(node *Node[K, V], key K, value V) *Node[K, V] {
	if node == nil {
		t.size++
		return newNode(key, value)
	}

	// 根据键值比较决定插入位置
	if t.less(key, node.key) {
		node.left = t.add(node.left, key, value)
	} else if t.less(node.key, key) {
		node.right = t.add(node.right, key, value)
	} else {
		// 键已存在，更新值
		node.value = value
		return node
	}

	// 重新计算当前节点的高度
	node.height = 1 + t.max(t.getHeight(node.left), t.getHeight(node.right))

	// 计算平衡因子
	balanceFactor := t.getBalanceFactor(node)

	// LL情况：左左不平衡，需要右旋转
	// 如果节点的高度因子（左高度-右高度）大于1
	// 并且，左子树的的高度因子>=0(也就是至少有一个元素)
	// 则进行右旋转
	if balanceFactor > 1 && t.getBalanceFactor(node.left) >= 0 {
		return t.rightRotate(node)
	}

	// RR情况：右右不平衡，需要左旋转
	// 如果节点的高度因子 （左高度-右高度）小于-1
	// 并且，右子树的高度因子<=0 (也就是至少有一个元素)
	// 则进行左旋转
	if balanceFactor < -1 && t.getBalanceFactor(node.right) <= 0 {
		return t.leftRotate(node)
	}

	// LR情况：左右不平衡，需要先左旋转再右旋转
	// 如果节点不平衡是因为左子节点的右子节点高度太高，则先将左子节点做一次左旋转，然后进行右旋转
	if balanceFactor > 1 && t.getBalanceFactor(node.left) < 0 {
		node.left = t.leftRotate(node.left)
		return t.rightRotate(node)
	}

	// RL情况：右左不平衡，需要先右旋转再左旋转
	// 如果节点不平衡是因为右子节点的左子节点高度太高，则先将右节点做一次右旋转，然后进行左旋转
	if balanceFactor < -1 && t.getBalanceFactor(node.right) > 0 {
		node.right = t.rightRotate(node.right)
		return t.leftRotate(node)
	}

	return node
}

// remove 从AVL树中删除节点
// 这是一个递归函数，在删除节点后会自动进行平衡操作
func (t *AVLTree[K, V]) remove(node *Node[K, V], key K) *Node[K, V] {
	if node == nil {
		return nil
	}

	var retNode *Node[K, V]
	if t.less(key, node.key) {
		node.left = t.remove(node.left, key)
		retNode = node
	} else if t.less(node.key, key) {
		node.right = t.remove(node.right, key)
		retNode = node
	} else {
		// 找到要删除的节点
		if node.left == nil {
			// 左子树为空，返回右子树
			tmpNode := node.right
			node.right = nil
			t.size--
			retNode = tmpNode
		} else if node.right == nil {
			// 右子树为空，返回左子树
			tmpNode := node.left
			node.left = nil
			t.size--
			retNode = tmpNode
		} else {
			// 左右子树都不为空
			// 找到比删除节点大的最小节点，这个节点应该在当前节点的右子树
			minNode := t.min(node.right)
			minNode.right = t.remove(node.right, minNode.key)
			minNode.left = node.left

			node.left = nil
			node.right = nil
			retNode = minNode
		}
	}

	if retNode == nil {
		return nil
	}

	// 重新计算高度
	retNode.height = 1 + t.max(t.getHeight(retNode.left), t.getHeight(retNode.right))

	// 计算平衡因子
	balanceFactor := t.getBalanceFactor(retNode)

	// LL情况
	if balanceFactor > 1 && t.getBalanceFactor(retNode.left) >= 0 {
		return t.rightRotate(retNode)
	}

	// RR情况
	if balanceFactor < -1 && t.getBalanceFactor(retNode.right) <= 0 {
		return t.leftRotate(retNode)
	}

	// LR情况
	if balanceFactor > 1 && t.getBalanceFactor(retNode.left) < 0 {
		retNode.left = t.leftRotate(retNode.left)
		return t.rightRotate(retNode)
	}

	// RL情况
	if balanceFactor < -1 && t.getBalanceFactor(retNode.right) > 0 {
		retNode.right = t.rightRotate(retNode.right)
		return t.leftRotate(retNode)
	}

	return retNode
}

// min 找到以指定节点为根的子树中的最小节点
func (t *AVLTree[K, V]) min(node *Node[K, V]) *Node[K, V] {
	if node.left == nil {
		return node
	}
	return t.min(node.left)
}

// getNode 根据键查找节点
func (t *AVLTree[K, V]) getNode(node *Node[K, V], key K) *Node[K, V] {
	if node == nil {
		return nil
	}

	if t.equal(key, node.key) {
		return node
	} else if t.less(key, node.key) {
		return t.getNode(node.left, key)
	} else {
		return t.getNode(node.right, key)
	}
}

// less 比较两个键的大小
// 这里使用类型断言，要求K类型实现了Comparable接口
func (t *AVLTree[K, V]) less(a, b K) bool {
	// 这里需要根据具体的K类型来实现比较逻辑
	// 在实际使用中，可能需要通过接口或泛型约束来处理
	// 为了简化，这里假设K是基本类型或实现了比较方法
	switch v := any(a).(type) {
	case string:
		return v < any(b).(string)
	case int:
		return v < any(b).(int)
	case int64:
		return v < any(b).(int64)
	case float64:
		return v < any(b).(float64)
	default:
		// 对于其他类型，可以尝试使用反射或要求实现特定接口
		panic("unsupported key type for comparison")
	}
}

// equal 判断两个键是否相等
func (t *AVLTree[K, V]) equal(a, b K) bool {
	return any(a) == any(b)
}

// max 返回两个整数中的较大值
func (t *AVLTree[K, V]) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
