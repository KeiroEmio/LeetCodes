package main

import "fmt"
import "dataStructureDev/QuickSort/QuickSorePkg"

func QuickSort(arr []int, low, high int) {
	if low < high {
		// 获取分区点
		pivotIndex := partition(arr, low, high)
		// 递归排序左右两部分
		QuickSort(arr, low, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, high)
	}
}

// partition 函数用于分区
func partition(arr []int, low, high int) int {
	// 选择最右边的元素作为基准
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			//为什么要i++
			i++
			// 交换 arr[i] 和 arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// 交换 arr[i+1] 和基准
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("排序前:", arr)
	QuickSorePkg.QuickSort1(arr, 0, len(arr)-1)
	fmt.Println("排序后:", arr)
}
