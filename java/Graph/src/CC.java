import java.util.ArrayList;

/**
 * 查询图的分量个数（分量是指一个图中相互不相连的集合的个数）
 * 通过将ccount赋值给visited可以标识出不同学的连通分量里都有哪些元素
 */
public class CC {
    Graph g;
    private int[] visited;
    private int ccount = 0;

    public CC(Graph g) {
        this.g = g;
        visited = new int[g.V()];
        for (int i=0; i<g.V(); i++) {
            visited[i] = -1;
        }
        for (int v = 0; v < g.V(); v++) 
            if (visited[v] == -1) {
                dfs(v, ccount);
                ccount++;
            }
    }

    private void dfs(int v, int cid) {
        visited[v] = cid;
        for (int w: g.adj(v)) 
            if (visited[w] == -1)
                dfs(w, cid);
    }

    public Integer getCCount() {
        return ccount;
    }

    public ArrayList<Integer>[] getComponents() {
        ArrayList<Integer>[] com = new ArrayList[ccount];
        for (int i=0; i<ccount; i++) {
            com[i] = new ArrayList<>();
        }
        
        for (int v=0; v<g.V(); v++) {
            com[visited[v]].add(v);
        }
        return com;
    }

    public static void main(String argv[]) {
        Graph g = new Graph("./Graph/g.txt");
        CC cc = new CC(g);
        System.out.println(cc.getCCount());

        ArrayList<Integer>[] tem = cc.getComponents();
        for (int i=0; i<cc.getCCount(); i++) {
            System.out.printf("%d: ",i);
            for (int w: tem[i]) {
                System.out.printf("%d ", w);
            }
            System.out.println();
        }
    }
}