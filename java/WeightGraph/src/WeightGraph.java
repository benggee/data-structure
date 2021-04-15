
import java.io.File;
import java.io.IOException;
import java.util.Scanner;
import java.util.TreeMap;
import java.util.Map;

/**
 * 带权图建图
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

    public WeightGraph(String filename) {
        File file = new File(filename);
        try(Scanner scanner = new Scanner(file)){
            V = scanner.nextInt();
            if (V < 0) 
                throw new IllegalArgumentException("V must be non-negative");

            adj = new TreeMap[V];
            for (int i = 0; i < V; i++) {
                adj[i] = new TreeMap<Integer,Integer>();
            }

            E = scanner.nextInt();
            if (E < 0)
                throw new IllegalArgumentException("E must be non-negative");

            for (int i=0; i < E; i++) {
                int a = scanner.nextInt();
                validateVertex(a);
                int b = scanner.nextInt(); 
                validateVertex(b);
                int weight = scanner.nextInt();

                if (a == b)
                    throw new IllegalArgumentException("Self Loop is Detected.");

                if (adj[a].containsKey(b)) 
                    throw new IllegalArgumentException("Parallel Edges are Detected.");

                adj[a].put(b, weight);
                adj[b].put(a, weight);
            }
        } catch(IOException e) {
            e.printStackTrace();
        }

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

    // 获取一个项点的度（有多少条邻边）
    public int degree(int v) {
        validateVertex(v);
        return adj[v].size();
    }

    public void validateVertex(int v) {
        if (v < 0 || v >= V) 
            throw new IllegalArgumentException("vertex "+v+" is invalid.");
    }

    public void removeEdge(int v, int w) {
        validateVertex(v);
        validateVertex(w);

        if (adj[v].contains(w)) E--;

        adj[v].remove(w);
    }

    @Override
    public Object clone() {
        try {
            Graph cloned = (Graph) super.clone();
            cloned.adj = new TreeSet[V];
            for (int v=0; v < V; v++) {
                cloned.adj[v] = new TreeSet<Integer>();
                for (int w: adj[v]) {
                    cloned.adj[v].add(w);
                }
            }
            return cloned;
        } catch (CloneNotSupportedException e) {
            e.printStackTrace();
        }
        return null;
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