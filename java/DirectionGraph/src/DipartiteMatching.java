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

    public static void main(String args[]) {
        Graph g = new Graph("DirectionGraph/g2.txt");
        DipartiteMatching bm = new DipartiteMatching(g);
        System.out.println(bm.maxMatching());

        Graph g2 = new Graph("DirectionGraph/g3.txt");
        DipartiteMatching bm2 = new DipartiteMatching(g2);
        System.out.println(bm2.maxMatching());
    }


}