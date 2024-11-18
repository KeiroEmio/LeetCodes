package InsertSortPkg

func InsertSort1(arr []int) {
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
