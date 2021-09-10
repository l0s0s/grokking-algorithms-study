package breadthfirstsearch

type Graph map[string][]string

func Search(member string, graph Graph, queue *Queue) bool {
	var searched []string

	for len(queue.queue) > 0 {
		person := queue.Get()
		if func() bool {
			for _, v := range searched {
				if person == v {
					return false
				}
			}
			return true
		}() {
			if person == member {
				return true
			}

			for _, v := range graph[person] {
				queue.Add(v)
			}
		}
	}
	return false
}
