package lists

type Stack struct {
	count int
	list  []int
}

// initialize a new stack
func NewStack(size int) *Stack {
	return &Stack{0, make([]int, size)}
}

// returns the Stack preallocated size
func (s *Stack) Size() int {
	return len(s.list)
}

// returns the # of items in a the Stack
func (s *Stack) Count() int {
	return s.count
}

// pushes a new item on the Stack
func (s *Stack) Push(x int) bool {
	if s.count == s.Size() {
		return false
	}

	s.list[s.count] = x
	s.count++
	return true
}

// pops an item off the stack
func (s *Stack) Pop() (int, bool) {
	if s.count == 0 {
		return 0, false
	}

	x := s.list[s.count-1]
	s.count--
	return x, true
}

// returns an array of the underlying data
func (s *Stack) ToArray() []int {
	a := make([]int, s.count)
	for i, pos := 0, s.count-1; i < s.count; i, pos = i+1, pos-1 {
		a[i] = s.list[pos]
	}
	return a
}
