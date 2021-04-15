public class LoopQueue<E> implements Queue<E> {
    private E[] data;
    private int front;
    private int tail;
    private int size;

    public LoopQueue(int capacity) {
        data = (E[])new Object[capacity + 1];
        this.front = 0;
        this.tail = 0;
        this.size = 0;
    }

    public LoopQueue() {
        this(10);
    }

    public int getCapacity() {
        return data.length - 1;
    }

    @Override
    public int size() {
        return size;
    }

    @Override
    public boolean empty() {
        return front == tail;
    }

    @Override
    public void enqueue(E e) {
        // 如果队尾下一个元素是队首说明队列满了，进行扩容
        if ((tail + 1) % data.length == front) 
            resize(2 * getCapacity());
        data[tail] = e;
        tail = (tail + 1) % data.length;
        size++;
    }

    @Override
    public E dequeue() {
        if (empty())
            throw new IllegalArgumentException("Cannot dequeue from empty queue.");
        E ret = data[front];
        data[front] = null;
        front = (front + 1) % data.length;
        size--;
        if (size == getCapacity()/4 && getCapacity()/2 != 0) 
            resize(getCapacity()/2);
        return ret;
    }

    @Override
    public E front() {
        if (empty())
            throw new IllegalArgumentException("Queue is empty.");
        return data[front];
    }

    private void resize(int newSize) {
        E[] newData = (E[])new Object[newSize + 1];
        for (int i=0; i<size; i++) {
            newData[i] = data[(i + front) % data.length];
        }
        data = newData;
        front = 0; 
        tail = size;
    }

    public String toString() {
        StringBuilder str = new StringBuilder();
        str.append(String.format("Queue: size=%d, capacity=%d\n", size, getCapacity()));
        str.append("front [");
        for (int i=front; i != tail; i = (i + 1) % data.length) {
            str.append(data[i]);
            if ((i + 1) % data.length != tail) 
                str.append(",");
        }
        str.append("]tail");
        return str.toString();
    }

    public static void main(String[] args) {
        LoopQueue<Integer> queue = new LoopQueue<>(5);
        for (int i=0; i<10; i++) {
            queue.enqueue(i);
            System.out.println(queue);

            if (i % 3 == 2) {
                queue.dequeue();
                System.out.println(queue);
            }
        }
    }
}