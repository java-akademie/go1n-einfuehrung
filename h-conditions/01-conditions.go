package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func doConditions() {
	myutils.H2("doConditions")

	myutils.Comment("in range 3...7")
	for a := 1; a <= 8; a++ {
		fmt.Printf("%2v in range 3...7 %v \n",
			a, a >= 3 && a <= 7)
	}
	myutils.Wait()

	myutils.Comment("out of range 3...7")
	for a := 1; a <= 8; a++ {
		fmt.Printf("%2v out of range 3...7 %v \n",
			a, a < 3 || a > 7)
	}
	myutils.Wait()

	myutils.Comment("gt 10 and mod 2")
	for a := 5; a <= 15; a++ {
		fmt.Printf("%2v gt 10 and even, %v \n",
			a, a >= 10 && a%2 == 0)
	}
	myutils.Wait()

	myutils.Comment("le 5 and mod 2 or ge 10 and mod 3")
	for a := 1; a <= 15; a++ {
		fmt.Printf("%2v le 5 and mod 2 or ge 10 amd mod 3, %v \n",
			a, a <= 5 && a%2 == 0 || a >= 10 && a%3 == 0)
	}
	myutils.Wait()
}
