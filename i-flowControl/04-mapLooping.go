package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func mapLooping() {
	myutils.H2("map Looping")

	m := map[string]string{"blau": "blue", "red": "rot", "yellow": "gelb"}
	for i := 1; i <= 5; i++ {
		m[fmt.Sprintf("color-%d", i)] = fmt.Sprintf("farbe-%d", i)
	}
	fmt.Printf("Typ: %T, len %v, %v \n", m, len(m), m)

	myutils.Comment("show key and value")
	for k, v := range m {
		fmt.Printf("key: %10v, value: %10v\n", k, v)
	}

	myutils.Comment("show only key")
	for k := range m {
		fmt.Printf("key: %10v\n", k)
	}

	myutils.Comment("show only value")
	for _, v := range m {
		fmt.Printf("value: %10v\n", v)
	}

	myutils.Wait()
}
