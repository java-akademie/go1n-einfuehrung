package module2

import "fmt"

// ExportedVar22 is a global variable ...
var ExportedVar22 = 22

var unexportedVar22 = 922

// Hello22 is a global function ...
func Hello22() {
	fmt.Println("----- Hello22 -----")

	fmt.Println(ExportedVar21) // aus mHello21 global
	fmt.Println(ExportedVar22)

	show21() // same package
	show22()
}

func show22() {
	fmt.Println("show22")
	fmt.Println(unexportedVar21) // same package
	fmt.Println(unexportedVar22)
}

func init() {
	fmt.Println("init Hello22")
}
