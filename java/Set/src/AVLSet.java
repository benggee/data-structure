public class AVLSet<E extends Comparable<E>, Object> implements Set<E> {
    private AVLTree<E, Object> avl;

    public AVLSet() {
        avl = new AVLTree<>();
    }

    @Override
    public int getSize() {
        return avl.size();
    }

    @Override
    public boolean isEmpty(){
        return avl.empty();
    }

    @Override
    public void add(E e) {
        avl.set(e, null);
    }

    @Override
    public boolean contains(E e) {
        return avl.contains(e);
    }

    @Override
    public void remove(E e) {
        avl.remove(e);
    }
    
}