package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func arrays1() {
	myutils.H1("Arrays1")
	myutils.Wait()

	var arr1 [3]int // 3 Elements: Index 0-2 (100,0,122)
	arr1[0] = 100
	arr1[2] = 122

	var arr2 = [6]int{22, 42, 13}        // 6 Elements: Index 0-5 (22,42,13,0,0,0)
	var arr3 = [...]string{"urs", "uwe"} // 2 Elements: Index 0-1
	arr4 := [4]float64{11.25, 21}        // 4 Elements: Index 0-3 (11.25, 21, 0, 0)
	arr5 := [...]int32{55, 21, 31}       // 3 Elements: Index 0-2

	arr6 := arr2

	fmt.Println("arr1", arr1)
	fmt.Println("arr2", arr2)
	fmt.Println("arr3", arr3)
	fmt.Println("arr4", arr4)
	fmt.Println("arr5", arr5)
	fmt.Println("arr6", arr6)

	arr1[1] = 111
	arr2[1] = 555000
	arr4[3] = 3.5
	arr6[1] = 666000

	myutils.Br()

	fmt.Println("arr1", arr1)
	fmt.Println("arr2", arr2)
	fmt.Println("arr4", arr4)
	fmt.Println("arr6", arr6)
}
