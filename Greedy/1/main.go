package main

import (
	"fmt"
	"sort"
)

// Activity 表示一个活动，包括开始时间和结束时间
type Activity struct {
	start  int
	finish int
}

// 按结束时间对活动排序
func sortActivities(activities []Activity) {
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].finish < activities[j].finish
	})
}

// activitySelection 实现活动选择问题的贪心算法
func activitySelection(start []int, finish []int) int {
	n := len(start)
	// 创建活动列表
	activities := make([]Activity, n)
	for i := 0; i < n; i++ {
		activities[i] = Activity{start[i], finish[i]}
	}

	// 按照结束时间排序活动
	sortActivities(activities)

	// 贪心选择活动

	// 第一个活动总是被选择
	count := 1
	lastFinishTime := activities[0].finish

	for i := 1; i < n; i++ {
		// 如果当前活动的开始时间大于等于最后一个选择的活动的结束时间，则选择这个活动
		if activities[i].start >= lastFinishTime {
			count++
			lastFinishTime = activities[i].finish
		}
	}

	return count
}

func main() {
	// start time
	start := []int{1, 3, 0, 5, 8, 5}
	// end time
	finish := []int{2, 4, 6, 7, 9, 9}
	result := activitySelection(start, finish)

	fmt.Printf("可以选择的最大活动数量是: %d\n", result)
}
