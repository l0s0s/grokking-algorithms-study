package breadthfirstsearch

type Queue struct {
	queue []string
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Get() string {
	member := q.queue[0]
	q.queue = q.queue[1:]
	return member
}

func (q *Queue) Add(newValue ...string) {
	q.queue = append(q.queue, newValue...)
}
