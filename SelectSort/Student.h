#ifndef SELECTSORT_STUDENT_H
#define SELECTSORT_STUDENT_H

#include <iostream>
#include <string>
using namespace std;

struct Student {
    string name;
    int score;

    // 重载小于符号
    bool operator<(const Student& otherStudent) {
        return score != otherStudent.score ?
                score < otherStudent.score :
                name > otherStudent.name;
    }

    // 重载<<符号
    friend ostream& operator<<(ostream &os, const Student &student) {
        os<<"Studnet"<<student.name<<" "<<student.score<<endl;
        return os;
    }
};

#endif 