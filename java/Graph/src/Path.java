import java.util.ArrayList;
import java.util.Collections;

// 求单源路径，即从一个固定点到指定点的路径（这里的路径是深度优先求出来的路径）
// 思想框架：
// 新加一个arrayList记录当前顶点上一个项点，例如：
// 1->2->5->6
// 对应pre ==> [6]=5  [5]=2  [2]=1
public class Path {
    Graph g;
    int s;
    int t; 
    private boolean[] visited;
    private int[] pre;

    public Path(Graph g, int s, int t) {
        g.validateVertex(s);
        g.validateVertex(t);

        this.g = g;
        this.s = s;
        this.t = t;

        visited = new boolean[g.V()];
        pre = new int[g.V()];
        for (int i=0; i<g.V(); i++) 
            pre[i] = -1;

        dfs(s, s);
    }

    private boolean dfs(int v, int parent) {
        visited[v] = true;
        pre[v] = parent;
        if (v == t) {
            return true;
        }
        for (int w: g.adj(v)) 
            if (!visited[w])
                if (dfs(w, v)) 
                    return true;
        return false;
    }

    public boolean isConnected() {
        return visited[t];
    }

    public Iterable<Integer> paths() {
        ArrayList<Integer> sspath = new ArrayList<>();
        if (!isConnected())
            return sspath;
        int cur = t;
        while(cur != s ) {
            sspath.add(cur);
            cur = pre[cur];
        }
        Collections.reverse(sspath);
        return sspath;
    }

    public static void main(String argv[]) {
        Graph g = new Graph("./Graph/g.txt");
        Path path = new Path(g, 0, 6);
        System.out.println("0 -> 6: " + path.paths());

        Path path2 = new Path(g, 0, 5);
        System.out.println("0 -> 5: " + path2.paths());
    }
}