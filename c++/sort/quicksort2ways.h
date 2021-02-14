//
// Created by HEADS on 2021/2/14.
//

#ifndef C___QUICKSORT2WAYS_H
#define C___QUICKSORT2WAYS_H

namespace QuickSort2Ways {
    // 对arr[l....r]部分进行partition操作
    // 返回p,使得arr[l...p-1] < arr[p]; arr[p+1...r] > arr[p]
    template<typename T>
    int _partition(T arr[], int l, int r) {
        // 随机在arr[l...r]的范围中，选择一个数值作为标定点
        swap(arr[l], arr[rand()%(r-l+1) + l]);

        T v = arr[l];
        int i = l+1, j=r;
        while(true) {
            // 注意这里的边界，arr[i]<v, 不能是arr[i]<=v
            while(i <= r && arr[i] < v)
                i++;
            // 注意这里的边界，arr[j] > v, 不能是arr[j] >= v
            while(j >= l+1 && arr[j] > v)
                j--;

            if (i > j)
                break;
            swap(arr[i], arr[j]);
            i++;
            j--;
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
