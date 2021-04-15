import java.util.ArrayList;
import java.util.LinkedList;
import java.util.Queue;

// 拓扑排序一般只针对无权有向图有有意义 
public class TopoSort {
    private Graph G; 
    private ArrayList<Integer> res;
    private boolean isCycle;
    private int[] indegrees;

    public TopoSort(Graph g) {
        this.G = g;
        if (!G.isDirection()) throw new RuntimeException("Topo sort only work on direction graph.");
        
        res = new ArrayList<>();
        indegrees = new int[G.V()];

        Queue<Integer> q = new LinkedList<>();
        for (int v=0; v<G.V(); v++) {
            indegrees[v] = g.indegree(v);
            if (indegrees[v] == 0) {
                q.add(v);
            }
        }

        while(!q.isEmpty()) {
            int cur = q.remove();
            res.add(cur);

            for (int w: G.adj(cur)) {
                indegrees[w] --;
                if (indegrees[w] == 0) {
                    q.add(w);
                }
            }
        }

        // 检测到环
        if (res.size() != G.V()) {
            isCycle = true;
            res.clear();
        }
    }

    public boolean isCycle() {
        return isCycle;
    }

    public Iterable<Integer> result() {
        return res;
    }
    
    public static void main(String args[]) {
        Graph g = new Graph("DirectionGraph/ug.txt", true);
        TopoSort ts = new TopoSort(g);
        System.out.println(ts.result());
    }
}