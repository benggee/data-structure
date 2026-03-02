package RBTree

/*
  红黑树实现 - Golang版本
  基于Java RBTree实现

  红黑树的5条性质：
  1. 每个节点要么是红色，要么是黑色
  2. 根节点是黑色
  3. 所有叶子节点（NIL）是黑色
  4. 如果一个节点是红色，那么它的两个子节点都是黑色
  5. 对每个节点，从该节点到其所有后代叶子节点的简单路径上，
     均包含相同数量的黑色节点
*/

const (
	RED   = true
	BLACK = false
)

// RBNode 红黑树节点
type RBNode[K comparable, V any] struct {
	key    K
	value  V
	left   *RBNode[K, V]
	right  *RBNode[K, V]
	color  bool
}

// NewRBNode 创建新节点（默认红色）
func NewRBNode[K comparable, V any](key K, value V) *RBNode[K, V] {
	return &RBNode[K, V]{
		key:   key,
		value: value,
		left:  nil,
		right: nil,
		color: RED, // 新节点默认为红色
	}
}

// RBTree 红黑树
type RBTree[K comparable, V any] struct {
	root *RBNode[K, V]
	size int
}

// NewRBTree 创建新的红黑树
func NewRBTree[K comparable, V any]() *RBTree[K, V] {
	return &RBTree[K, V]{
		root: nil,
		size:  0,
	}
}

func (t *RBTree[K, V]) Size() int {
	return t.size
}

func (t *RBTree[K, V]) IsEmpty() bool {
	return t.size == 0
}

func (t *RBTree[K, V]) Contains(key K) bool {
	return t.getNode(t.root, key) != nil
}

func (t *RBTree[K, V]) Get(key K) (V, bool) {
	node := t.getNode(t.root, key)
	if node == nil {
		var zero V
		return zero, false
	}
	return node.value, true
}

func (t *RBTree[K, V]) Set(key K, value V) {
	t.root = t.add(t.root, key, value)
}

// ===== 旋转操作 =====

// leftRotate 左旋转
/*
     node                    x
    /   \     左旋转       /   \
   T1   x    --------->  node  T3
      / \                /   \
     T2  T3             T1   T2
*/
func (t *RBTree[K, V]) leftRotate(node *RBNode[K, V]) *RBNode[K, V] {
	x := node.right

	node.right = x.left
	x.left = node

	x.color = node.color
	node.color = RED

	return x
}

// rightRotate 右旋转
/*
     node                     x
    /   \     右旋转        /   \
   x    T2   ------->   y   node
  / \                       /   \
 y  T1                     T1  T2
*/
func (t *RBTree[K, V]) rightRotate(node *RBNode[K, V]) *RBNode[K, V] {
	x := node.left

	node.left = x.right
	x.right = node

	x.color = node.color
	node.color = RED

	return x
}

// flipColor 颜色翻转
func (t *RBTree[K, V]) flipColor(node *RBNode[K, V]) {
	node.color = RED
	node.left.color = BLACK
	node.right.color = BLACK
}

// isRed 判断节点是否为红色（空节点视为黑色）
func (t *RBTree[K, V]) isRed(node *RBNode[K, V]) bool {
	if node == nil {
		return BLACK
	}
	return node.color
}

// ===== 核心操作 =====

// add 添加节点（递归）
func (t *RBTree[K, V]) add(node *RBNode[K, V], key K, value V) *RBNode[K, V] {
	if node == nil {
		t.size++
		return NewRBNode(key, value)
	}

	if key < node.key {
		node.left = t.add(node.left, key, value)
	} else if key > node.key {
		node.right = t.add(node.right, key, value)
	} else {
		// 键已存在，更新值
		node.value = value
		return node
	}

	// ===== 维护红黑树性质 =====

	// 情况1：右孩子是红色，左孩子是黑色 -> 左旋
	if t.isRed(node.right) && !t.isRed(node.left) {
		node = t.leftRotate(node)
	}

	// 情况2：左孩子是红色，左孩子的左孩子是红色 -> 右旋
	if t.isRed(node.left) && t.isRed(node.left.left) {
		node = t.rightRotate(node)
	}

	// 情况3：左右孩子都是红色 -> 颜色翻转
	if t.isRed(node.left) && t.isRed(node.right) {
		t.flipColor(node)
	}

	return node
}

// getNode 获取节点
func (t *RBTree[K, V]) getNode(node *RBNode[K, V], key K) *RBNode[K, V] {
	if node == nil {
		return nil
	}

	if key == node.key {
		return node
	} else if key < node.key {
		return t.getNode(node.left, key)
	} else {
		return t.getNode(node.right, key)
	}
}

// Keys 获取所有键（中序遍历）
func (t *RBTree[K, V]) Keys() []K {
	var result []K
	t.inOrder(t.root, &result)
	return result
}

func (t *RBTree[K, V]) inOrder(node *RBNode[K, V], result *[]K) {
	if node == nil {
		return
	}
	t.inOrder(node.left, result)
	*result = append(*result, node.key)
	t.inOrder(node.right, result)
}

// IsValidRBTree 验证是否是有效的红黑树
func (t *RBTree[K, V]) IsValidRBTree() bool {
	return t.isValidRBTreeRecursive(t.root)
}

func (t *RBTree[K, V]) isValidRBTreeRecursive(node *RBNode[K, V]) bool {
	if node == nil {
		return true
	}

	// 性质4：红色节点的子节点必须是黑色
	if node.color == RED {
		if (node.left != nil && node.left.color == RED) ||
			(node.right != nil && node.right.color == RED) {
			return false
		}
	}

	// 递归检查左右子树
	return t.isValidRBTreeRecursive(node.left) && t.isValidRBTreeRecursive(node.right)
}
