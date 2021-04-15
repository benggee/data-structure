
import java.io.File;
import java.io.IOException;
import java.util.Scanner;
import java.util.ArrayList;

/**
 * 邻接矩阵的实现方式
 * 在以下 V表示顶点数 E表示图的边数
 * 空间复杂度 O(V^2)  
 * 建图时间   O(E)
 * 两点是否相邻 O(1)
 * 查找所有邻边 O(V)   需要全部遍历一次
 */
class AdjMatrix {
    private int V; // 图的顶点数
    private int E; // 图的边数
    private int[][] adj; // 图方隈

    public AdjMatrix(String filename) {
        File file = new File(filename);
        try(Scanner scanner = new Scanner(file)){
            V = scanner.nextInt();
            if (V < 0) 
                throw new IllegalArgumentException("V must be non-negative");

            adj = new int[V][V];
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

                if (adj[a][b] == 1) 
                    throw new IllegalArgumentException("Parallel Edges are Detected.");

                adj[a][b] = 1;
                adj[b][a] = 1;
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
        return adj[v][w] == 1;
    }

    // 获取一个顶点的邻边
    public ArrayList<Integer> adj(int v) {
        validateVertex(v);
        ArrayList<Integer> res = new ArrayList<>();
        for (int i=0; i < V; i++) {
            if (adj[v][i] == 1) {
                res.add(i);
            }
        }

        return res;
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

        for (int i=0; i < V; i++) {
            for (int j=0; j < V; j++) {
                sb.append(String.format("%d ", adj[i][j]));
            }
            sb.append("\n");
        }
        return sb.toString();
    }

    public static void main(String[] args) {
        AdjMatrix adjMatrix = new AdjMatrix("./Graph/g.txt");
        System.out.println(adjMatrix);
    }

}