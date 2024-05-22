package binarySearch

func FirstOccurrence(arr []int, k int) int {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] < k {
			start = mid + 1
		} else if arr[mid] > k {
			end = mid - 1
		} else {
			if mid == 0 || arr[mid-1] != k {
				return mid
			}
			end = mid - 1
		}
	}
	return -1
}
