package base

func BubbleSort(arr []int) []int {
	end := len(arr) - 1
	for i := 0; i+1 <= end; end-- {
		for j := i; j+1 <= end; j++ {
			if arr[j] > arr[j+1] {
				pre := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = pre
			}
		}
	}
	return arr
}
