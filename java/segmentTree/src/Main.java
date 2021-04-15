public class Main{
    public void testQuery() {
        Integer[] arr = {-3, 44, 2, 1, 9, -9, -10, 99, 34, 66};
        SegmentTree<Integer> segmentTree = new SegmentTree<>(arr, (r, l)->r + l);
        
        System.out.println("Treeaaa:  " +  segmentTree);
        System.out.println(segmentTree.query(1, 2));
    }

    public static void main(String[] argv) {
        Main main = new Main();
        main.testQuery();
        // Integer[] nums = {-2, 0, 3, -5, 2, -1};
        // SegmentTree<Integer> segTree = new SegmentTree<>(nums,
        //        (a, b) -> a + b);
        // System.out.println(segTree);
    }
}