package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func main() {
	myutils.H1("Test MyUtils")
	myutils.Version()
	doGetString()
	doGetInteger()
	doGetRandom()
	doGetRandom2()
}

func doGetString() {
	myutils.H2("doGetString()")
	s1 := myutils.GetString("input string 1")
	s2 := myutils.GetString("input string 2")
	fmt.Println(s1, s2)
	myutils.Wait()
}

func doGetInteger() {
	myutils.H2("doGetInteger()")
	a := myutils.GetInteger("input int 1")
	b := myutils.GetInteger("input int 2")
	fmt.Println(a, b)
	myutils.Wait()
}

func doGetRandom() {
	myutils.H2("doGetRandom()")
	for i := 1; i <= 5; i++ {
		r := myutils.GetRandom()
		fmt.Println(r)
	}
	myutils.Wait()
}

func doGetRandom2() {
	myutils.H2("doGetRandom2(5,15)")
	for i := 1; i <= 5; i++ {
		r := myutils.GetRandom2(5, 15)
		fmt.Println(r)
	}
	myutils.Wait()
}
