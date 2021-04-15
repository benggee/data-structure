package leetcode;

// floodfill算法本质上也是图的遍历
// 详情见leetcode 695号题
// 数据如下：
// [[0,0,1,0,0,0,0,1,0,0,0,0,0],
//  [0,0,0,0,0,0,0,1,1,1,0,0,0],
//  [0,1,1,0,1,0,0,0,0,0,0,0,0],
//  [0,1,0,0,1,1,0,0,1,0,1,0,0],
//  [0,1,0,0,1,1,0,0,1,1,1,0,0],
//  [0,0,0,0,0,0,0,0,0,0,1,0,0],
//  [0,0,0,0,0,0,0,1,1,1,0,0,0],
//  [0,0,0,0,0,0,0,1,1,0,0,0,0]]
class FloodFileLeetcode695 {
    private int[][] grid;
    private boolean[][] visited;
    private int R,C;

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
        visited = new boolean[R][C];
        for (int i=0; i < R; i++) {
            for (int j=0; j<C; j++) {
                if (!visited[i][j] && grid[i][j] == 1) {
                    res = Math.max(res, dfs(i,j));
                }
            }
        }
        return res;
    }

    private int dfs(int i, int j) {
        visited[i][j] = true;
        int res = 1;
        for (int d=0; d<4; d++) {
            int nextx = i + dirs[d][0], nexty = j + dirs[d][1];
            if (isArea(nextx, nexty) && !visited[nextx][nexty] && grid[nextx][nexty] == 1) {
                res += dfs(nextx, nexty);
            }
        }
        return res;
    }

    private boolean isArea(int nextx, int nexty) {
        return (nextx>=0 && nextx<R) && (nexty >=0 && nexty < C);
    }
}