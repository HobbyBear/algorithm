package base

func BinarySearch(arr []int, target int) int {
	return binarySearchHelper(arr, 0, len(arr)-1, target)
}

func binarySearchHelper(arr []int, start, end, target int) int {
	if start > end {
		return -1
	}
	middle := (start + end) / 2
	obj := arr[middle]
	if obj < target {
		return binarySearchHelper(arr, middle+1, end, target)
	}
	if obj > target {
		return binarySearchHelper(arr, start, middle-1, target)
	}
	if obj == target {
		return middle
	}
	return -1
}

func BinarySearchLoop(arr []int, target int) int {
	var (
		start int
		end   int
	)
	end = len(arr) - 1
	for start <= end{
		middle := (start + end) / 2
		obj := arr[middle]
		if obj > target{
			end  = middle - 1
			continue
		}
		if obj < target{
			start = middle + 1
			continue
		}
		if obj == target{
			return middle
		}
	}
	return -1
}
