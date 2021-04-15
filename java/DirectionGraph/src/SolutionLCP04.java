import java.util.TreeSet;
import java.util.TreeMap;
import java.util.Collections;
import java.util.ArrayList;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Arrays;

// Leetcode LCP 04. 覆盖
class SolutionLCP04 {
    class Graph {
        private int V; // 图的顶点数
        private int E; // 图的边数
        private TreeSet<Integer>[] adj; // 图方隈
        private boolean direction = false; 
        private int[] indegrees; // 入度
        private int[] outdegress; // 出度

        public Graph(int v, boolean direction) {
            this.direction = direction;
            this.V = v;
            this.E = 0;

            adj = new TreeSet[v];
            for (int i=0; i<v; i++) {
                adj[i] = new TreeSet<Integer>();
            }
        }

        public void addEdge(int v, int w) {
            validateVertex(v);
            validateVertex(w);

            if (v == w)
                throw new IllegalArgumentException("Self Loop is Detected.");

            if (adj[v].contains(w)) 
                throw new IllegalArgumentException("Parallel Edges are Detected.");

            adj[v].add(w);
            if (direction) {
                indegrees[w] ++;
                outdegress[v] ++;
            }
            if (!direction)
                adj[w].add(v);
            
            this.E++;
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
        public Iterable<Integer> adj(int v) {
            validateVertex(v);
            return adj[v];
        }

        public void validateVertex(int v) {
            if (v < 0 || v >= V) 
                throw new IllegalArgumentException("vertex "+v+" is invalid.");
        }

        public boolean isDirection() {
            return direction;
        }

        // 获取一个项点的度（有多少条邻边）
        public int degree(int v) {
            if (isDirection()) throw new RuntimeException("degree only work on undiretion graph.");
            validateVertex(v);
            return adj[v].size();
        }

        // 获取入度
        public int indegree(int v) {
            if (!isDirection()) throw new RuntimeException("indegree only work on diretion graph.");
            validateVertex(v);
            return indegrees[v];
        }

        // 获取出度
        public int outdegree(int v) {
            if (!isDirection()) throw new RuntimeException("outdegree only work on diretion graph.");
            validateVertex(v);
            return outdegree(v);
        }

        public void removeEdge(int v, int w) {
            validateVertex(v);
            validateVertex(w);

            if (adj[v].contains(w)) {
                E--;
                if (isDirection()) {
                    indegrees[w] --;
                    outdegress[v] --;
                }
            } 

            adj[v].remove(w);
            if (!isDirection())
                adj[w].remove(v);
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
    }
    
    class WeightGraph {
        private int V; // 图的顶点数
        private int E; // 图的边数
        private TreeMap<Integer, Integer>[] adj; // 图方隈
        private boolean direction = false ;
    
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
    }
    
    public class BipartitionDetection {
        Graph g;
        private boolean[] visited;
        private int[] colors;
        private boolean isBipart = true;  // 是否是二分图
    
        public BipartitionDetection(Graph g) {
            this.g = g;
            colors = new int[g.V()];
            for (int i=0; i<g.V(); i++) {
                colors[i] = -1;
            }
    
            visited = new boolean[g.V()];
            for (int v = 0; v < g.V(); v++) {
                if (!visited[v]) {
                    if (!dfs(v, 0)) {
                        isBipart = false;
                        break;
                    }
                }
            }
        }
    
        private boolean dfs(int v, int color) {
            visited[v] = true;
            colors[v] = color;
            for (int w: g.adj(v)) {
                if (!visited[w]) {
                    if (!dfs(w, 1-color)) return false;
                } else {
                    if (colors[w] == colors[v]) return false;
                }
            }
            return true;
        }
    
        public int[] getColors() {
            return colors;
        }
    
        public boolean isBipart() {
            return isBipart;
        }
    }

    class MaxFlow {
        private WeightGraph net;
        private int s,t;
        
        private WeightGraph rG; 
        private int maxFlow; // 最大流
    
        public MaxFlow(WeightGraph net, int s, int t) {
            if (!net.isDirection())  throw new IllegalArgumentException("MaxFlow only work on direction graph.");
            if (net.V() < 2) throw new IllegalArgumentException("The graph v should least 2.");
    
            net.validateVertex(s);
            net.validateVertex(t);
    
            if (s == t) throw new IllegalArgumentException("s and t must be different.");
    
            this.net = net;
            this.s = s;
            this.t = t;
    
            this.rG = new WeightGraph(net.V(), true);
            for (int v=0; v<net.V(); v++) {
                for (int w: net.adj(v)) {
                    rG.addEdge(v, w, net.getWeight(v, w));
                    rG.addEdge(w, v, 0);
                }
            }
    
            while(true) {
                // 示一条增广路径
                ArrayList<Integer> augPath = getAugmentingPath();
                if (augPath.size() == 0) break;
    
                int f = Integer.MAX_VALUE;
                for (int i=1; i<augPath.size(); i++) {
                    int v = augPath.get(i - 1);
                    int w = augPath.get(i);
                    f = Math.min(f, rG.getWeight(v, w));
                }
                maxFlow += f;
    
                for (int i = 1; i<augPath.size(); i++) {
                    int v = augPath.get(i-1);
                    int w = augPath.get(i);
    
                    rG.setWeight(v, w, rG.getWeight(v, w) - f);
                    rG.setWeight(w, v, rG.getWeight(w, v) + f);
                }
            }
        }
    
        private ArrayList<Integer> getAugmentingPath() {
            Queue<Integer> q = new LinkedList<>();
            int[] pre = new int[net.V()];
            Arrays.fill(pre, -1);
    
            q.add(s);
            pre[s] = s;
            while(!q.isEmpty()) {
                int cur = q.remove();
                if (cur == t) break;
                for (int w: rG.adj(cur)) {
                    if (pre[w] == -1 && rG.getWeight(cur, w) > 0) {
                        pre[w] = cur;
                        q.add(w);
                    }
                }
            }
    
            ArrayList<Integer> res = new ArrayList<>();
            if (pre[t] == -1) return res;
    
            int cur = t;
            while(cur != s){
                res.add(cur);
                cur = pre[cur];
            }
            res.add(s);
    
            Collections.reverse(res);
            return res;
        }
    
        // 获取网络最大流
        public int result() {
            return maxFlow;
        }
        
        public int flow(int v, int w) {
            if (!net.hasEdge(v, w))  throw new IllegalArgumentException(String.format("No edge %d-%d", v, w));
            return rG.getWeight(w, v);
        }
    }
    
    public class DipartiteMatching {
        private Graph G;
        private int maxMatching;
    
        public DipartiteMatching(Graph G) {
            BipartitionDetection db = new BipartitionDetection(G);
            if (!db.isBipart()) 
                throw new IllegalArgumentException("dipartite matching only work on diprtite graph.");
    
            this.G = G;
    
            int[] colors = db.getColors();
    
            // 网络流模型建模
            // 源点：V  汇点：V+1
            WeightGraph net = new WeightGraph(G.V() + 2, true);
            for (int v=0; v<G.V(); v++) {
                if (colors[v] == 0) 
                    net.addEdge(G.V(), v, 1);  // 将二分图中的一个和源点建立一条边
                else 
                    net.addEdge(v, G.V() + 1, 1); // 将二分图中的另一个和汇点建立边
    
                for (int w: G.adj(v)) {
                    // 保证只建立一个方向的边
                    if (v < w) {
                        if (colors[v] == 0) 
                            net.addEdge(v, w, 1);
                        else 
                            net.addEdge(w, v, 1); 
                    }
                }
            }
    
            // 使用网络最大流求最大匹配数
            MaxFlow maxFlow = new MaxFlow(net, G.V(), G.V()+1);
            maxMatching = maxFlow.result();
        }
    
        public int maxMatching() {
            return maxMatching;
        }
    
        public boolean isPerfactMatching() {
            return maxMatching * 2 == G.V();
        }
    }

    public int domino(int n, int m, int[][] broken) {
        int[][] board = new int[n][m];
        for (int[] p: broken) {
            board[p[0]][p[1]] = 1;
        }

        Graph g = new Graph(m * n, false);



        for(int i = 0; i < n; i ++)
            for(int j = 0; j < m; j ++){
                if(j + 1 < m && board[i][j] == 0 && board[i][j + 1] == 0)
                    g.addEdge(i * m + j, i * m + (j + 1));
                if(i + 1 < n && board[i][j] == 0 && board[i + 1][j] == 0)
                    g.addEdge(i * m + j, (i + 1) * m + j);
            }

        // for (int i=0; i<n; i++) {
        //     for (int j=0; j<m; j++) {
        //         if (j + 1 < m && board[i][j] == 0 && board[i][j + 1] == 0) 
        //             g.addEdge(i * m + j, i * m + (j + 1));
        //         if (i + 1 < n && board[i][j] == 0 && board[i + 1][j] == 0)
        //             g.addEdge(i * m + j, (i + 1) * m + j);
        //     }
        // }

        DipartiteMatching bm = new DipartiteMatching(g);
        return bm.maxMatching();
    }

    public static void main(String args[]) {
        SolutionLCP04 s = new SolutionLCP04();
        int[][] re = new int[][]{{1,0},{1,1}};
        System.out.println(s.domino(2, 3, re));
    }
}