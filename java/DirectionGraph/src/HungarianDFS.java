import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.Queue;

// 匈牙利算法，求最大匹配数
// 深度优先遍历的实现
public class HungarianDFS {
    private Graph G;
    private int maxMatching = 0;
    private int[] matching;
    private boolean[] visited;

    public HungarianDFS(Graph g) {
        BipartitionDetection bd = new BipartitionDetection(g);
        if (!bd.isBipart()) 
            throw new IllegalArgumentException("Hungarian only works for bipartition graph.");
        this.G = g;

        int[] colors = bd.getColors();

        visited = new boolean[G.V()];

        matching = new int[G.V()];
        Arrays.fill(matching, -1);
        for (int v=0; v<G.V(); v++) {
            if (colors[v] == 0 && matching[v] == -1) {
                Arrays.fill(visited, false); // 重置visited， 每一轮循环重新记录
                if (dfs(v)) maxMatching++;
            }
        }
    }

    private boolean dfs(int v) {
        visited[v] = true;
        for (int w: G.adj(v)) {
            if(!visited[w]) {
                visited[w] = true;
                if (matching[w] == -1 || dfs(matching[w])) {
                    matching[w] = v;
                    matching[v] = w;
                    return true;
                }
            }
        }
        return false;
    }

    public int maxMatching() {
        return maxMatching;
    }

    public boolean isPerfactMatching() {
        return maxMatching * 2 == G.V();
    }

    public static void main(String args[]) {
        Graph g = new Graph("DirectionGraph/g2.txt");
        HungarianDFS hungarian = new HungarianDFS(g);
        System.out.println(hungarian.maxMatching());

        Graph g2 = new Graph("DirectionGraph/g3.txt");
        HungarianDFS hungarian2 = new HungarianDFS(g2);
        System.out.println(hungarian2.maxMatching());
    }
}