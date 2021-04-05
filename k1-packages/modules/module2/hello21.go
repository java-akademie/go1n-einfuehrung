package module2

import "fmt"

// ExportedVar21 is a global variable ...
var ExportedVar21 = 21

var unexportedVar21 = 921

// Hello21 is a global function ...
func Hello21() {
	fmt.Println("----- Hello21 -----")

	fmt.Println(ExportedVar21)
	fmt.Println(ExportedVar22) // aus Hello22 global

	show21()
	show22() // same package
}

func show21() {
	fmt.Println("show21")
	fmt.Println(unexportedVar21)
	fmt.Println(unexportedVar22) // same package
}

func init() {
	fmt.Println("init Hello21")
}
