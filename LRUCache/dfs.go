package main

import "fmt"

var directions = [][]int{
	{0, 1},  // 右
	{1, 0},  // 下
	{0, -1}, // 左
	{-1, 0}, // 上
}

// 检查当前位置是否越界或碰到墙
func isValid(x, y, rows, cols int, maze [][]int, visited [][]bool) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols && maze[x][y] == 0 && !visited[x][y]
}

// 深度优先搜索
func dfs(maze [][]int, visited [][]bool, x, y int, path *[][2]int) bool {
	rows := len(maze)
	cols := len(maze[0])

	// 如果到达终点，返回true
	if x == rows-1 && y == cols-1 {
		*path = append(*path, [2]int{x, y})
		return true
	}

	visited[x][y] = true
	*path = append(*path, [2]int{x, y})

	// 遍历四个方向
	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		if isValid(newX, newY, rows, cols, maze, visited) {
			if dfs(maze, visited, newX, newY, path) {
				return true
			}
		}
	}

	// 如果当前路径不通，回溯
	//*path = (*path)[:len(*path)-1]
	return false
}

func findPath(maze [][]int) ([][2]int, bool) {
	if len(maze) == 0 {
		return nil, false
	}

	rows := len(maze)
	cols := len(maze[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var path [][2]int
	if dfs(maze, visited, 0, 0, &path) {
		return path, true
	}

	return nil, false
}

func main() {
	maze := [][]int{
		{0, 0, 1, 0},
		{1, 0, 1, 0},
		{1, 0, 0, 0},
		{1, 1, 0, 0},
	}

	path, found := findPath(maze)
	if found {
		fmt.Println("存在路径，路径为：", path)
	} else {
		fmt.Println("没有可通行的路径")
	}
}
