//
// Created by HEADS on 2021/2/14.
//

#include "maxheap.h"
#include <iostream>
#include "indexmaxheap.h"
#include "testhelper.h"
using namespace std;

template<typename T>
void heapSortUsingIndexMaxHeap(T arr[], int n) {
    IndexMaxHeap<T> indexMaxHeap = IndexMaxHeap<T>(n);
    for (int i = 0; i < n; i ++) {
        indexMaxHeap.insert(i, arr[i]);
    }
    assert(indexMaxHeap.testIndexesAndReverseIndexes());

    for (int i = n - 1; i >= 0; i --)
        arr[i] = indexMaxHeap.popMax();
}

int main() {
//    maxheap<int> h = maxheap<int>(100);
//
//    srand(time(NULL));
//
//    for (int i = 0; i < 20; i ++) {
//        h.insert(rand() % 100);
//    }
//
//    h.testPrint();
//
//
//    while(!h.isEmpty()) {
//        cout << h.popMax() << " ";
//    }
//    cout << endl;

    int n = 10000;
    int *arr = SortTestHelper::generateRandomArr(n, 0, n);
    SortTestHelper::testSort("Heap Sort Using Index-Max-Heap", heapSortUsingIndexMaxHeap, arr, n);
    delete[] arr;

    return 0;
}