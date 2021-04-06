package main

import (
	"fmt"
	"time"

	"github.com/java-akademie/myutils"
)

func doChannel1() {
	myutils.H2("channel1")

	channel1 := make(chan string)
	channel2 := make(chan int)

	max := 5

	for i := 1; i <= max; i++ {
		go longLastingFunction21(i, channel1, channel2)
	}

	// no sleep necessary

	for i := 1; i <= max*2; i++ {
		// wartet bis der channel 1 liefert
		fmt.Println("ch1: ", <-channel1)
	}

	for i := 1; i <= max; i++ {
		// wartet bis der channel 2 liefert
		fmt.Println("ch2: ", <-channel2)
	}

	myutils.Wait()
}

func longLastingFunction21(i int, ch1 chan string, ch2 chan int) {
	d := time.Duration(myutils.GetRandom2(3, 9))
	ch1 <- fmt.Sprintf("Function %3v started: will  be running %2d Seconds", i, d)
	time.Sleep(d * time.Second)
	ch1 <- fmt.Sprintf("Function %3v stopped: has been running %2d Seconds", i, d)
	ch2 <- i
}
