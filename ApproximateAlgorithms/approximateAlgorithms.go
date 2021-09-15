package approximatealgorithms

type Stations map[string][]string

func Search(states []string, stations Stations) []string {
	finalStations := make([]string, 0, len(stations))

	for len(states) > 0 {
		bestStation := ""
		statesCovered := []string{}

		for station, statesFromStations := range stations {
			covered := FindCrossing(states, statesFromStations)

			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
		}

		states = Pop(states, statesCovered)

		finalStations = append(finalStations, bestStation)
	}

	return finalStations
}

func FindCrossing(arrs ...[]string) []string {
	repetitions := make(map[string]int)

	for _, v := range arrs {
		for _, val := range v {
			repetitions[val]++
		}
	}

	var crossing []string

	for key, val := range repetitions {
		if val > 1 {
			crossing = append(crossing, key)
		}
	}
	return crossing
}

func Pop(arr, superfluous []string) []string {
	var newArr []string

	for _, v := range arr {
		var coincidence int

		for _, val := range superfluous {
			if v == val && func() bool {
				for _, i := range newArr {
					if v == i {
						return false
					}
				}
				return true
			}() {
				coincidence++
			}
		}
		if coincidence == 0 {
			newArr = append(newArr, v)
		}
	}

	return newArr
}
