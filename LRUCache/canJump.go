package main

import "fmt"

// 给定一个非负整数数组 nums ，你最初位于数组的第一个位置。数组中的每个元素代表你在该位置可以跳跃的最大长度。判断你是否能够到达最后一个位置。
//
// 示例：
// 输入：nums = [2,3,1,1,4]
// 输出：true
// 解释：从位置 0 跳到 1，然后跳到最后一个位置。
func canJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tt := canJump([]int{1, 8, 9})
	fmt.Println("tt:", tt)
}
