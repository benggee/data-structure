import java.util.Random;

public class Main{
    public static double testQueue(Queue<Integer> q, int opCount) {
        long startTime = System.nanoTime();
        
        Random random = new Random();
        for (int i=0; i<opCount; i++) 
            q.enqueue(random.nextInt(Integer.MAX_VALUE));
        
        for (int i=0; i<opCount; i++) 
           q.dequeue();

        long endTime = System.nanoTime();
        double times = (endTime - startTime) / 1000000000.0;
        return times;
    }

    public static void main(String[] argv) {
        int opCount = 100000;
        ArrayQueue<Integer> aq = new ArrayQueue<>();

        double time1  = testQueue(aq, opCount);
        System.out.println("Array Queue time: " + time1);

        LoopQueue<Integer> lq = new LoopQueue<>();
        double time2 = testQueue(lq, opCount);
        System.out.println("Loop Queue time: " + time2);
    }
}