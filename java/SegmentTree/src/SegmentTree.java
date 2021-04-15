public class SegmentTree<E>{

    private E[] tree;
    private E[] data;
    private Merger<E> merger;
    public SegmentTree(E[] arr, Merger<E> merger) {
        this.merger = merger;

        data = (E[])new Object[arr.length];
        tree = (E[])new Object[4 * arr.length];
        for (int i=0; i<arr.length; i++) {
            data[i] = arr[i];
        }
        createSegmentTree(0, 0, arr.length-1);
    }

    public void createSegmentTree(int index, int l, int r) {
        if (l == r) { 
            tree[index] = data[l];
            return;
        }

        int leftTreeIndex = leftChild(index);
        int rightTreeIndex = rightChild(index);

        int mid = l + (r - l)/2;
        createSegmentTree(leftTreeIndex, l, mid);
        createSegmentTree(rightTreeIndex, mid+1, r);

        tree[index] = merger.merge(tree[leftTreeIndex], tree[rightTreeIndex]);
    }

    public E query(int qL, int qR) {
        if (qL < 0 || qL >= data.length || qR < 0 || qR >= data.length || qL > qR) 
            throw new IllegalArgumentException("Range error.");
        return query(0, 0, data.length - 1, qL, qR);
    }

    private E query(int index, int l, int r, int qL, int qR) {
        if (l == qL && r == qR) 
            return tree[index];
        
        int mid = l + (r - l) / 2;
        int leftIndex = leftChild(index);
        int rightIndex = rightChild(index);

        // 如果不在左区间
        if (qL > mid ) 
            return query(rightIndex, mid + 1, r, qL, qR);
        // 如果不在右区间
        if (qR <= mid) {
            System.out.println("aaaa");
            System.out.println(leftIndex + "=" + l + "=" + mid + "=" + qL + "=" + qR);
            return query(leftIndex, l, mid, qL, qR);
        }
        
        // 如果一部分在左
        E leftE = query(leftIndex, l, mid, qL, mid);
        // 如果一部分在右
        E rightE = query(rightIndex, mid + 1, r, mid + 1, qR);
        return merger.merge(leftE, rightE);
    }

    // 修改节点
    public void set(int index, E e) {
        if (index < 0 || index > data.length) 
            throw new IllegalArgumentException("Index out of range");
        set(0, 0, tree.length-1, index, e);
    }

    private void set(int index, int l, int r, int index2, E e) {
        if (l == r) {
            tree[index] = e;
            return;
        }

        int mid = l + (r - l) / 2;
        int leftIndex = leftChild(index);
        int rightIndex = rightChild(index);

        if (index2 < mid) 
            set(rightIndex, mid + 1, r, index2, e);
        else 
            set(leftIndex, l, mid, index2, e);
        
        // 这一步比较关键，表示将所有父节点重新溶合
        tree[index] = merger.merge(tree[leftIndex], tree[rightIndex]);
    }

    public int getSize() {
        return data.length;
    }

    public E get(int index) {
        if (index < 0 || index > data.length) 
            throw new IllegalArgumentException("Index out of range.");
        return data[index];
    }

    // 获取左子节点的索引值
    public int leftChild(int index) {
        return 2 * index + 1;
    }

    // 获取右子节点的索引值
    public int rightChild(int index) {
        return 2 * index + 2;
    }

    @Override
    public String toString() {
        StringBuilder str = new StringBuilder();
        str.append("[");
        for (int i=0; i<tree.length; i++) {
            if (tree[i] != null) 
                str.append(tree[i]);
            else  
                str.append("null");
            if (tree.length - 1 != i) 
                str.append(",");
        }
        str.append("]");
        return str.toString();
    }
}