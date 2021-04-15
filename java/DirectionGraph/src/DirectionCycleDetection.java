import java.util.ArrayList;

// 检测图是否有环
// DAG(Direction Acyclic Graph)  有向无环图
public class DirectionCycleDetection {
    Graph g;
    private boolean[] visited;
    private boolean[] inPath;   // 检测是否有环
    private boolean isCycle;

    public DirectionCycleDetection(Graph g) {
        this.g = g;
        if (!g.isDirection()) throw new RuntimeException("Cycel detection only world on undirection.");
        inPath = new boolean[g.V()];
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
        inPath[v] = true;
        for (int w: g.adj(v)) {
            if (!visited[w]) {
                if (dfs(w, v)) 
                    return true;
            } else if (inPath[w]) {
                return true;
            }
        }
        inPath[v] = false;
        return false;
    }

    public boolean isCycle() {
        return isCycle;
    }

    public static void main(String argv[]) {
        Graph g = new Graph("./DirectionGraph/ug.txt", true);
        DirectionCycleDetection cyc = new DirectionCycleDetection(g);
        System.out.println(cyc.isCycle());
    }
}