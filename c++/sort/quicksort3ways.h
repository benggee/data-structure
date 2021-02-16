//
// Created by HEADS on 2021/2/15.
//

#ifndef C___QUICKSORT3WAYS_H
#define C___QUICKSORT3WAYS_H
namespace QuickSort3Ways {
    template<typename T>
    void _quickSort(T arr[], int l, int r) {
        if (l >= r)
            return;

        swap(arr[l], arr[rand()%(r-l+1) + l]);
        T v = arr[l];
        int lt = l;     // arr[l+1....lt] < v
        int gt = r + 1; // arr[gt...r] > v
        int i = l + 1;  // arr[lt+1....i] == v

        while(i < gt) {
            if (arr[i] < v) {
                swap(arr[i], arr[lt + 1]);
                i++;
                lt++;
            } else if (arr[i] > v) {
                swap(arr[i], arr[gt - 1]);
                gt--;
            } else {
                i++;
            }
        }
        swap(arr[l], arr[lt]);

        _quickSort(arr, l, lt-1);
        _quickSort(arr, gt, r);
    }

    template<typename T>
    void quickSort(T arr[], int n) {
        srand(time( NULL));
        _quickSort(arr, 0, n - 1);
    }
}
#endif //C___QUICKSORT3WAYS_H
