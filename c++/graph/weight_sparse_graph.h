//
// Created by HEADS on 2021/2/16.
// 稀疏图  邻接表实现
// 带权图
//

#ifndef C___WEIGHT_SPARSE_GRAPH_H
#define C___WEIGHT_SPARSE_GRAPH_H

#include <iostream>
#include <vector>
#include "edge.h"
using namespace std;

template<typename Weight>
class WeightSparseGraph {
private:
    int n;    // 节点数
    int m;    // 边数
    bool directed;          // 是否是有向图
    vector<vector<Edge<Weight> *>> g;  // 邻接表

public:
    WeightSparseGraph(int n, bool directed) {
        this->n = n;
        this->m = 0;
        this->directed = directed;
        g = vector<vector<Edge<Weight> *>>(n, vector<Edge<Weight> *>());
    }

    ~WeightSparseGraph() {
        for (int i = 0; i < n; i ++)
            for (int j = 0; j < g[i].size(); j ++)
                delete g[i][j];
    }


    int V() {return n;}
    int E() {return m;}

    void addEdge(int v, int w, double weight) {
        assert(v >= 0 && v < n);
        assert(w >= 0 && w < n);

        g[v].push_back(new Edge<Weight>(v, w, weight));
        if (v != w && !directed)
            g[w].push_back(new Edge<Weight>(w, v, weight));
        m++;
    }

    bool hasEdge(int v, int w) {
        assert(v >= 0 && v < n);
        assert(w >= 0 && w < n);
        for (int i = 0; i < n; i ++)
            if (g[v][i] != NULL)
                return true;
        return false;
    }

    void show() {
        for (int i = 0; i < n; i ++) {
            cout << "vertex " << i << ":\t";
            for (int j = 0; j < g[i].size(); j ++)
                cout << "( to:" << g[i][j]->w() << ", wt"<< g[i][j]->wt()<< ")\t";
            cout << endl;
        }
    }

    // 迭代器
    class adjIterator {
    private:
        WeightSparseGraph &G;
        int v;   // 要迭代的顶点
        int idx; // 当前遍历到的索引
    public:
        adjIterator(WeightSparseGraph &graph, int v): G(graph) {
            this->v = v;
            this->idx = 0;
        }

        ~adjIterator(){}

        // 返回与顶点v相边的第一个顶点
        Edge<Weight>* begin() {
            idx = 0;
            if (G.g[v].size())
                return G.g[v][idx];
            return NULL;
        }

        // 返回图G中与顶点v相连的下一个顶点
        Edge<Weight>* next() {
            idx++;
            if (idx < G.g[v].size())
                return G.g[v][idx];
            return NULL;
        }

        // 查看是否已经迭代完了图G中与顶点v相连接的所有顶点
        bool end() {
            return idx >= G.g[v].size();
        }
    };
};

#endif //C___WEIGHT_SPARSE_GRAPH_H
