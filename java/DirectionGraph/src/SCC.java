import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;

/**
 * 强联通分量求解
 * 1. 计算图的反图
 * 2. 对反图进行深度优先遍历 
 * 3. 深度后序遍历结果求逆
 * 查询图的分量个数（分量是指一个图中相互不相连的集合的个数）
 * 通过将ccount赋值给visited可以标识出不同学的连通分量里都有哪些元素
 */
public class SCC {
    Graph g;
    private int[] visited;
    private int sccount = 0;

    public SCC(Graph g) {
        if (!g.isDirection()) throw new IllegalArgumentException("SCC only work on direction graph");
        this.g = g;
        visited = new int[g.V()];
        Arrays.fill(visited, -1);
        
        GraphDFS gd = new GraphDFS(g.reverseGraph());
        ArrayList<Integer>  order = new ArrayList<>();
        for (int v: gd.post()) {
            order.add(v);
        }

        Collections.reverse(order);

        for (int v: order) {
            if (visited[v] == -1) {
                dfs(v, sccount);
                sccount++;
            }
        }
    }

    private void dfs(int v, int scid) {
        visited[v] = scid;
        for (int w: g.adj(v)) 
            if (visited[w] == -1)
                dfs(w, scid);
    }

    public Integer getSCCount() {
        return sccount;
    }

    public ArrayList<Integer>[] getComponents() {
        ArrayList<Integer>[] com = new ArrayList[sccount];
        for (int i=0; i<sccount; i++) {
            com[i] = new ArrayList<>();
        }
        
        for (int v=0; v<g.V(); v++) {
            com[visited[v]].add(v);
        }
        return com;
    }

    public static void main(String argv[]) {
        Graph g = new Graph("./DirectionGraph/ug2.txt", true);
        SCC cc = new SCC(g);
        System.out.println(cc.getSCCount());

        ArrayList<Integer>[] tem = cc.getComponents();
        for (int i=0; i<cc.getSCCount(); i++) {
            System.out.printf("%d: ",i);
            for (int w: tem[i]) {
                System.out.printf("%d ", w);
            }
            System.out.println();
        }
    }
}