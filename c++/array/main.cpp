//
// Created by Fajun on 2023/4/19.
//

#include <iostream>
#include <assert.h>
#include "ArrayList.hpp"

using namespace std;

int main() {
    ArrayList<int> al(10);

    assert(al.Insert(1, 12)==true);
    assert(al.Insert(2, 13)==true);
    assert(al.Insert(3, 33)==true);
    // assert(al.Insert(11, 2)==false);

    al.Display();

    assert(al.Delete(2) == true);

    al.Display();

    int newval = 0;
    al.GetElem(1, newval);

    cout << "newval:" << newval << endl;
}