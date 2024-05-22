package binarySearch

func Recursive(arr []int, start int, end int, k int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if arr[mid] == k {
		return mid
	}
	if arr[mid] > k {
		return Recursive(arr, start, mid-1, k)
	} else {
		return Recursive(arr, mid+1, end, k)
	}
}
