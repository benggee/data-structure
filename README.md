# 数据结构与算法  

数据结构与算法是计算机科学基础中的基础，要掌握好它并不容易，其中的原因倒不是因为它有多难。结合过去的工作经历，我觉得大致有以下几个原因：
- 第一，在大学的课程里并没有讲明白一个问题，就是数据结构算法应该怎么样和实际场景结合。比如，老师讲到树形结构的时候会告诉你，树形效率很高，
它的插入和查找时间复杂度都是O(LogN)，是一种非常重要的数据结构，在计算机中应用非常广泛，讲到这里就讲完了,相信很多同学听到这里都是一头雾水。
- 第二，在实际工作中，大部分人都没有机会从零开始手写一个数据结构和算法，在主流的编程语言中都已经封装好了。
- 第三，数据结构算法是一门需要大量练习的技能，很多时候我们能说出队列是什么原理，栈是什么原理，树的旋转是怎么回事，
但如果让我们自己实现一个队列或者栈还是有点难度的， 尤其是在需要考虑一些性能问题的时候。

这个项目梳理了常用的数据结构和算法。其特点是比较完整，而且前后有一些关联性，可以帮助需要学习数据结构算法的人梳理出一个比较清晰的
学习路径。我为这个项目在公众号写了一个系列的文章，你可以搜索"程序员班吉"找到公众号，里面从数组到红黑树、图论、排序算法都有非常
详细的讲解，再配合项目里的代码，可以轻松入门甚至进阶数据结构与算法。

很多同学都是冲着面试去的，跟着文章学完这个系列可以应付面试吗？答案是：不能。这是因为面试中很少会直接让你实现一个链表，
而大概率是让你使用现有的数据结构与算法去解决一个个实际的问题，很多时候可能一个问题需要使用很多种数据结构来配合解决。
所以，对于需要面试的同学来说，还是需要去leetcode这类平台多刷一些题找找感觉，看看如何用掌握的知识去分析实际问题，
找到问题和知识之间的对应关系。那么，是不是说学习数据结构算法就没用了，直接去刷题就好了？答案也是否定的，因为你得先有储备，
遇到问题才能有相应的招式去解决。这就像是你和泰森打拳，可能你一顿王八拳确实也能击中他几拳，但是非常大的概率你是打不赢他的。

leetcode上面目前已经有近2000道题了，如果从头到尾刷一遍也不太现实，这里给了一份leetcode题目清单，
包含了从基础的数组到树到动态规划的题[点击查看](DOCS/leetcode.md)

下面是完整的目录，点击目录进去可以查看不同编程语言实现的版本


## 线性数组结构  

- [数组 Array](DOCS/line-array.md)
- [链表 Link List](DOCS/line-link-list.md)
- [队列 Queue](DOCS/line-queue.md)
- [堆栈（栈）Stack](DOCS/line-stack.md)
- [集合 Set](DOCS/line-set.md)
- [映射 Map](DOCS/line-map.md)
- [哈希表 Hash](DOCS/line-hash.md)



## 树形数组结构  

- [递归 Recursion](DOCS/tree-recursion.md)
- [二分搜索树 Binary Search Tree](DOCS/tree-bst.md)
- [堆 Heap](DOCS/tree-heap.md)
- [并查集 Union Find](DOCS/tree-union-find.md)
- [线段对 Segment Tree](DOCS/tree-segment-tree.md)
- [字典树 Trie](DOCS/tree-trie.md)
- [平衡二叉树 AVL Tree](DOCS/tree-avl-tree.md)
- [红黑树 Red Back Tree](DOCS/tree-red-back-tree.md)



### 图论与图论算法    

#### 无权图 

- [邻接矩阵的实现(矩阵) Matrix](DOCS/graph-matrix.md)
- [邻接表的实现(链表) LinkList](DOCS/graph-linklist.md)
- [邻接表的实现(哈希表) HashSet](DOCS/graph-hashset.md)
- [邻接表的实现(红黑树) TreeSet](DOCS/graph-treeset.md)  

#### 深度优先遍历   

- [图的深度优先遍历(前、后序) Graph DFS](DOCS/graph-dfs-order.md)
- [图的分量求解 Graph CC](DOCS/graph-cc.md)
- [图的单源路径 Single Source DFS](DOCS/singlesource-order.md)
- [是否是二分图 Bipartition Detetion DFS](DOCS/bipartition-detection.md)  

#### 广度优先遍历    

- [广度优先遍历基本实现 Graph BFS](DOCS/graph-bfs.md)    
- [检测环 Has Cycle BFS](DOCS/cycledetection.md)
- [单源路径求解 Single Path BFS](DOCS/single-source-path-bfs.md)
- [求路径长度 Source Length BFS](DOCS/ussspath-bfs.md)     

#### 无权图算法   

- [Floodfill算法基础实现 Floodfill](DOCS/floodfill.md)   
- [哈密尔顿回路 Hamilton Loop](DOCS/graph-hamiltonloop.md)
- [哈密尔顿回路(基于状态压缩的实现) Hamilton Loop Zip Status](DOCS/graph-hamiltonloop.md)
- [哈密尔顿路径 Hamilton Path](DOCS/graph-hamiltonpath.md)
- [寻找桥](DOCS/graph/find-bridge.md)
- [寻找割点](DOCS/graph/find-cut-points.md)
- [欧拉路径（Hierholzer算法）](DOCS/graph/euler-loop.md)

#### 有权图

- [建图(红黑树) TreeSet](DOCS/weight-graph-treeset.md)   

#### 有权图算法

- [最小生成树（Kruskal）](DOCS/weight-graph-kruskal.md)   
- [最小生成树（Prim）](DOCS/graph/prim.md)   
- [最短路径 (Dijkstra）](DOCS/graph/dijkstra.md)   
- [最短路径 (Bellman Ford)](DOCS/graph/bellman-ford.md)   
- [最短路径 (Floyed)](DOCS/graph/floyed.md)   

#### 有向图

- [建图](DOCS/graph/direction-graph.md)   

#### 有向图算法    

- [检测环](DOCS/graph/direction-cycle-detection.md)   
- [计算强联通分量(Kosaraju算法)](DOCS/graph/scc.md)   
- [拓扑排序](DOCS/graph/toposort.md)   
- [最大网络流（Edmonds Karp算法）](DOCS/graph/maxflow.md)   
- [最大匹配（Edmonds Karp算法）](DOCS/graph/dipartite-matching.md)   
- [最大匹配（匈牙利算法）](DOCS/graph/hungarian.md)   
- [欧拉路径（Hierholzer 算法）](DOCS/graph/direction-euler-loop.md)

## 排序算法

- [选择排序 Select Sort](DOCS/select-sort.md)
- [冒泡排序 Bubble Sort](DOCS/bubble-sort.md)
- [插入排序 Insertion Sort](DOCS/insertion-sort.md)
- [希尔排序 Shell Sort](DOCS/shell-sort.md)
- [归并排序 Merge Sort](DOCS/merge-sort.md)
- [快速排序 Quick Sort](DOCS/quick-sort.md)



数据结构系列文章可以关注我的公众号：“程序员班吉”

![wechat-qr](DOCS/wechat-qr.jpg)
