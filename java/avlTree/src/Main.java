import java.util.ArrayList;
import java.util.Collections;

public class Main{
    public void testAVLTree() {
        long startTime = System.nanoTime();

        System.out.println("傲慢与偏见");
        ArrayList<String> word1 = new ArrayList<>();
        if (FileOperation.readFile("Map/pride-and-prejudice.txt", word1)) {
            System.out.println("Total words: " + word1.size());

            AVLTree<String, Integer> map1 = new AVLTree<>();
            for (String word: word1) {
                if (map1.contains(word))
                    map1.set(word, map1.get(word)+1);
                else
                    map1.set(word, 1); 

            }

            System.out.println("Total different words: " + map1.size());
            System.out.println("Frequency of PRIDE: " + map1.get("pride"));  
            System.out.println("Frequency of PREJUDICE: " + map1.get("prejudice"));

            System.out.println("Is Binary Tree: " + map1.isBST());
            System.out.println("Is AVL tree: " + map1.isAVL());

            // 测试每次删除之后是否还是AVL树
            for (String word: word1) {
                map1.remove(word);
                if (!map1.isAVL())
                    throw new IllegalArgumentException("AVL Error.");
            }
        }
        long endTime = System.nanoTime();

        System.out.println("Times:" + (endTime-startTime)/1000000000.0);
    }

    public void testBranch() {
        
        System.out.println("傲慢与偏见");
        ArrayList<String> word1 = new ArrayList<>();
        if (FileOperation.readFile("Map/pride-and-prejudice.txt", word1)) {
            System.out.println("Total words: " + word1.size());
            
            // 排一下排，让二分搜索树退化成链表
            Collections.sort(word1);

            long time1 = System.nanoTime();
            AVLTree<String, Integer> map1 = new AVLTree<>();
            for (String word: word1) {
                if (map1.contains(word))
                    map1.set(word, map1.get(word)+1);
                else
                    map1.set(word, 1); 

            }

            for (String word: word1) {
                map1.contains(word);
            }

            long time2 = System.nanoTime();
            System.out.println("The AVLTree times: " + (time2-time1)/1000000000.0);

            BinarySearchTree<String, Integer> map2 = new BinarySearchTree<>();
            for (String word: word1) {
                if (map2.contains(word))
                    map2.set(word, map1.get(word)+1);
                else
                    map2.set(word, 1); 

            }

            for (String word: word1) {
                map2.contains(word);
            }

            long time3 = System.nanoTime();
            System.out.println("The BSTTree times: " + (time3-time2)/1000000000.0);
        }
    }

    public static void main(String[] argv) {
        Main main = new Main();
        main.testAVLTree();
    }
}