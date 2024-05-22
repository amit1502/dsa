package binarySearch

// unbound binary search
func InfiniteArray(arr []int, k int) int {
	if arr[0] == k {
		return 0
	}
	end := 1
	for arr[end] < k {
		end = end * 2
	}
	if arr[end] == k {
		return end
	} else {
		return Recursive(arr, (end/2)+1, end-1, k)
	}
}
