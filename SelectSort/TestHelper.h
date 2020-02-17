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
}
#endif