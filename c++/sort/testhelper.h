//
// Created by HEADS on 2021/2/14.
//
#include "maxheap.h"

namespace SortTestHelper {
    // 生成一个随机数组
    int *generateRandomArr(int n, int range_l, int range_r) {
        int *arr = new int[n];
        srand(time(NULL));

        for (int i = 0; i < n; i ++) {
            arr[i] = rand() % (range_r - range_l + 1) + range_l;
        }
        return arr;
    }

    // 生成一个近乎有序数组
    int *generateNearlyOrderArr(int n, int swapTimes) {
        int *arr = new int[n];
        for (int i = 0; i < n; i ++) {
            arr[i] = i;
        }

        srand(time(NULL));
        for( int i = 0; i < swapTimes; i ++) {
            int posx = rand() % n;
            int posy = rand() % n;
            swap(arr[posx], arr[posy]);
        }

        return arr;
    }

    // 拷贝整形数组a中的所有元素到一个新的数组，并返回新数组
    int *copyIntArray(int a[], int n) {
        int *arr = new int[n];
        copy(a, a + n, arr);
        return arr;
    }

    // 打印arr数组所有内容
    template<typename T>
    void printArr(T arr[], int n) {
        for (int i = 0; i < n; i ++)
            cout << arr[i] << ",";
        cout << endl;
        return;
    }

    // 判断数组是否有序
    template<typename T>
    bool isSorted(T arr[], int n) {
        for (int i = 0; i < n - 1; i ++)
            if (arr[i] > arr[i + 1])
                return false;
        return true;
    }

    // 测试sort排序算法排序arr数组所得到结果的正确性和算法运行时间
    // 将算法的运行时间打印在控制台上
    template<typename T>
    void testSort(const string &sortName, void (*sort)(T[], int), T arr[], int n) {

        clock_t startTime = clock();
        sort(arr, n);
        clock_t endTime = clock();
        cout << sortName << " : " << double(endTime - startTime) / CLOCKS_PER_SEC << " s"<<endl;

        assert(isSorted(arr, n));

        return;
    }

    // 测试sort排序算法排序arr数组所得到结果的正确性和算法运行时间
    // 将算法的运行时间以double类型返回, 单位为秒(s)
    template<typename T>
    double testSort(void (*sort)(T[], int), T arr[], int n) {

        clock_t startTime = clock();
        sort(arr, n);
        clock_t endTime = clock();

        assert(isSorted(arr, n));

        return double(endTime - startTime) / CLOCKS_PER_SEC;
    }
}