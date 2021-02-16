//
// Created by HEADS on 2021/2/16.
// 稠密图 邻接矩阵实现
// 带权图
//

#ifndef C___WEIGHT_DENSE_GRAPH_H
#define C___WEIGHT_DENSE_GRAPH_H

#include <iostream>
#include <vector>
#include "edge.h"

using namespace std;

template<typename Weight>
class WeightDenseGraph {
private:
    int n;          // 表示节点数
    int m;          // 表示边数
    bool directed;  // 是否是有向图
    vector<vector<Edge<Weight> *>> g;  // 邻接矩阵

public:
    WeightDenseGraph(int n, bool directed) {
        assert(n >= 0);

        this->n = n;
        this->m = 0;
        this->directed = directed;
        g = vector<vector<Edge<Weight> *>>(n, vector<Edge<Weight> *>(n, NULL));
    }

    ~WeightDenseGraph() {
        for (int i = 0; i < n; i ++)
            for (int j = 0; j < n; j ++)
                if (g[i][j] != NULL)
                    delete g[i][j];
    }

    // 获取节点数
    int V() { return n;}
    // 获取边数
    int E() { return m;}

    // 添加一条边
    void addEdge(int v, int w, Weight weight) {
        assert(v >= 0 && v < n);
        assert(w >= 0 && w < n);

        // 防止重复计算
        if (hasEdge(v, w)) {
            delete g[v][w];
            if (v != w && !directed)
                delete g[w][v];
            m--;
        }

        g[v][w] = new Edge<Weight>(v, w, weight);
        // 如果是无向图，需要互相连接
        if (v != w && !directed)
            g[w][v] = new Edge<Weight>(w, v, weight);
        m++;
    }

    bool hasEdge(int v, int w) {
        assert(v >= 0 && v < n);
        assert(w >= 0 && w < n);

        return g[v][w] != NULL;
    }

    void show() {
        for (int i = 0; i < n; i ++) {
            for (int j = 0; j < n; j ++) {
                if (g[i][j]) {
                    cout << g[i][j]->wt() << "\t";
                } else {
                    cout << "NULL\t";
                }
            }
            cout << endl;
        }
    }

    // 迭代器
    class adjIterator {
    private:
        WeightDenseGraph &G;
        int v;     // 要迭代的顶点
        int idx; // 当前迭代到的索引位置

    public:
        adjIterator(DenseGraph &graph, int v): G(graph) {
            this->v = v;
            this->idx = -1;
        }
        ~adjIterator(){}

        Edge<Weight>* begin() {
            idx = -1;
            return next();
        }

        Edge<Weight>* next() {
            for (idx += 1; idx < G.V(); idx ++) {
                if (G.g[v][idx])
                    return G.g[v][idx];
            }
            return NULL;
        }

        bool end() {
            return idx >= G.V();
        }
    };

};

#endif //C___WEIGHT_DENSE_GRAPH_H
