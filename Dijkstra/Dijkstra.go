package dijkstra

import (
	"math"
)

type Graph map[string]map[string]int

type Costs map[string]float64

type Parents map[string]string

var processed []string

func GetNeighbors(m map[string]int) []string {
	keys := make([]string, len(m))

	var i int

	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func Search(graph Graph, costs Costs, parents Parents) int {
	var newCost float64

	node := findLowestCostNode(costs)

	for node != "" {
		cost := costs[node]
		neighbours := graph[node]

		for _, n := range GetNeighbors(neighbours) {
			newCost = cost + float64(neighbours[n])

			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}

		processed = append(processed, node)
		node = findLowestCostNode(costs)
	}
	return int(newCost)
}

func findLowestCostNode(costs Costs) string {
	lowestCost := math.Inf(1)

	var lowestCostNode string

	for key := range costs {
		cost := costs[key]
		if cost < lowestCost && func() bool {
			for _, v := range processed {
				if key == v {
					return false
				}
			}
			return true
		}() {
			lowestCost = cost
			lowestCostNode = key
		}
	}
	return lowestCostNode
}
