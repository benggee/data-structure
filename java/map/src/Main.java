import java.util.ArrayList;

public class Main{
    public void testLinkListMap() {
        long startTime = System.nanoTime();

        System.out.println("傲慢与偏见");
        ArrayList<String> word1 = new ArrayList<>();
        if (FileOperation.readFile("Map/pride-and-prejudice.txt", word1)) {
            System.out.println("Total words: " + word1.size());

            LinkListMap<String, Integer> map1 = new LinkListMap<>();
            for (String word: word1) {
                if (map1.contains(word))
                    map1.set(word, map1.get(word)+1);
                else
                    map1.set(word, 1); 

            }
            System.out.println("Total different words: " + map1.size());
            System.out.println("Frequency of PRIDE: " + map1.get("pride"));  
            System.out.println("Frequency of PREJUDICE: " + map1.get("prejudice"));
        }

        System.out.println();

        System.out.println("双城记");

        ArrayList<String> word2 = new ArrayList<>();
        if (FileOperation.readFile("Map/a-tale-of-two-cities.txt", word2)) {
            System.out.println("Total words: " + word2.size());
            
            LinkListMap<String, Integer> map2 = new LinkListMap<>();
            for (String word: word2) {
                if (map2.contains(word)) 
                    map2.set(word, map2.get(word)+1);
                else  
                    map2.set(word, 1);
            }

            System.out.println("Total different words: " + map2.size());
            System.out.println("Frequency of PRIDE: " + map2.get("pride"));
            System.out.println("Frequency of PREJUDICE: " + map2.get("prejudice"));
        }

        long endTime = System.nanoTime();

        System.out.println("Link List Map Times:" + (endTime-startTime)/1000000000.0);
    }

    public void testBinarySearchTreeMap() {
        long startTime = System.nanoTime();

        System.out.println("傲慢与偏见");
        ArrayList<String> word1 = new ArrayList<>();
        if (FileOperation.readFile("Map/pride-and-prejudice.txt", word1)) {
            System.out.println("Total words: " + word1.size());

            BinarySearchTreeMap<String, Integer> map1 = new BinarySearchTreeMap<>();
            for (String word: word1) {
                if (map1.contains(word))
                    map1.set(word, map1.get(word)+1);
                else
                    map1.set(word, 1); 

            }
            System.out.println("Total different words: " + map1.size());
            System.out.println("Frequency of PRIDE: " + map1.get("pride"));  
            System.out.println("Frequency of PREJUDICE: " + map1.get("prejudice"));
        }

        System.out.println();

        System.out.println("双城记");

        ArrayList<String> word2 = new ArrayList<>();
        if (FileOperation.readFile("Map/a-tale-of-two-cities.txt", word2)) {
            System.out.println("Total words: " + word2.size());
            
            BinarySearchTreeMap<String, Integer> map2 = new BinarySearchTreeMap<>();
            for (String word: word2) {
                if (map2.contains(word)) 
                    map2.set(word, map2.get(word)+1);
                else  
                    map2.set(word, 1);
            }

            System.out.println("Total different words: " + map2.size());
            System.out.println("Frequency of PRIDE: " + map2.get("pride"));
            System.out.println("Frequency of PREJUDICE: " + map2.get("prejudice"));
        }

        long endTime = System.nanoTime();

        System.out.println("BST Map Times:" + (endTime-startTime)/1000000000.0);
    }

    public void testAVLMap() {
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
        }

        System.out.println();

        System.out.println("双城记");

        ArrayList<String> word2 = new ArrayList<>();
        if (FileOperation.readFile("Map/a-tale-of-two-cities.txt", word2)) {
            System.out.println("Total words: " + word2.size());
            
            AVLTree<String, Integer> map2 = new AVLTree<>();
            for (String word: word2) {
                if (map2.contains(word)) 
                    map2.set(word, map2.get(word)+1);
                else  
                    map2.set(word, 1);
            }

            System.out.println("Total different words: " + map2.size());
            System.out.println("Frequency of PRIDE: " + map2.get("pride"));
            System.out.println("Frequency of PREJUDICE: " + map2.get("prejudice"));
        }

        long endTime = System.nanoTime();

        System.out.println("AVL Tree Times:" + (endTime-startTime)/1000000000.0);
    }

    public static void main(String[] argv) {
        Main main = new Main();
        main.testLinkListMap();
        System.out.println();
        main.testBinarySearchTreeMap();
        System.out.println();
        main.testAVLMap();
    }
}