//
// Created by HEADS on 2021/2/16.
// 稀疏图  邻接表实现
//

#ifndef C___SPARSE_GRAPH_H
#define C___SPARSE_GRAPH_H

#include <iostream>
#include <vector>
using namespace std;

class SparseGraph {
private:
    int n;    // 节点数
    int m;    // 边数
    bool directed;          // 是否是有向图
    vector<vector<int>> g;  // 邻接表

public:
    SparseGraph(int n, bool directed) {
        this->n = n;
        this->m = 0;
        this->directed = directed;
        g = vector<vector<int>>(n, vector<int>());
    }

    ~SparseGraph() {}


    int V() {return n;}
    int E() {return m;}

    void addEdge(int v, int w) {
        assert(v >= 0 && v < n);
        assert(w >= 0 && w < n);

        g[v].push_back(w);
        if (v != w && !directed)
            g[w].push_back(v);
        m++;
    }

    bool hasEdge(int v, int w) {
        assert(v >= 0 && v < n);
        assert(w >= 0 && w < n);
        for (int i = 0; i < n; i ++)
            if (g[v][i] == w)
                return true;
        return false;
    }

    void show() {
        for (int i = 0; i < n; i ++) {
            cout << "vertex " << i << ":\t";
            for (int j = 0; j < g[i].size(); j ++)
                cout << g[i][j] << "\t";
            cout << endl;
        }
    }

    // 迭代器
    class adjIterator {
    private:
        SparseGraph &G;
        int v;   // 要迭代的顶点
        int idx; // 当前遍历到的索引
    public:
        adjIterator(SparseGraph &graph, int v): G(graph) {
            this->v = v;
            this->idx = 0;
        }

        ~adjIterator(){}

        // 返回与顶点v相边的第一个顶点
        int begin() {
            idx = 0;
            if (G.g[v].size())
                return G.g[v][idx];
            return -1;
        }

        // 返回图G中与顶点v相连的下一个顶点
        int next() {
            idx++;
            if (idx < G.g[v].size())
                return G.g[v][idx];
            return -1;
        }

        // 查看是否已经迭代完了图G中与顶点v相连接的所有顶点
        bool end() {
            return idx >= G.g[v].size();
        }
    };
};

#endif //C___SPARSE_GRAPH_H
