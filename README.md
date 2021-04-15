# 数据结构与算法  

数据结构与算法是基础中的基础，但要掌握好它确实不容易，这个项目梳理了常用的数据结构和部分图论算法和排序算法。其特点是比较完整，而且前后有一些关联性，可以帮助需要学习数据结构算法的人梳理出一个比较清晰的学习路径。这个项目配套的系列文章在我的公众号里面，你可以搜索"Mr Binary"找到订阅号，里面从数组到红黑树、图论、排序算法都有非常详细的讲解，再配全项目里的代码，可以轻松入门甚至进阶数据结构与算法。

很多同学都是冲着面试去的，跟着文章学完这个系列可以应付面试吗？答案是：不能。这是因为面试中都是从实际问题出发，很少会直接让你实现一个链表，而是通过一个个实际问题让你使用现有的数据结构去有解决，很多时候可能一个问题需要使用很多种数据结构来配合着解决。所以，对于需要面试的同学来说，还是需要去leetcode这类平台多刷一些题找找感觉，看看如何用掌握的知识去分析实际问题，找到问题和已经掌握的知识之间的映射关系。那么，是不是说学习数据结构算法就没用了，我直接去刷题就好了。答案也是否定的，因为你得先有储备，遇到问题才能有相应的招式去解决。这就像是你和泰森打拳，可能你一顿王八拳确实也能击中他几拳，但你非常大的概率是大不赢他的。

- ## 线性数组结构  
  
  - [数组 Array](DOCS/line-array.md)
  - [链表 Link List](DOCS/line-link-list.md)
  - [队列 Queue](DOCS/line-queue.md)
  - [堆栈（栈）Stack](DOCS/line-stack.md)
  - [集合 Set](DOCS/line-set.md)
  - [映射 Map](DOCS/line-map.md)
  - [哈希表 Hash](DOCS/line-hash.md)
  
  
- ## 树形数组结构  
  
  - [递归 Recursion](DOCS/tree-recursion.md)
  - [二分搜索树 Binary Search Tree](DOCS/tree-bst.md)
  - [堆 Heap](DOCS/tree-heap.md)
  - [并查集 Union Find](DOCS/tree-union-find.md)
  - [线段对 Segment Tree](DOCS/tree-segment-tree.md)
  - [字典树 Trie](DOCS/tree-trie.md)
  - [平衡二叉树 AVL Tree](DOCS/tree-avl-tree.md)
  - [红黑树 Red Back Tree](DOCS/tree-red-back-tree.md)
  
  
- ### 图论与图论算法    

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


- ## 排序算法
  
  - [选择排序 Select Sort](DOCS/select-sort.md)
  - [冒泡排序 Bubble Sort](DOCS/bubble-sort.md)
  - [插入排序 Insertion Sort](DOCS/insertion-sort.md)
  - [希尔排序 Shell Sort](DOCS/shell-sort.md)
  - [归并排序 Merge Sort](DOCS/merge-sort.md)
  - [快速排序 Quick Sort](DOCS/quick-sort.md)



数据结构系统文章可以关注我的公众号：“Mr Binary”

```
![lamphttp image](DOCS/wechat-qr.jpg)
```