//
// Created by HEADS on 2021/2/17.
//

#ifndef C___PRIM_MST_H
#define C___PRIM_MST_H

#include "edge.h"
#include "index_min_heap.h"

template<typename Graph, typename Weight>
class PrimMST {
    Graph &G;
    IndexMinHeap<Weight> ipq;
    vector<Edge<Weight>*> edgeTo;
    bool* marked;
    vector<Edge<Weight>> mst;
    Weight mstWeight;

    void visit(int v) {
        assert(!marked[v]);
        marked[v] = true;

        typename Graph::adjIterator adj(G, v);
        for (Edge<Weight>* e = adj.begin(); !adj.end(); e = adj.next()) {
            int w = e->other(v);
            if (!marked[w]) {
                if (!edgeTo[w]) {
                    edgeTo[w] = e;
                    ipq.insert(w, e->wt());
                }
                else if (e->wt() < edgeTo[w]->wt()) {
                    edgeTo[w] = e;
                    ipq.change(w, e->wt());
                }
            }
        }
    }

public:
    PrimMST(Graph &graph):G(graph), ipq(IndexMinHeap<double>(graph.V())) {
        assert(graph.E() >= 1);

        marked = new bool[G.V()];
        for(int i = 0; i < G.V(); i ++) {
            marked[i] = false;
            edgeTo.push_back(NULL);
        }
        mst.clear();

        visit(0);
        while(!ipq.isEmpty()) {
            int v = ipq.extractMinIndex();
            assert(edgeTo[v]);
            mst.push_back(*edgeTo[v]);
            visit(v);
        }

        mstWeight = mst[0].wt();
        for(int i = 0; i < mst.size(); i ++)
            mstWeight += mst[i].wt();
    }

    ~PrimMST() {
        delete[] marked;
    }

    vector<Edge<Weight>> mstEdges() {
        return mst;
    }

    Weight result() {
        return mstWeight;
    }
};

#endif //C___PRIM_MST_H
