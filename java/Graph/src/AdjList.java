
import java.io.File;
import java.io.IOException;
import java.util.Scanner;
import java.util.LinkedList;

/**
 * 链表数组的实现方式
 * 在以下 V表示顶点数 E表示图的边数
 * 空间复杂度 O(V + E)   这里V和E都是必要的，极端情况下可能E=0
 * 建图时间   O(E) 如果需要查重的话要遍历链表 时间复杂度为 O(V*E)
 * 两点是否相邻 O(degree(V)) 默认为顶点的度， 如果是完全图或者稠密图接近O(V)
 * 查找所有邻边 O(degree(V)) 默认为顶点的度， 如果是完全图或者稠密图接近O(V)
 */
class AdjList {
    private int V; // 图的顶点数
    private int E; // 图的边数
    private LinkedList<Integer>[] adj; // 图方隈

    public AdjList(String filename) {
        File file = new File(filename);
        try(Scanner scanner = new Scanner(file)){
            V = scanner.nextInt();
            if (V < 0) 
                throw new IllegalArgumentException("V must be non-negative");

            adj = new LinkedList[V];
            for (int i = 0; i < V; i++) {
                adj[i] = new LinkedList<Integer>();
            }

            E = scanner.nextInt();
            if (E < 0)
                throw new IllegalArgumentException("E must be non-negative");

            for (int i=0; i < E; i++) {
                int a = scanner.nextInt();
                validateVertex(a);
                int b = scanner.nextInt(); 
                validateVertex(b);

                if (a == b)
                    throw new IllegalArgumentException("Self Loop is Detected.");

                if (adj[a].contains(b)) 
                    throw new IllegalArgumentException("Parallel Edges are Detected.");

                adj[a].add(b);
                adj[b].add(a);
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
        return adj[v].contains(w);
    }

    // 获取一个顶点的邻边
    public LinkedList<Integer> adj(int v) {
        validateVertex(v);
        return adj[v];
    }

    // 获取一个项点的度（有多少条邻边）
    public int degree(int v) {
        return adj(v).size();
    }

    private void validateVertex(int v) {
        if (v < 0 || v >= V) 
            throw new IllegalArgumentException("vertex "+v+" is invalid.");
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append(String.format("V = %d, E = %d\n", V, E));

        for (int v=0; v < V; v++) {
            sb.append(String.format("%d: ", v));
            for (int w: adj[v]) {
                sb.append(String.format("%d ", w));
            }
            sb.append("\n");
        }
        return sb.toString();
    }

    public static void main(String[] args) {
        AdjList adjMatrix = new AdjList("./Graph/g.txt");
        System.out.println(adjMatrix);
    }

}