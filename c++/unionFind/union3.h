//
// Created by HEADS on 2021/2/14.
// 基于size的优化
//

#ifndef C___UNION3_H
#define C___UNION3_H

#include <assert.h>

namespace UF3 {
    class UnionFind {
    public:
        UnionFind(int n) {
            count = n;
            parent = new int[n];
            zs = new int[n];
            for (int i = 0; i < n; i ++) {
                parent[i] = i;
                zs[i] = 1;
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

            if (zs[pRoot] < zs[qRoot]) {
                parent[pRoot] = parent[qRoot];
                zs[qRoot] += zs[pRoot];
            } else {
                parent[qRoot] = parent[pRoot];
                zs[pRoot] += zs[qRoot];
            }
        }

        ~UnionFind() {
            delete[] parent;
            delete[] zs;
        }
    private:
        int* parent;
        int* zs;     // 存每个节点对应的size
        int count;
    };
}

#endif //C___UNION3_H
