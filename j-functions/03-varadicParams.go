package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func variadicParams() {
	myutils.H2("variadic params")

	foo1("bill", "vienna")
	foo1("jack", "berlin", 50)
	foo1("mike", "london", 20, 12, 2000)

	myutils.Wait()
}

func foo1(name, addr string, i ...int) {
	fmt.Println(name, addr, i)
}
