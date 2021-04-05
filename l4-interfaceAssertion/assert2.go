package main

import (
	"github.com/java-akademie/myutils"
)

var (
	s2 = "Hallo Welt"
	i2 = 42
	//
	toCheck21 interface{} = s2
	toCheck22 interface{} = i2
)

func doAssert2() {
	myutils.H2("assert2")
	myutils.ShowObjects(toCheck21, toCheck22)
	// c := toCheck21.(int)    // NOK - panic
	// d := toCheck22.(string) // NOK - panic
	// myutils.ShowObjects(c, d)
	myutils.Wait()
}
