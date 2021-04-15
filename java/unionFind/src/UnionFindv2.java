public class UnionFindv2 implements UF {
    public int[] parent;

    public UnionFindv2(int size) {
        parent = new int[size];
        for (int i=0; i<size; i++) {
            parent[i] = i;
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
        
        parent[qIN] = parent[pIN];
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