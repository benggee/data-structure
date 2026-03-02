# C++ 数据结构实现文档

## 目录
- [C++实现特点](#c实现特点)
- [模板与泛型](#模板与泛型)
- [STL兼容性](#stl兼容性)
- [核心数据结构](#核心数据结构)
- [代码示例](#代码示例)
- [现代C++特性](#现代c特性)

---

## C++实现特点

### 1. 模板泛型编程

C++使用模板实现类型安全的泛型：

```cpp
template<typename T>
class Array {
private:
    T* data;
    int size;
    int capacity;

public:
    Array(int capacity = 10) {
        this->data = new T[capacity];
        this->size = 0;
        this->capacity = capacity;
    }

    void add(int index, const T& e) {
        if (index < 0 || index > size) {
            throw std::out_of_range("Index out of range");
        }
        // ...
    }
};
```

**模板特化：**
```cpp
// 通用版本
template<typename T>
class Comparator {
public:
    int compare(const T& a, const T& b) {
        if (a < b) return -1;
        if (a > b) return 1;
        return 0;
    }
};

// 特化版本（针对指针）
template<typename T>
class Comparator<T*> {
public:
    int compare(T* a, T* b) {
        return Comparator<T>().compare(*a, *b);
    }
};
```

### 2. RAII与智能指针

C++使用RAII（Resource Acquisition Is Initialization）自动管理资源：

```cpp
template<typename T>
class ArrayList {
private:
    std::unique_ptr<T[]> data;  // 智能指针自动管理内存
    int size;
    int capacity;

public:
    ArrayList(int capacity = 10)
        : data(new T[capacity]), size(0), capacity(capacity) {
    }

    // 析构函数自动释放内存
    ~ArrayList() = default;

    // 移动构造（高效）
    ArrayList(ArrayList&& other) noexcept
        : data(std::move(other.data)),
          size(other.size),
          capacity(other.capacity) {
    }
};
```

### 3. 运算符重载

```cpp
template<typename T>
class Array {
public:
    // 下标运算符
    T& operator[](int index) {
        if (index < 0 || index >= size) {
            throw std::out_of_range("Index out of range");
        }
        return data[index];
    }

    const T& operator[](int index) const {
        if (index < 0 || index >= size) {
            throw std::out_of_range("Index out of range");
        }
        return data[index];
    }

    // 比较运算符
    bool operator==(const Array& other) const {
        if (size != other.size) return false;
        for (int i = 0; i < size; i++) {
            if (data[i] != other.data[i]) return false;
        }
        return true;
    }
};
```

---

## 模板与泛型

### 1. 类型约束（C++20 Concepts）

```cpp
// C++20 之前：使用static_assert
template<typename T>
class SortedContainer {
    static_assert(std::is_arithmetic<T>::value,
                  "T must be arithmetic type");
};

// C++20：使用Concepts
template<typename T>
concept Comparable = requires(T a, T b) {
    { a < b } -> std::convertible_to<bool>;
    { a == b } -> std::convertible_to<bool>;
};

template<Comparable T>
class SortedContainer {
    // T保证是可比较的
};
```

### 2. 完美转发

```cpp
template<typename T>
void wrapper(T&& arg) {
    // 完美转发保持值类别
    process(std::forward<T>(arg));
}

// 使用场景
template<typename T>
void ArrayList<T>::add(T&& e) {
    // 如果是右值，移动；如果是左值，复制
    data[size++] = std::forward<T>(e);
}
```

---

## STL兼容性

### 1. 迭代器支持

```cpp
template<typename T>
class ArrayList {
public:
    // 迭代器类型
    class Iterator {
    private:
        T* ptr;
    public:
        Iterator(T* p) : ptr(p) {}

        T& operator*() { return *ptr; }
        T* operator->() { return ptr; }

        Iterator& operator++() {
            ++ptr;
            return *this;
        }

        bool operator!=(const Iterator& other) const {
            return ptr != other.ptr;
        }
    };

    Iterator begin() { return Iterator(data); }
    Iterator end() { return Iterator(data + size); }
};

// 支持范围for循环
ArrayList<int> arr;
for (int& val : arr) {
    val *= 2;
}
```

### 2. 算法兼容

```cpp
#include <algorithm>
#include <vector>

// 可以配合STL算法使用
ArrayList<int> arr;
std::sort(arr.begin(), arr.end());

int sum = std::accumulate(arr.begin(), arr.end(), 0);

// 查找
auto it = std::find(arr.begin(), arr.end(), 42);
if (it != arr.end()) {
    // 找到了
}
```

---

## 核心数据结构

### 1. 动态数组 (ArrayList)

**文件：** `c++/array/ArrayList.hpp`

```cpp
template<typename T>
class ArrayList {
private:
    T* data;
    int size;
    int capacity;

    void resize(int newCapacity) {
        T* newData = new T[newCapacity];
        for (int i = 0; i < size; i++) {
            newData[i] = std::move(data[i]);  // 移动语义
        }
        delete[] data;
        data = newData;
        capacity = newCapacity;
    }

public:
    ArrayList(int capacity = 10)
        : data(new T[capacity]), size(0), capacity(capacity) {}

    ~ArrayList() {
        delete[] data;
    }

    void add(const T& e) {
        if (size == capacity) {
            resize(2 * capacity);
        }
        data[size++] = e;
    }

    void add(T&& e) {
        if (size == capacity) {
            resize(2 * capacity);
        }
        data[size++] = std::move(e);  // 移动语义
    }

    T& operator[](int index) { return data[index]; }
    int getSize() const { return size; }
};
```

### 2. AVL树

**文件：** `c++/tree/avl_tree.h`

```cpp
template<typename T>
class AVLTree {
private:
    struct Node {
        T val;
        Node* left;
        Node* right;
        int height;

        Node(T val) : val(val), left(nullptr), right(nullptr), height(1) {}
    };

    Node* root;
    int size;

    // 右旋转
    Node* rightRotate(Node* y) {
        Node* x = y->left;
        Node* T3 = x->right;

        // 执行旋转
        x->right = y;
        y->left = T3;

        // 更新高度
        y->height = std::max(getHeight(y->left), getHeight(y->right)) + 1;
        x->height = std::max(getHeight(x->left), getHeight(x->right)) + 1;

        return x;
    }

    // 左旋转
    Node* leftRotate(Node* y) {
        Node* x = y->right;
        Node* T2 = x->left;

        // 执行旋转
        x->left = y;
        y->right = T2;

        // 更新高度
        y->height = std::max(getHeight(y->left), getHeight(y->right)) + 1;
        x->height = std::max(getHeight(x->left), getHeight(x->right)) + 1;

        return x;
    }

    Node* add(Node* node, T val) {
        if (node == nullptr) {
            size++;
            return new Node(val);
        }

        if (val < node->val) {
            node->left = add(node->left, val);
        } else if (val > node->val) {
            node->right = add(node->right, val);
        } else {
            return node;  // 不重复添加
        }

        // 更新高度
        node->height = 1 + std::max(getHeight(node->left),
                                   getHeight(node->right));

        // 平衡因子
        int balanceFactor = getBalanceFactor(node);

        // LL情况
        if (balanceFactor > 1 && val < node->left->val) {
            return rightRotate(node);
        }

        // RR情况
        if (balanceFactor < -1 && val > node->right->val) {
            return leftRotate(node);
        }

        // LR情况
        if (balanceFactor > 1 && val > node->left->val) {
            node->left = leftRotate(node->left);
            return rightRotate(node);
        }

        // RL情况
        if (balanceFactor < -1 && val < node->right->val) {
            node->right = rightRotate(node->right);
            return leftRotate(node);
        }

        return node;
    }

public:
    AVLTree() : root(nullptr), size(0) {}

    void add(T val) {
        root = add(root, val);
    }

    bool isBST() {
        std::vector<T> vals;
        inOrder(root, vals);
        for (int i = 1; i < vals.size(); i++) {
            if (vals[i-1] > vals[i]) {
                return false;
            }
        }
        return true;
    }
};
```

### 3. 哈希表

**文件：** `c++/hash/hashtable.hpp`

```cpp
template<typename K, typename V>
class HashTable {
private:
    struct Node {
        K key;
        V value;
        Node* next;
        Node(const K& k, const V& v) : key(k), value(v), next(nullptr) {}
    };

    std::vector<std::unique_ptr<Node>> buckets;
    int size;
    int capacity;
    double loadFactor;

    int hash(const K& key) const {
        return std::hash<K>{}(key) % capacity;
    }

    void resize(int newCapacity) {
        std::vector<std::unique_ptr<Node>> newBuckets(newCapacity);

        for (int i = 0; i < capacity; i++) {
            Node* node = buckets[i].get();
            while (node != nullptr) {
                int newIndex = std::hash<K>{}(node->key) % newCapacity;
                auto newNode = std::make_unique<Node>(node->key, node->value);
                newNode->next = newBuckets[newIndex].release();
                newBuckets[newIndex].reset(newNode.release());
                node = node->next;
            }
        }

        buckets = std::move(newBuckets);
        capacity = newCapacity;
    }

public:
    HashTable(int initialCapacity = 16, double lf = 0.75)
        : size(0), capacity(initialCapacity), loadFactor(lf) {
        buckets.resize(capacity);
    }

    void put(const K& key, const V& value) {
        if ((double)size / capacity >= loadFactor) {
            resize(capacity * 2);
        }

        int index = hash(key);
        Node* node = buckets[index].get();

        while (node != nullptr) {
            if (node->key == key) {
                node->value = value;
                return;
            }
            node = node->next;
        }

        auto newNode = std::make_unique<Node>(key, value);
        newNode->next = buckets[index].release();
        buckets[index].reset(newNode.release());
        size++;
    }

    V* get(const K& key) {
        int index = hash(key);
        Node* node = buckets[index].get();

        while (node != nullptr) {
            if (node->key == key) {
                return &(node->value);
            }
            node = node->next;
        }
        return nullptr;
    }
};
```

### 4. 并查集 (UnionFind)

**文件：** `c++/unionFind/unionfind.hpp`

```cpp
class UnionFind {
private:
    std::vector<int> parent;
    std::vector<int> rank;

public:
    UnionFind(int size) {
        parent.resize(size);
        rank.resize(size, 1);
        for (int i = 0; i < size; i++) {
            parent[i] = i;
        }
    }

    int find(int p) {
        while (p != parent[p]) {
            // 路径压缩
            parent[p] = parent[parent[p]];
            p = parent[p];
        }
        return p;
    }

    bool isConnected(int p, int q) {
        return find(p) == find(q);
    }

    void unionElements(int p, int q) {
        int pRoot = find(p);
        int qRoot = find(q);

        if (pRoot == qRoot) {
            return;
        }

        // 按秩合并
        if (rank[pRoot] < rank[qRoot]) {
            parent[pRoot] = qRoot;
        } else if (rank[pRoot] > rank[qRoot]) {
            parent[qRoot] = pRoot;
        } else {
            parent[qRoot] = pRoot;
            rank[pRoot]++;
        }
    }
};
```

### 5. 排序算法

**文件：** `c++/sort/quicksort.h`

```cpp
template<typename T>
class QuickSort {
public:
    static void sort(std::vector<T>& arr) {
        quickSort(arr, 0, arr.size() - 1);
    }

private:
    static void quickSort(std::vector<T>& arr, int l, int r) {
        if (l >= r) return;

        int p = partition(arr, l, r);
        quickSort(arr, l, p - 1);
        quickSort(arr, p + 1, r);
    }

    static int partition(std::vector<T>& arr, int l, int r) {
        // 随机选择pivot
        int randomIndex = l + rand() % (r - l + 1);
        std::swap(arr[l], arr[randomIndex]);

        T pivot = arr[l];
        int i = l, j = r;

        while (i < j) {
            while (i < j && arr[j] >= pivot) j--;
            if (i < j) arr[i] = arr[j];
            while (i < j && arr[i] <= pivot) i++;
            if (i < j) arr[j] = arr[i];
        }
        arr[i] = pivot;
        return i;
    }
};
```

---

## 代码示例

### 示例1：使用动态数组

```cpp
#include <iostream>
#include "ArrayList.hpp"

int main() {
    ArrayList<int> arr;

    // 添加元素
    for (int i = 0; i < 10; i++) {
        arr.add(i);
    }

    // 访问元素
    std::cout << "Element at 5: " << arr[5] << std::endl;

    // 范围for循环（如果实现了迭代器）
    for (int i = 0; i < arr.getSize(); i++) {
        std::cout << arr[i] << " ";
    }

    return 0;
}
```

### 示例2：使用AVL树

```cpp
#include <iostream>
#include "avl_tree.h"

int main() {
    AVLTree<int> tree;

    // 插入数据
    tree.add(5);
    tree.add(3);
    tree.add(7);
    tree.add(2);
    tree.add(4);
    tree.add(6);
    tree.add(8);

    // 验证
    std::cout << "Is BST: " << (tree.isBST() ? "Yes" : "No") << std::endl;
    std::cout << "Is Balanced: " << (tree.isBalanced() ? "Yes" : "No") << std::endl;

    return 0;
}
```

### 示例3：使用哈希表

```cpp
#include <iostream>
#include <string>
#include "hashtable.hpp"

int main() {
    HashTable<std::string, int> ht;

    // 插入键值对
    ht.put("apple", 5);
    ht.put("banana", 3);
    ht.put("cherry", 8);

    // 查找
    if (int* value = ht.get("banana")) {
        std::cout << "banana = " << *value << std::endl;
    }

    return 0;
}
```

---

## 现代C++特性

### 1. 移动语义 (C++11)

```cpp
template<typename T>
class ArrayList {
public:
    // 移动构造函数
    ArrayList(ArrayList&& other) noexcept
        : data(other.data),
          size(other.size),
          capacity(other.capacity) {
        other.data = nullptr;
        other.size = 0;
        other.capacity = 0;
    }

    // 移动赋值运算符
    ArrayList& operator=(ArrayList&& other) noexcept {
        if (this != &other) {
            delete[] data;
            data = other.data;
            size = other.size;
            capacity = other.capacity;
            other.data = nullptr;
        }
        return *this;
    }
};
```

### 2. constexpr和编译期计算

```cpp
template<typename T>
class Array {
public:
    constexpr int size() const { return _size; }
    constexpr bool isEmpty() const { return _size == 0; }
};

// 编译期即可计算
constexpr int arrSize = arr.size();
```

### 3. 范围for循环

```cpp
ArrayList<int> arr;
for (const auto& elem : arr) {
    std::cout << elem << std::endl;
}

// 或提供begin()/end()
for (auto it = arr.begin(); it != arr.end(); ++it) {
    std::cout << *it << std::endl;
}
```

### 4. Lambda表达式

```cpp
// 在排序中使用自定义比较
std::sort(arr.begin(), arr.end(),
    [](const auto& a, const auto& b) {
        return a > b;  // 降序排序
    });

// 在哈希表中使用lambda
auto hash = [](const std::string& key) {
    return std::hash<std::string>{}(key) % capacity;
};
```

### 5. 结构化绑定 (C++17)

```cpp
std::pair<bool, int> result = map.insert(42);

if (auto [success, value] = map.insert(42); success) {
    std::cout << "Inserted: " << value << std::endl;
}
```

---

## 调试与最佳实践

### 1. 使用断言

```cpp
#include <cassert>

template<typename T>
class Array {
public:
    T& get(int index) {
        assert(index >= 0 && index < size && "Index out of range");
        return data[index];
    }
};
```

### 2. 异常处理

```cpp
template<typename T>
class Array {
public:
    T& get(int index) {
        if (index < 0 || index >= size) {
            throw std::out_of_range("Index out of range");
        }
        return data[index];
    }
};

// 使用
try {
    T value = arr.get(100);
} catch (const std::out_of_range& e) {
    std::cerr << "Error: " << e.what() << std::endl;
}
```

### 3. 智能指针使用

```cpp
// 独占所有权
std::unique_ptr<ArrayList<int>> arr =
    std::make_unique<ArrayList<int>>();

// 共享所有权
std::shared_ptr<ArrayList<int>> arr2 = arr;  // 拷贝但共享底层对象

// 观察者（不拥有所有权）
std::weak_ptr<ArrayList<int>> arr3 = arr2;
```

---

## 总结

C++数据结构实现的特点：

**优点：**
- ✅ 模板提供类型安全的泛型
- � STL提供丰富的基础设施
- � RAII自动管理资源
- ✅ 零开销抽象
- ✅ 完美的性能控制

**挑战：**
- ⚠️ 模板语法复杂
- ⚠️ 编译时间长
- ⚠️ 错误信息难懂
- ⚠️ 需要管理很多细节

**适用场景：**
- 高性能应用
- 游戏开发
- 系统编程
- 金融应用
- 竞赛编程

**学习建议：**
1. 先掌握STL的使用
2. 理解模板基础
3. 学习RAII和智能指针
4. 熟悉现代C++特性
5. 多练习，阅读优秀代码

---

## 代码实现链接

- [C++完整实现](../c++/)
- [动态数组](../c++/array/ArrayList.hpp)
- [AVL树](../c++/tree/avl_tree.h)
- [哈希表](../c++/hash/)
- [并查集](../c++/unionFind/)
- [排序算法](../c++/sort/)
- [链表](../c++/linkList/)
- [栈](../c++/stack/)
- [队列](../c++/queue/)
- [图论](../c++/graph/)
- [字典树](../c++/trie/)
- [线段树](../c++/segmentTree/)
