package quicksort_test

import (
	"testing"

	quicksort "github.com/l0s0s/JokesAPI/QuickSort"
	"github.com/stretchr/testify/assert"
)

func TestSelectionSortTest(t *testing.T) {
	TestList := []int{5, 3, 6, 2, 10}
	actual := quicksort.Quicksort(TestList)
	assert.Equal(t, actual, []int{2, 3, 5, 6, 10})
}
