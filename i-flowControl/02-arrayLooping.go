package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func arrayLooping() {
	myutils.H2("array Looping")

	a := [...]int{7667, 2881, 1232, 8543}
	fmt.Printf("Typ: %T, %v \n", a, a)

	myutils.Comment("show index and value")
	for i, v := range a {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}

	myutils.Comment("show only value")
	for _, v := range a {
		fmt.Printf("value: %v\n", v)
	}
	myutils.Wait()
}
