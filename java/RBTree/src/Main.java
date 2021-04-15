import java.util.ArrayList;
public class Main{
    public static void testRBTree(ArrayList<String> word1) {
        long startTime = System.nanoTime();

        System.out.println("傲慢与偏见");
        System.out.println("Total words: " + word1.size());

        RBTree<String, Integer> map1 = new RBTree<>();
        for (String word: word1) {
            if (map1.contains(word))
                map1.set(word, map1.get(word)+1);
            else
                map1.set(word, 1); 

        }

        System.out.println("Total different words: " + map1.size());
        System.out.println("Frequency of PRIDE: " + map1.get("pride"));  
        System.out.println("Frequency of PREJUDICE: " + map1.get("prejudice"));


        long endTime = System.nanoTime();

        System.out.println("Times:" + (endTime-startTime)/1000000000.0);
    }

    public static void main(String[] argv) {
        ArrayList<String> word1 = new ArrayList<>();

        if (!FileOperation.readFile("Set/pride-and-prejudice.txt", word1)) 
            throw new IllegalArgumentException("File open error.");


        testRBTree(word1);

    }
}