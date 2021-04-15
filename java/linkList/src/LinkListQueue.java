public class LinkListQueue<E>{
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

    private Node head;
    private Node tail;
    private int size;

    public LinkListQueue() {
        this.head = null;
        this.tail = null;
        this.size = 0;
    }

    public void enqueue(E e) {
        if (tail == null) {
            tail = new Node(e);
            head = tail;
        } else {
            tail.next = new Node(e);
            tail = tail.next;
        }
        size++;
    }

    public E dequeue() {
        if (size <= 0) 
            throw new IllegalArgumentException("The Queue is empty");
        Node tmpNode = head;
        head = head.next;
        // 注意这个地方容易写成 tmpNode = null 是不对的
        // 这是因为，tmpNode是作用域仅在函数里，这样赋值只是表示tmpNode指向的位置是一个空地址
        // 使用tmpNode.next = null 实际上是将原head指向的next那个地方清空，当函数结束时，tmpNode也将被销毁，达到释放内存的目地
        tmpNode.next = null; 
        if (head == null) 
            tail = null;

        size--;
        return tmpNode.e;
    }

    public E getFront() {
        if (size <= 0) {
            throw new IllegalArgumentException("The queue is empty");
        }
        return head.e;
    }

    public String toString() {
        StringBuilder ret = new StringBuilder();
        ret.append("Queue:");
        Node cur = head;
        while(cur != null) {
            ret.append(cur.e + "->");
            cur = cur.next;
        }
        ret.append("NULL");
        return ret.toString();
    }


    public static void main(String[] argc) {
        LinkListQueue<Integer> queue = new LinkListQueue<>();
        for (int i=0; i<10; i++) {
            queue.enqueue(i);
            System.out.println(queue);
        }

        queue.dequeue();
        queue.dequeue();
        queue.dequeue();
        queue.dequeue();
        System.out.println(queue);
    }
}