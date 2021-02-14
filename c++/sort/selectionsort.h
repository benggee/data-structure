//
// Created by HEADS on 2021/2/14.
//

#ifndef C___SELECTIONSORT_H
#define C___SELECTIONSORT_H

namespace SelectionSort {
    template<typename T>
    void selectionSort(T arr[], int n) {
        for (int i = 0; i < n; i ++) {
            // 寻找[i,n)区间最小值
            int min = i;
            for (int j = i + 1;j < n; j ++) {
                if (arr[j] < arr[min]) {
                    min = j;
                }
            }
            swap(arr[i], arr[min]);
        }
    }
}

#endif //C___SELECTIONSORT_H
