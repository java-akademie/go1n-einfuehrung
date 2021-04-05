package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func doSwitch() {
	myutils.H2("doSwitch")
	switch1(1)
	switch1(3)
	switch1(4)
	myutils.Wait()
	switch2(4, 5)
	switch2(5, 5)
	switch2(6, 5)
	myutils.Wait()
	switch3()
	myutils.Wait()

}

func switch1(s int) {
	myutils.Comment(fmt.Sprintf("switch1 s: %v", s))
	switch s {
	case 1:
		fmt.Println("one  ", s)
	case 2:
		fmt.Println("two  ", s)
	case 3:
		fmt.Println("three", s)
	default:
		fmt.Println("???  ", s)
	}
}

func switch2(a int, b int) {
	myutils.Comment(fmt.Sprintf("switch2 a:%v, b:%v", a, b))
	switch {
	case a < b:
		fmt.Printf("%v < %v \n", a, b)
	case a > b:
		fmt.Printf("%v > %v \n", a, b)
	default:
		fmt.Printf("%v = %v \n", a, b)
	}

}

func switch3() {
	myutils.Comment("switch3")
	for i := 1; i <= 10; i++ {
		switch s4 := myutils.GetRandom2(0, 5); s4 {
		case 1:
			fmt.Println(s4, "eins")
		case 2:
			fmt.Println(s4, "zwei")
		case 3:
			fmt.Println(s4, "drei")
		default:
			fmt.Println(s4, "??? ")
		}
	}
}
