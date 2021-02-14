//
// Created by HEADS on 2021/2/14.
//

#ifndef C___UNION5_H
#define C___UNION5_H

#include <assert.h>

namespace UF5 {
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
            while(p != parent[p]) {
                parent[p] = parent[parent[p]];
                p = parent[p];
            }

            // 递归的方式
//            if (p != parent[p])
//                parent[p] = find(parent[p]);
//            return parent[p];

            return p;

        }

        bool isConnect(int p, int q) {
            return find(p) == find(q);
        }

        void unionElm(int p, int q) {
            int pRoot = parent[p];
            int qRoot = parent[q];
            if (pRoot == qRoot)
                return;
            if (rank[pRoot] < rank[qRoot]) {
                parent[pRoot] = parent[qRoot];
            } else if (rank[qRoot] < rank[pRoot]) {
                parent[qRoot] = parent[pRoot];
            } else {
                parent[qRoot] = parent[pRoot];
                rank[pRoot] += rank[qRoot];
            }
        }

        ~UnionFind() {
            delete[] parent;
            delete[] rank;
        }
    private:
        int *parent;
        int *rank;
        int count;
    };
}

#endif //C___UNION5_H
