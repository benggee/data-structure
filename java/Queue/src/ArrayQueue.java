public class ArrayQueue<E> implements Queue <E> {
    private Array<E> array;

    public ArrayQueue(int capacity) {
        array = new Array<>(capacity);
    }

    public ArrayQueue() {
        array = new Array<>();
    }

    @Override
    public int size() {
        return array.size();
    }

    @Override
    public boolean empty() {
        return array.isEmpty();
    }

    @Override
    public void enqueue(E e){
        array.addLast(e);
    }

    @Override
    public E dequeue() {
        return array.removeFirst();
    }

    @Override
    public E front() {
        return array.get(0);
    }

    @Override
    public String toString() {
        StringBuilder str = new StringBuilder();
        str.append("front:[");
        for (int i=0; i<array.size(); i++) {
            str.append(array.get(i));
            if (array.size()-1 > i)
                str.append(",");
        }
        str.append("]tail");
        return str.toString();
    }
}