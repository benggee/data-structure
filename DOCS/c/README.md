# C 语言数据结构实现文档

## 目录
- [C语言实现特点](#c语言实现特点)
- [内存管理策略](#内存管理策略)
- [核心数据结构](#核心数据结构)
- [代码示例](#代码示例)
- [调试与最佳实践](#调试与最佳实践)

---

## C语言实现特点

### 1. 手动内存管理

C语言最显著的特点是需要**手动管理所有内存**：

```c
// 分配内存
ListElmt *new_element = (ListElmt *)malloc(sizeof(ListElmt));
if (new_element == NULL) {
    return -1;  // 内存分配失败
}

// 使用内存
new_element->data = (void *)data;

// 释放内存（必须手动调用）
free(new_element);
new_element = NULL;  // 避免悬空指针
```

**内存管理原则：**
1. 每个malloc必须有对应的free
2. 释放后立即置NULL
3. 检查malloc返回值是否为NULL
4. 避免内存泄漏和重复释放

### 2. 泛型编程（void指针）

C语言没有模板，使用`void*`实现泛型：

```c
typedef struct ListElmt_ {
    void *data;              // 可以指向任何类型
    struct ListElmt_ *next;
} ListElmt;

// 存储整数
int value = 42;
element->data = &value;

// 存储字符串
char *str = "hello";
element->data = str;

// 使用时需要转换
int *ptr = (int *)element->data;
printf("%d\n", *ptr);
```

**void指针的限制：**
- 不能直接解引用
- 不能进行指针运算
- 使用前必须转换为具体类型

### 3. 函数指针机制

C使用函数指针实现回调：

```c
typedef struct List_ {
    int size;
    int (*match)(const void *key1, const void *key2);    // 比较函数
    void (*destroy)(void *data);                        // 销毁函数
    ListElmt *head;
    ListElmt *tail;
} List;

// 初始化时传入函数指针
void list_init(List *list, void (*destroy)(void *data)) {
    list->destroy = destroy;
    // ...
}

// 使用回调函数
if (list->destroy != NULL) {
    list->destroy(data);
}
```

---

## 内存管理策略

### 1. 三步分配模式

```c
// 1. 分配内存
List *list = (List *)malloc(sizeof(List));
if (list == NULL) {
    return -1;
}

// 2. 初始化
list_init(list, free_function);

// 3. 使用
// ... 使用list

// 4. 清理
list_destroy(list);
free(list);
list = NULL;
```

### 2. 嵌套结构内存释放

```c
// 释放树结构（递归）
static void destroy_left(BisTree *tree, BiTreeNode *node) {
    BiTreeNode **position;

    if (bitree_size(tree) == 0)
        return;

    if (node == NULL)
        position = &tree->root;
    else
        position = &node->left;

    if (*position != NULL) {
        // 先递归释放子树
        destroy_left(tree, *position);
        destroy_right(tree, *position);

        // 释放数据
        if (tree->destroy != NULL) {
            tree->destroy(((AvlNode *)(*position)->data)->data);
        }

        // 释放节点
        free((*position)->data);
        free(*position);
        *position = NULL;

        tree->size--;
    }
}
```

**释放顺序很重要：**
1. 先释放子节点（后序遍历）
2. 再释放当前节点
3. 最后释放根节点

---

## 核心数据结构

### 1. 链表 (List)

**文件：** `c/list.c`, `c/list.h`

**数据结构：**
```c
// 节点
typedef struct ListElmt_ {
    void *data;
    struct ListElmt_ *next;
} ListElmt;

// 链表
typedef struct List_ {
    int size;
    int (*match)(const void *key1, const void *key2);
    void (*destroy)(void *data);
    ListElmt *head;
    ListElmt *tail;
} List;
```

**插入操作：**
```c
int list_ins_next(List *list, ListElmt *element, const void *data) {
    ListElmt *new_element;

    // 分配新节点
    if ((new_element = (ListElmt *)malloc(sizeof(ListElmt))) == NULL)
        return -1;

    new_element->data = (void *)data;

    // 在element后插入
    if (element == NULL) {
        // 在头部插入
        if (list_size(list) == 0)
            list->tail = new_element;

        new_element->next = list->head;
        list->head = new_element;
    } else {
        // 在中间插入
        if (element->next == NULL)
            list->tail = new_element;

        new_element->next = element->next;
        element->next = new_element;
    }

    list->size++;
    return 0;
}
```

### 2. 双向链表 (DList)

**文件：** `c/dlist.c`, `c/dlist.h`

**数据结构：**
```c
typedef struct DListElmt_ {
    void *data;
    struct DListElmt_ *prev;
    struct DListElmt_ *next;
} DListElmt;

typedef struct DList_ {
    int size;
    int (*match)(const void *key1, const void *key2);
    void (*destroy)(void *data);
    DListElmt *head;
    DListElmt *tail;
} DList;
```

**优势：**
- 可以双向遍历
- O(1)时间删除给定节点

### 3. 二叉搜索树/AVL树 (BisTree)

**文件：** `c/bistree.c`, `c/bistree.h`

**AVL节点定义：**
```c
typedef enum AvlNode_ {
    AvlNodeLeftHeavy,
    AvlNodeBalanced,
    AvlNodeRightHeavy
} AvlNodeFactor;

typedef struct AvlNode_ {
    void *data;
    AvlNodeFactor factor;  // 平衡因子
    char hidden;           // 标记删除
} AvlNode;
```

**旋转操作：**

```c
// 左旋转
static void rotate_left(BiTreeNode **node) {
    BiTreeNode *left, *grandchild;

    left = bitree_left(*node);
    if (((AvlNode *)bitree_data(left))->factor == AVL_LFT_HEAVY) {
        // LL情况：直接左旋
        bitree_left(*node) = bitree_right(left);
        bitree_right(left) = *node;
        ((AvlNode *)bitree_data(*node))->factor = AVL_BALANCED;
        ((AvlNode *)bitree_data(left))->factor = AVL_BALANCED;
        *node = left;
    } else {
        // LR情况：先右旋左子节点，再左旋
        grandchild = bitree_right(left);
        bitree_right(left) = bitree_left(grandchild);
        bitree_left(grandchild) = left;
        bitree_left(*node) = bitree_right(grandchild);
        bitree_right(grandchild) = *node;

        // 更新平衡因子
        switch (((AvlNode *)bitree_data(grandchild))->factor) {
            case AVL_LFT_HEAVY:
                ((AvlNode *)bitree_data(*node))->factor = AVL_RGT_HEAY;
                ((AvlNode *)bitree_data(left))->factor = AVL_BALANCED;
                break;
            case AVL_BALANCED:
                ((AvlNode *)bitree_data(*node))->factor = AVL_BALANCED;
                ((AvlNode *)bitree_data(left))->factor = AVL_BALANCED;
                break;
            case AVL_RGT_HEAY:
                ((AvlNode *)bitree_data(*node))->factor = AVL_BALANCED;
                ((AvlNode *)bitree_data(left))->factor = AVL_BALANCED;
                break;
        }

        ((AvlNode *)bitree_data(grandchild))->factor = AVL_BALANCED;
        *node = grandchild;
    }
}
```

### 4. 哈希表 (Chtbl)

**文件：** `c/chtbl.c`, `c/chtbl.h`

**链地址法实现：**
```c
typedef struct ChtblElmt_ {
    void *data;
    struct ChtblElmt_ *next;
} ChtblElmt;

typedef struct Chtbl_ {
    int buckets;
    int (*h)(const void *key);                     // 哈希函数
    int (*match)(const void *key1, const void *key2);
    void (*destroy)(void *data);
    int size;
    ListElmt **table;                               // 链表数组
} Chtbl;
```

**哈希表插入：**
```c
int chtbl_insert(Chtbl *htbl, const void *data) {
    int bucket, retval;

    // 找到bucket
    bucket = htbl->h(data) % htbl->buckets;

    // 检查是否已存在
    if (chtbl_lookup(htbl, data) != NULL)
        return 1;

    // 插入到链表头部
    if ((retval = list_ins_next(htbl->table[bucket], NULL, data)) == 0)
        htbl->size++;

    return retval;
}
```

### 5. 栈 (Stack) 和队列 (Queue)

**文件：** `c/stack.c`, `c/queue.c`

**栈（基于链表）：**
```c
typedef struct Stack_ {
    List list;
} Stack;

// 入栈
int stack_push(Stack *stack, const void *data) {
    return list_ins_next(&stack->list, NULL, data);
}

// 出栈
int stack_pop(Stack *stack, void **data) {
    return list_rem_next(&stack->list, NULL, data);
}
```

**队列（基于链表）：**
```c
typedef struct Queue_ {
    List list;
} Queue;

// 入队
int queue_enqueue(Queue *queue, const void *data) {
    return list_ins_next(&queue->list, list_tail(&queue->list), data);
}

// 出队
int queue_dequeue(Queue *queue, void **data) {
    return list_rem_next(&queue->list, NULL, data);
}
```

### 6. 堆和优先队列 (Heap, PQueue)

**文件：** `c/heap.c`, `c/pqueue.c`

**堆结构：**
```c
typedef struct Heap_ {
    int size;
    int (*compare)(const void *key1, const void *key2);
    void (*destroy)(void *data);
    void **tree;                                       // 动态数组
} Heap;
```

**堆插入（上浮）：**
```c
int heap_insert(Heap *heap, const void *data) {
    void *temp;
    int ipos, ppos;

    // 插入到最后
    if ((heap->tree[heap->size] = malloc(sizeof(void *))) == NULL)
        return -1;

    heap->tree[heap->size] = (void *)data;
    ipos = heap->size;
    ppos = (ipos - 1) / 2;

    // 上浮
    while (ipos > 0 && heap->compare(heap->tree[ppos], heap->tree[ipos]) < 0) {
        temp = heap->tree[ppos];
        heap->tree[ppos] = heap->tree[ipos];
        heap->tree[ipos] = temp;

        ipos = ppos;
        ppos = (ipos - 1) / 2;
    }

    heap->size++;
    return 0;
}
```

### 7. 集合 (Set)

**文件：** `c/set.c`, `c/set.h`

**基于链表实现：**
```c
typedef struct Set_ {
    int (*match)(const void *key1, const void *key2);
    void (*destroy)(void *data);
    List list;
} Set;

int set_insert(Set *set, const void *data) {
    // 检查是否已存在
    if (set_is_member(set, data))
        return 1;

    // 插入
    return list_ins_next(&set->list, NULL, data);
}
```

### 8. 图 (Graph)

**文件：** `c/graph.c`, `c/graph.h`

**邻接表实现：**
```c
typedef struct AdjList_ {
    List adjacent;     // 邻接顶点列表
} AdjList;

typedef struct Graph_ {
    int vcount;        // 顶点数
    int ecount;        // 边数
    int (*match)(const void *key1, const void *key2);
    void (*destroy)(void *data);
    List adjlists;     // 邻接表
} Graph;
```

---

## 代码示例

### 示例1：使用链表

```c
#include <stdio.h>
#include <stdlib.h>
#include "list.h"

// 销毁函数
void destroy(void *data) {
    free(data);
}

int main() {
    List list;
    int *data;

    // 初始化
    list_init(&list, destroy);

    // 插入元素
    for (int i = 0; i < 10; i++) {
        data = (int *)malloc(sizeof(int));
        *data = i;
        list_ins_next(&list, NULL, data);
    }

    // 遍历
    ListElmt *element = list_head(&list);
    while (element != NULL) {
        printf("%d ", *(int *)list_data(element));
        element = list_next(element);
    }

    // 销毁
    list_destroy(&list);

    return 0;
}
```

### 示例2：使用哈希表

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "chtbl.h"

// 哈希函数
int hash(const void *key) {
    const char *str = (const char *)key;
    int hash = 0;
    for (int i = 0; str[i] != '\0'; i++) {
        hash = (hash << 5) + str[i];
    }
    return hash;
}

// 匹配函数
int match(const void *key1, const void *key2) {
    return strcmp((const char *)key1, (const char *)key2) == 0;
}

int main() {
    Chtbl htbl;
    char *key, *value;

    // 初始化哈希表
    chtbl_init(&htbl, 16, hash, match, free);

    // 插入键值对
    chtbl_insert(&htbl, "one", "1");
    chtbl_insert(&htbl, "two", "2");
    chtbl_insert(&htbl, "three", "3");

    // 查找
    if ((value = chtbl_lookup(&htbl, "two")) != NULL) {
        printf("Found: %s\n", value);
    }

    // 销毁
    chtbl_destroy(&htbl);

    return 0;
}
```

### 示例3：使用AVL树

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "bistree.h"

// 比较函数
int compare(const void *key1, const void *key2) {
    return strcmp((const char *)key1, (const char *)key2);
}

int main() {
    BisTree tree;

    // 初始化
    bistree_init(&tree, compare, free);

    // 插入
    bistree_insert(&tree, "apple");
    bistree_insert(&tree, "banana");
    bistree_insert(&tree, "cherry");

    // 查找
    void *data;
    if (bistree_lookup(&tree, &data) == 0) {
        printf("Found: %s\n", (char *)data);
    }

    // 销毁
    bistree_destroy(&tree);

    return 0;
}
```

---

## 调试与最佳实践

### 1. 内存调试

**使用Valgrind检测内存泄漏：**

```bash
gcc -g program.c -o program
valgrind --leak-check=full ./program
```

**常见内存错误：**
```c
// 错误1：内存泄漏
void *ptr = malloc(100);
// 忘记free(ptr);

// 错误2：重复释放
void *ptr = malloc(100);
free(ptr);
free(ptr);  // 崩溃！

// 错误3：悬空指针
void *ptr = malloc(100);
free(ptr);
*ptr = 42;  // 未定义行为

// 正确做法
void *ptr = malloc(100);
if (ptr != NULL) {
    // 使用ptr
    free(ptr);
    ptr = NULL;
}
```

### 2. 防御性编程

```c
// 检查指针
if (list == NULL || list->head == NULL) {
    return -1;
}

// 检查返回值
if ((new_element = malloc(sizeof(ListElmt))) == NULL) {
    return -1;
}

// 使用const修饰不改变的数据
int list_size(const List *list) {
    return list->size;
}

// 使用宏定义避免魔法数字
#ifndef MAX_SIZE
#define MAX_SIZE 1000
#endif
```

### 3. 错误处理策略

```c
// 返回0表示成功，-1表示失败
int operation(...) {
    if (error_condition) {
        return -1;
    }
    return 0;
}

// 使用时检查
if (operation(...) != 0) {
    // 处理错误
}
```

### 4. 调试宏

```c
#ifdef DEBUG
#define DEBUG_PRINT(fmt, args...) \
    printf("[DEBUG] " fmt "\n", ##args)
#else
#define DEBUG_PRINT(fmt, args...)
#endif

// 使用
DEBUG_PRINT("Current size: %d", list->size);
```

---

## 总结

C语言数据结构实现的特点：

**优点：**
- ✅ 完全掌控内存和性能
- ✅ 理解底层实现原理
- ✅ 可移植性强
- ✅ 适合嵌入式和系统编程

**挑战：**
- ⚠️ 手动内存管理容易出错
- ⚠️ 没有泛型，使用void*
- ⚠️ 代码量大，易出错
- ⚠️ 调试困难

**适用场景：**
- 系统编程
- 嵌入式开发
- 性能关键应用
- 学习底层原理

**学习建议：**
1. 熟练掌握指针和内存管理
2. 使用工具检测内存错误
3. 理解计算机底层原理
4. 多练习，多调试

---

## 代码实现链接

- [C完整实现](../c/)
- [链表](../c/list.c)
- [双向链表](../c/dlist.c)
- [二叉搜索树](../c/bistree.c)
- [哈希表](../c/chtbl.c)
- [栈](../c/stack.c)
- [队列](../c/queue.c)
- [堆](../c/heap.c)
- [优先队列](../c/pqueue.c)
- [集合](../c/set.c)
- [图](../c/graph.c)
