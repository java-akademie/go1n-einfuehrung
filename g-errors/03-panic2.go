package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func doPanic2() {
	myutils.H2("Panic2 - divide with panic")
	doDivideWithPanic(144, 12, false)
	doDivideWithPanic(144, 0, true)
	doDivideWithPanic(222, 5, false)
	doDivideWithPanic(222, 0, false)
	myutils.Wait()
}

func doDivideWithPanic(dnd int, dor int, handle bool) {
	if handle {
		// panic is handled
		defer func() {
			err := recover()
			if err != nil {
				fmt.Printf("%v / %v = err: %v \n", dnd, dor, err)
			}
		}()
	}

	erg := divideWithPanic(dnd, dor) // not save, panic() could occur
	fmt.Printf("%v / %v = %v \n", dnd, dor, erg)
}

func divideWithPanic(dnd int, dor int) int {
	if dor == 0 {
		panic("divisor may not be null")
	}
	return dnd / dor
}
