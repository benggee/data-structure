import java.util.ArrayList;
import java.util.Queue;
import java.util.Collections;
import java.util.LinkedList;
import java.util.Arrays;


// 网络最大流
// Ford-Fullkerson思想
// Edmonds-Karp 算法
// 网络最大流量的核心步骤：
// 1. 残量图建模
// 2. 广度优先求增文路径
// 3. 更新残量图的流量权值
// 4. 重复2、3步，直到没有增广路径
class MaxFlow {
    private WeightGraph net;
    private int s,t;
    
    private WeightGraph rG; 
    private int maxFlow; // 最大流

    public MaxFlow(WeightGraph net, int s, int t) {
        if (!net.isDirection())  throw new IllegalArgumentException("MaxFlow only work on direction graph.");
        if (net.V() < 2) throw new IllegalArgumentException("The graph v should least 2.");

        net.validateVertex(s);
        net.validateVertex(t);

        if (s == t) throw new IllegalArgumentException("s and t must be different.");

        this.net = net;
        this.s = s;
        this.t = t;

        this.rG = new WeightGraph(net.V(), true);
        for (int v=0; v<net.V(); v++) {
            for (int w: net.adj(v)) {
                rG.addEdge(v, w, net.getWeight(v, w));
                rG.addEdge(w, v, 0);
            }
        }

        while(true) {
            // 示一条增广路径
            ArrayList<Integer> augPath = getAugmentingPath();
            if (augPath.size() == 0) break;

            int f = Integer.MAX_VALUE;
            for (int i=1; i<augPath.size(); i++) {
                int v = augPath.get(i - 1);
                int w = augPath.get(i);
                f = Math.min(f, rG.getWeight(v, w));
            }
            maxFlow += f;

            for (int i = 1; i<augPath.size(); i++) {
                int v = augPath.get(i-1);
                int w = augPath.get(i);

                rG.setWeight(v, w, rG.getWeight(v, w) - f);
                rG.setWeight(w, v, rG.getWeight(w, v) + f);
            }
        }
    }

    private ArrayList<Integer> getAugmentingPath() {
        Queue<Integer> q = new LinkedList<>();
        int[] pre = new int[net.V()];
        Arrays.fill(pre, -1);

        q.add(s);
        pre[s] = s;
        while(!q.isEmpty()) {
            int cur = q.remove();
            if (cur == t) break;
            for (int w: rG.adj(cur)) {
                if (pre[w] == -1 && rG.getWeight(cur, w) > 0) {
                    pre[w] = cur;
                    q.add(w);
                }
            }
        }

        ArrayList<Integer> res = new ArrayList<>();
        if (pre[t] == -1) return res;

        int cur = t;
        while(cur != s){
            res.add(cur);
            cur = pre[cur];
        }
        res.add(s);

        Collections.reverse(res);
        return res;
    }

    // 获取网络最大流
    public int result() {
        return maxFlow;
    }
    
    public int flow(int v, int w) {
        if (!net.hasEdge(v, w))  throw new IllegalArgumentException(String.format("No edge %d-%d", v, w));
        return rG.getWeight(w, v);
    }


    public static void main(String args[]) {
        WeightGraph net = new WeightGraph("DirectionGraph/network1.txt", true);
        MaxFlow mf = new MaxFlow(net, 0, 3);
        System.out.println(mf.result());
        for (int v=0; v<net.V(); v++) 
            for (int w: net.adj(v)) 
                System.out.println(String.format("%d-%d: %d/%d", v, w, mf.flow(v, w), net.getWeight(v, w)));


        WeightGraph net2 = new WeightGraph("DirectionGraph/network2.txt", true);
        MaxFlow mf2 = new MaxFlow(net2, 0, 5);
        System.out.println(mf2.result());
        for (int v=0; v<net2.V(); v++) 
            for (int w: net2.adj(v)) 
                System.out.println(String.format("%d-%d: %d/%d", v, w, mf2.flow(v, w), net2.getWeight(v, w)));
    }

}