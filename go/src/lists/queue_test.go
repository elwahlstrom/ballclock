package lists

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	size := 5
	q := buildQueue(size, 0)

	//make sure the queue is the proper size
	if q.Size() != size {
		t.Errorf("NewQueue(%v) failed, queue.Size=%v which is not the expected value of %v", size, q.Size(), size)
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := buildQueue(5, 5)

	//make sure the queue contains the proper # of items
	if q.Count() != 5 {
		t.Errorf("Queue contains %v item when we only expected 5", q.Count())
	}

	//make sure the queue contains the expected values
	a := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(q.ToArray(), a) {
		t.Errorf("Queue %v does not match %v", q.ToArray(), a)
	}
}

func TestQueueEnequeueFull(t *testing.T) {
	q := buildQueue(5, 6)

	//make sure the queue contains the proper # of items
	if q.Count() != 5 {
		t.Errorf("Queue contains %v item when we only expected 5", q.Count())
	}

	//make sure the queue contains the expected values
	a := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(q.ToArray(), a) {
		t.Errorf("Queue %v does not match %v", q.ToArray(), a)
	}
}

func TestQueueWrap(t *testing.T) {
	q := buildQueue(5, 5)

	//remote an item from the tail so we can add an item that wraps to the begining of the internal array
	q.Dequeue()
	q.Enqueue(5)

	//make sure the queue contains the proper # of items
	if q.Count() != 5 {
		t.Errorf("Queue contains %v item when we only expected 5", q.Count())
	}

	//make sure the queue contains the expected values
	a := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(q.ToArray(), a) {
		t.Errorf("Queue %v does not match %v", q.ToArray(), a)
	}
}

func TestQueueDequeue(t *testing.T) {
	q := buildQueue(5, 1)

	//pop the item and make sure it's the expected value
	val, success := q.Dequeue()
	if !success || val != 0 {
		t.Errorf("Unexpected dequeue value %v, expected 0", val)
	}

	//make sure the queue is now empty
	if q.Count() != 0 {
		t.Errorf("Expected an empty queue got %v", q.Count())
	}
}

func TestQueueDequeueEmpty(t *testing.T) {
	q := buildQueue(5, 0)

	//pop the item and make sure it's the expected value
	_, success := q.Dequeue()
	if success {
		t.Error("Queue was empty but got a 'success' on dequeue")
	}

	//make sure the queue is now empty
	if q.Count() != 0 {
		t.Errorf("Expected an empty queue got %v", q.Count())
	}
}

func TestQueueDequeueWrap(t *testing.T) {
	q := buildQueue(5, 5)

	//pop 3 items
	for i := 0; i < 3; i++ {
		val, success := q.Dequeue()
		if !success || val != i {
			t.Errorf("Unexpected dequeue value %v, expected 0", val)
		}
	}

	q.Enqueue(5)
	q.Enqueue(6)
	q.Enqueue(7)

	//pop 3 more items
	for i := 3; i < 6; i++ {
		val, success := q.Dequeue()
		if !success || val != i {
			t.Errorf("Unexpected dequeue value %v, expected 0", val)
		}
	}

	//make sure the queue contains the proper # of items
	if q.Count() != 2 {
		t.Errorf("Queue contains %v item when we only expected 2", q.Count())
	}

	//make sure the queue contains the expected values
	a := []int{6, 7}
	if !reflect.DeepEqual(q.ToArray(), a) {
		t.Errorf("Queue %v does not match %v", q.ToArray(), a)
	}
}

func buildQueue(size int, count int) *Queue {
	q := NewQueue(size)
	for i := 0; i < count; i++ {
		q.Enqueue(i)
	}
	return q
}
