package main

import (
	"fmt"

	tools "github.com/java-akademie/myutils"
)

func variable1() {
	tools.H2("Variable1 (varname mit Nullwerten)")
	tools.Version()

	var counter int      // 0
	var price float64    // 0.0
	var message string   // ""
	var ok bool          // false
	var dash byte        // 0
	var chineseSign rune // 0

	fmt.Println(counter)
	fmt.Println(price)
	fmt.Println(message)
	fmt.Println(ok)
	fmt.Println(dash)
	fmt.Println(chineseSign)
	tools.Wait()
}

func variable2() {
	tools.H2("Variable2 (var varname = init)")

	var counter = 42
	var price = 12.45
	var message = "this is a string"
	var ok = false
	ok = true
	var dash = 'A'
	var chineseSign = '世'

	fmt.Println(counter)
	fmt.Println(price)
	fmt.Println(message)
	fmt.Println(ok)
	fmt.Println(string(dash), dash)
	fmt.Println(string(chineseSign), chineseSign)
	tools.Wait()

	var gesellschaftlicheGruppe = '世'
	var generationen = '界'
	var welt = "世界"
	fmt.Println(string(gesellschaftlicheGruppe), gesellschaftlicheGruppe)
	fmt.Println(string(generationen), generationen)
	fmt.Println("welt", welt, len(welt))
	tools.Wait()
}

func variable3() {
	tools.H2("Variable3 (varname := init)")

	counter := 42
	price := 12.45
	message := "this is a string"
	ok := false
	ok = true
	dash := 'A'
	chineseSign := '世'

	fmt.Println(counter)
	fmt.Println(price)
	fmt.Println(message)
	fmt.Println(ok)
	fmt.Println(dash)
	fmt.Println(string(chineseSign))
	tools.Wait()

	gesellschaftlicheGruppe := '世'
	generationen := '界'
	welt := "世界"
	fmt.Println(string(gesellschaftlicheGruppe), gesellschaftlicheGruppe)
	fmt.Println(string(generationen), generationen)
	fmt.Println("welt", welt, len(welt))
	tools.Wait()
}
