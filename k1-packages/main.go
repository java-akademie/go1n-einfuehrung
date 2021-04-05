package main

import (
	"fmt"

	mods "github.com/java-akademie/go1n-einfuehrung/k1-packages/modules"
)

func init() {
	fmt.Println("init main 1")
}

func main() {
	fmt.Println("=== main ===")
	mods.Module1()
	mods.Module2()
}

func init() {
	fmt.Println("init main 2")
}
