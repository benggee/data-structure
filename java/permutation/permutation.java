import java.util.ArrayList;
import java.util.Arrays;


public class permutation {
    ArrayList<ArrayList<Integer>> res = new ArrayList<>();
    int num = 0;

    public ArrayList<ArrayList<Integer>> compile(ArrayList<Integer> nums, int k) {
        num = k;
        permutate(nums, new ArrayList<Integer>());
        return res;
    }

    // unPermutateNums表示未被排列的数字
    // permutatedNums表示已经排列的数字
    private void permutate(ArrayList<Integer> unPermutateNums, ArrayList<Integer> permutatedNums) {
        if (unPermutateNums.size() == 0) {
            res.add(permutatedNums);
            return;
        }

        for (int i = 0; i < unPermutateNums.size(); i++) {
            ArrayList<Integer> tmp = (ArrayList<Integer>)(permutatedNums.clone());
            tmp.add(unPermutateNums.get(i));

            ArrayList<Integer> res = (ArrayList<Integer>)unPermutateNums.clone();
            res.remove(i);

            permutate(res, tmp);
        }
    }


    public static void main(String[] argc) {
        int[] a = new int[2];
        System.out.println(a);

        ArrayList<Integer> nums = new ArrayList<>(Arrays.asList(1,2,3,4));

        permutation p = new permutation();
        ArrayList<ArrayList<Integer>> res = p.compile(nums, 2);
        System.out.println(res);
    }
}


