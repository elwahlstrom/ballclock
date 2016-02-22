package lists

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	size := 5
	s := buildStack(size, 0)

	//make sure the stack is the proper size
	if s.Size() != size {
		t.Errorf("NewStack(%v) failed, stack.Size=%v which is not the expected value of %v", size, s.Size(), size)
	}
}

func TestStackPush(t *testing.T) {
	s := buildStack(5, 5)

	//make sure the stack contains the proper # of items
	if s.Count() != 5 {
		t.Errorf("Stack contains %v items when we only expected 5", s.Count())
	}

	//make sure the Stack contains the expected values
	a := []int{4, 3, 2, 1, 0}
	if !reflect.DeepEqual(s.ToArray(), a) {
		t.Errorf("Stack %v does not match %v", s.ToArray(), a)
	}
}

func TestStackPushFull(t *testing.T) {
	s := buildStack(5, 6)

	//make sure the Stack contains the proper # of items
	if s.Count() != 5 {
		t.Errorf("Stack contains %v item when we only expected 5", s.Count())
	}

	//make sure the stack contains the expected values
	a := []int{4, 3, 2, 1, 0}
	if !reflect.DeepEqual(s.ToArray(), a) {
		t.Errorf("Stack %v does not match %v", s.ToArray(), a)
	}
}

func TestStackPopEmpty(t *testing.T) {
	s := buildStack(5, 0)

	//pop the item and make sure it's the expected value
	_, success := s.Pop()
	if success {
		t.Error("Success was empty but got a 'success' on Pop")
	}

	//make sure the stack is now empty
	if s.Count() != 0 {
		t.Errorf("Expected an empty stack got %v", s.Count())
	}
}

func TestStackPopAll(t *testing.T) {
	s := buildStack(5, 5)

	//pop all items off the stack
	for i := 4; i >= 0; i-- {
		val, _ := s.Pop()
		if val != i {
			t.Errorf("Unexpected stack value %v, expected %v", val, i)
		}
	}

	//make sure the stack is now empty
	if s.Count() != 0 {
		t.Errorf("Expected an empty stack got %v", s.Count())
	}
}

func buildStack(size int, count int) *Stack {
	s := NewStack(size)
	for i := 0; i < count; i++ {
		s.Push(i)
	}
	return s
}
