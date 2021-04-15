public class LinkListMap<K, V> implements Map<K, V>{
    private class Node {
        public K key;
        public V value;
        public Node next;

        public Node(K key, V value, Node next) {
            this.value = value;
            this.key = key;
            this.next = next;
        }

        public Node(K key) {
            this(key, null, null);
        }

        public Node() {
            this(null, null, null);
        }

        @Override
        public String toString() {
            return key.toString() + ":" + value.toString();
        }
    }

    private Node dummyHead;
    private int size;

    public LinkListMap() {
        dummyHead = new Node();
        size = 0;
    }

    private Node getNode(K key) {
        Node cur = dummyHead.next;
        while (cur != null) {
            if (cur.key.equals(key))
                return cur;
            cur = cur.next;
        }
        return null;
    }

    @Override
    public int size() {
        return size;
    }

    @Override
    public boolean empty() {
        return size == 0;
    }

    @Override 
    public boolean contains(K key) {
        return getNode(key) != null;
    }

    @Override 
    public V get(K key) {
        Node tmpNode = getNode(key);
        return tmpNode==null ? null : tmpNode.value;
    }

    @Override
    public void set(K key, V value) {
        Node tmpNode = getNode(key);
        if (tmpNode != null) 
            tmpNode.value = value;
        else {
            dummyHead.next = new Node(key, value, dummyHead.next);
            size++;
        }
    }

    @Override
    public V remove(K key) {
        Node prev = dummyHead;
        while(prev.next!=null) {
            if (prev.key.equals(key))
                break;
            prev = prev.next;
        }

        if (prev != null) {
            Node tmpNode = prev.next;
            prev.next = tmpNode.next;
            tmpNode.next = null;
            size--;
            return tmpNode.value;
        }
        return null;
    }
}