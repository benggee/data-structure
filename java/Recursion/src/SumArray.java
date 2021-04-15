public class SumArray {
    
    public int sum(int[] arr) {
        return this.sum(arr, 0);
    }

    // 这个问题的递推公式可以写成下面这样：
    // f(N) = arr[N] + arr[N+1]
    // N != arr.length 
    public int sum(int[] arr, int level) {
        if (level == arr.length) 
            return 0;
        return arr[level] + sum(arr, level + 1);
    }

    public static void main(String argv[]) {
        int[] arr = {1,2,3,4,5,6,7,8,9};
        SumArray sum = new SumArray();
        int ret = sum.sum(arr);
        System.out.println("Result:"+ ret);
     }
}