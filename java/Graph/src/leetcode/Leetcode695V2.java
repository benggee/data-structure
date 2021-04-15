package leetcode;
import java.util.HashSet;

// leetcode 695号题
class Solution {
    private int R;
    private int C;

    private HashSet<Integer>[] G;

    private int[][] grid;
    private boolean[] visited;

    // 代表方向 
    private int[][]dirs = {{-1,0},{0,1},{1,0},{0,-1}};

    public int maxAreaOfIsland(int[][] grid) {
        if (grid == null) 
            return 0;
        R = grid.length;
        if (R == 0) 
            return 0;
        C = grid[0].length;
        if (C == 0) 
            return 0;

        this.grid = grid;

        int res = 0;
        G = constructGraph();
        visited = new boolean[G.length];
        for (int v=0; v < G.length; v++) {
            int x = v/C, y = v % C;
            if (!visited[v] && grid[x][y] == 1) {
                res = Math.max(res, dfs(v));
            }
        }
        return res;
    }

    private int dfs(int v) {
        visited[v] = true;
        int res = 1;
        for (int w: G[v]) {
            if (!visited[w]) {
                res += dfs(w);
            }
        }
        return res;
    }

    private HashSet<Integer>[] constructGraph() {
        // R*C是顶点数
        HashSet<Integer>[] g = new HashSet[R * C];
        for (int i=0; i < g.length; i++) {
            g[i] = new HashSet<>();
        }
        for (int v=0; v<g.length; v++) {
            int x = v / C, y = v % C;
            if (grid[x][y] == 1) {
                for (int d=0; d<4; d++) {
                    int nextx = x + dirs[d][0];
                    int nexty = y + dirs[d][1];
                    if (isArea(nextx, nexty) && grid[nextx][nexty] == 1) {
                        int next = nextx * C + nexty;
                        g[v].add(next);
                        g[next].add(v);
                    }
                }
            }
        }
        return g;
    }

    private boolean isArea(int nextx, int nexty) {
        return (nextx>=0 && nextx<R) && (nexty >=0 && nexty < C);
    }
}