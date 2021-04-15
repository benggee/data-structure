import java.util.HashSet;

// 寻找图中的割点
public class FindCutPoints {
    private Graph G;
    private boolean[] visited;

    private int[] ord;  // 第几个被遍历的
    private int[] low;  // 记录当前顶点能到达的最小的ord对应的值
    private HashSet<Integer> res;  // 找到的桥
    private int count;

    public FindCutPoints(Graph g) {
        this.G = g;
        visited = new boolean[G.V()];
        ord = new int[G.V()];
        low = new int[G.V()];
        res = new HashSet<>();
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

        int childByV = 0;

        for (int w: G.adj(v)) {
            if (!visited[w]) {
                dfs(w, v);
                low[v] = Math.min(low[v], low[w]);
                
                if (v != parent && low[w] >= ord[v]) {
                    res.add(v);
                }

                childByV ++;
                // 如果是根节点，且孩子节点大于1（注意这个孩子节点表示的不是有邻边数，而是DFS遍历树）
                if (v == parent && childByV > 1) {
                    res.add(v);
                }   
            } else if (w != parent) {
                low[v] = Math.min(low[v], low[w]);
            }
        }

    }

    public HashSet<Integer> result() {
        return res;
    }

    public static void main(String args[]) {
        Graph g = new Graph("Graph/bridge01.txt");
        FindCutPoints fcp = new FindCutPoints(g);
        System.out.println(fcp.result());


        Graph g2 = new Graph("Graph/bridge02.txt");
        FindCutPoints fcb2 = new FindCutPoints(g2);
        System.out.println(fcb2.result());

        Graph g3 = new Graph("Graph/bridge03.txt");
        FindCutPoints fcb3 = new FindCutPoints(g3);
        System.out.println(fcb3.result());

        Graph g4 = new Graph("Graph/bridge04.txt");
        FindCutPoints fcb4 = new FindCutPoints(g4);
        System.out.println(fcb4.result());

        Graph tree = new Graph("Graph/tree.txt");
        FindCutPoints fcb5 = new FindCutPoints(tree);
        System.out.println(fcb5.result());
    }

}