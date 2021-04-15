import java.util.ArrayList;
import java.util.Collections;
import java.util.Stack;

// 求解欧拉回路
// Hierholzer算法实现
public class DirectionEulerLoop {
    private Graph G;

    public DirectionEulerLoop(Graph G) {
        if (!G.isDirection()) throw new IllegalArgumentException("DirectionEulerLoop only work on undirection graph.");
        this.G = G;
    }

    public boolean hasEulerLoop() {
        // TODO 判断是否连通
        // CC cc = new CC(G);
        // if (cc.getCCount() > 1) 
        //     return false;
        for (int v=0; v<G.V(); v++) {
            // 如果顶点的度不是偶数，则直接返回
            if (G.indegree(v) != G.outdegree(v)) {
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
        while(!st.isEmpty()) {
            if (g.outdegree(curv) != 0) {
                st.push(curv);
                int w = g.adj(curv).iterator().next(); // 随便找一条边
                g.removeEdge(curv, w); // 将找到的这边条删掉
                curv = w;
            } else {
                res.add(curv);
                curv = st.pop();
            }
        }
        Collections.reverse(res);
        return res;
    }

    public static void main(String args[]) {
        Graph g = new Graph("DirectionGraph/euler01.txt", true);
        DirectionEulerLoop el = new DirectionEulerLoop(g);
        System.out.println(el.result());


        Graph g2 = new Graph("DirectionGraph/euler02.txt", true);
        DirectionEulerLoop el2 = new DirectionEulerLoop(g2);
        System.out.println(el2.result());
    }
}