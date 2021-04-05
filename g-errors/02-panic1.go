package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

var a = []int{10, 11, 12, 13}

func doPanic1() {
	myutils.H2("Panic1")
	show(0)
	show(1)
	show(10) // not save
	show(2)
	show(3)
	myutils.Wait()
}

func show(i int) {
	// panic handler
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("foo(%v): err: %v \n", i, err)
		}
	}()

	fmt.Printf("foo(%d): %d \n", i, a[i])
}
