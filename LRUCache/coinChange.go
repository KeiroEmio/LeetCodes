package main

import "fmt"

// 给你一个整数数组 coins ，表示不同面额的硬币；另一个整数 amount ，表示总金额。请你计算并返回可以凑成总金额的最少硬币个数。
// 如果没有任何一种硬币组合能组成总金额，返回 -1。你可以认为每种硬币的数量是无限的。
// 示例：
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1

func coinChange1(coins []int, amount int) int {
	var n int

	if len(coins) == 1 && amount%coins[0] != 0 {
		return -1
	}
	nn := amount / coins[len(coins)-1]
	mm := amount % coins[len(coins)-1]
	coins = coins[:len(coins)-2]

	n = nn + n
	if len(coins) == 1 && amount%coins[0] == 0 {
		return n
	}

	return coinChange1(coins, mm)
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	num := coinChange([]int{1, 2, 3, 4}, 100)
	fmt.Println("num:", num)
}
