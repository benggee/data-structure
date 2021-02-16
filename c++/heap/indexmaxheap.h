//
// Created by HEADS on 2021/2/15.
//

#ifndef C___INDEXMAXHEAP_H
#define C___INDEXMAXHEAP_H
#include <iostream>
using namespace std;

template<typename T>
class IndexMaxHeap {
public:
    IndexMaxHeap(int capacity) {
        count = 0;
        data = new T[capacity + 1];
        indexes = new int[capacity + 1]; // 索引堆的索引
        reverse = new int[capacity + 1];
        for (int i = 0; i <= capacity; i ++)
            reverse[i] = 0; // 真实索引是从1开始的，这里将索引置为0表示没有任何索引
        this->capacity = capacity;
    }

    int size() {
        return count;
    }

    bool isEmpty() {
        return count == 0;
    }

    // 判断索引是否存在
    bool contain(int i) {
        assert(i >= 0 && i < capacity);
        return reverse[i + 1] != 0;
    }

    void insert(int idx, T item) {
        assert(capacity > count);

        assert(!contain(idx));

        idx += 1;
        data[idx] = item;
        indexes[count + 1] = idx;
        reverse[idx] = count + 1;

        count++;
        shiftUp(count);
    }

    T popMax() {
        assert(count > 0);
        T ret = data[indexes[1]];
        swap(indexes[1], indexes[count]);
        reverse[indexes[count]] = 0;
        count--;

        if (count) {
            reverse[indexes[1]] = 1;
            shiftDown(1);
        }

        return ret;
    }

    int popMaxIndex() {
        assert(count > 0);

        int ret = indexes[1] - 1; // 对于用户而言索引是从0开始的
        swap(indexes[1], indexes[count]);
        count--;

        if (count) {
            reverse[indexes[1]] = 1;
            shiftDown(1);
        }
        return ret;
    }

    // 获取堆中最大的堆顶元素
    int getMax() {
        assert(count > 0);
        return data[indexes[1]];
    }

    // 获取堆中最大元素的索引
    int getMaxIndex() {
        assert(count > 0);
        return indexes[1] - 1;
    }

    // 获取堆中索引为i的元素
    int getItem(int i) {
        assert(contain(i));
        return data[i];
    }

    // 修改指针索引的值
    int change(int i, T item) {
        assert(contain(i));

        i += 1;
        data[i] = item;

        // 找到indexes[j] == i, j表示data[i]在堆中的位置
        // 之后shifUp(j), 再shiftDown(j)
//        for (int j = 1; j <= count; j ++) {
//            if (indexes[j] == i) {
//                shiftUp(j);
//                shiftDown(j);
//                return;
//            }
//        }
        shiftUp(reverse[i]);
        shiftDown(reverse[i]);
    }

    ~IndexMaxHeap() {
        delete[] data;
        delete[] indexes;
        delete[] reverse;
    }

    // 测试索引堆中的索引数组index和反向数组reverse
    // 注意:这个测试在向堆中插入元素以后, 不进行extract操作有效
    bool testIndexesAndReverseIndexes(){

        int *copyIndexes = new int[count+1];
        int *copyReverseIndexes = new int[count+1];

        for( int i = 0 ; i <= count ; i ++ ){
            copyIndexes[i] = indexes[i];
            copyReverseIndexes[i] = reverse[i];
        }

        copyIndexes[0] = copyReverseIndexes[0] = 0;
        std::sort(copyIndexes, copyIndexes + count + 1);
        std::sort(copyReverseIndexes, copyReverseIndexes + count + 1);

        // 在对索引堆中的索引和反向索引进行排序后,
        // 两个数组都应该正好是1...count这count个索引
        bool res = true;
        for( int i = 1 ; i <= count ; i ++ )
            if( copyIndexes[i-1] + 1 != copyIndexes[i] ||
                copyReverseIndexes[i-1] + 1 != copyReverseIndexes[i] ){
                res = false;
                break;
            }

        delete[] copyIndexes;
        delete[] copyReverseIndexes;

        if( !res ){
            cout<<"Error!"<<endl;
            return false;
        }

        for( int i = 1 ; i <= count ; i ++ )
            if( reverse[ indexes[i] ] != i ){
                cout<<"Error 2"<<endl;
                return false;
            }

        return true;
    }

private:
    T *data;      // 最大索引堆的数据
    int *indexes; // 最大索引堆中的索引，indexes[x]=i 表示索引i在x的位置
    int *reverse; // 最大索引堆的反向索引，reverse[i]=x表示索引i在x的位置

    int count;    // 堆当前有多少元素
    int capacity; // 堆的容量

    // 索引堆中，数据之间的比较根据data的大小进行比较，但实际操作的是索引
    void shiftUp(int k) {
        while(k > 1 && data[indexes[k]] > data[indexes[k/2]]) {
            swap(indexes[k], indexes[k/2]);
            reverse[indexes[k]] = k;
            reverse[indexes[k/2]] = k/2;
            k /= 2;
        }
    }

    // 索引堆中，数据之间的比较根据data的大小进行，但实际操作的是索引
    void shiftDown(int k) {
        while(2*k < count) {
            int j = k * 2;
            if (j + 1 <= count && data[indexes[j]] < data[indexes[j + 1]])
                j += 1;

            if (data[indexes[k]] >= data[indexes[j]])
                break;

            swap(indexes[k], indexes[j]);
            reverse[indexes[k]] = k;
            reverse[indexes[j]] = j;
            k = j;
        }
    }
};
#endif //C___INDEXMAXHEAP_H
