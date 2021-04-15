public interface Map<K, V>  {
    V remove(K key);
    boolean contains(K key);
    V get(K key);
    void set(K key, V value);
    int size();
    boolean empty();
}