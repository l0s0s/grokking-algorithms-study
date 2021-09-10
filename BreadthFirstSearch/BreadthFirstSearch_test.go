package breadthfirstsearch_test

import (
	"testing"

	breadthfirstsearch "github.com/l0s0s/JokesAPI/BreadthFirstSearch"
	"github.com/stretchr/testify/assert"
)

func NewMockGraph() breadthfirstsearch.Graph {
	return breadthfirstsearch.Graph{
		"you":    []string{"alice", " ЬоЬ", "claire"},
		"bob":    []string{"anuj", "peggy"},
		"alice":  []string{"peggy"},
		"claire": []string{"thom", "jonny"},
		"anuj":   []string{},
		"peggy":  []string{},
		"thom":   []string{},
		"jonny":  []string{},
	}
}

func TestSearch(t *testing.T) {
	testGraph := NewMockGraph()
	queue := breadthfirstsearch.NewQueue()
	queue.Add(testGraph["you"]...)
	assert.Equal(t, breadthfirstsearch.Search("jonny", testGraph, queue), true)
}
