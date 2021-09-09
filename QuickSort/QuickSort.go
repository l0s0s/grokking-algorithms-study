package QuickSort

func Quicksort(array []int) []int{
	if len(array) < 2{
		return array
	} 

	pivot := array[(len(array)-1)/2]
	
	var less []int
	var greater []int

	for _, i := range array {
		if i < pivot {
			less = append(less, i)
		} 
		if i > pivot {
			greater = append(greater, i)
		} 
	} 
	return append(append(Quicksort(less), pivot), Quicksort(greater)...)
}
