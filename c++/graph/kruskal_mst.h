//
// Created by HEADS on 2021/2/17.
//

#ifndef C___KRUSKAL_MST_H
#define C___KRUSKAL_MST_H

#include <iostream>
#include <vector>
#include "uf.h"
#include "minheap.h"
#include "edge.h"

template<typename Graph, typename Weight>
class KruskalMST {
    vector<Edge<Weight>> mst;  // 最小生成树所包含的所有边
    Weight mstWeight;          // 最小生成树的权值

public:
    KruskalMST(Graph &graph) {
        // 将图中的所有边存放到一个最小值，达到排序的效果
        MinHeap<Edge<Weight>> pq(graph.E());

        for (int i = 0; i < graph.V(); i ++) {
            typename Graph::adjIterator adj(graph, i);
            for (Edge<Weight> *e = adj.begin(); !adj.end(); e = adj.next())
                if (e->v() < e->w())
                    pq.insert(*e);
        }

        // 创建一个并查集，来查看已经访问过的节点是存在环
        UnionFind uf = UnionFind(graph.V());
        while(!pq.isEmpty() && mst.size() < graph.V() - 1) {
            // 从最小堆中依次从小到大取出所有的边
            Edge<Weight> e = pq.extractMin();
            // 如果该边的两个端点是联通的，说明加入这条边将产生环，扔掉这条边
            if (uf.isConnect(e.v(), e.w()))
                continue;

            // 否则，将这条边添加进最小生成树，同时标记边的两个元素的连通
            mst.push_back(e);
            uf.unionElm(e.v(), e.w());
        }

        mstWeight = mst[0].wt();
        for (int i = 0; i < mst.size(); i ++)
            mstWeight += mst[i].wt();
    }

    ~KruskalMST(){}

    vector<Edge<Weight>> mstEdges() {
        return mst;
    }

    Weight result() {
        return mstWeight;
    }
};

#endif //C___KRUSKAL_MST_H
