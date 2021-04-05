package main

import (
	"fmt"
	"math"

	"github.com/java-akademie/myutils"
)

func constants() {
	myutils.H2("constants")
	const1() // expression
	const2() // iota
	const3() // iota
	myutils.Wait()
}

func const1() {
	myutils.H2("const1")
	const (
		u = 40_000
		d = u / math.Pi
		a = 444
		b
		c
	)
	fmt.Printf("Erdumfang: %v, Erddurchmesser: %v \n", u, d)
	fmt.Printf("a:%v, b:%v, c:%v \n", a, b, c)
}

func const2() {
	myutils.H2("const2 (iota)")
	const (
		a = iota + 111
		b
		c
		d
		e
	)
	fmt.Println(a, b, c, d, e)
}

func const3() {
	myutils.H2("const3 (iota)")
	const (
		_ = iota + 110
		a
		b
		c
		d
		e
	)
	fmt.Println(a, b, c, d, e)
}
