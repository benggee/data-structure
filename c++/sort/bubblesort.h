//
// Created by HEADS on 2021/2/14.
//

#ifndef C___BUBBLESORT_H
#define C___BUBBLESORT_H

namespace BubbleSort {
    template<typename T>
    void bubbleSort(T arr[], int n) {
        bool swapped;
        do {
            swapped = false;
            for (int i = 1; i < n; i ++) {
                if (arr[i-1] > arr[i]) {
                    swap(arr[i-1], arr[i]);
                    swapped = true;
                }
            }

            // 优化，每一次bubble Sort都将最大的元素放在了最后的位置
            // 所以下一次排序，最后的元素可以不考虑
            n--;
        }while(swapped);
    }

    template<typename T>
    void bubbleSort2(T arr[], int n) {
        int newn; // 使用newn进行优化
        do {
            newn = 0;
            for (int i = 1; i < n; i ++) {
                if (arr[i-1] > arr[i]) {
                    swap(arr[i - 1], arr[i]);
                    newn = i; // 记录最后一次交换的位置，在此之后的元素在下一轮扫描中均不考虑
                }
            }
            n = newn;
        } while (newn > 0);
    }
}

#endif //C___BUBBLESORT_H
