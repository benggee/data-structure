import java.util.ArrayList;

public class Main{
    public static void main(String[] argv) {
        long startTime = System.nanoTime();

        System.out.println("傲慢与偏见");
        ArrayList<String> word1 = new ArrayList<>();
        if (FileOperation.readFile("Trie/pride-and-prejudice.txt", word1)) {
            System.out.println("Total words: " + word1.size());

            Trie t1 = new Trie();
            for (String word: word1) {
                t1.add(word);
            }

            for (String word: word1) {
                t1.contains(word);
            }
            System.out.println("Total different words: " + t1.size());
        }
        long endTime = System.nanoTime();

        System.out.println("Times:" + (endTime-startTime)/1000000000.0);
    }
}