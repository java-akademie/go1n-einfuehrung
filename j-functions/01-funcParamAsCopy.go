package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func paramAsCopy() {
	myutils.H2("param as copy")

	counter := 42

	myutils.Comment("add 100 to counter 42, parameter as copy")
	addCounter1(counter)
	fmt.Println("after function", counter)

	myutils.Wait()
}

func addCounter1(c int) {
	c = c + 100
	fmt.Println("in function", c)
}
