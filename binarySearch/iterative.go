package binarySearch

func Iterative(arr []int, k int) int {

	start, end := 0, len(arr)-1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == k {
			return mid
		}
		if arr[mid] < k {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
