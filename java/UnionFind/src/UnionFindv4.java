public class UnionFindv4 implements UF {
    public int[] parent;
    public int[] rank;

    public UnionFindv4(int size) {
        parent = new int[size];
        rank = new int[size];
        for (int i=0; i<size; i++) {
            parent[i] = i;
            rank[i] = 1;
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

        // rank优化， 尽量保持深度一致
        if (rank[pIN] > rank[qIN]) {
            parent[qIN] = parent[pIN];
        } else if (rank[pIN] < rank[qIN]){
            parent[pIN] = parent[qIN];
        } else {
            parent[pIN] = parent[qIN];
            rank[qIN]++;
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