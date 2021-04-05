package main

import (
	"github.com/java-akademie/myutils"
)

var (
	s1 = "Hallo Welt"
	i1 = 42
	//
	toCheck11 interface{} = s1
	toCheck12 interface{} = i1
)

func doAssert1() {
	myutils.H2("assert1")
	myutils.ShowObjects(toCheck11, toCheck12)
	a := toCheck11.(string) // OK
	b := toCheck12.(int)    // OK
	myutils.ShowObjects(a, b)
	myutils.Wait()
}
