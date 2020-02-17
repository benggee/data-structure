#include <iostream>
#include "Student.h"
#include "TestHelper.h"

using namespace std;

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

int main() {
    int n = 10000;
    int *arr = TestHelper::generateRandArray(n , 0, n);
    selectSort(arr, 13);
    // TestHelper::printArray(arr, n);
    TestHelper::testSort("Select Sort", selectSort, arr, n);

    delete[] arr;



    // 对结构体的排序
    // Student d[4] = {{"D", 94}, {"C", 50}, {"B", 60}, {"A", 95}};
    // selectSort(d, 4);
    // TestHelper::printArray(d, 4);

    // return 0;
}