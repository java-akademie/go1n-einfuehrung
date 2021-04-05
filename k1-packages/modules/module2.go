package modules

import (
	"fmt"

	mod2 "github.com/java-akademie/go1n-einfuehrung/k1-packages/modules/module2"
)

func Module2() {
	fmt.Println("----- Module2 -----")
	fmt.Println(mod2.ExportedVar21)
	fmt.Println(mod2.ExportedVar22)
	mod2.Hello21()
	mod2.Hello22()
}

func init() {
	fmt.Println("init Module2")
}
