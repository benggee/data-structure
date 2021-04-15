import java.util.ArrayList;
import java.util.Queue;
import java.util.LinkedList;

public class GraphBFS {

    private boolean[] visited;
    private ArrayList<Integer> order = new ArrayList<>();
    private Graph G;

    public GraphBFS(Graph G) {
        this.G = G;
        visited = new boolean[G.V()];
        for (int v = 0; v<G.V(); v++) {
            if (!visited[v])
                bfs(v);
        }
    }

    public void bfs(int s) {
        Queue<Integer> queue = new LinkedList<>();
        visited[s] = true;
        queue.add(s);

        while(!queue.isEmpty()) {
            int v = queue.remove();
            order.add(v);

            for (int w: G.adj(v)) {
                if (!visited[w]) {
                    queue.add(w);
                    visited[w] = true;
                }
            }
        }
    }

    public Iterable<Integer> order() {
        return order;
    }


    public static void main(String argc[]) {
        Graph g = new Graph("graph/g.txt");
        GraphBFS graphBFS = new GraphBFS(g);
        System.out.println("BFS result: "+ graphBFS.order());
    }
}