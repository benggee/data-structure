public class UnionFindv3 implements UF {
    public int[] parent;
    public int[] sz;

    public UnionFindv3(int size) {
        parent = new int[size];
        sz = new int[size];
        for (int i=0; i<size; i++) {
            parent[i] = i;
            sz[i] = 1;
        }
    }

    @Override 
    public int size() {
        return parent.length;
    }

    @Override 
    public boolean unionContains(int p, int q) {
        return find(p) == find(q);
    }

    @Override 
    public void union (int p, int q) {
        int pIN = find(p);
        int qIN = find(q);

        if (pIN == qIN) 
            return;

        // size优化， 尽量处理成一个平衡树，避免退化成链表
        if (sz[pIN] > sz[qIN]) {
            parent[qIN] = parent[pIN];
            sz[pIN] += sz[qIN];
        } else {
            parent[pIN] = parent[qIN];
            sz[qIN] += sz[pIN];
        }
    }

    // 时间复杂度为h, h表示树的高度
    private int find(int i) {
        if (i < 0 || i >= parent.length) 
            throw new IllegalArgumentException("Index out of range.");

        while(i != parent[i])
            i = parent[i];
        
        return i;
    }
}