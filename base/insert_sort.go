package base

func InsertSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		start := i
		cur := arr[i]
		for j := start - 1; j >= 0; j-- {
			if arr[j] > cur {
				arr[start] = arr[j]
				arr[j] = cur
				start = j
			}
		}
	}
	return arr
}
