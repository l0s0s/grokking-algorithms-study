package dijkstra_test

import (
	"math"
	"testing"

	dijkstra "github.com/l0s0s/JokesAPI/Dijkstra"
	"github.com/stretchr/testify/assert"
)

func NewMockGraph() dijkstra.Graph {
	return dijkstra.Graph{
		"start": {
			"a": 6,
			"b": 2,
		},
		"a": {
			"fin": 1,
		},
		"b": {
			"a":   3,
			"fin": 5,
		},
		"fin": {},
	}
}

func NewMockCosts() dijkstra.Costs {
	return dijkstra.Costs{
		"a":   6,
		"b":   2,
		"fin": math.Inf(1),
	}
}

func NewMockParents() dijkstra.Parents {
	return dijkstra.Parents{
		"a":  "start",
		"b":  "start",
		"in": "",
	}
}

func TestSearch(t *testing.T) {
	actual := dijkstra.Search(NewMockGraph(), NewMockCosts(), NewMockParents())
	assert.Equal(t, 6, actual)
}
