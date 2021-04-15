import java.util.ArrayList;
import java.util.PriorityQueue;
import java.util.Queue;

public class Prim {
    private WeightGraph G;
    private ArrayList<WeightGraphEdges> mst;

    public Prim(WeightGraph g) {
        this.G = g;
        mst = new ArrayList<WeightGraphEdges>();
        CC1 c = new CC1(g);
        if (c.getCCount()>1) {
            return;
        }

        // Prim算法
        boolean[] visited = new boolean[G.V()];
        visited[0] = true;

        // 优先队列的实现
        // 时间复杂度：O(E*LogE)
        Queue<WeightGraphEdges> q = new PriorityQueue<WeightGraphEdges>(); // java里优先队列默认是小顶堆
        for (int w: G.adj(0)) {
            q.add(new WeightGraphEdges(0, w, G.getWeight(0, w)));
        }

        while(!q.isEmpty()) {
            WeightGraphEdges minEdge = (WeightGraphEdges) q.remove();
            if (visited[minEdge.V()] && visited[minEdge.W()]) 
                continue;
            
            mst.add(minEdge);

            int newW = visited[minEdge.V()] ? minEdge.W() : minEdge.V();
            visited[newW] = true;
            for (int w: G.adj(newW)) {
                if (!visited[w]) {
                    q.add(new WeightGraphEdges(newW, w, G.getWeight(newW, w)));
                }
            }
        }


        // 暴力法实现
        // 时间复杂度 O(V*E)
        // for (int i=1; i<G.V(); i++) {
        //     WeightGraphEdges minEdge = new WeightGraphEdges(-1, -1, Integer.MAX_VALUE);
        //     for (int v=0; v<G.V(); v++) {
        //         if (visited[v]) {
        //             for (int w: G.adj(v)) {
        //                 if (!visited[w] && G.getWeight(v, w) < minEdge.weight()) {
        //                     minEdge = new WeightGraphEdges(v, w, G.getWeight(v, w));
        //                 }
        //             }
        //         }
        //     }
        //     mst.add(minEdge);
        //     visited[minEdge.V()] = true;
        //     visited[minEdge.W()] = true;
        // }
    }

    public ArrayList<WeightGraphEdges> result() {
        return mst;
    }

    public static void main(String args[]) {
        WeightGraph g = new WeightGraph("WeightGraph/g.txt");
        Prim p = new Prim(g);
        System.out.println(p.result());
    }

}