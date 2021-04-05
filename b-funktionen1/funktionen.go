package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func main() {
	myutils.H1("start main")
	myutils.Version()
	printHalloWelt()
	fmt.Println(halloWelt())

	sum, diff := compute(17, 4)
	fmt.Println(sum, diff)
	fmt.Println(compute(12, 4))
}

func printHalloWelt() {
	fmt.Println("hallo Welt 1")
}

func halloWelt() string {
	return "hallo Welt 2"
}

func compute(a int, b int) (int, int) {
	return a + b, a - b
}
