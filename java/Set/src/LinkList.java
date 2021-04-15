import java.lang.instrument.IllegalClassFormatException;

public class LinkList<E> {
    private class Node {
        public E e;
        public Node next;

        public Node(E e, Node next) {
            this.e = e;
            this.next = next;
        }

        public Node(E e) {
            this(e, null);
        }

        public Node() {
            this(null, null);
        }

        @Override
        public String toString() {
            return e.toString();
        }
    }

    private Node dummyHead;
    private int size;
    
    public LinkList() {
        dummyHead = new Node();
        size = 0;
    }

    // 返回长度
    public int getSize() {
        return size;
    }

    // 是否为空
    public boolean isEmpty() {
        return size == 0;
    }

    // 添加到指定索引位置 
    public void add(int index, E e) {
        if (index < 0 || index > size) 
            throw new IllegalArgumentException("Index out of range.");
        Node pre = dummyHead;
        for (int i=0; i<index; i++) 
            pre = pre.next;
    
        Node node = new Node(e);
        node.next = pre.next;
        pre.next = node;
        size++;
    }

     // 添加到头部
     public void addFirst(E e) {
        add(0, e);
    }

    // 添加到末尾
    public void addLast(E e) {
        add(size, e);
    }

    // 查找 指定第几个
    public E get(int index) {
       if (index < 0 || index > size) 
            throw new IllegalArgumentException("Index out of range.");
        Node cur = dummyHead.next;
        for (int i=0; i<index; i++) {
            cur = cur.next;
        }
        return cur.e;
    }

    // 获取第一个
    public E getFirst() {
        return get(0);
    }

    // 获取最后一个
    public E getLast() {
        return get(size-1);
    }

    // 设置
    public void set(int index, E e) {
        if (index < 0 || index > size) 
            throw new IllegalArgumentException("Index out of range.");
        Node cur = dummyHead.next;
        for (int i=0; i<size; i++) {
            cur = cur.next;
        }
        cur.e = e;
    }

    // 查找元素
    public boolean find(E e) {
        Node cur = dummyHead.next;
        while(cur != null) {
            if (cur.e.equals(e))
                return true;
            cur = cur.next;
        }
        return false;
    }

    // 删除元素
    public E del(int index) {
        if (index < 0 || index > size) 
            throw new IllegalArgumentException("Index out of range.");
        Node pre = dummyHead;
        for (int i=0; i<index; i++) {
            pre = pre.next;
        }
        Node delNode = pre.next;
        pre.next = delNode.next;
        delNode.next = null;
        return delNode.e;
    }

    // 删除元素
    public void removeElement(E e) {
        Node prev = dummyHead;
        while(prev.next != null) {
            if (prev.next.e.equals(e)) 
                break;
            prev = prev.next;
        }

        if (prev.next != null) {
            Node tmpNode = prev.next;
            prev.next = tmpNode.next;
            tmpNode.next = null;
            size--;
        }
    }

    // 删除第一个元素
    public E delFirst() {
        return del(0);
    }

    // 删除最后一个元素
    public E delLast() {
        return del(size);
    }

    @Override
    public String toString() {
        StringBuilder res = new StringBuilder();
        Node cur = dummyHead.next;
        while(cur != null) {
            res.append(cur + "->");
            cur = cur.next;
        }
        res.append("NULL");
        return res.toString();
    }
}