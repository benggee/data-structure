import java.util.ArrayList;

public class GraphDFS {
    Graph g;
    private ArrayList<Integer> pre = new ArrayList<>();
    private ArrayList<Integer> post = new ArrayList<>();
    private boolean[] visited;

    public GraphDFS(Graph g) {
        this.g = g;
        visited = new boolean[g.V()];
        for (int v = 0; v < g.V(); v++) 
            if (!visited[v])
                dfs(v);
    }

    private void dfs(int v) {
        visited[v] = true;
        // 深度优先前序遍历 
        pre.add(v);

        for (int w: g.adj(v)) 
            if (!visited[w])
                dfs(w);
        // 深度优先后序遍历 
        post.add(v);
    }

    public Iterable<Integer> pre() {
        return pre;
    }

    public Iterable<Integer> post() {
        return post;
    }

    public static void main(String argv[]) {
        Graph g = new Graph("./Graph/g.txt");
        GraphDFS dfs = new GraphDFS(g);
        System.out.println(dfs.pre());
        System.out.println(dfs.post());
    }
}