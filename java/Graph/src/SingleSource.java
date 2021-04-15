import java.util.ArrayList;
import java.util.Collections;

// 求单源路径，即从一个固定点到指定点的路径（这里的路径是深度优先求出来的路径）
// 思想框架：
// 新加一个arrayList记录当前顶点上一个项点，例如：
// 1->2->5->6
// 对应pre ==> [6]=5  [5]=2  [2]=1
public class SingleSource {
    Graph g;
    int s;
    private boolean[] visited;
    private int[] pre;

    public SingleSource(Graph g, int s) {
        g.validateVertex(s);

        this.g = g;
        this.s = s;

        visited = new boolean[g.V()];
        pre = new int[g.V()];
        for (int i=0; i<g.V(); i++) 
            pre[i] = -1;

        dfs(s, s);
    }

    private void dfs(int v, int parent) {
        visited[v] = true;
        pre[v] = parent;
        for (int w: g.adj(v)) 
            if (!visited[w])
                dfs(w, v);
    }

    public boolean isConnected(int s) {
        g.validateVertex(s);
        return visited[s];
    }

    public Iterable<Integer> paths(int t) {
        ArrayList<Integer> sspath = new ArrayList<>();
        if (!isConnected(t))
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
        SingleSource sspath = new SingleSource(g, 0);
        System.out.println("0 -> 5: " + sspath.paths(5));
        System.out.println("0 -> 6: " + sspath.paths(6));
    }
}