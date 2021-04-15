import java.util.ArrayList;
import java.util.Queue;
import java.util.LinkedList;
import java.util.Collections;

// 广度优化求解单源路径 
public class SingleSourceBFS {

    private boolean[] visited;
    private int[] pre;
    private Graph G;
    private int s;

    public SingleSourceBFS(Graph G, int s) {
        this.G = G;
        this.s = s;

        pre = new int[G.V()];
        for (int i=0; i<G.V(); i++) {
            pre[i] = -1;
        }

        visited = new boolean[G.V()];

        bfs(s, s);
    }

    public void bfs(int s, int parent) {
        Queue<Integer> queue = new LinkedList<>();
        visited[s] = true;
        queue.add(s);
        pre[s] = parent;

        while(!queue.isEmpty()) {
            int v = queue.remove();

            for (int w: G.adj(v)) {
                if (!visited[w]) {
                    queue.add(w);
                    visited[w] = true;
                    pre[w] = v;
                }
            }
        }
    }

    public boolean isConnectTo(int t) {
        G.validateVertex(t);
        return visited[t];
    }

    public Iterable<Integer> path(int t) {
        ArrayList<Integer> ret = new ArrayList<>();
        if (!isConnectTo(t)) 
            return ret;

        int cur = t;
        while(cur != s) {
            ret.add(cur);
            cur = pre[cur];
        }
        ret.add(s);

        Collections.reverse(ret);
        return ret;
    }

    public static void main(String argc[]) {
        Graph g = new Graph("graph/g.txt");
        SingleSourceBFS singleSourceBFS = new SingleSourceBFS(g, 0);
        System.out.println("BFS source: "+singleSourceBFS.path(6));
    }
}