//
// Created by HEADS on 2021/2/16.
// 从文件读取内容构建图
// 带权图
//

#ifndef C___WEIGHT_READ_GRAPH_H
#define C___WEIGHT_READ_GRAPH_H

#include <iostream>
#include <string>
#include <fstream>
#include <sstream>
#include <cassert>

using namespace std;

template<typename Graph, typename Weight>
class WeightReadGraph {
public:
    WeightReadGraph(Graph &graph, const string &filename) {
        ifstream file(filename);
        string line;
        int V, E;

        assert(file.is_open());

        // 第一行读取图中的节点个数和边的个数
        assert(getline(file, line));
        stringstream ss(line);
        ss>>V>>E;

        assert(V == graph.V());

        // 读取边的信息
        for (int i = 0; i < E; i ++) {
            assert(getline(file, line));
            stringstream ss(line);

            int a, b;
            Weight w;
            ss>>a>>b>>w;
            assert(a >= 0 && a < V);
            assert(b >= 0 && b < V);
            graph.addEdge(a, b, w);
        }
    }
};

#endif //C___WEIGHT_READ_GRAPH_H
