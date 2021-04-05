package main

import (
	"github.com/java-akademie/myutils"
)

func doStructs() {
	myutils.H2("structs")
	myutils.Wait()

	var s8 = struct {
		pnr  int
		name string
	}{7, "bob"}
	s8.pnr = 8
	s8.name = "holger"

	s9 := s8
	s9.pnr = 9
	s9.name = "george"

	myutils.ShowObjects(s8, s9)
}
