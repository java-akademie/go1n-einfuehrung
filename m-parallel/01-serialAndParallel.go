package main

import (
	"fmt"
	"time"

	"github.com/java-akademie/myutils"
)

func doSerial() {
	myutils.H2("serial")

	start := time.Now().UnixNano()

	for i := 1; i <= 5; i++ {
		// each will be running 3 to 9 seconds
		longLastingFunction1(i, 3, 9)
	}

	// has been running about 30 seconds
	stop := time.Now().UnixNano()
	duration := float64((stop - start)) / 1000000000
	fmt.Println("Duration: ", duration)

	myutils.Wait()
}

func doParallel() {
	myutils.H2("parallel")

	for i := 1; i <= 100; i++ {
		// each will be running 3 to 9 seconds
		// as they run parallel the whole loop
		// will not be running longer than 9 seconds
		go longLastingFunction1(i, 3, 9)
	}

	// longLastingFunction should be able to come to an end
	// so main() has to sleep a little longer than 9 seconds
	time.Sleep(11 * time.Second)

	myutils.Wait()
}

func longLastingFunction1(i, min, max int) {
	d := time.Duration(myutils.GetRandom2(min, max))
	fmt.Printf("longlastingFunction %3v start: will  be running %2d Seconds\n", i, d)
	time.Sleep(d * time.Second) // simulation of doing something important
	fmt.Printf("longlastingFunction %3v stop : has been running %2d Seconds\n", i, d)
}
