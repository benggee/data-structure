public class RBTree<K extends Comparable<K>, V> {
    public static final boolean RED = true;
    public static final boolean BLACK = false;

    private class Node {
        private K key;
        private V value;
        private Node left,right;
        private boolean color;

        public Node(K key, V value) {
            this.key = key;
            this.value = value;
            this.left = null;
            this.right = null;
            this.color = RED;
        }
    }

    private Node root;
    private int size;

    public RBTree() {
        root = null;
        size = 0;
    }

    public int size() {
        return size;
    }

    public boolean empty() {
        return size == 0;
    }

    public boolean contains(K key){
        return !(getNode(root, key) == null);
    }

    public V get(K key) {
        Node tmpNode = getNode(root, key);
        return tmpNode==null ? null : tmpNode.value;
    }

    public void set(K key, V value) {
        root = add(root, key, value);
    }

    // 左旋转  
    //   node                     x
    //  /   \     左旋转         /  \
    // T1   x   --------->   node   T3
    //     / \              /   \
    //    T2 T3            T1   T2
    public Node leftRotate(Node node) {
        Node x = node.right;

        node.right = x.left;
        x.left = node;

        x.color = node.color;
        node.color = RED;

        return x;
    }

    // 右旋转
    //     node                   x
    //    /   \     右旋转       /  \
    //   x    T2   ------->   y   node
    //  / \                       /  \
    // y  T1                     T1  T2
    public Node rightRotate(Node node) {
        Node x = node.left;

        node.left = x.right;
        x.right = node;

        x.color = node.color;
        node.color = RED;
        
        return x;
    }

    // 颜色翻转
    private void flipColor(Node node) {
        node.color = RED;
        node.left.color = BLACK;
        node.right.color = BLACK;
    }

    private boolean isRed(Node node) {
        if (node == null)
            return BLACK;
        return node.color;
    }

    private Node add(Node node, K key, V value) {
        if (node == null) {
            size++;
            return new Node(key, value);
        }

        if (key.compareTo(node.key) < 0) 
            node.left = add(node.left, key, value);
        else if (key.compareTo(node.key) > 0) 
            node.right = add(node.right, key, value);
        else 
            node.value = value;

        if (isRed(node.right) && !isRed(node.left)) 
            node = leftRotate(node);
        
        if (isRed(node.left) && isRed(node.left.left)) 
            node = rightRotate(node);
        
        if (isRed(node.left) && isRed(node.right))
            flipColor(node);

        return node;
    }
    private Node getNode(Node node,K key) {
        
        if (node == null) 
            return null;
        
        if (key.equals(node.key)) 
            return node;
        else if (key.compareTo(node.key) < 0) 
            return getNode(node.left, key);
        else    
            return getNode(node.right, key);
    }
}