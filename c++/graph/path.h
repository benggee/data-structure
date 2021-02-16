//
// Created by HEADS on 2021/2/16.
//

#ifndef C___PATH_H
#define C___PATH_H

#include <iostream>
#include <vector>
#include <stack>

template<typename Graph>
class Path {
private:
    Graph &G;
    bool *visited;
    int s; // 到s的路径
    int *from;

    void dfs(int v) {
        visited[v] = true;
        typename Graph::adjIterator adj(G, v);
        for (int i = adj.begin(); !adj.end(); i = adj.next()) {
            if (!visited[i]) {
                from[i] = v;
                dfs(i);
            }
        }
    }

public:
    Path(Graph &graph, int s): G(graph) {
        this->s = s;
        visited = new bool[G.V()];
        from = new int[G.V()];
        for (int i = 0; i < G.V(); i ++) {
            visited[i] = false;
            from[i] = -1;
        }

        dfs(s);
    }

    ~Path(){
        delete[] visited;
        delete[] from;
    }

    bool hasPath(int w) {
        assert(s >= 0 && s < G.V());
        return visited[w];
    }

    void path(int w, vector<int> &vec) {
        assert(hasPath(w));

        int p = w;
        stack<int> s;
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

    void printPath(int w) {
        assert(hasPath(w));

        vector<int> source;
        path(w, source);

        for (int i = 0; i < source.size(); i ++) {
            cout << source[i];
            if (i == source.size() - 1) {
                cout << endl;
            } else {
                cout << "->";
            }
        }
    }
};

#endif //C___PATH_H
