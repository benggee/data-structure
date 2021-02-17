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
#include "lazy_prim_mst.h"
#include "prim_mst.h"

using namespace std;

void graphTest1();
void readGraph();               // 从文件读取内容建图
void testComponent();           // 图的实现
void testPath();                // 求路径
void testShortPathBFS();        // 最短路径-广度优先
void testWeightDenseGraph();    // 带权图
void testLazyPrim();            // 最小生成树
void testPrim();                // 最小生成树，索引堆实现

int main() {
    // readGraph();
    // testComponent();
    // testPath();
    // testShortPathBFS();
    // testWeightDenseGraph();
    // testLazyPrim();
    testPrim();

    return 0;
}

void testPrim() {
    string filename1 = "/Users/HEADS/app/data-structure/c++/graph/test-data/prim1.txt";
    int V1 = 8;

    string filename2 = "/Users/HEADS/app/data-structure/c++/graph/test-data/prim2.txt";
    int V2 = 250;

    string filename3 = "/Users/HEADS/app/data-structure/c++/graph/test-data/prim3.txt";
    int V3 = 1000;

    string filename4 = "/Users/HEADS/app/data-structure/c++/graph/test-data/prim4.txt";
    int V4 = 10000;

//    string filename5 = "testG5.txt";
//    int V5 = 1000000;


    // 文件读取
    WeightSparseGraph<double> g1 = WeightSparseGraph<double>(V1, false);
    WeightReadGraph<WeightSparseGraph<double>,double> readGraph1(g1, filename1);
    cout<<filename1<<" load successfully."<<endl;

    WeightSparseGraph<double> g2 = WeightSparseGraph<double>(V2, false);
    WeightReadGraph<WeightSparseGraph<double>,double> readGraph2(g2, filename2);
    cout<<filename2<<" load successfully."<<endl;

    WeightSparseGraph<double> g3 = WeightSparseGraph<double>(V3, false);
    WeightReadGraph<WeightSparseGraph<double>,double> readGraph3(g3, filename3);
    cout<<filename3<<" load successfully."<<endl;

    WeightSparseGraph<double> g4 = WeightSparseGraph<double>(V4, false);
    WeightReadGraph<WeightSparseGraph<double>,double> readGraph4(g4, filename4);
    cout<<filename4<<" load successfully."<<endl;

//    SparseGraph<double> g5 = SparseGraph<double>(V5, false);
//    ReadGraph<SparseGraph<double>,double> readGraph5(g5, filename5);
//    cout<<filename5<<" load successfully."<<endl;

    cout<<endl;


    clock_t startTime, endTime;

    // Test Lazy Prim MST
    cout<<"Test Lazy Prim MST:"<<endl;

    startTime = clock();
    LazyPrimMST<WeightSparseGraph<double>, double> lazyPrimMST1(g1);
    endTime = clock();
    cout<<"Test for G1: "<<(double)(endTime-startTime)/CLOCKS_PER_SEC<<" s."<<endl;

    startTime = clock();
    LazyPrimMST<WeightSparseGraph<double>, double> lazyPrimMST2(g2);
    endTime = clock();
    cout<<"Test for G2: "<<(double)(endTime-startTime)/CLOCKS_PER_SEC<<" s."<<endl;

    startTime = clock();
    LazyPrimMST<WeightSparseGraph<double>, double> lazyPrimMST3(g3);
    endTime = clock();
    cout<<"Test for G3: "<<(double)(endTime-startTime)/CLOCKS_PER_SEC<<" s."<<endl;

    startTime = clock();
    LazyPrimMST<WeightSparseGraph<double>, double> lazyPrimMST4(g4);
    endTime = clock();
    cout<<"Test for G4: "<<(double)(endTime-startTime)/CLOCKS_PER_SEC<<" s."<<endl;

//    startTime = clock();
//    LazyPrimMST<SparseGraph<double>, double> lazyPrimMST5(g5);
//    endTime = clock();
//    cout<<"Test for G5: "<<(double)(endTime-startTime)/CLOCKS_PER_SEC<<" s."<<endl;

    cout<<endl;


    // Test Prim MST
    cout<<"Test Prim MST:"<<endl;

    startTime = clock();
    PrimMST<WeightSparseGraph<double>, double> PrimMST1(g1);
    endTime = clock();
    cout<<"Test for G1: "<<(double)(endTime-startTime)/CLOCKS_PER_SEC<<" s."<<endl;

    startTime = clock();
    PrimMST<WeightSparseGraph<double>, double> PrimMST2(g2);
    endTime = clock();
    cout<<"Test for G2: "<<(double)(endTime-startTime)/CLOCKS_PER_SEC<<" s."<<endl;

    startTime = clock();
    PrimMST<WeightSparseGraph<double>, double> PrimMST3(g3);
    endTime = clock();
    cout<<"Test for G3: "<<(double)(endTime-startTime)/CLOCKS_PER_SEC<<" s."<<endl;

    startTime = clock();
    PrimMST<WeightSparseGraph<double>, double> PrimMST4(g4);
    endTime = clock();
    cout<<"Test for G4: "<<(double)(endTime-startTime)/CLOCKS_PER_SEC<<" s."<<endl;

//    startTime = clock();
//    PrimMST<SparseGraph<double>, double> PrimMST5(g5);
//    endTime = clock();
//    cout<<"Test for G5: "<<(double)(endTime-startTime)/CLOCKS_PER_SEC<<" s."<<endl;

    cout<<endl;
}

void testLazyPrim() {
    string filename = "/Users/HEADS/app/data-structure/c++/graph/test-data/lazy-prim.txt";
    int v = 8;

    cout << fixed << setprecision(2);  // 设置精度保留两位小数

    WeightSparseGraph<double> g = WeightSparseGraph<double>(v, false);
    WeightReadGraph<WeightSparseGraph<double>, double> readGraph(g, filename);

    cout << "Lazy Prim MST:" << endl;
    LazyPrimMST<WeightSparseGraph<double>, double> lazyPrimMst(g);
    vector<Edge<double>> mst = lazyPrimMst.mstEdges();
    for (int i = 0; i < mst.size(); i ++) {
        cout << mst[i] << endl;
    }

    cout << "The MST weight is:" << lazyPrimMst.result()<<endl;
    cout << endl;
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