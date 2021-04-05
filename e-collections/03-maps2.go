package main

import (
	"github.com/java-akademie/myutils"
)

func maps2() {
	myutils.H1("Maps2")

	copyMapByAssignment()
	copyMapData()
}

func copyMapByAssignment() {
	myutils.H2("copy map by assignment")
	myutils.Wait()

	var c = map[int]string{}
	c[10] = "black"
	c[20] = "white"
	c[30] = "red"
	myutils.ShowObject("c", c)

	var d = c
	d[30] = "yellow"
	d[40] = "green"
	d[50] = "blue"

	myutils.ShowObject("c", c)
	myutils.ShowObject("d", d)
}

func copyMapData() {
	myutils.H2("copyMapData")
	myutils.Wait()

	var c = map[int]string{}
	c[10] = "black"
	c[20] = "white"
	c[30] = "red"

	var d = map[int]string{}
	for key, value := range c {
		d[key] = value
	}
	d[30] = "yellow"
	d[40] = "green"
	d[50] = "blue"

	myutils.ShowObject("c", c)
	myutils.ShowObject("d", d)
}
