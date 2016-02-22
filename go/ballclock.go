package main

import (
	"./lists"
	"fmt"
)

type BallClock struct {
	NumBalls    int
	Days        float64
	InSeq       bool
	_1minTrack  *lists.Stack
	_5minTrack  *lists.Stack
	_1hourTrack *lists.Stack
	_queue      *lists.Queue
}

// initializes a new ball clock
func NewBallClock(numBalls int) (*BallClock, error) {
	if numBalls < 27 || numBalls > 127 {
		return nil, fmt.Errorf("'%v' is not a valid number of balls, must range from 27 to 127!", numBalls)
	}
	
	bc := &BallClock{numBalls, 0, false, lists.NewStack(4), lists.NewStack(11), lists.NewStack(11), lists.NewQueue(numBalls)}
	for i := 1; i <= numBalls; i++ {
		bc._queue.Enqueue(i)
	}
	return bc, nil
}

// run the clock until all balls are in sequence
func (bc *BallClock) Run() {
	for {
		bc.Tick()
		if bc.InSeq {
			break
		}
	}
}

// cycle the ball clock N times
func (bc *BallClock) TickN(n int) {
	for i := 0; i < n; i++ {
		bc.Tick()
	}
}

// cycle the ball clock with a ball in the queue
func (bc *BallClock) Tick() {
	ball, ok := bc._queue.Dequeue()
	if !ok {
		return
	}

	if bc._1minTrack.Count() < 4 {
		bc._1minTrack.Push(ball)
	} else {
		bc.emptyTrack(bc._1minTrack)
		if bc._5minTrack.Count() < 11 {
			bc._5minTrack.Push(ball)
		} else {
			bc.emptyTrack(bc._5minTrack)
			if bc._1hourTrack.Count() < 11 {
				bc._1hourTrack.Push(ball)
			} else {
				bc.emptyTrack(bc._1hourTrack)
				bc._queue.Enqueue(ball)

				bc.Days += .5
				bc.InSeq = bc.checkSeqOrder()
			}
		}
	}
}

// empty the balls from the track and put them back in the queue
func (bc *BallClock) emptyTrack(track *lists.Stack) {
	for track.Count() > 0 {
		ball, ok := track.Pop()
		if ok {
			bc._queue.Enqueue(ball)
		}
	}
}

// determine if the balls inthe queue are in sequence
func (bc *BallClock) checkSeqOrder() bool {
	i := 1
	for _, ball := range bc._queue.ToArray() {
		if i != ball {
			return false
		}
		i++
	}
	return true
}