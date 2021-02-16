//
// Created by HEADS on 2021/2/16.
// 稠密图 邻接矩阵实现
//

#ifndef C___DENSE_GRAPH_H
#define C___DENSE_GRAPH_H

#include <iostream>
#include <vector>

using namespace std;

class DenseGraph {
private:
    int n;          // 表示节点数
    int m;          // 表示边数
    bool directed;  // 是否是有向图
    vector<vector<bool>> g;  // 邻接矩阵

public:
    DenseGraph(int n, bool directed) {
        this->n = n;
        this->m = 0;
        this->directed = directed;

        for (int i = 0; i < n; i ++)
            g.push_back(vector<bool>(n, false));

    }

    ~DenseGraph() {}

    // 获取节点数
    int V() { return n;}
    // 获取边数
    int E() { return m;}

    // 添加一条边
    void addEdge(int v, int w) {
        assert(v >= 0 && v < n);
        assert(w >= 0 && w < n);

        if (hasEdge(v, w))
            return;

        g[v][w] = true;
        // 如果是无向图，需要互相连接
        if (!directed)
            g[w][v] = true;
        m++;
    }

    bool hasEdge(int v, int w) {
        assert(v >= 0 && v < n);
        assert(w >= 0 && w < n);

        return g[v][w];
    }

    void show() {
        for (int i = 0; i < n; i ++) {
            for (int j = 0; j < n; j ++) {
                cout << g[i][j] << "\t";
            }
            cout << endl;
        }
    }

    // 迭代器
    class adjIterator {
    private:
        DenseGraph &G;
        int v;     // 要迭代的顶点
        int idx; // 当前迭代到的索引位置

    public:
        adjIterator(DenseGraph &graph, int v): G(graph) {
            this->v = v;
            this->idx = 0;
        }
        ~adjIterator(){}

        int begin() {
            idx = 0;
            return next();
        }

        int next() {
            for (idx += 1; idx < G.V(); idx ++) {
                if (G.g[v][idx])
                    return idx;
            }
            return -1;
        }

        bool end() {
            return idx >= G.V();
        }
    };

};

#endif //C___DENSE_GRAPH_H
