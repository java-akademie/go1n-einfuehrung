package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func arrays2() {
	myutils.H1("Arrays2")
	myutils.Wait()

	var arr1 [3]string
	arr1[0] = "urs"
	arr1[1] = "lutz"
	arr1[2] = "beat"

	myutils.Comment("change array - param as copy")
	fmt.Println("before", arr1)
	changeArr1(arr1, 1, "lorenz")
	fmt.Println("after ", arr1)

	myutils.Comment("change array - param as reference")
	fmt.Println("before", arr1)
	changeArr2(&arr1, 1, "lorenz")
	fmt.Println("after ", arr1)

	myutils.Comment("show more information of an object")
	myutils.ShowObject("arr1", arr1)
	myutils.ShowObject("int", 111)
	myutils.ShowObject("char", 'a')
	myutils.ShowObject("float", -2.5)

	myutils.Wait()
}

func changeArr1(a [3]string, index int, value string) {
	// argument 'a' is a copy of the Param
	// parameter is not changed
	a[index] = value
}

func changeArr2(a *[3]string, index int, value string) {
	// argument 'a' is a reference to the Param
	// parameter is changed
	(*a)[index] = value
}
