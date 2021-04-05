package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func anonymousFunctions() {

	myutils.H2("anonymous functions")

	func() { fmt.Println("anonymous function 1") }()

	f2 := func() { fmt.Println("anonymous function 2") }
	f2()

	f3 := func() { fmt.Println("anonymous function 3") }
	foo2(f3)

	myutils.Wait()
}

func foo2(function3 func()) {
	function3()
}
