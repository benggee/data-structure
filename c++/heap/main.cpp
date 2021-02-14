//
// Created by HEADS on 2021/2/14.
//

#include "maxheap.h"
#include <iostream>
using namespace std;


int main() {
    maxheap<int> h = maxheap<int>(100);

    srand(time(NULL));

    for (int i = 0; i < 20; i ++) {
        h.insert(rand() % 100);
    }

    h.testPrint();


    while(!h.isEmpty()) {
        cout << h.popMax() << " ";
    }
    cout << endl;

    return 0;
}