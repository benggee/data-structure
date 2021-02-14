//
// Created by HEADS on 2021/2/14.
// 基于rank（层）的优化
//

#ifndef C___UNION4_H
#define C___UNION4_H

#include <assert.h>

namespace UF4 {
    class UnionFind {
    public:
        UnionFind(int n) {
            count = n;
            parent = new int[n];
            rank = new int[n];
            for (int i = 0; i < n; i ++) {
                parent[i] = i;
                rank[i] = 1;
            }
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

            if (rank[pRoot] < rank[qRoot]) {
                parent[pRoot] = parent[qRoot];
            } else if (rank[qRoot] < rank[pRoot]) {
                parent[qRoot] = parent[pRoot];
            } else {
                parent[pRoot] = parent[qRoot];
                rank[qRoot] += rank[pRoot];
            }
        }
    private:
        int *parent;
        int *rank;
        int count;
    };
}

#endif //C___UNION4_H
