//
// Created by HEADS on 2021/2/14.
//

#ifndef C___HEAPSORT_H
#define C___HEAPSORT_H
#include "maxheap.h"

namespace HeapSort {
    // shift down
    // 表示在剩下的n个元素的数组进行shiftdown, k 表示从哪开始shiftdown
    template<typename T>
    void __shiftDown(T arr[], int n, int k) {
        while( 2*k + 1 < n) {
            int j = 2 * k + 1;
            if (j + 1 < n && arr[j + 1] > arr[j]) {
                j += 1;
            }
            if (arr[k] >= arr[j])
                break;
            swap(arr[k], arr[j]);
            k = j;
        }
    }

    // 原地堆排序
    template<typename T>
    void heapSort(T arr[], int n) {
        for (int i = (n-1-1)/2; i >= 0; i --) {
            __shiftDown(arr, n-1, i);
        }

        for (int i = n-1; i > 0; i --) {
            swap(arr[i], arr[0]);
            __shiftDown(arr, i, 0);
        }
    }

    // 堆排序
    template<typename T>
    void heapSort1(T arr[], int n) {
        maxheap<T> h = maxheap<T>(n);
        for (int i = 0; i < n; i ++)
            h.insert(arr[i]);

        for (int i = n - 1; i >= 0; i --)
            arr[i] = h.popMax();

    }

    template<typename T>
    void heapSort2(T arr[], int n) {
        maxheap<T> h = maxheap<T>(arr, n);
        for (int i = n-1; i >= 0; i --)
            arr[i] = h.popMax();
    }
}

#endif //C___HEAPSORT_H
