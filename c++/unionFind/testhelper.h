//
// Created by HEADS on 2021/2/13.
//

#ifndef DATA_STRUCTURE_TESTHELPER_H
#define DATA_STRUCTURE_TESTHELPER_H

#include <iostream>
#include <ctime>
#include "union1.h"
#include "union2.h"
#include "union3.h"
#include "union4.h"
#include "union5.h"

using namespace std;

namespace testhelper {
    // test1
    void testUF1(int n) {
        srand(time(NULL));
        UF1::UnionFind uf = UF1::UnionFind(n);

        time_t start_time = clock();

        // n次并操作
        for (int i = 0; i < n; i ++) {
            int a = rand()%n;
            int b = rand()%n;
            uf.unionElm(a, b);
        }

        // n次查操作
        for (int i = 0; i < n; i ++) {
            int a = rand() % n;
            int b = rand() % n;
            uf.isConnect(a,b);
        }

        time_t end_time = clock();

        cout << "UF1, " << 2 * n << " ops," << double(end_time-start_time)/CLOCKS_PER_SEC << " s" << endl;
    }

    // test2
    void testUF2(int n) {
        srand(time(NULL));
        UF2::UnionFind uf = UF2::UnionFind(n);

        time_t start_time = clock();

        // n次并操作
        for (int i = 0; i < n; i ++) {
            int a = rand()%n;
            int b = rand()%n;
            uf.unionElm(a, b);
        }

        // n次查操作
        for (int i = 0; i < n; i ++) {
            int a = rand() % n;
            int b = rand() % n;
            uf.isConnect(a,b);
        }

        time_t end_time = clock();

        cout << "UF2, " << 2 * n << " ops," << double(end_time-start_time)/CLOCKS_PER_SEC << " s" << endl;
    }

    // test3
    void testUF3(int n) {
        srand(time(NULL));
        UF3::UnionFind uf = UF3::UnionFind(n);

        time_t start_time = clock();

        // n次并操作
        for (int i = 0; i < n; i ++) {
            int a = rand()%n;
            int b = rand()%n;
            uf.unionElm(a, b);
        }

        // n次查操作
        for (int i = 0; i < n; i ++) {
            int a = rand() % n;
            int b = rand() % n;
            uf.isConnect(a,b);
        }

        time_t end_time = clock();

        cout << "UF3, " << 2 * n << " ops," << double(end_time-start_time)/CLOCKS_PER_SEC << " s" << endl;
    }

    // test4
    void testUF4(int n) {
        srand(time(NULL));
        UF4::UnionFind uf = UF4::UnionFind(n);

        time_t start_time = clock();

        // n次并操作
        for (int i = 0; i < n; i ++) {
            int a = rand()%n;
            int b = rand()%n;
            uf.unionElm(a, b);
        }

        // n次查操作
        for (int i = 0; i < n; i ++) {
            int a = rand() % n;
            int b = rand() % n;
            uf.isConnect(a,b);
        }

        time_t end_time = clock();

        cout << "UF4, " << 2 * n << " ops," << double(end_time-start_time)/CLOCKS_PER_SEC << " s" << endl;
    }


    // test5
    void testUF5(int n) {
        srand(time(NULL));
        UF5::UnionFind uf = UF5::UnionFind(n);

        time_t start_time = clock();

        // n次并操作
        for (int i = 0; i < n; i ++) {
            int a = rand()%n;
            int b = rand()%n;
            uf.unionElm(a, b);
        }

        // n次查操作
        for (int i = 0; i < n; i ++) {
            int a = rand() % n;
            int b = rand() % n;
            uf.isConnect(a,b);
        }

        time_t end_time = clock();

        cout << "UF5, " << 2 * n << " ops," << double(end_time-start_time)/CLOCKS_PER_SEC << " s" << endl;
    }
}

#endif //DATA_STRUCTURE_TESTHELPER_H