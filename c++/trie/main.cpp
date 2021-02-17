//
// Created by HEADS on 2021/2/17.
//
#include <iostream>
#include <string>
#include <vector>
#include "trie.h"

using namespace std;

vector<string> generateWords(int n);

int main() {

    int n = 1000000;
    vector<string> words = generateWords(n);

    Trie e = Trie();

    clock_t start = clock();
    for (int i = 0; i < words.size(); i ++) {
        e.add(words[i]);
    }

    for (int i = 0; i < words.size(); i ++) {
        e.contains(words[i]);
    }
    clock_t end = clock();

    cout << "Trie add and contais time:" << double(end-start)/CLOCKS_PER_SEC << endl;
}

vector<string> generateWords(int n) {
    string c = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";

    srand(time(NULL));

    vector<string> vec;
    for (int i = 0; i < n; i ++) {
        int len = rand() % (c.size()/2);
        char w[len];
        for (int j = 0; j < len - 1; j ++) {
            int idx = rand() % (c.size() - 1);
            w[j] = c[idx];
        }
        if (sizeof(w)/sizeof(char) > 1) {
            vec.push_back(w);
        }
    }

    return vec;
}

