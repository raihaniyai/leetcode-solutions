func numIslands(grid [][]byte) int {
    totalIslands := 0
    totalRows := len(grid)
    totalCols := len(grid[0])

    for i := 0; i < totalRows; i++ {
        for j := 0; j < totalCols; j++ {
            if grid[i][j] == '1' {
                totalIslands++
                dfs(grid, i, j, totalRows, totalCols)
            }
        }
    }

    return totalIslands
}

func dfs(grid [][]byte, i, j, m, n int) {
    if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] != '1' {
        return
    }

    var emptyByte byte
    grid[i][j] = emptyByte
    dfs(grid, i-1, j, m, n)
    dfs(grid, i+1, j, m, n)
    dfs(grid, i, j+1, m, n)
    dfs(grid, i, j-1, m, n)
}