public interface Stack<E> {
    int size();
    boolean empty();
    void push(E e);
    E pop();
    E peek();
}