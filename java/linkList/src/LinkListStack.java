import java.util.Random;

public class LinkListStack<E> {
    private LinkList<E> list;
    
    public LinkListStack() {
        list = new LinkList<>();
    }

    public int getSize() {
        return list.getSize();
    }

    public boolean isEmpty() {
        return list.isEmpty();
    }

    public void push(E e) {
        list.addFirst(e);
    }

    public E pop() {
        return list.delFirst();
    }

    public E peek() {
        return list.getFirst();
    }

    @Override
    public String toString() {
        StringBuilder ret = new StringBuilder();
        ret.append("Stack: top");
        ret.append(list);
        return ret.toString();
    }

    public static double testStack(LinkListStack<Integer> stack, int opCount) {
        long startTime = System.nanoTime();
        Random random = new Random();
        for (int i=0; i<opCount; i++) {
            stack.push(random.nextInt(Integer.MAX_VALUE));
        }

        for (int i=0; i<opCount; i++) {
            stack.pop();
        }

        long endTime = System.nanoTime();
        
        return (endTime - startTime) / 1000000000.0;
    }

    public static void main(String[] argv) {
        LinkListStack<Integer> stack = new LinkListStack<>();
        System.out.println("Time:" + testStack(stack, 1000000));

        // LinkListStack<Integer> stack = new LinkListStack<>();

        // for (int i=0; i<5; i++) {
        //     stack.push(i);
        //     System.out.println(stack);
        // }

        // stack.pop();
        // System.out.println(stack);
    }
}