import java.util.Random;
public class Main{

    public double testHeapify(Integer[] heapData, boolean isHeapify) {
        long startTime = System.nanoTime();
        
        MaxHeap<Integer> heap;
        if (isHeapify) {
            System.out.println("aaaaa");
            heap = new MaxHeap<>(heapData);
        } else {
            heap = new MaxHeap<>();
            System.out.println("bbbbb");
            for (int i=0; i<heapData.length;i++) 
                heap.add(heapData[i]);
        }

        int[] arr = new int[heapData.length];
        for (int i=0; i<arr.length; i++) {
            arr[i] = heap.extractMax();
        }

        for (int i=0; i<heapData.length - 1; i++) {
            if (arr[i]<arr[i+1]) 
                throw new IllegalArgumentException("error.");
        }

        long endTime = System.nanoTime();

        return (endTime-startTime) / 1000000000.0;
    }

    public void testBase() {
        long startTime = System.nanoTime();

        int n = 1000000;
        MaxHeap<Integer> heap = new MaxHeap<>();
        Random random = new Random();
        for (int i=0; i<n; i++) {
            heap.add(random.nextInt(Integer.MAX_VALUE));
        }

        int[] arr = new int[n];
        for (int i=0; i<n; i++) {
            arr[i] = heap.extractMax();
            System.out.println(arr[i]);

        }

        for (int i=0; i<n - 1; i++) {
            if (arr[i]<arr[i+1]) 
                throw new IllegalArgumentException("error.");
        }

        long endTime = System.nanoTime();

        System.out.println("Times:" + (endTime - startTime)/1000000000.0);
    }
    
    public static void main(String[] argv) {
        int N = 10000000;
        Integer[] testData = new Integer[N];
        Random random = new Random();
        for (int i=0; i<testData.length; i++) {
            testData[i] = random.nextInt(Integer.MAX_VALUE);
        }
        
        Main main = new Main();
        double times1 = main.testHeapify(testData, true);
        System.out.println("Times1:" + times1);

        double times2 = main.testHeapify(testData, false);
        System.out.println("Times2:" + times2);
    }
}