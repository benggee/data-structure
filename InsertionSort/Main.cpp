#include <iostream>
#include "TestHelper.h"

using namespace std;

// 选择排序
template<typename T>
void selectSort(T arr[], int n) {
    for (int i=0; i<n; i++) {
        int minIndex = i;
        for (int j = i + 1; j<n; j++) {
            if (arr[j] < arr[minIndex]) {
                minIndex = j;
            }
        }

        swap(arr[i], arr[minIndex]);
    }
}

// 插入排序
template<typename T>
void insertionSort(T arr[], int n) {
    // 写法1
    // for (int i=0; i<n; i++) {
    //     for (int j=i; j>0; j--) {
    //         if (arr[j] < arr[j-1]) {
    //             swap(arr[j], arr[j-1]);
    //         } else {
    //             break;
    //         }
    //     }
    // }

    // 写法2
    // for (int i=0; i<n; i++) {
    //     for (int j=i; j>0 && arr[j]<arr[j-1]; j--) {
    //         swap(arr[j], arr[j-1]);
    //     }
    // }

    // 写法3 
    for (int i=0; i<n; i++) {
        T e = arr[i];
        int j;
        for (j=i; j>0 && arr[j-1] > e; j--) {
            arr[j] = arr[j-1];
        }
        arr[j] = e;
    }
}

int main() {
    int n = 10000;
    //int *arr = TestHelper::generateRandArray(n , 0, n);
    int *arr = TestHelper::generateNearlyOrderedArray(n, 100);
    int *arr2 = TestHelper::copyIntArray(arr, n);

    TestHelper::testSort("Select Sort", selectSort, arr, n);

    cout << endl;

    TestHelper::testSort("Insertion Sort", insertionSort, arr2, n);

    delete[] arr;
    delete[] arr2;


    // 对结构体的排序
    // Student d[4] = {{"D", 94}, {"C", 50}, {"B", 60}, {"A", 95}};
    // selectSort(d, 4);
    // TestHelper::printArray(d, 4);

    // return 0;
}