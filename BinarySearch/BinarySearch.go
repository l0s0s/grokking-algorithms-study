package BinarySearch

func Search(list []int, item int) *int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		if list[mid] == item {
			return &mid
		}
		if list[mid] > item {
			high = mid + mid/2
		} else {
			low = mid + mid/2
		}
	}
	return nil
}
