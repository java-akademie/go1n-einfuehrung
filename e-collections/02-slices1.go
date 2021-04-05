package main

import (
	"strconv"

	"github.com/java-akademie/myutils"
)

func slices1() {
	myutils.H1("slices1")

	declareSlice()
	growSlice()
}

func declareSlice() {
	myutils.H2("declare")
	myutils.Wait()

	myutils.Comment("declare an empty string-slice (len=0, cap=0)")
	var e []string
	myutils.ShowSlice("e", e, cap(e), len(e))

	myutils.Comment("declare an empty int-slice (len=10, cap=10)")
	f := make([]int, 10)
	myutils.ShowSlice("f", f, cap(f), len(f))

	myutils.Comment("declare an empty int-slice (len=5, cap=100) and populate it")
	g := make([]int, 5, 100)
	g[0] = 42
	g[4] = 114
	myutils.ShowSlice("g", g, cap(g), len(g))

	myutils.Comment("slice literals")
	s := []int{1, 2, 3}
	myutils.ShowSlice("s", s, cap(s), len(s))
	s = []int{}
	myutils.ShowSlice("s", s, cap(s), len(s))
	s = []int{11, 22, 33, 44}
	myutils.ShowSlice("s", s, cap(s), len(s))

	t := []string{"x"}
	t = append(t, "f")
	myutils.ShowSlice("t", t, cap(t), len(t))

	t = []string{"a", "b", "c", "d", "e"}
	t = append(t, "f")
	myutils.ShowSlice("t", t, cap(t), len(t))
}

func growSlice() {
	myutils.H2("how a slice grows")
	myutils.Wait()

	var n []string
	myutils.ShowSlice("n", n, cap(n), len(n))
	for i := 0; i < 10; i++ {
		n = append(n, "n"+strconv.Itoa(i))
		myutils.ShowSlice("n", n, cap(n), len(n))
	}
}
