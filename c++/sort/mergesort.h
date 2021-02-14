//
// Created by HEADS on 2021/2/14.
//

#ifndef C___MERGESORT_H
#define C___MERGESORT_H

namespace MergeSort {
    template<typename T>
    void insertSort(T arr[], int l, int r) {
        for (int i = l; i <= r; i ++) {
            T e = arr[i];
            int j;
            for (j = i + 1; j >= 0 && arr[j] > e; j --) {
                arr[j] = arr[j-1];
            }
            arr[j] = e;
        }
    }

    template<typename T>
    void _merge(T arr[], int l, int mid, int r) {
        T aux[r-l+1];
        // 临时空间
        for (int i = l; i <= r; i ++)
            aux[i-l] = arr[i];

        int i = l, j = mid + 1;
        for (int k = l; k <= r; k ++) {
            if (i > mid) {
                arr[k] = aux[j-l]; j++;
            } else if (j > r) {
                arr[k] = aux[i-l]; i++;
            } else if (aux[i-l] < aux[j-l]) {
                arr[k] = aux[i-l]; i++;
            } else {
                arr[k] = aux[j-l]; j++;
            }
        }
    }

    template<typename T>
    void _mergeSort(T arr[], int l, int r) {
        // 优化，当数据量小的时候，退货成插入排序
        if (r-l <= 15) {
            insertSort(arr, l, r);
            return;
        }

//        if (l >= r)
//            return;

        //int mid = (l+r)/2;
        int mid = l + (r-l)/2;
        _mergeSort(arr, l, mid);
        _mergeSort(arr, mid+1, r);
        // 优化， 只有当右边的元素小于左边的元素的时候才进行归并
        if (arr[mid] > arr[mid+1])
            _merge(arr, l, mid, r);
    }

    template<typename T>
    void mergeSort(T arr[], int n) {
        _mergeSort(arr, 0, n-1);
    }

    // 自底向上的排序
    template<typename T>
    void mergeSortBU(T arr[], int n) {
        bool breakFlag = false;
        for (int sz = 1; sz <= n; sz += sz) {
            for (int i = 0; i + sz < n; i += sz + sz) {
                int mid = i + sz -1;
                // 优化,只有左边元素比右边大的时候才归并
                if (arr[mid] > arr[mid + 1])
                    // 对arr[i...i_sz-1]和arr[i+sz...i+2*sz-1]进行归并
                    _merge(arr, i, mid, min(i + sz + sz -1, n -1));
            }
            if (breakFlag) break;
        }
    }
}

#endif //C___MERGESORT_H
