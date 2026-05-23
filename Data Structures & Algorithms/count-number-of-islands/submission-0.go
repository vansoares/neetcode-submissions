
func dfs(grid [][]byte, row, column int, visited [][]bool){
	directions := [][2]int{{-1,0}, {1,0}, {0,-1}, {0,1}}
	rows := len(grid)
	columns := len(grid[0])

	visited[row][column] = true

	for _, v := range directions {
		currentRow, currentColumn := row+v[0], column+v[1]
		if currentRow >= rows || currentRow < 0 || currentColumn >= columns || currentColumn < 0 {
			continue
		}

		if grid[currentRow][currentColumn] == '0' || visited[currentRow][currentColumn] {
			continue
		}

		dfs(grid, currentRow, currentColumn, visited)
	}

}

func numIslands(grid [][]byte) int {
    //matriz visitados
	visited := make([][]bool, len(grid))
	//resultado
	result := 0

	for k := range visited {
		visited[k] = make([]bool, len(grid[0]))
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' || visited[i][j] {
				continue
			}

			dfs(grid, i, j, visited)
			result+=1
		}
	}

	return result

}