package kit

type CircularQueue struct {
	items []string
	n     int
	head  int
	tail  int
}

func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{items: make([]string, 0), n: n}
}

func (s *CircularQueue) Enqueue(item string) bool {
	if (s.tail+1)%s.n == s.head {
		return false
	}
	s.items = append(s.items, item)
	s.tail = (s.tail + 1) % s.n
	return true
}

func (s *CircularQueue) Dequeue() string {
	if s.head == s.tail {
		return ""
	}
	tmp := s.items[s.head]
	s.head = (s.head + 1) % s.n
	return tmp
}
