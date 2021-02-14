//
// Created by HEADS on 2021/2/14.
//

#ifndef C___SHELLSORT_H
#define C___SHELLSORT_H

namespace ShellSort {
    template<typename T>
    void shellSort(T arr[], int n) {
        // 计算increment sequence:1,4,13,40,121,364,1093...
        int h = 1;
        while(h < n/3)
            h = 3 * h + 1;

        while(h >= 1) {
            for (int i = h; i < n; i ++) {
                T e = arr[i];
                int j;
                for (j = i; j >= h && e < arr[j-h]; j -= h) {
                    arr[j] = arr[j-h];
                }
                arr[j] = e;
            }

            h /= 3;
        }
    }
}

#endif //C___SHELLSORT_H
