import java.util.ArrayList;
import java.util.Collections;

// 最小生成树：
// Kruskal最小生成树算法实现
// 使用了一个额外类WeightGraphEdges类来记录带权的节点
// 使用并查集检测是否有环
class Kruskal {
    private WeightGraph G;
    private ArrayList<WeightGraphEdges> mst; 

    public Kruskal(WeightGraph G) {
        this.G = G;
        mst = new ArrayList<WeightGraphEdges>();

        CC1 cc = new CC1(G);
        if (cc.getCCount() > 1) return;  // 如果不是连通的图直接返回

        ArrayList<WeightGraphEdges> edges = new ArrayList<>();
        for (int v=0; v < G.V(); v++) {
            for (int w: G.adj(v)){
                // 保存同一条边只处理一次
                if (v < w)  
                    edges.add(new WeightGraphEdges(v, w, G.getWeight(v, w)));
            }
        }

        Collections.sort(edges);

        UnionFind uf = new UnionFind(G.V());
        for (WeightGraphEdges ed: edges) {
            // 如果不在一个集合里，则说明有最小生成树的边
            if (!uf.unionContains(ed.V(), ed.W())) {
                mst.add(ed);
                uf.union(ed.V(), ed.W());
            }
        }
    }

    public ArrayList<WeightGraphEdges> getResult() {
        return mst;
    }

    public static void main(String[] args){
        WeightGraph g = new WeightGraph("WeightGraph/g.txt");
        Kruskal k = new Kruskal(g);
        System.out.println(k.getResult());
    }
}

