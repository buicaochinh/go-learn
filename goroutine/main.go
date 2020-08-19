package main

import (
	"sync"
	"runtime"
	// "time"
	"fmt"
)

func g1() {
	fmt.Println("g1")
	waitGroup.Done()
}

func g2() {
	fmt.Println("g2")
	waitGroup.Done()
}

func demoGoroutine() {
	// go func()
	// go g1()

	// numGo := runtime.NumGoroutine()
	// fmt.Println(numGo)
	// time.Sleep(time.Second)

	// Synchronized goroutines
	/**
		logic 1

		go g1()
		go g2()

		logic 2
	*/

	fmt.Println("Begin")
	waitGroup.Add(2)

	go g1()
	go g2()

	waitGroup.Wait()

	fmt.Println("End")
	runtime.GOMAXPROCS(2)

	numP := runtime.NumCPU()
	fmt.Println(numP)

	numP1 := runtime.GOMAXPROCS(0)
	fmt.Println(numP1)
}


var waitGroup sync.WaitGroup

func main() {
	queue := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			queue <- i
		}
		done <- true
	}()

	for {
		select {
		case v := <-queue:
			fmt.Println(v)
		case <-done:
			fmt.Println("done")
			return
		}
	}
}
