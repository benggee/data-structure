//
// Created by HEADS on 2021/2/14.
//

#ifndef C___MAXHEAP_H
#define C___MAXHEAP_H

#include <iostream>
using namespace std;

template<typename T>
class maxheap {
public:
    maxheap(int n) {
        count = 0;
        capacity = n;
        data = new T[n + 1];
    }

    maxheap(T arr[], int n) {
        capacity = n;
        data = new int[n + 1];
        count = n;
        for (int i = 1; i < n; i ++) {
            data[i + 1] = arr[i];
        }

        for (int i = count/2; i >= 1; i --) {
            shiftDown(i);
        }
    }

    int size() {
        return count;
    }

    bool isEmpty() {
        return count == 0;
    }

    void insert(T item) {
        assert(capacity >= count + 1);
        data[++count] = item;
        shiftUp(count);
    }

    T popMax() {
        assert(count > 0);
        T ret = data[1];
        swap(data[1], data[count]);
        count--;
        shiftDown(1);
        return ret;
    }

    ~maxheap() {
        delete[] data;
    }

private:
    T *data;
    int count;
    int capacity;

    void shiftUp(int k) {
        while(k > 1 && data[k] > data[k/2]) {
            swap(data[k], data[k/2]);
            k /= 2;
        }
    }

    void shiftDown(int k) {
        while(2*k < count) {
            int j = k * 2;
            if (j + 1 <= count && data[j] < data[j + 1])
                j += 1;

            if (data[k] >= data[j])
                break;

            swap(data[k], data[j]);
            k = j;
        }
    }
};
#endif //C___MAXHEAP_H


