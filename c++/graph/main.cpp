//
// Created by HEADS on 2021/2/16.
//

#include <iostream>
#include <string>
#include <filesystem>
#include "sparse_graph.h"
#include "dense_graph.h"
#include "read_graph.h"
#include "compnent.h"
#include "path.h"
#include "short_path_bfs.h"
#include "weight_dense_graph.h"
#include "weight_read_graph.h"
#include "weight_sparse_graph.h"

using namespace std;

void graphTest1();
void readGraph();
void testComponent();
void testPath();
void testShortPathBFS();
void testWeightDenseGraph();

int main() {
    // readGraph();
    // testComponent();
    // testPath();
    // testShortPathBFS();
    testWeightDenseGraph();

    return 0;
}

void testWeightDenseGraph() {
    string filename = "/Users/HEADS/app/data-structure/c++/graph/test-data/weight-graph1.txt";

    int v = 8;
    cout << fixed << setprecision(2);  // 设置精度保留两位小数

    WeightDenseGraph<double> g1(v, false);
    WeightReadGraph<WeightDenseGraph<double>, double> readGraph1(g1, filename);
    g1.show();
    cout << endl;

    WeightSparseGraph<double> g2(v, false);
    WeightReadGraph<WeightSparseGraph<double>, double> readGraph2(g2, filename);
    g2.show();
    cout << endl;
}

void testShortPathBFS() {
    string filename = "/Users/HEADS/app/data-structure/c++/graph/test-path.txt";
    SparseGraph g1(7, false);
    ReadGraph<SparseGraph> readGraph1(g1, filename);
    ShortPathBFS<SparseGraph> path(g1, 0);
    g1.show();
    cout << "BFS Paht: " << endl;
    path.showPath(6);
}


void testPath() {
    string filename = "/Users/HEADS/app/data-structure/c++/graph/test-path.txt";
    SparseGraph g1(7, false);
    ReadGraph<SparseGraph> readGraph1(g1, filename);
    Path<SparseGraph> path(g1, 0);
    g1.show();
    cout << "DFS Path:" << endl;
    path.printPath(6);
}

void testComponent() {
    string filename = "/Users/HEADS/app/data-structure/c++/graph/testg1.txt";
    SparseGraph g1(13, false);
    ReadGraph<SparseGraph> readGraph1(g1, filename);
    Component<SparseGraph> component(g1);
    cout << "testg1.txt, Using Sparse Graph, Component Count:" << component.count() << endl;

    DenseGraph g2(13, false);
    ReadGraph<DenseGraph> readGraph2(g2, filename);
    Component<DenseGraph> component1(g2);
    cout << "testg1.txt, Using Dense Graph, Component Count:" << component1.count() << endl;

    cout << endl;

    filename = "/Users/HEADS/app/data-structure/c++/graph/testg2.txt";
    SparseGraph g3(6, false);
    ReadGraph<SparseGraph> readGraph3(g3, filename);
    Component<SparseGraph> component2(g3);
    cout << "testg2.txt, Using Sparse Graph, Component Count:" << component2.count() << endl;


    DenseGraph g4(6, false);
    ReadGraph<DenseGraph> readGraph4(g4, filename);
    Component<DenseGraph> component3(g4);
    cout << "testg2.txt, Using Dense Graph, Component Count:" << component3.count() << endl;
}

void readGraph() {
    string filename = "/Users/HEADS/app/data-structure/c++/graph/testg1.txt";
    SparseGraph g1(13, false);
    ReadGraph<SparseGraph> readGraph1(g1, filename);
    cout << "test g1 in sparce graph:" << endl;
    g1.show();

    cout << endl;

    DenseGraph g2(13, false);
    ReadGraph<DenseGraph> readGraph2(g2, filename);
    cout << "test g1 in dense graph:" << endl;
    g2.show();

    cout << endl;


    filename = "/Users/HEADS/app/data-structure/c++/graph/testg2.txt";
    SparseGraph g3(6, false);
    ReadGraph<SparseGraph> readGraph3(g3, filename);
    cout << "test g2 in sparse graph:" << endl;
    g3.show();

    cout << endl;

    DenseGraph g4(6, false);
    ReadGraph<DenseGraph> readGraph4(g4, filename);
    cout << "test g2 in dense graph:" << endl;
    g4.show();

}


void graphTest1() {
    int N = 20;
    int M = 100;

    srand(time(NULL));

    SparseGraph g1(N, false);
    for (int i = 0; i < M; i ++) {
        int a = rand()%N;
        int b = rand()%N;
        g1.addEdge(a, b);
    }

    for (int v = 0; v < N; v ++) {
        cout << ":";
        SparseGraph::adjIterator adj(g1, v);
        for (int w = adj.begin(); !adj.end(); w = adj.next())
            cout << w << " ";
        cout << endl;
    }

    cout << endl;

    DenseGraph g2(N, false);
    for (int i = 0; i < M; i ++) {
        int a = rand() % N;
        int b = rand() % N;
        g2.addEdge(a, b);
    }

    for (int v = 0; v < N; v ++) {
        cout << ":" ;
        DenseGraph::adjIterator adj(g2, v);
        for (int w = adj.begin(); !adj.end(); w = adj.next())
            cout << w << " ";
        cout << endl;
    }
}