package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func sliceLooping() {
	myutils.H2("slice Looping")

	s := []int{1881, 2232, 3543, 4564, 5535}
	s = append(s, 6884, 7745)
	fmt.Printf("Typ: %T, len %v, cap: %v, %v \n",
		s, len(s), cap(s), s)

	myutils.Comment("show index and value")
	for i, v := range s {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}

	myutils.Comment("show only value")
	for _, v := range s {
		fmt.Printf("value: %v\n", v)
	}
	myutils.Wait()
}
