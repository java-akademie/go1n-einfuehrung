package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func doStructWithoutTypedef() {
	myutils.H2("struct without tyedef")
	myutils.Wait()

	var s5 struct {
		pnr  int
		name string
	}
	s5.pnr = 5
	s5.name = "bill"

	var s6 = struct {
		pnr  int
		name string
	}{6, "greg"}

	s7 := struct {
		pnr  int
		name string
	}{7, "mike"}

	fmt.Println(s5, s6, s7)
	myutils.ShowObjects(s5, s6, s7)

}
