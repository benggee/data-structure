import java.util.ArrayList;
import java.util.Collections;

public class Main{
    // 搜索树Set测试
    private void testBSTSet(ArrayList<String> word1, ArrayList<String> word2) {
        long startTime = System.nanoTime();

        System.out.println("傲慢与偏见");
        BinarySearchTreeSet<String> set1 = new BinarySearchTreeSet<>();
        for (String word: word1) {
            //System.out.println(word);
            set1.add(word);
        }
        System.out.println("Total different words: " + set1.getSize());

        System.out.println();

        System.out.println("双城记");

        BinarySearchTreeSet<String> set2 = new BinarySearchTreeSet<>();
        for (String word: word2) {
            set2.add(word);
        }

        System.out.println("Total different words: " + set2.getSize());

        long endTime = System.nanoTime();

        System.out.println("BST Times:" + (endTime-startTime)/1000000000.0);
    }

    // 链表Set测试
    private void testLinkListSet(ArrayList<String> word1, ArrayList<String> word2) {
        long startTime = System.nanoTime();

        System.out.println("傲慢与偏见");
        LinkListSet<String> l1 = new LinkListSet<>();
        for (String word: word1) {
            if (!l1.contains(word)) {
                l1.add(word);
            }
        }
        System.out.println("Total different words: " + l1.getSize());

        System.out.println();

        System.out.println("双城记");

        LinkListSet<String> l2 = new LinkListSet<>();
        for (String word: word2) {
            if (!l2.contains(word)) {
                l2.add(word);
            }
        }

        System.out.println("Total different words: " + l2.getSize());

        long endTime = System.nanoTime();

        System.out.println("Link List Times:" + (endTime-startTime)/1000000000.0);
    }
    
    // 搜索树Set测试
    private void testAVLSet(ArrayList<String> word1, ArrayList<String> word2) {
        long startTime = System.nanoTime();

        System.out.println("傲慢与偏见");
        AVLSet<String, Object> set1 = new AVLSet<>();
        for (String word: word1) {
            //System.out.println(word);
            set1.add(word);
        }
        System.out.println("Total different words: " + set1.getSize());

        System.out.println();

        System.out.println("双城记");

        AVLSet<String, Object> set2 = new AVLSet<>();
        for (String word: word2) {
            set2.add(word);
        }

        System.out.println("Total different words: " + set2.getSize());

        long endTime = System.nanoTime();

        System.out.println("AVL Times:" + (endTime-startTime)/1000000000.0);
    }

    public static void main(String argv[]) {
        ArrayList<String> word1 = new ArrayList<>();
        ArrayList<String> word2 = new ArrayList<>();

        if (!FileOperation.readFile("Set/pride-and-prejudice.txt", word1)) 
            throw new IllegalArgumentException("File open error.");
        if (!FileOperation.readFile("Set/a-tale-of-two-cities.txt", word2)) 
            throw new IllegalArgumentException("File open error.");

        Collections.sort(word1);
        Collections.sort(word2);

        Main main = new Main();
        main.testLinkListSet(word1, word2);
        System.out.println();
        main.testBSTSet(word1, word2);
        System.out.println();
        main.testAVLSet(word1, word2);
    }
        
}