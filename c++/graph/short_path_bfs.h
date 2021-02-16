//
// Created by HEADS on 2021/2/16.
//

#ifndef C___SHORT_PATH_BFS_H
#define C___SHORT_PATH_BFS_H

#include <iostream>
#include <vector>
#include <queue>
#include <stack>
#include <cassert>

template<typename Graph>
class ShortPathBFS {
private:
    Graph &G;
    bool *visited; // 是否有访问过
    int *from;     // 记录路径
    int *ord;      // s -> 目标的路径长度
    int s;         // 记录源点

public:
    ShortPathBFS(Graph &graph, int s):G(graph) {
        assert(s >= 0 && s < graph.V());
        this->s = s;
        visited = new bool[graph.V()];
        from = new int[graph.V()];
        ord = new int[graph.V()];

        for (int i = 0; i < graph.V(); i ++) {
            visited[i] = false;
            from[i] = -1;
            ord[i] = -1;
        }

        queue<int> q;
        q.push(s);
        visited[s] = true;
        ord[s] = 0;
        while(!q.empty()) {
            int v = q.front();
            q.pop();

            typename Graph::adjIterator adj(G, v);
            for (int i = adj.begin(); !adj.end(); i = adj.next()) {
                if (!visited[i]) {
                    q.push(i);
                    visited[i] = true;
                    from[i] = v;
                    ord[i] = ord[v] + 1;
                }
            }
        }
    }
    ~ShortPathBFS() {
        delete[] visited;
        delete[] from;
        delete[] ord;
    }

    bool hasPath(int w) {
        assert(w >= 0 && w < G.V());
        return visited[w];
    }

    void path(int w, vector<int> &vec) {
        assert(w >= 0 && w < G.V());

        stack<int> s;
        int p = w;
        while (p != -1) {
            s.push(p);
            p = from[p];
        }

        vec.clear();
        while(!s.empty()) {
            vec.push_back(s.top());
            s.pop();
        }
    }

    void showPath(int w) {
        assert(w >= 0 && w < G.V());

        vector<int> vec;
        path(w, vec);

        for (int i = 0; i < vec.size(); i ++) {
            cout << vec[i];
            if (i == vec.size()-1) {
                cout << endl;
            } else {
                cout << "->";
            }
        }
    }

    int length(int w) {
        assert(w >= 0 && w < G.V());
        return ord[w];
    }
};

#endif //C___SHORT_PATH_BFS_H
