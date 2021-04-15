import java.util.Arrays;
import java.util.ArrayList;
import java.util.Collections;

public class BellmanFord {
    private WeightGraph G;
    private int[] dis;
    private int[] pre;
    private int s;
    private boolean isNegativeCycle;

    public BellmanFord(WeightGraph g, int s) {
        this.G = g;
        g.validateVertex(s);
        this.s = s;

        this.isNegativeCycle = false;

        dis = new int[g.V()];
        Arrays.fill(dis, Integer.MAX_VALUE);
        dis[s] = 0;

        pre = new int[g.V()];
        Arrays.fill(pre, -1);
        pre[s] = s;
        
        // 进行V-1转的松弛操作
        for (int p = 1; p < g.V(); p ++) {
            for (int v = 0; v < g.V(); v++) {
                for (int w: g.adj(v)) {
                    if (dis[v] != Integer.MAX_VALUE && dis[v] + G.getWeight(v, w) < dis[w]) {
                        dis[w] = dis[v] + G.getWeight(v, w);
                        pre[w] = v;
                    }
                }
            }
        }

        // 判断是否有负权环
        for (int v = 0; v < g.V(); v++) {
            for (int w: g.adj(v)) {
                if (dis[v] != Integer.MAX_VALUE && dis[v] + G.getWeight(v, w) < dis[w]) {
                    isNegativeCycle = true;
                }
            }
        }
    }

    // 最短路径长度
    public int disTo(int v) {
        G.validateVertex(v);
        if (!isNetCycle()) throw new RuntimeException("exsits net cycle.");
        return dis[v];
    }

    // 是否能达
    public boolean isConnectTo(int v) {
        G.validateVertex(v);
        return dis[v] != Integer.MAX_VALUE;  // 如果还是max-value说明 没有被访问过
    }

    // 是否有负权环
    public boolean isNetCycle() {
        return isNegativeCycle;
    }

    public Iterable<Integer> path(int t) {
        G.validateVertex(t);
        ArrayList<Integer> res = new ArrayList<>();
        int cur = t;
        while(cur != s) {
            res.add(cur);
            cur = pre[cur];
        }
        Collections.reverse(res);
        return res;
    }


    public static void main(String args[]) {
        WeightGraph g = new WeightGraph("WeightGraph/g.txt");
        BellmanFord b = new BellmanFord(g, 0);
        if (!b.isNegativeCycle) {
            for (int i=0; i<b.dis.length; i++) {
                System.out.print(" " + b.dis[i]);
            }
            System.out.println();

            System.out.println(b.path(4));
        }
        
    }
}