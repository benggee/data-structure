public class LinkListSet<E> implements Set<E>{
    private LinkList<E> list;
    
    public LinkListSet() {
        list = new LinkList<>();
    }

    @Override
    public int getSize(){
        return list.getSize();
    }

    @Override
    public boolean isEmpty() {
        return list.isEmpty();
    }

    @Override
    public void add(E e) {
        list.addFirst(e);
    }

    @Override
    public boolean contains(E e) {
        return list.find(e);
    }

    @Override 
    public void remove(E e) {
        list.removeElement(e);
    }
}