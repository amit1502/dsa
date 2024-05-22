package binarySearch

func LastOccurrence(arr []int, k int) int {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := (start + end) //2
		if arr[mid] > k {
			end = mid - 1
		} else if arr[mid] < k {
			start = mid + 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] != k {
				return mid
			}
		}
	}
	return -1
}
