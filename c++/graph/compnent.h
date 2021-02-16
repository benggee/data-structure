//
// Created by HEADS on 2021/2/16.
//

#ifndef C___COMPNENT_H
#define C___COMPNENT_H

#include <iostream>
#include <cassert>
using namespace std;

template<typename Graph>
class Component {
private:
    Graph &G;
    int ccount;   // 连通分量
    int *id;      // 并查集，两个节点是否连通
    bool *visited; // 是否被访问过

    void dfs(int v) {
        assert(v >= 0 && v < G.V());

        visited[v] = true;
        id[v] = ccount; // 如果节点的联通分量是一样的，说明一定是联通的
        typename Graph::adjIterator adj(G, v);
        for (int i = adj.begin(); !adj.end(); i = adj.next())
            if (!visited[i])
                dfs(i);
    }
public:
    Component(Graph &graph):G(graph) {
        visited = new bool[graph.V()];
        id = new int[graph.V()];

        ccount = 0;
        for (int i = 0; i < G.V(); i ++) {
            visited[i] = false;
            id[i] = -1;
        }

        // 求联通分量
        for (int i = 0; i < G.V(); i ++) {
            if (!visited[i]) {
                dfs(i);
                ccount++;
            }
        }
    }

    ~Component() {
        delete[] visited;
        delete[] id;
    }

    int count() {
        return ccount;
    }

    bool isConnect(int v ,int w) {
        assert(v >= 0 && v < G.V());
        assert(w >= 0 && w < G.V());

        return id[v] == id[w];
    }
};

#endif //C___COMPNENT_H
