package sortbyselection_test

import (
	"testing"

	sortbyselection "github.com/l0s0s/JokesAPI/SortBySelection"
	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	selection := []int{5, 3, 6, 2, 10}
	actual := sortbyselection.SelectionSort(selection)
	assert.Equal(t, actual, []int{2, 3, 5, 6, 10})
}
