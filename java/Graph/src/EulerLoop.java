import java.util.ArrayList;
import java.util.Stack;

// 求解欧拉回路
// Hierholzer算法实现
public class EulerLoop {
    private Graph G;

    public EulerLoop(Graph G) {
        this.G = G;
    }

    public boolean hasEulerLoop() {
        CC cc = new CC(G);
        if (cc.getCCount() > 1) 
            return false;

        for (int v=0; v<G.V(); v++) {
            // 如果顶点的度不是偶数，则直接返回
            if (G.degree(v) % 2 == 1) {
                return false;
            }
        }

        return true;
    }

    public ArrayList<Integer> result() {
        ArrayList<Integer> res = new ArrayList<>();
        if (!hasEulerLoop()) 
            return res;

        Graph g = (Graph)G.clone();
        Stack<Integer> st = new Stack<>();
        int curv = 0;
        st.push(curv);
        while(!st.empty()) {
            if (g.degree(curv) != 0) {
                st.push(curv);
                int w = g.adj(curv).iterator().next(); // 随便找一条边
                g.removeEdge(curv, w); // 将找到的这边条删掉
                curv = w;
            } else {
                res.add(curv);
                curv = st.pop();
            }
        }

        return res;
    }

    public static void main(String args[]) {
        Graph g = new Graph("Graph/euler01.txt");
        EulerLoop el = new EulerLoop(g);
        System.out.println(el.result());


        Graph g2 = new Graph("Graph/euler02.txt");
        EulerLoop el2 = new EulerLoop(g2);
        System.out.println(el2.result());
    }
}