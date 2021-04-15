import java.util.ArrayList;
import java.util.Collections;

// 哈密尔顿回路问题：
// 基于压缩状态的实现
// 通过位运算的状态压缩有一个限制，比如我们在64位系统用64位系统的int型最大可以表示63位（有一位符号位）
// 从某个顶点开始，访问图中所有顶点，且每个顶点只访问一次，最终可以回到最开始的顶点
public class HamiltonLoopByZipStatus {
    Graph g;
    private int[] pre;
    private int end;

    public HamiltonLoopByZipStatus(Graph g) {
        this.g = g;
        pre = new int[g.V()];
        end = -1;

        int visited = 0;
        dfs(visited, 0, 0, g.V());
    }

    private boolean dfs(int visited, int v, int parent, int left) {
        visited += (1 << v); // 将第v位设置为1
        pre[v] = parent;
        left--;

        // 第二种写法
        if (left == 0 && g.hasEdge(v, 0)) {
            end = v;
            return true;
        }
        for (int w: g.adj(v)) 
            if ((visited & (1 << w)) == 0)   // 判断第w位是否==0 
                if (dfs(visited, w, v, left)) return true;

        visited -= (1 << v); // 将第v设置回0
        return false;
    }

    private Iterable<Integer> paths() {
        ArrayList<Integer> res = new ArrayList<>();
        if (end == -1) return res;
        int cur = end;
        while(cur !=0) {
            res.add(cur);
            cur = pre[cur];
        }
        res.add(0);
        Collections.reverse(res);
        return res;
    }

    public static void main(String argv[]) {
        Graph g = new Graph("./Graph/hamiltongraph.txt");
        HamiltonLoopByZipStatus dfs = new HamiltonLoopByZipStatus(g);
        System.out.println(dfs.paths());

        Graph g2 = new Graph("./Graph/hamiltongraph02.txt");
        HamiltonLoopByZipStatus dfs2 = new HamiltonLoopByZipStatus(g2);
        System.out.println(dfs2.paths());
    }
}