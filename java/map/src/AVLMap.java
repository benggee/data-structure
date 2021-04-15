public class AVLMap<K extends Comparable<K>, V> implements Map<K, V> {
    private AVLTree<K,V> avl;
    
    public AVLMap() {
        avl = new AVLTree<>();
    }

    @Override
    public int size() {
        return avl.size();
    }

    @Override
    public boolean empty() {
        return avl.empty();
    }

f   @Override
    public V remove(K key) {
        return avl.remove(k);
    }

    @Override
    public boolean contains(K key) {
        return avl.contains(k);
    }

    @Override
    public V get(K key) {
        return avl.get(key);
    }

    @Override
    public void set(K key, V value) {
        avl.set(key, value);
    }
 
} 