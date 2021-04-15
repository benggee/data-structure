
import java.io.File;
import java.io.IOException;
import java.util.Scanner;
import java.util.TreeMap;
import java.util.Map;

/**
 * 带权图（带权、有向图）
 * 红黑树的实现方式（图的最终实现方式）
 * 在以下 V表示顶点数 E表示图的边数
 * 空间复杂度 O(V + E)   这里V和E都是必要的，极端情况下可能E=0
 * 建图时间   O(ElogV) 
 * 两点是否相邻 O(logV) 
 * 查找所有邻边 O(degree(V)) 默认为顶点的度， 如果是完全图或者稠密图接近O(V)
 */
class WeightGraph {
    private int V; // 图的顶点数
    private int E; // 图的边数
    private TreeMap<Integer, Integer>[] adj; // 图方隈
    private boolean direction = false ;

    public WeightGraph(String filename, boolean direction) {
        this.direction = direction;
        File file = new File(filename);
        try(Scanner scanner = new Scanner(file)){
            V = scanner.nextInt();
            if (V < 0) 
                throw new IllegalArgumentException("V must be non-negative");

            adj = new TreeMap[V];
            for (int i = 0; i < V; i++) {
                adj[i] = new TreeMap<Integer,Integer>();
            }

            int e = scanner.nextInt();
            if (e < 0)
                throw new IllegalArgumentException("E must be non-negative");

            for (int i=0; i < e; i++) {
                int a = scanner.nextInt();
                int b = scanner.nextInt(); 
                int weight = scanner.nextInt();

                addEdge(a, b, weight);
            }
        } catch(IOException e) {
            e.printStackTrace();
        }

    }

    public WeightGraph(String filename) {
        this(filename, false);
    }

    public WeightGraph(int V, boolean direction) {
        this.V = V;
        this.direction = direction;
        this.E = 0;

        adj = new TreeMap[V];
        for (int i=0; i<V; i++) {
            adj[i] = new TreeMap<Integer, Integer>();
        }
    }

    public void setWeight(int v, int w, int weight) {
        if (!hasEdge(v, w)) throw new IllegalArgumentException(String.format("No edge %d-%d", v, w));

        adj[v].put(w, weight);
        if (!direction) 
            adj[w].put(v, weight);
    }

    public void addEdge(int a, int b, int v) {
        validateVertex(a);
        validateVertex(b);

        if (a == b)
        throw new IllegalArgumentException("Self Loop is Detected.");

        if (adj[a].containsKey(b)) 
            throw new IllegalArgumentException("Parallel Edges are Detected.");

        adj[a].put(b, v);
        if (!direction) 
            adj[b].put(a, v);
        E++;
    }

    public int V() {
        return V;
    }

    public int E() {
        return E; 
    }

    // 两个顶点是否有边
    public  boolean hasEdge(int v, int w) {
        validateVertex(v);
        validateVertex(w);
        return adj[v].containsKey(w);
    }

    // 获取一个顶点的邻边
    public Iterable<Integer> adj(int v) {
        validateVertex(v);
        return adj[v].keySet();
    }

    // 获取边的权值
    public int getWeight(int v, int w) {
        if (hasEdge(v, w))
            return adj[v].get(w);
        throw new IllegalArgumentException("v and w is out of range.");
    }

    public void validateVertex(int v) {
        if (v < 0 || v >= V) 
            throw new IllegalArgumentException("vertex "+v+" is invalid.");
    }

    // 是否是有向图
    public boolean isDirection() {
        return direction;
    }

    public void removeEdge(int v, int w){
        validateVertex(v);
        validateVertex(w);

        if(adj[v].containsKey(w)) E --;

        adj[v].remove(w);
        if(!direction)
            adj[w].remove(v);
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append(String.format("V = %d, E = %d\n", V, E));

        for (int v=0; v < V; v++) {
            sb.append(String.format("%d: ", v));
            for (Map.Entry<Integer, Integer> entry: adj[v].entrySet()) {
                sb.append(String.format("(%d: %d)", entry.getKey(), entry.getValue()));
            }
            sb.append("\n");
        }
        return sb.toString();
    }

    public static void main(String[] args) {
        WeightGraph adj = new WeightGraph("./WeightGraph/g.txt");
        System.out.println(adj);
    }

}