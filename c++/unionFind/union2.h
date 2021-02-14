//
// Created by HEADS on 2021/2/14.
//

#ifndef C___UNION2_H
#define C___UNION2_H

#include <assert.h>

namespace UF2 {
    class UnionFind {
    public:
        UnionFind(int n) {
            count = n;
            parent = new int[n];
            for (int i = 0; i < n; i ++)
                parent[i] = i;
        }

        int find(int p) {
            assert(p >= 0 && p < count);
            while(p != parent[p])
                p = parent[p];

            return p;
        }

        bool isConnect(int p, int q) {
            return find(p) == find(q);
        }

        void unionElm(int p, int q) {
            int pRoot = find(p);
            int qRoot = find(q);
            if (pRoot == qRoot)
                return;

            parent[pRoot] = qRoot;
        }

        ~UnionFind() {
            delete[] parent;
        }
    private:
        int *parent;
        int count;
    };
}

#endif //C___UNION2_H
