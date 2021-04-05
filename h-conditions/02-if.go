package main

import (
	"fmt"
	"math"

	"github.com/java-akademie/myutils"
)

func doIf() {
	myutils.H2("doIf")

	if1(2)
	if1(4)
	if1(22)

	if2(2)
	if2(3)
	if2(5)

	if3(3, 3, 400)
	if3(6, 4, 400)

}

func if1(a int) {
	myutils.Comment(fmt.Sprintf("if1 a: %v", a))

	if a == 1 {
		fmt.Println("1")
	} else if a == 2 {
		fmt.Println("2")
	} else if a == 3 {
		fmt.Println("3")
		return
	} else {
		fmt.Println("not 1,2,3")
		if a == 4 {
			fmt.Println("4")
		} else if a == 5 {
			fmt.Println("5")
		} else {
			fmt.Println("???")
		}
	}
	myutils.Wait()
}

func if2(a int) {
	myutils.Comment(fmt.Sprintf("if2 a: %v", a))

	if a == 1 {
		fmt.Println("1")
	} else {
		fmt.Println("not 1")
		if a == 2 {
			fmt.Println("2")
		} else {
			fmt.Println("not 2")
			if a == 3 {
				fmt.Println("3")
			} else {
				fmt.Println("not 3")
				fmt.Println("???")
			}
		}
	}
	myutils.Wait()
}

func if3(b, e, limit float64) {
	myutils.Comment(
		fmt.Sprintf(
			"if32 %v hoch %v, limit %v",
			b, e, limit,
		),
	)

	if v := math.Pow(b, e); v < limit {
		fmt.Println(v)
	} else {
		fmt.Println(limit, v)
	}
	myutils.Wait()
}
