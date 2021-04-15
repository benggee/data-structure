import java.util.Stack;
import java.util.LinkedList;
import java.util.Queue;
public class BinarySearchTree<E extends Comparable<E>> {
    private class Node {
        public E e;
        public Node left,right;
        public Node(E e) {
            this.e = e;
            this.left = null;
            this.right = null;
        }
    }

    private Node root;
    private int size;

    public BinarySearchTree() {
        this.root = null;
        this.size = 0;
    }

    public int size() {
        return size;
    }

    public boolean isEmpty() {
        return size == 0;
    }

    public void add(E e) {
        root = this.add(root, e);
    }

    private Node add(Node node, E e) {
        if (node == null) {
            size++;
            return new Node(e);
        }

        if (e.compareTo(node.e) < 0)
            node.left = this.add(node.left, e);
        else if (e.compareTo(node.e) > 0)
            node.right = this.add(node.right, e);
        return node;
    }

    public boolean contains(E e) {
        return contains(root, e);
    }

    private boolean contains(Node node, E e) {
        if (node == null)
            return false;

        if (e.compareTo(node.e) == 0) {
            return true;
        }  else if (e.compareTo(node.e) < 0) {
            return contains(node.left, e);
        } else {
            return contains(node.right, e);
        }
    }

    // 前序遍历
    public void preOrder() {
        preOrder(root);
    } 

    private void preOrder(Node node) {
        if (node == null)
            return;

        System.out.println(node.e);
        preOrder(node.left);
        preOrder(node.right);
    }

    // 非递归前序遍历
    public void preOrderNR() {
        if (root == null)
            return;
        
        Stack<Node> stack = new Stack<>();
        stack.push(root);
        while(!stack.isEmpty()) {
            Node cur = stack.pop();
            System.out.println(cur.e);
            if (cur.right != null)
                stack.push(cur.right);
            if (cur.left != null)
                stack.push(cur.left);
        }
    }

    // 中序遍历
    public void inOrder() {
        inOrder(root);
    }

    private void inOrder(Node node) {
        if (node == null) 
            return;
        inOrder(node.left);
        System.out.println(node.e);
        inOrder(node.right);
    }

    // 后序遍历 
    public void postOrder() {
        postOrder(root);
    }

    private void postOrder(Node node) {
        if (node == null)
            return;
        postOrder(node.left);
        postOrder(node.right);
        System.out.println(node.e);
    }

    // 层序遍历 
    public void levelOrder() {
        if (root == null)
            return;
        Queue<Node> q = new LinkedList<>();
        q.add(root);
        while(!q.isEmpty()) {
            Node cur = q.remove();
            
            System.out.println(cur.e);
            if (cur.left != null)
                q.add(cur.left);
            if (cur.right != null) 
                q.add(cur.right);
        }
    }

    // 获取最小值
    public E min() {
        return min(root).e;
    }

    private Node min(Node node) {
        if (node.left == null) 
            return node;

        return min(node.left);
    }

    // 获取最大值
    public E max() {
        return max(root).e;
    }

    private Node max(Node node) {
        if (node.left == null) 
            return node;
        return max(node.right);
    }

    // 删除最小值 
    public E delMin() {
        E e = min();
        root = delMin(root);
        return e;
    }

    // 删除最小值
    private Node delMin(Node node) {
        if (node.left == null) {
            Node tmpNode = node.right;
            node.right = null;
            size--;
            return tmpNode;
        }

        node.left = delMin(node.left);
        return node;
    }

    // 删除最大值
    public E delMax() {
        E e = max();
        root = delMax(root);
        return e;
    }

    private Node delMax(Node node) {
        if (node.right == null) {
            Node tmpNode = node.left;
            node.left = null;
            size--;
            return tmpNode;
        }
        node.right = delMax(node.right);
        return node;
    }

    // 删除任意节点
    public void del(E e) {
        root = del(root, e);   
    }

    private Node del(Node node, E e) {
        if (node == null) 
            return null;
        if (e.compareTo(node.e) < 0) {
            node.left = del(node.left, e);
            return node;
        } else if (e.compareTo(node.e) > 0) {
            node.right = del(node.right, e);
            return node;
        } else { // e.compareTo(node.e) == 0
            // 左子树为空的情况
            if (node.left == null) {
                Node tmpNode = node.right;
                node.right = null;
                size--;
                return tmpNode;
            }
            // 右子树为空的情况
            if (node.right == null) {
                Node tmpNode = node.left;
                node.left = null;
                size--;
                return tmpNode;
            }
            // 左右子树都不为空的情况
            // 找到最右子树最小的节点来补充到要删除的位置
            Node successor = min(node.right);
            successor.right = delMin(node.right);
            successor.left = node.left;

            node.right = node.left = null;

            return successor;
        }
    }

    @Override
    public String toString() {
        StringBuilder str = new StringBuilder();
        // str.append("null");
        
        generaTreeStr(root, 0, str);
        return str.toString();
    }
    private void generaTreeStr(Node node, int depth, StringBuilder str) {
        if (node == null) {
            str.append(generaString(depth)+"null\n");
            return;
        }
        str.append(generaString(depth) + node.e + "\n");
        generaTreeStr(node.left, depth++, str);
        generaTreeStr(node.right, depth++, str);
    }
    private String generaString(int depth) {
        StringBuilder str = new StringBuilder();
        for (int i=0; i<depth; i++) {
            str.append("--");
        }
        return str.toString();
    }


    public static void main(String[] argv) {
        BinarySearchTree<Integer> bst = new BinarySearchTree<>();
        int[] tree = {3,4,5,12,343,8,10,22};
        for (int i=0; i<tree.length; i++) {
            bst.add(tree[i]);
        }

        bst.preOrder();
        System.out.println();
        bst.preOrderNR();

       // System.out.println(bst);
    }
}