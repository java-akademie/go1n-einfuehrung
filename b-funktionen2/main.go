package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func main() {
	myutils.H1("b-function2.main()")
	myutils.Version()
	test1()
	test2()
}
func test1() {
	myutils.H2("test1")
	printHelloWorld()
	fmt.Println(helloWorld())
	fmt.Println(compute(17, 4))
	add, diff := compute(22, 32)
	fmt.Println(add, diff)
}
func test2() {
	myutils.H2("test2")
	show(11)
	show(11.25)
	show("uwe")
	show(true)

	r := rune(64000)
	a := [5]int{1, 2, 3}
	s := []int{4, 5, 6}
	m := map[int]string{1: "ann", 2: "sue"}
	p := struct {
		pnr  int
		name string
	}{42, "bob"}

	show(r)
	show(a)
	show(s)
	show(m)
	show(p)
}
