package BinarySearch_test

import (
	"testing"

	"github.com/l0s0s/JokesAPI/BinarySearch"
	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	TestList := []int{1, 3, 5, 7, 9}
	actual := BinarySearch.Search(TestList, 9)
	assert.Equal(t, *actual, 4)
}
