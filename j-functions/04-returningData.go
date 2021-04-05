package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func returningData() {
	myutils.H2("returning data")

	var a, b, sum, diff, prod, quot float64

	a, b = 17, 4
	sum, diff, prod, quot = compute1(a, b)
	ausgabe(a, b, sum, diff, prod, quot)

	a, b = 11, 14
	sum, diff, prod, quot = compute2(a, b)
	ausgabe(a, b, sum, diff, prod, quot)

	a, b = 321, 45
	sum, diff, prod, quot = compute3(a, b)
	ausgabe(a, b, sum, diff, prod, quot)

	myutils.Wait()
}
func ausgabe(a, b, s, d, p, q float64) {
	fmt.Printf("args: (%v,%v), sum=%v, diff=%v, prod=%v, quot=%v \n",
		a, b, s, d, p, q)
}

// return values
func compute1(a, b float64) (float64, float64, float64, float64) {
	myutils.Comment("return Values")
	return a + b, a - b, a * b, a / b
}

func compute2(a, b float64) (sum float64, diff float64, prod float64, quot float64) {
	myutils.Comment("named returns but return values")
	return a + b, a - b, a * b, a / b
}

// named returns
func compute3(a, b float64) (sum float64, diff float64, prod float64, quot float64) {
	myutils.Comment("named returns")
	sum, diff, prod, quot = a+b, a-b, a*b, a/b
	return
}
