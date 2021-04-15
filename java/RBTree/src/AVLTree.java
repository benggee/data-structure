import java.util.ArrayList;

public class AVLTree<K extends Comparable<K>,V> {
    private class Node {
        private K key;
        private V value;
        private Node left,right;
        // 树的高度
        private int height;

        public Node(K key, V value) {
            this.key = key;
            this.value = value;
            this.left = null;
            this.right = null;
            this.height = 1;
        }
    }

    private Node root;
    private int size;

    public AVLTree() {
        this.root = null;
        this.size = 0;
    }

    public V remove(K key) {
        Node tmpNode = getNode(root, key);
        if (tmpNode != null) {
            remove(root, key);
            return tmpNode.value;
        }
        return null;
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

    public int size() {
        return size;
    }

    public boolean empty() {
        return size == 0;
    }

    // 判断是否是二分搜索树
    // 可以利用二分搜索树的特点 中序遍历得到的是一个从小到大的有序列表
    public boolean isBST() {
        ArrayList<K>  list = new ArrayList<>(size);
        isBST(root, list);
        for (int i=1; i<list.size(); i++) {
            if (list.get(i-1).compareTo(list.get(i)) > 0) 
                return false;
        }
        return true;
    }
    
    private void isBST(Node node, ArrayList<K> list) {
        if (node == null)
            return;
        isBST(node.left, list);
        list.add(node.key);
        isBST(node.right, list);
    }

    // 判断是否是AVL树
    // 利用平衡因子不能大于1的特点
    public boolean isAVL() {
        return isAVL(root);
    }

    private boolean isAVL(Node node) {
        if (node == null)
            return true;
        if (Math.abs(getAVLFactor(node))>1) 
            return false; 

        return isAVL(node.left) && isAVL(node.right);
    }

    // 获取节点的高度
    private int getHeight(Node node) {
        if (node == null)
            return 0;
        return node.height;
    }
    
    // 计算树的平衡因子  abs(左子树高度-右子树高度)
    private int getAVLFactor(Node node) {
        if (node == null) 
            return 0;
        return getHeight(node.left) - getHeight(node.right);
    }

    // 右旋转
    //        y                              x
    //       / \                           /   \
    //      x   T4     向右旋转 (y)        z     y
    //     / \       - - - - - - - ->    / \   / \
    //    z   T3                       T1  T2 T3 T4
    //   / \
    // T1   T2
    private Node rightRotate(Node y) {
        Node x = y.left;
        Node T3 = x.right;

        x.right = y;
        y.left = T3;

        y.height = Math.max(getHeight(y.right), getHeight(y.left)) + 1; 
        x.height = Math.max(getHeight(x.right), getHeight(x.left)) + 1;

        return x;
    }

    // 左旋转
    // 对节点y进行向左旋转操作，返回旋转后新的根节点x
    //    y                             x
    //  /  \                          /   \
    // T1   x      向左旋转 (y)       y     z
    //     / \   - - - - - - - ->   / \   / \
    //   T2  z                     T1 T2 T3 T4
    //      / \
    //     T3 T4
    private Node leftRotate(Node y) {
        Node x = y.right;
        Node T2 = x.left;

        x.left = y;
        y.right = T2;

        y.height = Math.max(getHeight(y.left), getHeight(y.right)) + 1;
        x.height = Math.max(getHeight(x.left), getHeight(x.right)) + 1;

        return x;
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

        // 重新计算高度
        node.height = 1 + Math.max(getHeight(node.left), getHeight(node.right));

        // 计算高度因子
        int factor = getAVLFactor(node);
        // LL
        // 如果节点的高度因子（左高度-右高度）大于1 
        // 并且，左子树的的高度因子>=0(也就是至少有一个元素) 
        // 则进行右旋转
        if (factor > 1 && getAVLFactor(node.left)>=0) 
            return rightRotate(node);

        // RR
        // 如果节点的高度因子 （左高度-右高度）小于-1 
        // 并且，右子树的高度因子<=0 (也就是至少有一个元素)
        // 则进行左旋转
        if (factor < -1 && getAVLFactor(node.right) <= 0)
            return leftRotate(node);

        // LR
        // 如果节点不平衡是因为左子节点的右子节点高度太高，则先将左子节点做一次左旋转，然后进行右旋转
        if (factor > 1 && getAVLFactor(node.left) < 0) {
            node.left = leftRotate(node.left);
            return rightRotate(node);
        }
        // RL
        // 如果节点不平衡是因为右子节点的左子节点高度太高，则先将右节点做一次右旋转，然后进行左旋转
        if (factor < -1 && getAVLFactor(node.right) > 0) {
            node.right = rightRotate(node.right);
            return leftRotate(node);
        }

        return node;
    }

    private Node remove(Node node, K key) {
        if (node == null)
            return null;

        Node retNode;
        if (key.compareTo(node.key) < 0) {
            node.left = remove(node.left, key);
            retNode =  node;
        } else if (key.compareTo(node.key) > 0) {
            node.right = remove(node.right, key);
            retNode = node;
        } else { // node.key.compareTo(key) ==0 的情况
            if (node.left == null) {
                Node tmpNode = node.right;
                node.right = null;
                size--;
                retNode = tmpNode;
            } else if (node.right == null) {
                Node tmpNode = node.left;
                node.left = null;
                size--;
                retNode = tmpNode;
            } else {
                // 找到比删除节点大的最小节点，这个节点应该在当前节点的右子树
                Node min = min(node.right);
                min.right = remove(node.right, min.key);
                min.left = node.left;

                node.left = node.right = null;
                retNode = min;
            }
        }
        if (retNode == null)
            return null;

        // 重新计算高度
        retNode.height = 1 + Math.max(getHeight(retNode.left), getHeight(retNode.right));

        // 计算高度因子
        int factor = getAVLFactor(retNode);
        // LL
        // 如果节点的高度因子（左高度-右高度）大于1 
        // 并且，左子树的的高度因子>=0(也就是至少有一个元素) 
        // 则进行右旋转
        if (factor > 1 && getAVLFactor(retNode.left)>=0) 
            return rightRotate(retNode);

        // RR
        // 如果节点的高度因子 （左高度-右高度）小于-1 
        // 并且，右子树的高度因子<=0 (也就是至少有一个元素)
        // 则进行左旋转
        if (factor < -1 && getAVLFactor(retNode.right) <= 0)
            return leftRotate(retNode);

        // LR
        // 如果节点不平衡是因为左子节点的右子节点高度太高，则先将左子节点做一次左旋转，然后进行右旋转
        if (factor > 1 && getAVLFactor(retNode.left) < 0) {
            retNode.left = leftRotate(retNode.left);
            return rightRotate(retNode);
        }
        // RL
        // 如果节点不平衡是因为右子节点的左子节点高度太高，则先将右节点做一次右旋转，然后进行左旋转
        if (factor < -1 && getAVLFactor(retNode.right) > 0) {
            retNode.right = rightRotate(retNode.right);
            return leftRotate(retNode);
        }

        return retNode;
    }

    private Node min(Node node) {
        if (node.left == null) 
            return node;
        return min(node.left);
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