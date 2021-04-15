public class Array<E>{
    private E[] arr;
    private int size;

    public Array(int capacity) {
        arr = (E[]) new Object[capacity];
        size = 0;
    }

    public Array() {
        this(10);
    }

    public Array(E[] e) {
        arr = (E[]) new Object[e.length];
        for (int i=0; i<e.length; i++) 
            arr[i] = e[i];
        size = e.length;
    }

    // 获取数组容量
    public int capacity() {
        return arr.length;
    }

    // 获取元素数量
    public int size() {
        return size;
    }

    // 是否为空
    public boolean isEmpty() {
        return size == 0;
    }

    // 插入元素
    public void add(int index, E e) {
        if (index < 0 || index > size) 
            throw new IllegalArgumentException("Add fiald. index must be required index>0  index<=size");

        if (size == arr.length) 
            reSize(2 * arr.length);
        
        for (int i=size - 1; i >= index; i--)
            arr[i + 1] = arr[i];

        arr[index] = e;
        size++;
    }

    // 在数组开头添加一个元素
    public void addFirst(E e) {
        add(0, e);
    }

    // 在数组结尾添加一个元素
    public void addLast(E e) {
        add(size, e);
    }

    // 获取指定索引元素
    public E get(int index) {
        if (index < 0 || index > size) 
            throw new IllegalArgumentException("The index out of range size");
        return arr[index];
    }

    // 设置某个索引的值
    public void set(int index, E e) {
        if (index < 0 || index > size) 
            throw new IllegalArgumentException("The index out of range size");
        arr[index] = e;
    }

    // 交换元素
    public void swap(int i, int j) {
        if (i < 0 || i > size || j < 0 || j > size)
            throw new IllegalArgumentException("The i or j out of range.");
        E t = arr[i];
        arr[i] = arr[j];
        arr[j] = t;
    }

    // 查找元素所在索引位置
    public int find(E e) {
        for (int i = 0; i < size; i++) {
            if (e.equals(arr[i])) 
                return i;
        }
        return -1;
    }

    // 判断元素是否在数组内
    public boolean contains(E e) {
        for (int i = 0; i < size; i++) {
            if (e.equals(arr[i])
                return true;
        }

        return false;
    }

    // 删除元素
    public E remove(int index) {
        if (index < 0 || index > size) 
            throw new IllegalArgumentException("Index out of range.");
        E ret = arr[index];
        for (int i=index+1; i<size; i++) {
            arr[i-1] = arr[i];
        }
        size--;
        arr[size] = null;

        if (size == arr.length/2) {
            reSize(arr.length/2);
        }
        return ret;
    }

    // 删除数组最后一个元素
    public E removeLast() {
        return remove(size-1);
    }

    // 删除数组第一个元素
    public E removeFirst() {
        return remove(0);
    }

    // 删除某个元素
    public void removeE(E e) {
        int ret = find(e);
        if (ret!=-1) 
            remove(ret);
    }

    @Override
    public String toString() {
        StringBuilder str = new StringBuilder();
        str.append("[");
        for (int i=0; i<size; i++) {
            str.append(arr[i]);
            if (i != size-1) {
                str.append(",");
            }
        }
        str.append("]");
        return str.toString();
    }

    private void reSize(int newCapacity) {
        E[] newArr = (E[]) new Object[newCapacity];
        for (int i = 0; i < size; i++) {
            newArr[i] = arr[i];
        }
        arr = newArr;
    }
}