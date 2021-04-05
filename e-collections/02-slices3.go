package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func slices3() {
	myutils.H1("Slices3")

	copyScliceOnlyPointer()
	copySliceNewArray()
	copyPartOfSlice()
}

func copyScliceOnlyPointer() {
	myutils.H2("copy a slice (only the pointer)")
	myutils.Wait()

	h := make([]int, 3, 5)
	h[0] = 42
	h[1] = 43
	i := h
	h[1] = 44
	i[2] = 45
	fmt.Println(cap(h), len(h), h[:cap(h)])
	fmt.Println(cap(i), len(i), i[:cap(i)])
	h = append(h, 46)
	h[0] = 1
	i[0] = 2
	fmt.Println(cap(h), len(h), h[:cap(h)])
	fmt.Println(cap(i), len(i), i[:cap(i)])
	i = append(i, 47)
	i[0] = 2
	h[0] = 1
	fmt.Println(cap(h), len(h), h[:cap(h)])
	fmt.Println(cap(i), len(i), i[:cap(i)])
	h = append(h, 48)
	h[0] = 1
	i[0] = 2
	fmt.Println(cap(h), len(h), h[:cap(h)])
	fmt.Println(cap(i), len(i), i[:cap(i)])
	i = append(i, 49)
	i[0] = 2
	h[0] = 1
	fmt.Println(cap(h), len(h), h[:cap(h)])
	fmt.Println(cap(i), len(i), i[:cap(i)])
	myutils.Br()
	h = append(h, 50)
	h[0] = 1
	i[0] = 2
	fmt.Println(cap(h), len(h), h[:cap(h)])
	fmt.Println(cap(i), len(i), i[:cap(i)])
	i = append(i, 51)
	i[0] = 2
	h[0] = 1
	fmt.Println(cap(h), len(h), h[:cap(h)])
	fmt.Println(cap(i), len(i), i[:cap(i)])
}

func copySliceNewArray() {
	myutils.H2("copy a slice (new array)")
	myutils.Wait()

	h := []int{22, 42, 124, 512}
	i := make([]int, len(h))
	copy(i[:], h[:])
	i[0] = 44
	h[0] = 33
	fmt.Println(cap(h), len(h), h[:cap(h)])
	fmt.Println(cap(i), len(i), i[:cap(i)])
}

func copyPartOfSlice() {
	myutils.H2("copy part of slice")
	myutils.Wait()

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("s", s)
	fmt.Println(":", s[:])
	fmt.Println("0:", s[0:])
	//fmt.Println(":len(s)", s[:len(s)])
	//fmt.Println("0:len(s)", s[0:len(s)])
	fmt.Printf("index 3-8         %v \n", s[3:9])
	fmt.Printf("first 13 (0-12)   %v \n", s[0:13])
	fmt.Printf("first             %v \n", s[0:1])
	fmt.Printf("last              %v \n", s[len(s)-1:])
	fmt.Printf("last 5            %v \n", s[len(s)-5:])
	fmt.Printf("index 11. und 12. %v \n", s[11:13])
	fmt.Printf("index 3           %v \n", s[3:4])
}
