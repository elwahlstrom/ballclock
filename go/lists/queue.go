package lists

type Queue struct {
	len, count int
	head, tail int
	list       []int
}

// initalizes a new Queue with a specific max size
func NewQueue(size int) *Queue {
	return &Queue{size, 0, 0, 0, make([]int, size)}
}

// returns the Queue preallocated size
func (q *Queue) Size() int {
	return q.len
}

// returns the # of items in a the Queue
func (q *Queue) Count() int {
	return q.count
}

// adds an item to the head of the Queue
func (q *Queue) Enqueue(x int) bool {
	if q.count == q.len {
		return false
	}

	if q.count > 0 {
		q.head = (q.head + 1) % q.len
	}

	q.list[q.head] = x
	q.count++
	return true
}

// removes an item from the back of the Queue
func (q *Queue) Dequeue() (int, bool) {
	if q.count == 0 {
		return 0, false
	}

	x := q.list[q.tail]
	q.list[q.tail] = 0
	q.count--

	if q.count > 0 {
		q.tail = (q.tail + 1) % q.len
	}

	return x, true
}

// returns an array of the underlying data
func (q *Queue) ToArray() []int {
	a := make([]int, q.count)
	for i := 0; i < q.count; i++ {
		a[i] = q.list[(q.tail+i)%q.len]
	}
	return a
}
