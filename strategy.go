package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func strategy(errType string) {
	switch errType {
	case "exit":
		fmt.Println("Crash")
		os.Exit(2)
	case "throw":
		fmt.Println("Throw an error")
		panic("Forced a panic attack")
	case "cpu":
		fmt.Println("Using max cpu cycles")
		loopForever()
	case "heap":
		fmt.Println("Recursively consume heap")
		var dummy []int
		abuseHeap(dummy)
	default:
		fmt.Println("No strategic match")
	}
}

func runFib() {
	f := fibonacci()
	j := 0
	var i int64 = 0
	for  ; i < 100000000000; i++ {
		j = f()
	}
	fmt.Println(j)
}

func fibonacci() func() int {
	x, y := 0, 1
	return func() (r int) {
		r = x
		x, y = y, x + y
		return
	}
}

// An "endless" loop which does nothing except check if it's time to abort.
//  see: https://stackoverflow.com/questions/41079492/how-to-artificially-increase-cpu-usage
func loopForever() {
	done := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * 10)
	close(done)
}

func abuseHeap(buf []int) {
	var slice = make([]int,4096*4096)
	abuseHeap(slice)

}