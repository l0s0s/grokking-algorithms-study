package main

import "fmt"

const expected = 99

var MyList = []int{1, 3, 5, 7, 9}

var steps = 0

func search(list []int, item int) *int {
	steps++

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
			low = mid - mid/2
		}
	}
	return nil
}

func main() {
	fmt.Println(*search(MyList, 5))
}
