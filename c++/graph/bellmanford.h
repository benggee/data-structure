//
// Created by HEADS on 2021/2/17.
//

#ifndef C___BELLMANFORD_H
#define C___BELLMANFORD_H

#include <iostream>
#include <vector>
#include "edge.h"
#include <stack>

using namespace std;

template<typename Graph, typename Weight>
class BellmanFord{
private:
    Graph &G;                   // 图
    int s;                      // 起始点
    Weight* distTo;             // distTo[i]存储从起始点s到i的最短路径长度
    vector<Edge<Weight>*> from; // from[i]记录最短路径中，到达i点的边是哪一条

    bool hasNegativeCycle;      // 标记图中是否有负权环

    bool detectNegativeCycle() {
        for (int i = 0; i < G.V(); i ++) {
            typename Graph::adjIterator adj(G, i);
            for (Edge<Weight>* e = adj.begin(); !adj.end(); e = adj.next()) {
                if (from[e->v()] && distTo[e->v()] + e->wt() < distTo[e->w()])
                    return true;
            }
        }
        return false;
    }

public:
    BellmanFord(Graph &graph, int s):G(graph) {
        this->s = s;
        distTo = new Weight[G.V()];
        // 初始化所有的节点s都不可达，由from数组来表示
        for (int i = 0; i < G.V(); i ++)
            from.push_back(NULL);

        // 设置distTo[s] = 0, 并且from[s]不为NULL，表示初始节点可达且距离为0
        distTo[s] = Weight();
        // 这里我们from[s]的内容是new出来的，注意要在析构函数delete
        from[s] = new Edge<Weight>(s, s, Weight());

        // Bellman-Ford的过程
        // 进行v-1次循环，每一次循环求出从起点到其余所有点，最多使用pass步可到达的最短距离
        for (int pass = 1; pass < G.V(); pass ++) {
            // 每次循环中对所有边进行一遍松弛操作
            // 遍历所有边的方式是先遍历所有的顶点，然后遍历和所有顶点相邻的所有边
            for (int i = 0; i < G.V(); i ++) {
                typename Graph::adjIterator adj(G, i);
                for (Edge<Weight>* e = adj.begin(); !adj.end(); e = adj.next()) {
                    // 对于每一个边首先判断e->v()可达
                    // 之后看如果e->w()以前没有到达过，显然我们可以更新distTo[e->w()]
                    // 或者e->w()以前虽然到过，但是通过这个e我们可以获得一个更短的距离，即可以进行一次松弛操作，我们可以更新distTo[e->w()]
                    if (from[e->v()] && (!from[e->w()] || distTo[e->v()] + e->wt() < distTo[e->w()])) {
                        distTo[e->w()] = distTo[e->v()] + e->wt();
                        from[e->w()] = e;
                    }
                }
            }
        }

        hasNegativeCycle = detectNegativeCycle();
    }

    ~BellmanFord(){
        delete[] distTo;
        delete from[s];
    }

    bool negativeCycle() {
        return hasNegativeCycle;
    }

    Weight shortPathTo(int w) {
        assert(w >= 0 && w < G.V());
        assert(!hasNegativeCycle);
        assert(hasPathTo(w));
        return distTo[w];
    }

    bool hasPathTo(int w) {
        assert(w >= 0 && w < G.V());
        return from[w] != NULL;
    }

    void shortPath(int w, vector<Edge<Weight>> &vec) {
        assert(w >= 0 && w < G.V());
        assert(!hasNegativeCycle);
        assert(hasPathTo(w));

        stack<Edge<Weight>*> s;
        Edge<Weight> *e = from[w];
        while(e->v() != this->s) {
            s.push(e);
            e = from[e->v()];
        }
        s.push(e);

        while(!s.empty()) {
            e = s.top();
            vec.push_back(*e);
            s.pop();
        }
    }

    void showPath(int w) {
        assert(w >= 0 && w < G.V());
        assert(!hasNegativeCycle);
        assert(hasPathTo(w));

        vector<Edge<Weight>> vec;
        shortPath(w, vec);
        for (int i = 0; i < vec.size(); i ++) {
            cout << vec[i].v()<< "->";
            if (i == vec.size()-1)
                cout << vec[i].w()<< endl;
        }
    }
};

#endif //C___BELLMANFORD_H
