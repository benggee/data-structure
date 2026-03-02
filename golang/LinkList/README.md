# Go语言链表实现

这是一个用Go语言实现的链表数据结构，包含链表、栈、队列的完整实现，包含详细的中文注释。

## 项目结构

```
LinkList/
├── linklist.go           # 链表核心实现
├── linklist_stack.go     # 基于链表的栈实现
├── linklist_queue.go     # 基于链表的队列实现
├── main.go               # 主程序文件，包含测试代码
└── README.md             # 项目说明文档
```

## 功能特性

### 链表特性
- **虚拟头节点**: 使用虚拟头节点简化操作
- **泛型支持**: 支持任意类型的数据
- **动态大小**: 自动管理内存，无需手动扩容
- **完整操作**: 支持增删改查等所有基本操作

### 支持的操作
- `Add(index, e)`: 在指定位置插入元素
- `AddFirst(e)`: 在头部插入元素
- `AddLast(e)`: 在尾部插入元素
- `Get(index)`: 获取指定位置的元素
- `GetFirst()`: 获取第一个元素
- `GetLast()`: 获取最后一个元素
- `Set(index, e)`: 设置指定位置的元素
- `Find(e)`: 查找元素是否存在
- `Del(index)`: 删除指定位置的元素
- `DelFirst()`: 删除第一个元素
- `DelLast()`: 删除最后一个元素
- `Remove(e)`: 删除指定元素
- `GetSize()`: 返回链表大小
- `IsEmpty()`: 判断是否为空

### 高级功能
- `Reverse()`: 反转链表
- `ToSlice()`: 转换为切片
- `Clear()`: 清空链表
- `HasCycle()`: 检测是否有环

## 栈实现

基于链表实现的栈，支持以下操作：
- `Push(e)`: 压入元素
- `Pop()`: 弹出元素
- `Peek()`: 查看栈顶元素
- `GetSize()`: 返回栈大小
- `IsEmpty()`: 判断是否为空

## 队列实现

基于链表实现的队列，支持以下操作：
- `Enqueue(e)`: 入队
- `Dequeue()`: 出队
- `GetFront()`: 查看队首元素
- `GetSize()`: 返回队列大小
- `IsEmpty()`: 判断是否为空

## 编译和运行

### 编译
```bash
cd /Users/mrbinary/app/data-structure/golang/LinkList
go build -o linklist main.go
```

### 运行
```bash
./linklist
```

或者直接运行：
```bash
go run main.go
```

## 测试功能

程序包含以下测试：

### 1. 链表基本功能测试
- 添加元素（头部、尾部、指定位置）
- 查找元素
- 获取元素
- 删除元素
- 设置元素

### 2. 链表高级功能测试
- 反转链表
- 转换为切片
- 清空链表
- 环检测

### 3. 栈功能测试
- 压栈操作
- 出栈操作
- 查看栈顶元素
- 性能测试

### 4. 队列功能测试
- 入队操作
- 出队操作
- 查看队首元素
- 性能测试

### 5. 性能对比测试
- 链表栈性能
- 链表队列性能

## 代码结构说明

### 链表实现
```go
type LinkList[E any] struct {
    dummyHead *Node[E] // 虚拟头节点
    size      int      // 链表大小
}

type Node[E any] struct {
    e    E        // 节点数据
    next *Node[E] // 下一个节点指针
}
```

### 栈实现
```go
type LinkListStack[E any] struct {
    list *LinkList[E] // 使用链表作为底层存储
}
```

### 队列实现
```go
type LinkListQueue[E any] struct {
    head *Node[E] // 队列头部
    tail *Node[E] // 队列尾部
    size int      // 队列大小
}
```

## 使用示例

### 链表使用
```go
// 创建链表
list := NewLinkList[int]()

// 添加元素
list.AddFirst(1)
list.AddLast(2)
list.Add(1, 3)

// 查找元素
if list.Find(2) {
    fmt.Println("找到元素2")
}

// 删除元素
deleted := list.Del(1)
fmt.Printf("删除的元素: %d\n", deleted)

// 获取元素
first := list.GetFirst()
last := list.GetLast()
```

### 栈使用
```go
// 创建栈
stack := NewLinkListStack[int]()

// 压栈
stack.Push(1)
stack.Push(2)
stack.Push(3)

// 查看栈顶
top := stack.Peek()
fmt.Printf("栈顶元素: %d\n", top)

// 出栈
popped := stack.Pop()
fmt.Printf("弹出元素: %d\n", popped)
```

### 队列使用
```go
// 创建队列
queue := NewLinkListQueue[int]()

// 入队
queue.Enqueue(1)
queue.Enqueue(2)
queue.Enqueue(3)

// 查看队首
front := queue.GetFront()
fmt.Printf("队首元素: %d\n", front)

// 出队
dequeued := queue.Dequeue()
fmt.Printf("出队元素: %d\n", dequeued)
```

## 算法复杂度

| 操作 | 时间复杂度 | 空间复杂度 |
|------|------------|------------|
| 头部插入 | O(1) | O(1) |
| 尾部插入 | O(1) | O(1) |
| 指定位置插入 | O(n) | O(1) |
| 头部删除 | O(1) | O(1) |
| 尾部删除 | O(n) | O(1) |
| 指定位置删除 | O(n) | O(1) |
| 查找 | O(n) | O(1) |
| 访问 | O(n) | O(1) |
| 反转 | O(n) | O(1) |

## 与Java版本的对比

### 主要差异
1. **泛型实现**: Go使用泛型语法 `[E any]`
2. **错误处理**: Go使用panic处理异常
3. **内存管理**: Go自动垃圾回收
4. **接口设计**: Go更简洁的接口设计

### 相同点
1. **算法逻辑**: 核心链表算法完全相同
2. **数据结构**: 使用相同的节点结构
3. **操作接口**: 提供相同的操作方法

## 注意事项

1. **类型限制**: 支持任意类型，但复杂类型需要实现自定义比较方法
2. **并发安全**: 当前实现不是并发安全的
3. **内存使用**: 每个节点都有额外的指针开销

## 扩展建议

1. **双向链表**: 实现双向链表支持反向遍历
2. **循环链表**: 实现循环链表
3. **并发安全**: 添加读写锁支持并发访问
4. **序列化**: 支持链表的序列化和反序列化
5. **迭代器**: 实现迭代器模式
6. **性能优化**: 使用对象池减少内存分配

## 许可证

本项目仅供学习和研究使用。 