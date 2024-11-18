package main

import "fmt"

//1. 题目：搜索旋转排序数组
//给定一个升序排列的整数数组 nums ，它的每个元素都是唯一的。数组在某个未知的点上进行了旋转（例如，[0,1,2,4,5,6,7] 可能变成 [4,5,6,7,0,1,2]）。
//请你编写一个函数来判断目标值 target 是否存在于数组中，如果存在，返回它的索引，否则返回 -1。
//你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
//
//示例：
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// 判断哪边是有序的
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 6, 5}
	targe := 2
	targeInt := search(nums, targe)
	fmt.Println("targeInt:", targeInt)
}
