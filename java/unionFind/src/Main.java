import java.util.Random;

public class Main{
    private static double testUF(UF uf, int num) {
        int size = uf.size();
        Random random = new Random();
        long startTime = System.nanoTime();

        for (int i=0; i<num; i++) {
            int a = random.nextInt(size);
            int b = random.nextInt(size);
            uf.union(a, b);
        }

        for (int i=0; i<num; i++) {
            int a = random.nextInt(size);
            int b = random.nextInt(size);
            uf.unionContains(a, b);
        }

        long endTime = System.nanoTime();

        return (endTime - startTime) / 1000000000.0;
    }

    public static void main(String[] argv) {
        int size = 10000000;
        int m = 10000000;

        // UnionFindv1 u1 = new UnionFindv1(size);
        // System.out.println("Unionfind1: " + testUF(u1, m));

        // UnionFindv2 u2 = new UnionFindv2(size);
        // System.out.println("Unionfind2: " + testUF(u2, m));

        UnionFindv3 u3 = new UnionFindv3(size);
        System.out.println("Unionfind3: " + testUF(u3, m));

        UnionFindv4 u4 = new UnionFindv4(size);
        System.out.println("Unionfind4: " + testUF(u4, m));

        UnionFindv5 u5 = new UnionFindv5(size);
        System.out.println("Unionfind5: " + testUF(u5, m));

        UnionFindv6 u6 = new UnionFindv6(size);
        System.out.println("Unionfind5: " + testUF(u6, m));
    }


}