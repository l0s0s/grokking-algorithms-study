package approximatealgorithms_test

import (
	"testing"

	approximatealgorithms "github.com/l0s0s/JokesAPI/ApproximateAlgorithms"
	"github.com/stretchr/testify/assert"
)

func NewMockStates() []string {
	return []string{"mt", "wa", "or", "id", "nv", "ut", "са", "az"}
}

func NewMockStations() approximatealgorithms.Stations {
	return approximatealgorithms.Stations{
		"kone":   []string{"id", "nv", "ut"},
		"ktwo":   []string{"wa", "id", "mt"},
		"kthree": []string{"or", "nv", "са"},
		"kfour":  []string{"nv", "ut"},
		"kfive":  []string{"ca", "az"},
	}
}

func TestSearch(t *testing.T) {
	actual := approximatealgorithms.Search(NewMockStates(), NewMockStations())
	assert.Equal(t, []string{"kone", "ktwo", "kthree", "kfive"}, actual)
}
