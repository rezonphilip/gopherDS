package recursion

func BinarySearch(list []int, val int, low, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if list[mid] == val {
		return mid
	}
	if list[mid] < val {
		return BinarySearch(list, val, mid+1, high)
	}
	return BinarySearch(list, val, low, mid-1)
}
