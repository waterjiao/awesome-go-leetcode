package kit

type ArrayQueue struct {
	items []string
	n     int
	head  int
	tail  int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{items: make([]string, n), n: n}
}

func (s *ArrayQueue) Enqueue(item string) bool {
	if s.tail == s.n {
		if s.head == 0 {
			return false
		}
		for i := s.head; i < s.tail; i++ {
			s.items[i-s.head] = s.items[i]
		}
		s.tail -= s.head
		s.head = 0
	}
	s.items = append(s.items, item)
	s.tail++
	return true
}

func (s *ArrayQueue) Dequeue() string {
	if s.head == s.tail {
		return ""
	}
	tmp := s.items[s.head]
	s.head++
	return tmp
}
