package SortBySelection

func FindSmallest(arr []int) map[string]int {
	smallest := map[string]int{
		"index": 0,
		"value": arr[0],
	}
	for i, v := range arr {
		if v < smallest["value"] {
			smallest["value"] = v
			smallest["index"] = i
		}
	}
	return smallest
}

func SelectionSort(arr []int) []int {
	newArr := make([]int, len(arr))
	length := len(arr)
	for i := 0; i < length; i++ {
		smallest := FindSmallest(arr)
		newArr[i] = smallest["value"]
		arr = append(arr[:smallest["index"]], arr[smallest["index"]+1:]...)
	}
	return newArr
}
