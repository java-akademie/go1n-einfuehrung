package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func pointer() {
	myutils.H2("pointer")
	var counter int
	var counterPtr *int
	fmt.Println("pointer of the counter (reference) before init:", counterPtr)
	counter = 25
	counterPtr = &counter // reference of the counter

	fmt.Println("counter                                       :", counter)
	fmt.Println("pointer of the counter (reference) after init :", counterPtr)
	fmt.Println("de-referenced pointer of the counter          :", *counterPtr)

	f1(counter) // copy of variable
	fmt.Println("counter not changed", counter)
	f2(&counter) // copy of reference
	fmt.Println("counter     changed", counter)
	myutils.Wait()
}

func f1(c int) {
	c = c + 1
	fmt.Println(c)
}

func f2(c *int) {
	*c = *c + 1
	fmt.Println(*c)
}
