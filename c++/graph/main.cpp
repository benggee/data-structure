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

using namespace std;

void graphTest1();
void readGraph();
void testComponent();

int main() {
    // readGraph();
    testComponent();

    return 0;
}

void testComponent() {
    string filename = "/Users/HEADS/app/data-structure/c++/graph/testg1.txt";
    SparseGraph g1(13, false);
    ReadGraph<SparseGraph> readGraph1(g1, filename);
    Component<SparseGraph> component(g1);
    cout << "testg1.txt, Using Sparse Graph, Component Count:" << component.count() << endl;

    cout << endl;
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