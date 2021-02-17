//
// Created by HEADS on 2021/2/17.
//

#ifndef DATA_STRUCTURE_TRIE_H
#define DATA_STRUCTURE_TRIE_H

#include <iostream>
#include <map>
#include <string>

using namespace std;

struct Node {
    bool isWord;
    map<char, Node*> next;
    Node(): isWord(false) {}
};

class Trie {
private:
    Node* root;
    int size;

public:
    Trie() {
        root = new Node();
        size = 0;
    }
    ~Trie(){
//        Node *cur = root;
//        while(cur != NULL) {
//            cur = cur
//        }
    }

    int getSize() {
        return size;
    }

    // 向trie树中添加一个新的单词
    void add(string &word) {
        Node* cur = root;
        for (int i = 0; i < word.size(); i ++) {
            char c = word[i];
            if (cur->next[c] == NULL)
                cur->next[c] = new Node();
            cur = cur->next[c];
        }

        if (!cur->isWord) {
            cur->isWord = true;
            size ++;
        }
    }

    // 查询单词word是否在trie中
    bool contains(string word) {
        Node* cur = root;
        for (int i = 0; i < word.size(); i ++) {
            char c = word[i];
            if (cur->next[c] == NULL)
                return false;
            cur = cur->next[c];
        }
        return cur->isWord;
    }

    bool isPrefix(string word) {
        Node* cur = root;
        for (int i = 0; i < word.size(); i ++) {
            char c = word[i];
            if (cur->next[c] == NULL)
                return false;
            cur = cur->next[c];
        }
        return true;
    }
};


#endif //DATA_STRUCTURE_TRIE_H
