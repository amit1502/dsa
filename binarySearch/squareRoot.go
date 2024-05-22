package binarySearch

func SquareRoot(n int) int {
	start, end, ans := 1, n, -1
	for start <= end {
		mid := (start + end) / 2
		sqr := mid * mid

		if sqr == n {
			return mid
		}
		if sqr > n {
			end = mid - 1
		} else {
			ans = mid
			start = mid + 1
		}
	}
	return ans
}
