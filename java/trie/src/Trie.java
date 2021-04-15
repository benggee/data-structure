import java.util.TreeMap;

public class Trie {
    private class Node {
        public boolean isWord;
        TreeMap<Character, Node> next;

        public Node(boolean isWord) {
            this.isWord = isWord;
            next = new TreeMap<>();
        }

        public Node() {
            this(false);
        }
    }

    private Node root;
    private int size;

    public Trie() {
        root = new Node();
        size = 0;
    }

    public int size() {
        return size;
    }
    
    // 添加
    public void add(String word) {
        Node cur = root;
        for (int i=0; i<word.length(); i++) {
            char c = word.charAt(i);
            if (cur.next.get(c) == null) {
                cur.next.put(c, new Node());
            }
            cur = cur.next.get(c);
        }
        if (!cur.isWord) {
            cur.isWord = true;
            size++;
        }
    }
    
    // 查询
    public boolean contains(String word) {
        Node cur = root;
        for (int i=0; i<word.length(); i++) {
            char c = word.charAt(i);
            if (cur.next.get(c) == null)
                return false;
            cur = cur.next.get(c);
        }
        return cur.isWord;
    }

    // 前缀查询
    public boolean prefixSearch(String prefix) {
        Node cur = root;
        for (int i=0; i<prefix.length(); i++) {
            char c = prefix.charAt(i);
            if (cur.next.get(c) == null) 
                return false;
        }
        return true;
    }
}