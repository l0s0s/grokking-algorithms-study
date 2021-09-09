package BinarySearch_test

import (
	"testing"

	"github.com/l0s0s/JokesAPI/BinarySearch"
	"github.com/stretchr/testify/assert"
)

func SelectionSortTest(t *testing.T) {
	TestList := []int{5, 3, 6, 2, 10}
	actual := BinarySearch.Search(TestList, 9)
	assert.Equal(t, *actual, []int{2, 3, 5, 6, 10})
}