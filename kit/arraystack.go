package kit

type ArrayStack struct {
	items []string
	count int
	n     int
}

func NewArrayStack(n int) *ArrayStack {
	return &ArrayStack{items: make([]string, n), n: n}
}

func (s *ArrayStack) Push(item string) bool {
	if s.count == s.n {
		return false
	}
	s.items = append(s.items, item)
	s.count++
	return true
}

func (s *ArrayStack) Pop() string {
	if s.count == 0 {
		return ""
	}
	tmp := s.items[s.count-1]
	s.count--
	return tmp
}
