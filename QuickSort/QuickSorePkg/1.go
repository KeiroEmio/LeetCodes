package QuickSorePkg

func QuickSort1(arr []int, low, high int) {
	if low < high {
		//获取分区区点
		n := partition1(arr, low, high)
		//递归左右排序
		QuickSort1(arr, low, n-1)
		QuickSort1(arr, n+1, high)
	}
}

// partition 函数用于分区
func partition1(arr []int, low, high int) int {
	//以最右为基点
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	//互换基点
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}
