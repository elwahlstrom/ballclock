package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	line := ""
	fmt.Println("Enter the number of clock balls one per line ranging from 27 to 127 (0 to stop):")

	for line != "0" {
		fmt.Scanln(&line)
		if line == "0" {
			break
		}

		numBalls, err := strconv.Atoi(line)
		if err != nil || numBalls < 27 || numBalls > 127 {
			fmt.Printf("'%v' is not a valid number of balls!\n", line)
			continue
		}

		stime := time.Now()
		clock, err := NewBallClock(numBalls)
		if err != nil {
			fmt.Println(err)
			continue
		}

		clock.Run()
		fmt.Printf("%v balls cycle after %v days. time=%v\n", clock.NumBalls, clock.Days, time.Now().Sub(stime))
	}
}
