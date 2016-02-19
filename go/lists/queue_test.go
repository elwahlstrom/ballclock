package lists

import "testing"

func TestNewQueue(t *testing.T) {
	size := 5
	q := NewQueue(size);
	if q.Size() != size {
		t.Errorf("NewQueue(%v) failed, queue.Size=%v which is not the expected value of %v", size, q.Size(), size)
	}
}

func TestEnqueue(t *testing.T) {

}