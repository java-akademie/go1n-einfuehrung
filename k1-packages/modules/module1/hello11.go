package module1

import "fmt"

// ExportedVar11 is a global variable ...
var ExportedVar11 = 11

var unexportedVar11 = 911

// Hello11 is a global function ...
func Hello11() {
	fmt.Println("----- Hello11 -----")

	fmt.Println(ExportedVar11)
	fmt.Println(ExportedVar12) // aus Hello12 global

	show11()
	show12() // same package
}

func show11() {
	fmt.Println("show11")
	fmt.Println(unexportedVar11)
	fmt.Println(unexportedVar12) // same package
}

func init() {
	fmt.Println("init Hello11")
}
