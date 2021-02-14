//
// Created by HEADS on 2021/2/13.
//

#ifndef DATA_STRUCTURE_UNION1_H
#define DATA_STRUCTURE_UNION1_H

#include <iostream>
#include <cassert>

#endif //DATA_STRUCTURE_UNION1_H

namespace UF1 {
    class UnionFind {
    public:
        int *id;
        int count;

        UnionFind(int n) {
            count = n;
            id = new int[n];
            for (int i = 0; i < n; i++)
                id[i] = i;
        }

        ~UnionFind() {
            delete[] id;
        }

        int find(int p) {
            assert(p >= 0 && p < count);
            return id[p];
        }

        bool isConnect(int p, int q) {
            return id[p] == id[q];
        }

        void unionElm(int p, int q) {
            int pId = find(p);
            int qId = find(q);
            if (pId == qId)
                return;

            for (int i = 0; i < count; i++)
                if (id[i] == pId)
                    id[i] = qId;
        }
    };
}