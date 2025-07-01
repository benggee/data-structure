package avltree

// BinarySearchTree 二叉搜索树
// 二叉搜索树是一种特殊的二叉树，其中每个节点的值大于其左子树中所有节点的值，
// 小于其右子树中所有节点的值
type BinarySearchTree[K comparable, V any] struct {
	root *BSTNode[K, V] // 根节点
	size int            // 树中节点的数量
}

// BSTNode 二叉搜索树的节点结构
type BSTNode[K comparable, V any] struct {
	key   K              // 节点的键
	value V              // 节点的值
	left  *BSTNode[K, V] // 左子节点
	right *BSTNode[K, V] // 右子节点
}

// NewBinarySearchTree 创建一个新的二叉搜索树
func NewBinarySearchTree[K comparable, V any]() *BinarySearchTree[K, V] {
	return &BinarySearchTree[K, V]{
		root: nil,
		size: 0,
	}
}

// newBSTNode 创建一个新的二叉搜索树节点
func newBSTNode[K comparable, V any](key K, value V) *BSTNode[K, V] {
	return &BSTNode[K, V]{
		key:   key,
		value: value,
		left:  nil,
		right: nil,
	}
}

// Remove 从二叉搜索树中删除指定键的节点
// 返回被删除节点的值，如果键不存在则返回零值
func (t *BinarySearchTree[K, V]) Remove(key K) (V, bool) {
	var zero V
	node := t.getNode(t.root, key)
	if node == nil {
		return zero, false
	}

	oldValue := node.value
	t.root = t.remove(t.root, key)
	return oldValue, true
}

// Contains 检查二叉搜索树中是否包含指定的键
func (t *BinarySearchTree[K, V]) Contains(key K) bool {
	return t.getNode(t.root, key) != nil
}

// Get 获取指定键对应的值
// 如果键不存在，返回零值和false
func (t *BinarySearchTree[K, V]) Get(key K) (V, bool) {
	var zero V
	node := t.getNode(t.root, key)
	if node == nil {
		return zero, false
	}
	return node.value, true
}

// Set 设置指定键的值，如果键已存在则更新值
func (t *BinarySearchTree[K, V]) Set(key K, value V) {
	t.root = t.add(t.root, key, value)
}

// Size 返回二叉搜索树中节点的数量
func (t *BinarySearchTree[K, V]) Size() int {
	return t.size
}

// Empty 检查二叉搜索树是否为空
func (t *BinarySearchTree[K, V]) Empty() bool {
	return t.size == 0
}

// add 向二叉搜索树中添加节点
// 这是一个递归函数，根据键值比较决定插入位置
func (t *BinarySearchTree[K, V]) add(node *BSTNode[K, V], key K, value V) *BSTNode[K, V] {
	if node == nil {
		t.size++
		return newBSTNode(key, value)
	}

	// 根据键值比较决定插入位置
	if t.less(key, node.key) {
		node.left = t.add(node.left, key, value)
	} else if t.less(node.key, key) {
		node.right = t.add(node.right, key, value)
	} else {
		// 键已存在，更新值
		node.value = value
	}
	return node
}

// remove 从二叉搜索树中删除节点
// 这是一个递归函数，处理三种删除情况：
// 1. 要删除的节点是叶子节点
// 2. 要删除的节点只有一个子节点
// 3. 要删除的节点有两个子节点
func (t *BinarySearchTree[K, V]) remove(node *BSTNode[K, V], key K) *BSTNode[K, V] {
	if node == nil {
		return nil
	}

	if t.less(key, node.key) {
		node.left = t.remove(node.left, key)
		return node
	} else if t.less(node.key, key) {
		node.right = t.remove(node.right, key)
		return node
	} else {
		// 找到要删除的节点
		if node.left == nil {
			// 左子树为空，返回右子树
			tmpNode := node.right
			node.right = nil
			t.size--
			return tmpNode
		}
		if node.right == nil {
			// 右子树为空，返回左子树
			tmpNode := node.left
			node.left = nil
			t.size--
			return tmpNode
		}

		// 左右子树都不为空
		// 找到比删除节点大的最小节点，这个节点应该在当前节点的右子树
		minNode := t.min(node.right)
		minNode.right = t.removeMin(node.right)
		minNode.left = node.left

		node.left = nil
		node.right = nil
		return minNode
	}
}

// removeMin 删除以指定节点为根的子树中的最小节点
// 这是一个递归函数，用于辅助删除操作
func (t *BinarySearchTree[K, V]) removeMin(node *BSTNode[K, V]) *BSTNode[K, V] {
	if node == nil {
		return nil
	}
	if node.left == nil {
		// 找到最小节点（最左节点）
		tmpNode := node.right
		node.right = nil
		t.size--
		return tmpNode
	}
	node.left = t.removeMin(node.left)
	return node
}

// min 找到以指定节点为根的子树中的最小节点
func (t *BinarySearchTree[K, V]) min(node *BSTNode[K, V]) *BSTNode[K, V] {
	if node == nil {
		return node
	}
	return t.min(node.left)
}

// getNode 根据键查找节点
func (t *BinarySearchTree[K, V]) getNode(node *BSTNode[K, V], key K) *BSTNode[K, V] {
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
func (t *BinarySearchTree[K, V]) less(a, b K) bool {
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
func (t *BinarySearchTree[K, V]) equal(a, b K) bool {
	return any(a) == any(b)
}
