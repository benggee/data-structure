import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.PriorityQueue;
import java.util.Queue;

// 基于优先队列对Dijkstra的实现
// 时间复杂度：O(V*E)
public class DijkstraByPriorityQueue {
    private WeightGraph G;
    private int s;
    private boolean[] visited;
    private int[] dis;  // 保存的是从原点开始遍历到的路径顶点对应的最短路径
    private int[] pre;  // 最短路径 


    private class Node implements Comparable<Node> {
        private int v,dis;
        public Node(int v, int dis) {
            this.v = v;
            this.dis = dis;
        }

        @Override
        public int compareTo(Node o) {
            return dis-o.dis;
        }
    }

    public DijkstraByPriorityQueue(WeightGraph g, int s) {
        this.G = g;
        g.validateVertex(s);
        this.s = s;

        visited = new boolean[g.V()];
        dis = new int[g.V()];
        Arrays.fill(dis, Integer.MAX_VALUE);
        pre = new int[g.V()];
        Arrays.fill(pre, -1);
        pre[s] = 0;

        Queue<Node> pq = new PriorityQueue<Node>();
        // 第一次的时候路径长度为0
        pq.add(new Node(s, 0));
        dis[s] = s;

        while(!pq.isEmpty()) {
            int cur = pq.remove().v;
            if (visited[cur]) continue;

            visited[cur] = true;
            for (int w: G.adj(cur)) {
                if (!visited[w]) {
                    if (dis[cur] + G.getWeight( cur, w) < dis[w]) {
                        dis[w] = dis[cur] + G.getWeight( cur, w);
                        pq.add(new Node(w, dis[w]));
                        pre[w] = cur; // 记录最短路径 
                    }
                }
            }
        }
    }

    // 返回是否连通
    public boolean isConnectTo(int v) {
        G.validateVertex(v);
        return visited[v];
    }

    // 返回最短路径长度
    public int disTo(int v) {
        G.validateVertex(v);
        return dis[v];
    }

    public int[] result() {
        return dis;
    }
 
    // 获取最短路径 
    public Iterable<Integer> disPath(int v) {
        ArrayList<Integer> res = new ArrayList<>();
        if (!isConnectTo(v)) return res;
        int cur = v;
        while(cur != s) {
            res.add(cur);
            cur = pre[cur];
        }

        Collections.reverse(res);
        return res;
    }

    public static void main(String args[]) {
        WeightGraph g = new WeightGraph("WeightGraph/g.txt");
        DijkstraByPriorityQueue d = new DijkstraByPriorityQueue(g, 0);
        for (int i = 0; i<d.dis.length; i++) {
            System.out.print(" "+d.dis[i]);
        }
        System.out.println();
        System.out.println("dis len:" + d.disTo(4));

        // 获取最短路径 
        System.out.println("dis path:" + d.disPath(4));
    }
}