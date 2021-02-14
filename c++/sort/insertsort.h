//
// Created by HEADS on 2021/2/14.
//

#ifndef C___INSERTSORT_H
#define C___INSERTSORT_H

namespace InsertSort{
    template<typename T>
    void insertSort(T arr[], int n) {
        for (int i = 0; i < n; i ++) {
            T e = arr[i];
            int j;
            for (j = i; j > 0 && e < arr[j - 1]; j--)
                arr[j] = arr[j - 1];
            arr[j] = e;
        }
    }
}

#endif //C___INSERTSORT_H
