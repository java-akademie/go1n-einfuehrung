package module1

import "fmt"

// ExportedVar12 is a global variable ...
var ExportedVar12 = 12

var unexportedVar12 = 912

// Hello12 is a global function ...
func Hello12() {
	fmt.Println("----- Hello12 -----")

	fmt.Println(ExportedVar11)
	fmt.Println(ExportedVar12)

	show11() // same package
	show12()
}

func show12() {
	fmt.Println("show12")
	fmt.Println(unexportedVar11) // same package
	fmt.Println(unexportedVar12)
}

func init() {
	fmt.Println("init Hello12")
}
