package binarySearch

func Peak(arr []int) int {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := (start + end) / 2
		if (mid == 0 || arr[mid] >= arr[mid-1]) && (mid < len(arr)-1 ||
			arr[mid] >= arr[mid+1]) {
			return mid
		}
		if mid > 0 && arr[mid-1] >= arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
