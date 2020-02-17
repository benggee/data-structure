#ifndef SELECTSORT_TESTHELPER_H
#define SELECTSORT_TESTHELPER_H

#include <iostream>
#include <ctime>
#include <cassert>

namespace TestHelper {
    int *generateRandArray(int n, int l, int r) {
        
        assert(l <= r);

        int *tmpArr = new int[n];
        // 使用时间作为随机生成的种子
        srand(time(NULL));
        for (int i=0;i<n; i++) 
            tmpArr[i] = rand() % (r - l + 1) + l;

        return tmpArr;
    }

    // 生成一个近乎有序的数组
    // 首先生成一个含有[0...n-1]的完全有序数组, 之后随机交换swapTimes对数据
    // swapTimes定义了数组的无序程度:
    // swapTimes == 0 时, 数组完全有序
    // swapTimes 越大, 数组越趋向于无序
    int *generateNearlyOrderedArray(int n, int swapTimes){

        int *arr = new int[n];
        for(int i = 0 ; i < n ; i ++ )
            arr[i] = i;

        srand(time(NULL));
        for( int i = 0 ; i < swapTimes ; i ++ ){
            int posx = rand()%n;
            int posy = rand()%n;
            std::swap( arr[posx] , arr[posy] );
        }

        return arr;
    }

    template<typename T>
    void printArray(T arr[], int n) {
        for (int i=0; i<n; i++) 
            std::cout << arr[i] << " ";
        std::cout << std::endl;
    }

    template<typename T> 
    bool isSort(T arr[], int n) {
        for (int i=0; i<n-1; i++) {
            if (arr[i] > arr[i+1]) 
                return false;
        }
        return true;
    }

    // 排序测试
    template<typename T> 
    void testSort(const std::string &sortName, void(*sort)(T arr[], int n), T arr[], int n) {
        clock_t startTime = clock();
        sort(arr, n);
        clock_t endTime = clock();

        isSort(arr, n);

        std::cout<<sortName<<" TIMES: "<< double(endTime-startTime) / CLOCKS_PER_SEC << "s" << std::endl;
        return;
    }

    int* copyIntArray(int a[], int n) {
        int *arr = new int[n];
        std::copy(a, a+n, arr);
        return arr;
    }
}
#endif