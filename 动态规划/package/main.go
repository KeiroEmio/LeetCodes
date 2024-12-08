package main

import "fmt"

// knapsack 使用动态规划解决0/1背包问题
func knapsack(W int, wt []int, val []int, n int) int {
	// 创建dp数组，dp[i][j]表示前i个物品，背包容量为j时的最大价值
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	// 填充dp数组
	for i := 1; i <= n; i++ {
		for w := 1; w <= W; w++ {
			// 如果不选择第i个物品
			dp[i][w] = dp[i-1][w]
			// 如果选择第i个物品（前提是w >= wt[i-1]，即物品i可以放入背包）
			if wt[i-1] <= w {
				dp[i][w] = max(dp[i][w], dp[i-1][w-wt[i-1]]+val[i-1])
			}
		}
	}

	// 返回最大价值
	return dp[n][W]
}

// max 返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	wt := []int{2, 3, 4, 5}

	val := []int{3, 4, 5, 6}

	W := 5
	// 物品的数量
	n := len(wt)

	// 调用knapsack函数计算最大价值
	result := knapsack(W, wt, val, n)
	fmt.Printf("最大价值: %d\n", result)
}
