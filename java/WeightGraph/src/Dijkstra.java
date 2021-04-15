import java.util.Arrays;

// Kijkstra算法基本实现
// 算法的基本步骤：
// 1. 找到当前没有访问到的最短路径的节点
// 2. 确认这个节点的最短路径就是当前大小 
// 3. 根据这个节点的最短路径大小， 更新其他节点的路径长度
public class Dijkstra {
    private WeightGraph G;
    private int s;
    private boolean[] visited;
    private int[] dis;

    public Dijkstra(WeightGraph g, int s) {
        this.G = g;
        g.validateVertex(s);
        this.s = s;

        visited = new boolean[G.V()];
        dis = new int[G.V()];
        Arrays.fill(dis, Integer.MAX_VALUE);
        dis[s] = 0;

        while(true) {
            int curdis = Integer.MAX_VALUE;
            int cur = -1; // 当前最后访问到的顶点

            for (int v=0; v < G.V(); v++) {
                if (!visited[v] && dis[v] < curdis) {
                    curdis = dis[v];
                    cur = v;
                }
            }

            if (cur == -1) {
                break;
            }

            visited[cur] = true;
            for (int w: G.adj(cur)) {
                if (!visited[w]) {
                    if (dis[cur] + G.getWeight(cur, w) < dis[w]) {
                        dis[w] = dis[cur] + G.getWeight(cur, w);
                    }
                }
            }
        }

    }

    // 是否和指定顶点连通
    public boolean isConnectedTo(int v) {
        G.validateVertex(v);
        return visited[v];
    }


    // 到某个顶点的最短路径
    public int distTo(int v) {
        G.validateVertex(v);
        return dis[v];
    }


    public static void main(String args[]) {
        WeightGraph g = new WeightGraph("WeightGraph/g.txt");
        Dijkstra d = new Dijkstra(g, 1);
        System.out.println("is connect:"+ d.isConnectedTo(5));
        for (int i=0; i<d.dis.length; i++) {
            System.out.print(" " + d.dis[i]);
        }
        System.out.println();
        System.out.println("dis path:" + d.distTo(5));
    }
}