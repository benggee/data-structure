# Go语言AVL树实现

这是一个用Go语言实现的AVL平衡二叉搜索树，包含详细的中文注释。

## 项目结构

```
AVLTree/
├── main.go                    # 主程序文件，包含所有实现
├── pride-and-prejudice.txt    # 测试用的文本文件
└── README.md                  # 项目说明文档
```

## 功能特性

### AVL树特性
- **自平衡**: 通过旋转操作自动保持树的平衡
- **平衡因子**: 定义为左子树高度 - 右子树高度
- **平衡条件**: 每个节点的平衡因子绝对值不超过1
- **时间复杂度**: 查找、插入、删除操作都是O(log n)

### 支持的操作
- `Set(key, value)`: 插入或更新键值对
- `Get(key)`: 获取指定键的值
- `Contains(key)`: 检查是否包含指定键
- `Remove(key)`: 删除指定键的节点
- `Size()`: 返回树中节点的数量
- `Empty()`: 检查树是否为空
- `IsBST()`: 验证是否是二叉搜索树
- `IsAVL()`: 验证是否是AVL树

### 旋转操作
- **右旋转 (LL情况)**: 修复左左不平衡
- **左旋转 (RR情况)**: 修复右右不平衡
- **左右旋转 (LR情况)**: 先左旋转再右旋转
- **右左旋转 (RL情况)**: 先右旋转再左旋转

## 编译和运行

### 编译
```bash
cd /Users/mrbinary/app/data-structure/golang/AVLTree
go build -o avltree main.go
```

### 运行
```bash
./avltree
```

或者直接运行：
```bash
go run main.go
```

## 测试功能

程序包含以下测试：

### 1. 基本操作测试
- 插入操作
- 查找操作
- 更新操作
- 删除操作

### 2. 边界情况测试
- 空树操作
- 单个节点操作
- 重复插入
- 删除不存在的节点

### 3. AVL树完整测试
- 读取《傲慢与偏见》文件
- 统计单词频率
- 验证AVL树属性
- 测试删除后的平衡性

### 4. 性能对比测试
- AVL树与二叉搜索树性能对比
- 使用排序数据测试退化情况

## 代码结构说明

### AVL树实现
```go
type AVLTree[K comparable, V any] struct {
    root *Node[K, V] // 根节点
    size int         // 树中节点的数量
}

type Node[K comparable, V any] struct {
    key    K           // 节点的键
    value  V           // 节点的值
    left   *Node[K, V] // 左子节点
    right  *Node[K, V] // 右子节点
    height int         // 节点的高度
}
```

### 二叉搜索树实现
```go
type BinarySearchTree[K comparable, V any] struct {
    root *BSTNode[K, V] // 根节点
    size int            // 树中节点的数量
}
```

### 文件操作
- `ReadFile(filename, words)`: 读取文件并进行分词
- `tokenizeLine(line, words)`: 对单行文本进行分词
- `firstCharacterIndex(s, start)`: 查找第一个字母字符位置

## 使用示例

```go
// 创建AVL树
avlTree := NewAVLTree[string, int]()

// 插入数据
avlTree.Set("apple", 1)
avlTree.Set("banana", 2)
avlTree.Set("cherry", 3)

// 查找数据
if count, exists := avlTree.Get("apple"); exists {
    fmt.Printf("apple: %d\n", count)
}

// 检查是否包含
if avlTree.Contains("banana") {
    fmt.Println("包含banana")
}

// 删除数据
if value, removed := avlTree.Remove("cherry"); removed {
    fmt.Printf("删除cherry: %d\n", value)
}

// 验证树的性质
fmt.Printf("是否是AVL树: %t\n", avlTree.IsAVL())
fmt.Printf("是否是BST: %t\n", avlTree.IsBST())
```

## 算法复杂度

| 操作 | 时间复杂度 | 空间复杂度 |
|------|------------|------------|
| 查找 | O(log n) | O(1) |
| 插入 | O(log n) | O(log n) |
| 删除 | O(log n) | O(log n) |
| 遍历 | O(n) | O(n) |

## 与Java版本的对比

### 主要差异
1. **泛型实现**: Go使用泛型语法 `[K comparable, V any]`
2. **错误处理**: Go使用多返回值处理错误
3. **内存管理**: Go自动垃圾回收，无需手动管理
4. **接口设计**: Go更简洁的接口设计

### 相同点
1. **算法逻辑**: 核心AVL树算法完全相同
2. **平衡策略**: 使用相同的旋转操作
3. **测试用例**: 使用相同的测试数据和方法

## 注意事项

1. **键类型限制**: 目前支持 `string`, `int`, `int64`, `float64` 类型
2. **并发安全**: 当前实现不是并发安全的
3. **内存使用**: 对于大数据集，考虑内存使用情况

## 扩展建议

1. **支持更多类型**: 实现通用的比较接口
2. **并发安全**: 添加读写锁支持并发访问
3. **序列化**: 支持树的序列化和反序列化
4. **可视化**: 添加树的可视化功能
5. **性能优化**: 进一步优化内存使用和性能

## 许可证

本项目仅供学习和研究使用。 