import java.util.Arrays;

// Floyed算法求做生意两点的最短路径 
// 时间复杂度：O(V*V*V)
public class Floyed {
    private WeightGraph G;
    private int[][] dis;
    private boolean hasNegativeCycle;

    public Floyed(WeightGraph g) {
        this.G = g;
        this.hasNegativeCycle = false;
        dis = new int[G.V()][G.V()];
        for (int i=0; i<G.V(); i++) {
            Arrays.fill(dis[i], Integer.MAX_VALUE);
        }

        for (int v=0; v<G.V(); v++) {
            dis[v][v] = 0;
            for (int w: G.adj(v)) {
                dis[v][w] = G.getWeight(v, w);
            }
        }

        for (int t=0; t<G.V(); t++) {
            for (int v=0; v<G.V(); v++) {
                for (int w=0; w<G.V(); w++) {
                    if (dis[v][t] != Integer.MAX_VALUE && dis[t][w] != Integer.MAX_VALUE && dis[v][t] + dis[t][w]  < dis[v][w]) {
                        dis[v][w] = dis[v][t] + dis[t][w];
                    }
                }
            }
        }

        // 判断是否存在负权环
        for (int v=0; v<G.V(); v++) {
            if (dis[v][v] < 0) {
                hasNegativeCycle = true;
            }
        }
    }

    public int disTo(int v, int w) {
        G.validateVertex(v);
        G.validateVertex(w);

        return dis[v][w];        
    }

    public boolean hasNetCycle() {
        return hasNegativeCycle;
    }

    public boolean isConnectedTo(int v, int w) {
        G.validateVertex(v);
        G.validateVertex(w);
        return dis[v][w] == Integer.MAX_VALUE;
    }

    public static void main(String args[]) {
        WeightGraph g = new WeightGraph("WeightGraph/g.txt");
        Floyed f = new Floyed(g);
        if (f.hasNegativeCycle) {
            System.out.println("exsits net cycle.");
        } else {
            for (int v=0; v<g.V(); v++) {
                for (int w=0; w<g.V(); w++) {
                    System.out.print(f.disTo(v, w) + " ");
                }
                System.out.println();
            }
        }
    }
    
}