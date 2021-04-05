package main

import (
	"strconv"

	"github.com/java-akademie/myutils"
)

func slices2() {
	myutils.H1("slices2")

	copyByAssignment()
	copyByCopy()
}

func copyByAssignment() {
	myutils.H2("copy by assignment")
	myutils.Wait()

	var n []string
	myutils.ShowSlice("n", n, cap(n), len(n))
	for i := 0; i < 10; i++ {
		n = append(n, "n"+strconv.Itoa(i))
		//	showSlice("n", n, cap(n), len(n))
	}

	m := n[0:3] // 0, 1, 2
	n[1] = "11"
	myutils.ShowSlice("n", n, cap(n), len(n))
	myutils.ShowSlice("m", m, cap(m), len(m))

	m = append(m, "a1", "b1", "c1", "d1", "e1")
	n[2] = "22"
	myutils.ShowSlice("n", n, cap(n), len(n))
	myutils.ShowSlice("m", m, cap(m), len(m))

	m = append(m, "x2", "y2", "z2")
	n[3] = "33"
	myutils.ShowSlice("n", n, cap(n), len(n))
	myutils.ShowSlice("m", m, cap(m), len(m))

	m = append(m, "f3", "g3", "h3", "i3", "j3")
	n[4] = "44"
	myutils.ShowSlice("n", n, cap(n), len(n))
	myutils.ShowSlice("m", m, cap(m), len(m))
}

func copyByCopy() {
	myutils.H2("copy by copy()")
	myutils.Wait()

	var n []int
	myutils.ShowSlice("n", n, cap(n), len(n))
	for i := 0; i < 10; i++ {
		n = append(n, i)
	}

	myutils.ShowSlice("n", n, cap(n), len(n))

	c := make([]int, 20)
	copy(c[:], n[:])
	c[0] = 44
	n[0] = 33
	myutils.ShowSlice("n", n, cap(n), len(n))
	myutils.ShowSlice("c", c, cap(c), len(c))

	myutils.Comment("shrink to 5 elements")

	s := make([]int, 5)
	copy(s[:], c[3:8])
	c[0] = 44
	s[0] = 33
	myutils.ShowSlice("c", c, cap(c), len(c))
	myutils.ShowSlice("s", s, cap(s), len(s))
}
