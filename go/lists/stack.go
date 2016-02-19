package lists

type Stack struct {
	len, count int
	list       []int
}

func NewStack(n int) *Stack {
	return &Stack{n, 0, make([]int, n)}
}

func (s *Stack) Count() int {
	return s.count
}

func (s *Stack) Push(x int) bool {
	if s.count == s.len {
		return false
	}

	s.list[s.count] = x
	s.count++
	return true
}

func (s *Stack) Pop() (int, bool) {
	if s.count == 0 {
		return 0, false
	}

	x := s.list[s.count-1]
	s.count--
	return x, true
}
