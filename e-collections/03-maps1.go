package main

import (
	"github.com/java-akademie/myutils"
)

func maps1() {
	myutils.H1("Maps1")
	declareMap()
	populateMap()
}

func declareMap() {
	myutils.H2("declareMap")
	myutils.Wait()

	var p1 = map[int]string{}
	p1[11] = "p1"
	var p2 = make(map[int]string)
	p2[22] = "p2"
	p3 := map[int]string{1: "ann", 2: "sue"}
	p3[33] = "p3"

	myutils.ShowObject("p1", p1)
	myutils.ShowObject("p2", p2)
	myutils.ShowObject("p3", p3)
}

func populateMap() {
	myutils.H2("populateMap")
	myutils.Wait()

	var c = map[int]string{}
	myutils.ShowObject("c", c)
	c[10] = "black"
	c[20] = "white"
	c[30] = "red"
	c[30] = "green"
	c[40] = "yellow"
	myutils.ShowObject("c", c)

	myutils.Comment("delete c key:30")
	delete(c, 30)
	myutils.ShowObject("c", c)

	myutils.Comment("insert a few c")
	for i := 1; i <= 10; i++ {
		c[i] = "c" + myutils.IntegerToString(i)
	}
	myutils.ShowObject("c", c)
}
