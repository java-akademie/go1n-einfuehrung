package modules

import (
	"fmt"

	mod1 "github.com/java-akademie/go1n-einfuehrung/k1-packages/modules/module1"
)

func Module1() {
	fmt.Println("----- Module1 -----")
	fmt.Println(mod1.ExportedVar11)
	fmt.Println(mod1.ExportedVar12)
	mod1.Hello11()
	mod1.Hello12()
}

func init() {
	fmt.Println("init Module1")
}
