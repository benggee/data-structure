import java.util.HashSet;
import java.util.HashMap;
import java.util.ArrayList;

public class Main{

    public static void testHashTable() {
        System.out.println("Pride and Prejudice");

        ArrayList<String> words = new ArrayList<>();
        if(FileOperation.readFile("Hash/pride-and-prejudice.txt", words)) {
            System.out.println("Total words: " + words.size());
            // Test HashTable
            long startTime = System.nanoTime();

            HashTable<String, Integer> ht = new HashTable<>();
            //HashTable<String, Integer> ht = new HashTable<>(131071);
            for (String word : words) {
                if (ht.contains(word))
                    ht.set(word, ht.get(word) + 1);
                else
                    ht.add(word, 1);
            }

            for(String word: words)
                ht.contains(word);

            long endTime = System.nanoTime();

            double time = (endTime - startTime) / 1000000000.0;
            System.out.println("HashTable: " + time + " s");
        }
    }

    public static void testHashCode() {
        int a = 42; 
        System.out.println(((Integer)a).hashCode());

        int b = -42;
        System.out.println(((Integer)b).hashCode());

        double p = 3.1415926;
        System.out.println(((Double)p).hashCode());

        String d = "seepre";
        System.out.println(d.hashCode());


        Student student = new Student(12, 4, "seepre", "aaa");
        System.out.println(student.hashCode());

        HashSet<Student> set = new HashSet<>();
        set.add(student);

        HashMap<Student, Integer> map = new HashMap<>();
        map.put(student, 100);
    }

    public static void main(String[] argv) {
        
        testHashTable();
    }
}