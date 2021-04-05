package main

import (
	"fmt"

	utils "github.com/java-akademie/myutils"
)

func main() {

	utils.H1("halloWelt")
	utils.Version()

	var a = 5.0
	var b = 7.0

	var sum = a + b
	var diff = a - b
	var prod = a * b
	var quot = a / b

	fmt.Println("sum : ", sum)
	fmt.Println("diff: ", diff)
	fmt.Println("prod: ", prod)
	fmt.Println("quot: ", quot)

	utils.Wait()
}
