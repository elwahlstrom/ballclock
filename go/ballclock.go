package main

import (
	"fmt"
	"lists"
	"math"
	"strconv"
	"time"
)

type BallClock struct {
	NumBalls      int
	Days          float64
	AreBallsInSeq bool
	_1minTrack    *lists.Stack
	_5minTrack    *lists.Stack
	_1hourTrack   *lists.Stack
	_queue        *lists.Queue
}

func NewBallClock(numBalls int) *BallClock {
	bc := &BallClock{numBalls, 0, false, lists.NewStack(4), lists.NewStack(11), lists.NewStack(11), lists.NewQueue(numBalls)}
	for i := 1; i <= numBalls; i++ {
		bc._queue.Enqueue(i)
	}
	return bc
}

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
				if math.Trunc(bc.Days) == bc.Days {
					bc.AreBallsInSeq = bc.checkSeqOrder()
				}
			}
		}
	}
}

func (bc *BallClock) emptyTrack(track *lists.Stack) {
	for track.Count() > 0 {
		ball, ok := track.Pop()
		if ok {
			bc._queue.Enqueue(ball)
		}
	}
}

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

func main() {
	fmt.Println("Enter the number of clock balls one per line ranging from 27 to 127 (0 to stop):")
	input := []string{}

	line := ""
	for line != "0" {
		line = ""
		fmt.Scanln(&line)
		if line != "0" {
			input = append(input, line)
		}
	}

	stime := time.Now()
	for _, bc := range input {
		numBalls, err := strconv.Atoi(bc)
		if err != nil || numBalls < 27 || numBalls > 127 {
			fmt.Printf("'%v' is not a valid number of balls!\n", bc)
			continue
		}

		clock := NewBallClock(numBalls)
		for {
			clock.Tick()
			if clock.AreBallsInSeq {
				fmt.Printf("%v balls cycle after %v days.\n", clock.NumBalls, clock.Days)
				break
			}
		}
	}
	fmt.Printf("Total execution time: %v", time.Now().Sub(stime))
}
