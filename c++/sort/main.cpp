//
// Created by HEADS on 2021/2/14.
//

#include "maxheap.h"
#include <iostream>
#include "testhelper.h"
#include "heapsort.h"
#include "selectionsort.h"
#include "insertsort.h"
#include "bubblesort.h"
#include "shellsort.h"
#include "mergesort.h"
#include "quicksort.h"
#include "quicksort2ways.h"
#include "quicksort3ways.h"

using namespace std;



int main() {
    int n = 100000;

    cout << "Test for random arr, size=" << n << ", random range [0," << n << "]" << endl;
    int *arr = SortTestHelper::generateRandomArr(n, 0, n);
    int *arr1 = SortTestHelper::copyIntArray(arr, n);
    int *arr2 = SortTestHelper::copyIntArray(arr, n);
    int *arr3 = SortTestHelper::copyIntArray(arr, n);
    int *arr4 = SortTestHelper::copyIntArray(arr, n);
    int *arr5 = SortTestHelper::copyIntArray(arr, n);
    int *arr6 = SortTestHelper::copyIntArray(arr, n);
    int *arr7 = SortTestHelper::copyIntArray(arr, n);
    int *arr8 = SortTestHelper::copyIntArray(arr, n);
    int *arr9 = SortTestHelper::copyIntArray(arr, n);
    int *arr10 = SortTestHelper::copyIntArray(arr, n);
    int *arr11 = SortTestHelper::copyIntArray(arr, n);
    int *arr12 = SortTestHelper::copyIntArray(arr, n);
    int *arr13 = SortTestHelper::copyIntArray(arr, n);

    // O(NLogN)复杂度算法
//    SortTestHelper::testSort("Heap Sort 1", HeapSort::heapSort1, arr1, n);
//    SortTestHelper::testSort("Heap Sort 2", HeapSort::heapSort2, arr2, n);
//    SortTestHelper::testSort("Heap Sort 3", HeapSort::heapSort, arr3, n);
    SortTestHelper::testSort("Merge Sort", MergeSort::mergeSort, arr9, n);
    SortTestHelper::testSort("Merge Sort BU", MergeSort::mergeSortBU, arr10, n);
    SortTestHelper::testSort("Quick Sort", QuickSort::quickSort, arr11, n);
    SortTestHelper::testSort("Quick Sort 2 ways", QuickSort2Ways::quickSort, arr12, n);
    SortTestHelper::testSort("Quick Sort 3 ways", QuickSort3Ways::quickSort, arr13, n);

    // O(n^2)复杂度算法
//    SortTestHelper::testSort("Selection Sort", SelectionSort::selectionSort, arr4, n);
//    SortTestHelper::testSort("Insert Sort", InsertSort::insertSort, arr5, n);
//    SortTestHelper::testSort("Bubble Sort", BubbleSort::bubbleSort, arr6, n);
//    SortTestHelper::testSort("Bubble Sort2", BubbleSort::bubbleSort2, arr7, n);
//    SortTestHelper::testSort("Shell Sort", ShellSort::shellSort, arr8, n);


    delete[] arr;
    delete[] arr1;
    delete[] arr2;
    delete[] arr3;
    delete[] arr4;
    delete[] arr5;
    delete[] arr6;
    delete[] arr7;
    delete[] arr8;
    delete[] arr9;
    delete[] arr10;
    delete[] arr11;
    delete[] arr12;

    cout << "Test for order arr, size=" << n << ", order range [0," << n << "]" << endl;
    int *oarr = SortTestHelper::generateNearlyOrderArr(n, 0);
    int *oarr1 = SortTestHelper::copyIntArray(oarr, n);
    int *oarr2 = SortTestHelper::copyIntArray(oarr, n);
    int *oarr3 = SortTestHelper::copyIntArray(oarr, n);
    int *oarr4 = SortTestHelper::copyIntArray(oarr, n);

    SortTestHelper::testSort("Merge Sort", MergeSort::mergeSort,oarr1, n);
    SortTestHelper::testSort("Quick Sort", QuickSort::quickSort, oarr2, n);
    SortTestHelper::testSort("Quick Sort 2 Ways", QuickSort2Ways::quickSort, oarr3, n);
    SortTestHelper::testSort("Quick Sort 3 Ways", QuickSort3Ways::quickSort, oarr4, n);

    delete[] oarr1;
    delete[] oarr2;
    delete[] oarr3;
    delete[] oarr4;

    cout << "Test for repeat arr, size=" << n << ", order range [0," << n << "]" << endl;
    int *rarr = SortTestHelper::generateRandomArr(n, 0, 10);
    int *rarr1 = SortTestHelper::copyIntArray(rarr, n);
    int *rarr2 = SortTestHelper::copyIntArray(rarr, n);
    int *rarr3 = SortTestHelper::copyIntArray(rarr, n);
    int *rarr4 = SortTestHelper::copyIntArray(rarr, n);

    SortTestHelper::testSort("Merge Sort", MergeSort::mergeSort, rarr1, n);
    // SortTestHelper::testSort("Quick Sort", QuickSort::quickSort, rarr2, n);
    SortTestHelper::testSort("Quick Sort 2 ways", QuickSort2Ways::quickSort, rarr3, n);
    SortTestHelper::testSort("Quick Sort 3 ways", QuickSort3Ways::quickSort, rarr4, n);

    delete[] rarr1;
    delete[] rarr2;
    delete[] rarr3;
    delete[] rarr4;


    return 0;
}