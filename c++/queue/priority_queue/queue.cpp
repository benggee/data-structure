//
// Created by HEADS on 2021/2/7.
//
#include <iostream>
#include <queue>
using namespace std;

// 按个位数排序
bool mySort(int x, int y) {
    return x % 10 > y % 10;
}

int main() {
    srand(time(NULL)); // 随机种子

    // 默认是大顶堆
    priority_queue<int> pq;
    // 随机生成10个100以内的数放入优先队列
    for (int i = 0; i < 10; i ++) {
        int num = rand() % 100;
        pq.push(num);
        cout << "insert:" << num << " in priority queue." << endl;
    }

    // 遍历
    while(!pq.empty()) {
        cout << pq.top() << " ";
        pq.pop();
    }
    cout << endl << endl;

    // 使用小顶堆
    // vector<int> 表示实现的底层数据结构使用vector
    // greater<int> 表示我们要实现一个小顶堆
    priority_queue<int, vector<int>, greater<int>> pq2;
    for (int i = 0; i < 10; i ++) {
        int num = rand() % 100;
        pq2.push(num);
        cout << "insert:" << num << " inpriority queue." << endl;
    }
    while(!pq2.empty()) {
        cout << pq2.top() << " ";
        pq2.pop();
    }
    cout << endl << endl;

    // 自定义排序规则
    priority_queue<int, vector<int>, function<bool(int, int)>> pq3(mySort);
    for (int i = 0; i < 10; i ++) {
        int num = rand() % 100;
        pq3.push(num);
        cout << "insert:" << num << " in priority queue." << endl;
    }
    while(!pq3.empty()) {
        cout << pq3.top() << " ";
        pq3.pop();
    }
    cout << endl << endl;
}