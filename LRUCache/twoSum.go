package main

import "fmt"

//输入：nums = [2, 7, 11, 15], target = 9
//输出：[0, 1]
//解释：因为 nums[0] + nums[1] = 2 + 7 = 9，返回 [0, 1]。

func tt(nums []int, target int) []int {
	var result []int
	var ok bool
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ok = true
				result = append(result, i, j)
			}
		}
		if ok {
			return result
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("tt:", tt(nums, target))
}
