//
// Created by HEADS on 2021/2/16.
//

#ifndef C___EDGE_H
#define C___EDGE_H

#include <iostream>
#include <cassert>
using namespace std;

template<typename Weight>
class Edge {
private:
    int a, b;  // 边的两个端点
    Weight weight;

public:
    Edge(int a, int b, Weight weight) {
        this->a = a;
        this->b = b;
        this->weight = weight;
    }
    Edge(){}
    ~Edge(){}

    int v() {return a;} // 返回第一个顶点
    int w() {return b;} // 返回第二个顶点
    Weight wt(){return weight;} // 返回权值

    // 给定一个顶点，获取另外一个顶点
    int other(int x) {
        assert(x == a || x == b);
        return x == a ? b : a;
    }

    // 输出边的信息
    friend ostream& operator<<(ostream &os, const Edge &e){
        os<<e.a<<"-"<<e.b<<": "<<e.weight;
        return os;
    }

    // 比较运算符重载
    bool operator<(Edge<Weight>& e){
        return weight < e.wt();
    }
    bool operator<=(Edge<Weight>& e){
        return weight <= e.wt();
    }
    bool operator>(Edge<Weight>& e){
        return weight > e.wt();
    }
    bool operator>=(Edge<Weight>& e){
        return weight >= e.wt();
    }
    bool operator==(Edge<Weight>& e){
        return weight == e.wt();
    }
};

#endif //C___EDGE_H
