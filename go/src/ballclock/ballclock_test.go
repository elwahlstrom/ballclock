package main

import (
	"reflect"
	"testing"
)

func TestNewBallClock(t *testing.T) {
	_, err1 := NewBallClock(26)
	if err1 == nil {
		t.Errorf("NewBallClock(26) lower bound check should of failed, min ball count is 27")
	}

	_, err2 := NewBallClock(128)
	if err2 == nil {
		t.Error("NewBallClock(128) upper bound check should of failed, max ball count is 127")
	}

	bc, err3 := NewBallClock(30)
	if err3 != nil || bc.NumBalls != 30 {
		t.Error("NewBallClock(30) should of succeeded")
	}
}

func Test1minTrack(t *testing.T) {
	bc, err := NewBallClock(27)
	if err != nil {
		t.Errorf("NewBallClock(27) failed, reason: %v", err)
	}

	bc.TickN(4)

	a := []int{4, 3, 2, 1}
	if !reflect.DeepEqual(bc._1minTrack.ToArray(), a) {
		t.Errorf("Clock 1 minute track %v does not match %v", bc._1minTrack.ToArray(), a)
	}
}

func Test5minTrack(t *testing.T) {
	bc, err := NewBallClock(27)
	if err != nil {
		t.Errorf("NewBallClock(27) failed, reason: %v", err)
	}

	bc.TickN(5)

	if bc._1minTrack.Count() > 0 {
		t.Errorf("Clock 1 minute track %v should be empty", bc._1minTrack.ToArray())
	}

	a := []int{5}
	if !reflect.DeepEqual(bc._5minTrack.ToArray(), a) {
		t.Errorf("Clock 5 minute track %v does not match %v", bc._5minTrack.ToArray(), a)
	}
}

func Test1hourTrack(t *testing.T) {
	bc, err := NewBallClock(27)
	if err != nil {
		t.Errorf("NewBallClock(27) failed, reason: %v", err)
	}

	bc.TickN(12 * 5)

	if bc._1minTrack.Count() > 0 {
		t.Errorf("Clock 1 minute track %v should be empty", bc._1minTrack.ToArray())
	}

	if bc._5minTrack.Count() > 0 {
		t.Errorf("Clock 5 minute track %v should be empty", bc._5minTrack.ToArray())
	}

	if bc._1hourTrack.Count() != 1 {
		t.Errorf("Clock 1 hour track %v should contain only 1 ball", bc._1hourTrack.ToArray())
	}
}

func Test12hourPeriod(t *testing.T) {
	bc, err := NewBallClock(27)
	if err != nil {
		t.Errorf("NewBallClock(27) failed, reason: %v", err)
	}

	bc.TickN(12 * 60)

	if bc._1minTrack.Count() > 0 {
		t.Errorf("Clock 1 minute track %v should be empty", bc._1minTrack.ToArray())
	}

	if bc._5minTrack.Count() > 0 {
		t.Errorf("Clock 5 minute track %v should be empty", bc._5minTrack.ToArray())
	}

	if bc._1hourTrack.Count() > 0 {
		t.Errorf("Clock 1 hour track %v should be empty", bc._5minTrack.ToArray())
	}

	if bc._queue.Count() != 27 {
		t.Errorf("Clock queue %v should contain 27 balls", bc._queue.ToArray())
	}
}

func Test30BallClock(t *testing.T) {
	bc, err := NewBallClock(30)
	if err != nil {
		t.Errorf("NewBallClock(30) failed, reason: %v", err)
	}

	bc.Run()
	if bc.Days != 15 {
		t.Errorf("NewBallClock(30).Run() got %v, expected 15 days", bc.Days)
	}
}

func Test45BallClock(t *testing.T) {
	bc, err := NewBallClock(45)
	if err != nil {
		t.Errorf("NewBallClock(45) failed, reason: %v", err)
	}

	bc.Run()
	if bc.Days != 378 {
		t.Errorf("NewBallClock(45).Run() got %v, expected 378 days", bc.Days)
	}
}
