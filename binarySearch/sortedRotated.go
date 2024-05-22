package binarySearch

func SortedRotated(arr []int, k int) int {
	start, end := 0, len(arr)-1

	for start <= end {
		mid := (start + end) / 2

		if arr[mid] == k {
			return mid
		}
		if arr[mid] >= arr[start] {
			// left half is sorted
			if k >= arr[start] && k < arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			// right half is sorted
			if k > arr[mid] && k <= arr[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
