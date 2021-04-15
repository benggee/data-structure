import java.util.ArrayList;
import java.util.Collections;
import java.util.LinkedList;
import java.util.Queue;

// 深度优先遍历实现拓扑排序
// 深度优先不能检测是否有环
public class TopoSortByDFS {
    private Graph G; 
    private ArrayList<Integer> res;
    private boolean isCycle;

    public TopoSortByDFS(Graph g) {
        this.G = g;
        if (!G.isDirection()) throw new RuntimeException("Topo sort only work on direction graph.");
        
        res = new ArrayList<>();

        isCycle = (new DirectionCycleDetection(g)).isCycle();

        GraphDFS gd = new GraphDFS(g);
        for (int w: gd.post()) {
            res.add(w);
        }
        Collections.reverse(res);
    }

    public boolean isCycle() {
        return isCycle;
    }

    public Iterable<Integer> result() {
        return res;
    }
    
    public static void main(String args[]) {
        Graph g = new Graph("DirectionGraph/ug.txt", true);
        TopoSortByDFS ts = new TopoSortByDFS(g);
        System.out.println(ts.result());
    }
}