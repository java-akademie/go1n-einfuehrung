package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func doStructWithTypedef() {
	myutils.H2("struct with typedef")
	myutils.Wait()

	// define the type
	type person struct {
		pnr  int
		name string
	}

	// var declaration
	var s1 person
	var s2 person
	fmt.Println(s1, s2)

	// filling the var (the elements)
	s1.pnr = 41
	s1.name = "fred"

	// filling the var (literal)
	s2 = person{42, "barney"}

	fmt.Println(s1, s2)

	// var declaration and filling
	var s3 = person{pnr: 43, name: "wilma"}
	s4 := person{pnr: 44, name: "betty"}
	fmt.Println(s3, s4)

	myutils.Comment("ShowObjects (one line)")
	myutils.ShowObjects(s1, s2, s3, s4)

	myutils.Comment("ShowObject (four lines)")
	myutils.ShowObject("s1", s1)
	myutils.ShowObject("s2", s2)
	myutils.ShowObject("s3", s3)
	myutils.ShowObject("s4", s4)
}
