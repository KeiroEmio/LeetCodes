package main

import "fmt"

func bubbleSort(array []int) {
	n := len(array)
	for i := 0; i < len(array)-1; i++ {
		t := false
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				m := array[j]
				array[j] = array[j+1]
				array[j+1] = m
				t = true
			}
		}

		if !t {
			break
		}
	}
}

func main() {
	// 示例数组
	arr := []int{5, 2, 9, 1, 5, 6}
	fmt.Println("原始数组:", arr)

	// 调用冒泡排序
	bubbleSort(arr)

	// 输出排序后的数组
	fmt.Println("排序后的数组:", arr)
}
