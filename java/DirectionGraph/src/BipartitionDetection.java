
// 判断是否是二分图
public class BipartitionDetection {
    Graph g;
    private boolean[] visited;
    private int[] colors;
    private boolean isBipart = true;  // 是否是二分图

    public BipartitionDetection(Graph g) {
        this.g = g;
        colors = new int[g.V()];
        for (int i=0; i<g.V(); i++) {
            colors[i] = -1;
        }

        visited = new boolean[g.V()];
        for (int v = 0; v < g.V(); v++) {
            if (!visited[v]) {
                if (!dfs(v, 0)) {
                    isBipart = false;
                    break;
                }
            }
        }
    }

    private boolean dfs(int v, int color) {
        visited[v] = true;
        colors[v] = color;
        for (int w: g.adj(v)) {
            if (!visited[w]) {
                if (!dfs(w, 1-color)) return false;
            } else {
                if (colors[w] == colors[v]) return false;
            }
        }
        return true;
    }

    public int[] getColors() {
        return colors;
    }

    public boolean isBipart() {
        return isBipart;
    }

    public static void main(String argv[]) {
        Graph g = new Graph("./Graph/g.txt");
        BipartitionDetection bipartitionDetection = new BipartitionDetection(g);
        System.out.println(bipartitionDetection.isBipart());
    }
}