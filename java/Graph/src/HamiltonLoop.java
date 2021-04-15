import java.util.ArrayList;
import java.util.Collections;

// 哈密尔顿回路问题：
// 从某个顶点开始，访问图中所有顶点，且每个顶点只访问一次，最终可以回到最开始的顶点
public class HamiltonLoop {
    Graph g;
    private boolean[] visited;
    private int[] pre;
    private int end;

    public HamiltonLoop(Graph g) {
        this.g = g;
        visited = new boolean[g.V()];
        pre = new int[g.V()];
        end = -1;
        dfs(0, 0, g.V());
    }

    private boolean dfs(int v, int parent, int left) {
        visited[v] = true;
        pre[v] = parent;
        left--;

        // 第二种写法
        if (left == 0 && g.hasEdge(v, 0)) {
            end = v;
            return true;
        }
        for (int w: g.adj(v)) 
            if (!visited[w]) 
                if (dfs(w, v, left)) return true;

        // 第一种写法
        // for (int w: g.adj(v)) 
        //     if (!visited[w]) {
        //         if (dfs(w, v, left)) return true;
        //     } else if (w == 0 && left == 0()) {
        //         end = v;
        //         return true;
        //     }

        visited[v] = false;  // 如果没有找到需要找其它的相邻顶点，进行回溯
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
        HamiltonLoop dfs = new HamiltonLoop(g);
        System.out.println(dfs.paths());

        Graph g2 = new Graph("./Graph/hamiltongraph02.txt");
        HamiltonLoop dfs2 = new HamiltonLoop(g2);
        System.out.println(dfs2.paths());
    }
}