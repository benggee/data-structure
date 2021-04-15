public class UnionFind implements UF {
    public int[] id;

    public UnionFind(int size) {
        id = new int[size];
        for (int i=0; i<size; i++) {
            id[i] = i;
        }
    }

    @Override 
    public int size() {
        return id.length;
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
        
        for (int i=0; i<id.length; i++) 
            if (id[i] == pIN) 
                id[i] = qIN;
    }

    private int find(int i) {
        if (i < 0 || i >= id.length) 
            throw new IllegalArgumentException("Index out of range.");
        return id[i];
    }

}