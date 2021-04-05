package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func doAssert3() {
	myutils.H2("assert3")
	assert3("hallo welt")
	assert3(42)
	assert3(12.5)
}

func assert3(x interface{}) {
	myutils.ShowObjects(x)

	defer myutils.Wait()

	// string expected
	v1, ok := x.(string)
	if ok {
		fmt.Printf("variable ist ein string: %v \n", v1)
		return
	} else {
		fmt.Printf("variable ist kein string  sondern ein %8T, %v \n", x, x)
	}

	// int expected
	v2, ok := x.(int)
	if ok {
		fmt.Printf("variable ist ein int: %v \n", v2)
		return
	} else {
		fmt.Printf("variable ist kein int     sondern ein %8T, %v \n", x, x)
	}

	// float64 expected
	v3, ok := x.(float64)
	if ok {
		fmt.Printf("variable ist ein float64: %v \n", v3)
		return
	} else {
		fmt.Printf("variable ist kein float64 sondern ein %8T, %v \n", x, x)
	}

}
