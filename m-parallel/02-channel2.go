package main

import (
	"fmt"
	"time"

	"github.com/java-akademie/myutils"
)

func doChannel2() {
	myutils.H2("channel2")

	channel1 := make(chan string)
	channel2 := make(chan int)

	max := 5

	for i := 1; i <= max; i++ {
		go longLastingFunction22(i, channel1, channel2)
	}

	// no sleep necessary

	for i := 1; i <= max*3; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("ch1: ", msg1)
		case msg2 := <-channel2:
			fmt.Println("ch2: ", msg2)
		}
	}

	myutils.Wait()
}

func longLastingFunction22(i int, ch1 chan string, ch2 chan int) {
	d := time.Duration(myutils.GetRandom2(3, 9))
	ch1 <- fmt.Sprintf("Function %3v started: will  be running %2d Seconds", i, d)
	time.Sleep(d * time.Second)
	ch1 <- fmt.Sprintf("Function %3v stopped: has been running %2d Seconds", i, d)
	ch2 <- i
}
