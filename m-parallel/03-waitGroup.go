package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/java-akademie/myutils"
)

func doWaitGroup() {
	myutils.H2("waitGroup")

	start := time.Now().UnixNano()

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go longLastingFunction3(i, &wg)
	}

	wg.Wait()

	duration := float64(time.Now().UnixNano()-start) / 1000000000
	fmt.Println("Duration:", duration, " Seconds")

	myutils.Wait()
}

func longLastingFunction3(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	d := time.Duration(myutils.GetRandom2(3, 9))
	fmt.Printf("Function %3v lasts %2d Seconds - start \n", i, d)
	time.Sleep(d * time.Second) // do whatever the function should do
	fmt.Printf("Function %3v lasts %2d Seconds - stop \n", i, d)
}
