//
// Created by HEADS on 2021/2/14.
//

#ifndef C___QUICKSORT_H
#define C___QUICKSORT_H

namespace QuickSort {
    // 对arr[l....r]部分进行partition操作
    // 返回p,使得arr[l...p-1] < arr[p]; arr[p+1...r] > arr[p]
    template<typename T>
    int _partition(T arr[], int l, int r) {
        // 优化，随机化标定点
        swap(arr[l], arr[rand()%(r-l+1) + l]);

        T v = arr[l];
        int j = l;

        for (int i = l + 1; i <= r; i ++) {
            if (arr[i] < v) {
                j++;
                swap(arr[j], arr[i]);
            }
        }
        swap(arr[l], arr[j]);
        return j;
    }

    // 对arr[l...r]部分进行快速排序
    template<typename T>
    void _quickSort(T arr[], int l, int r) {
        if (l >= r)
            return;
        int p = _partition(arr, l, r);
        _quickSort(arr, l, p - 1);
        _quickSort(arr, p + 1, r);
    }


    template<typename T>
    void quickSort(T arr[], int n) {
        srand(time(NULL));
        _quickSort(arr, 0, n-1);
    }
}

#endif //C___QUICKSORT_H
