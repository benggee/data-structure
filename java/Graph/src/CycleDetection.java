import java.util.ArrayList;

// 检测图是否有环
public class CycleDetection {
    Graph g;
    private boolean[] visited;
    private boolean isCycle;

    public CycleDetection(Graph g) {
        this.g = g;
        visited = new boolean[g.V()];
        for (int v = 0; v < g.V(); v++) 
            if (!visited[v])
                if (dfs(v, v)) {
                    isCycle = true;
                    break;
                } 
    }

    private boolean dfs(int v, int parent) {
        visited[v] = true;
        for (int w: g.adj(v)) {
            if (!visited[w]) {
                if (dfs(w, v)) 
                    return true;
            } else if (w != parent) {
                return true;
            }
        }
        return false;
    }

    public boolean isCycle() {
        return isCycle;
    }

    public static void main(String argv[]) {
        Graph g = new Graph("./Graph/g.txt");
        CycleDetection cyc = new CycleDetection(g);
        System.out.println(cyc.isCycle());
       

        Graph g2 = new Graph("./Graph/g2.txt");
        CycleDetection cyc2 = new CycleDetection(g2);
        System.out.println(cyc2.isCycle());
    }
}