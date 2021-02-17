//
// Created by HEADS on 2021/2/17.
// 最小生成树，prim算法实现
//

#ifndef C___LAZY_PRIM_MST_H
#define C___LAZY_PRIM_MST_H

#include <iostream>
#include <vector>
#include <cassert>
#include "minheap.h"
#include "edge.h"

using namespace std;

template<typename Graph, typename Weight>
class LazyPrimMST {
private:
    Graph &G;                  // 图的引用
    MinHeap<Edge<Weight>> pq;  // 最小堆
    bool *marked;              // 标记数组，在运行过程中标记是否被访问过
    vector<Edge<Weight>> mst;  // 最小生成树包含的所有边
    Weight mstWeight;          // 最小生成树的权值

    void visit(int v) {
        assert(!marked[v]);
        marked[v] = true;

        typename Graph::adjIterator adj(G, v);
        for (Edge<Weight>* e = adj.begin(); !adj.end(); e = adj.next())
            pq.insert(*e);
    }

public:
    LazyPrimMST(Graph &graph): G(graph), pq(MinHeap<Edge<Weight>>(graph.E())) {
        marked = new bool[G.V()];
        for (int i = 0; i < G.V(); i ++)
            marked[i] = false;
        mst.clear();

        visit(0);
        while(!pq.isEmpty()) {
            // 使用最小堆找出已经访问的边中权值最小的边
            Edge<Weight> e = pq.extractMin();
            // 如果这条边的两端都已经访问过了，则扔掉这条边
            if (marked[e.v()] == marked[e.w()])
                continue;
            // 否则，这条边则应该存在最小生成树中
            mst.push_back(e);

            // 访问和这条边连接的还没有被访问过的节点
            if (!marked[e.v()])
                visit(e.v());
            else
                visit(e.w());
        }

        // 计算最小生成树的权值
        mstWeight = mst[0].wt();
        for(int i = 1; i < mst.size(); i ++)
            mstWeight += mst[i].wt();
    }
    ~LazyPrimMST() {
        delete[] marked;
    }

    vector<Edge<Weight>> mstEdges() {
        return mst;
    }

    Weight result() {
        return mstWeight;
    }
};

#endif //C___LAZY_PRIM_MST_H
