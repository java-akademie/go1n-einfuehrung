package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func paramAsPointer() {
	myutils.H2("param as pointer")

	counter := 42

	myutils.Comment("add 100 to counter 42, parameter as copy")
	addCounter2(&counter)
	fmt.Println("after function", counter)

	myutils.Wait()
}

func addCounter2(c *int) {
	*c = *c + 100
	fmt.Println("in function", *c)
}
