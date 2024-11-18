package main

import "fmt"
import "dataStructureDev/InsertSort/InsertSortPkg"

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("排序前:", arr)
	InsertSortPkg.InsertSort1(arr)
	fmt.Println("排序后:", arr)
}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		// 将 arr[i] 插入到已排序部分 arr[0...i-1] 中
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
