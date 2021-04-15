import java.util.ArrayList;

// 寻找图中的桥
public class FindBrigdes {
    private Graph G;
    private boolean[] visited;

    private int[] ord;  // 第几个被遍历的
    private int[] low;  // 记录当前顶点能到达的最小的ord对应的值
    private ArrayList<BridgeEdge> res;  // 找到的桥
    private int count;

    public FindBrigdes(Graph g) {
        this.G = g;
        visited = new boolean[G.V()];
        ord = new int[G.V()];
        low = new int[G.V()];
        res = new ArrayList<>();
        count = 0;

        for (int v=0; v<G.V(); v++) 
            if (!visited[v])
                dfs(v, v);
    }

    private void dfs(int v, int parent) {
        visited[v] = true;
        ord[v] = count;
        low[v] = ord[v];
        count ++;

        for (int w: G.adj(v)) {
            if (!visited[w]) {
                dfs(w, v);
                low[v] = Math.min(low[v], low[w]);
                if (low[w] > low[v]) 
                    res.add(new BridgeEdge(v, w));
                    
            } else if (w != parent) {
                low[v] = Math.min(low[v], low[w]);
            }
        }

    }

    public ArrayList<BridgeEdge> result() {
        return res;
    }

    public static void main(String args[]) {
        Graph g = new Graph("Graph/bridge01.txt");
        FindBrigdes fb = new FindBrigdes(g);
        System.out.println(fb.result());


        Graph g2 = new Graph("Graph/bridge02.txt");
        FindBrigdes fb2 = new FindBrigdes(g2);
        System.out.println(fb2.result());

        Graph g3 = new Graph("Graph/bridge03.txt");
        FindBrigdes fb3 = new FindBrigdes(g3);
        System.out.println(fb3.result());

        Graph g4 = new Graph("Graph/bridge04.txt");
        FindBrigdes fb4 = new FindBrigdes(g4);
        System.out.println(fb4.result());

        Graph tree = new Graph("Graph/tree.txt");
        FindBrigdes fb5 = new FindBrigdes(tree);
        System.out.println(fb5.result());
    }

}