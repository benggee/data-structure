public class MaxHeap<E extends Comparable<E>> {
    Array<E> data;

    public MaxHeap(int capacity) {
        data = new Array<>(capacity);
    }

    public MaxHeap() {
        data = new Array<>();
    }

    public MaxHeap(E[] e) {
        data = new Array<>(e);
        // 从最后一个父节点开始下沉
        for (int i=parent(e.length - 1); i>=0; i--) 
            siftDown(i);
    }

    public int size() {
        return data.size();
    }

    public boolean isEmpty() {
        return data.isEmpty();
    }

    // 添加元素
    public void add(E e) {
        data.addLast(e);
        siftUp(data.size()-1);
    } 

    // 取出当前最大元素
    public E max() {
        if (data.size() == 0) 
            throw new IllegalArgumentException("Arr error.");
        return data.get(0);
    }

    // 取出一个元素
    public E extractMax() {
        E ret = max();
        data.swap(0, data.size() - 1);
        data.removeLast();
        siftDown(0);
        return ret;
    }

    // 替换堆顶元素并返回原来的值
    public E replace(E e) {
        E headMax = max();
        data.set(0, e);
        siftDown(0);
        return headMax;
    }

    // 交换子节点和父节点
    private void siftUp(int i) {
        while(i > 0 && data.get(parent(i)).compareTo(data.get(i)) < 0) {
            data.swap(parent(i), i);
            i = parent(i);
        }
    }

    // 元素下沉
    private void siftDown(int i) {
        while (left(i) < data.size()) {
            int j = left(i);
            if (j + 1 < data.size() && data.get(j+1).compareTo(data.get(j)) > 0) {
                j = right(i);
            }

            if (data.get(i).compareTo(data.get(j)) >= 0) {
                break;
            }

            data.swap(i, j);
            i = j;
        }
    }

    private int parent(int index) {
        if (index <= 0) 
            throw new IllegalArgumentException("Index out of range.");
        return (index-1) / 2;
    }

    private int left(int index) {
        return index * 2 + 1;
    }

    private int right(int index) {
        return index * 2 + 2;
    }
}