package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func endOfLoop() {
	myutils.H2("end of Loop")

	a := [...]int{883, 232, 543, 564, 536, 222, 912, 817, 135, 888, 666}
	fmt.Printf("Typ: %T, %v \n", a, a)

	myutils.Description(
		"mark multiples of 3",
		"end the loop when the value is a multiple of 5",
	)

	for i, v := range a {
		if v%5 == 0 {
			fmt.Printf("index: %v, %v is a multiple of 5 -> break \n",
				i, v)
			break
		}
		if v%3 == 0 {
			fmt.Printf("index: %v, %v is a multiple of 3 -> continue\n",
				i, v)
			continue
		}
		fmt.Printf("index: %v, do something with value: %v\n",
			i, v)
	}

	myutils.Wait()
}
