package main

import (
	"fmt"
	"sort"

	"github.com/java-akademie/myutils"
)

func slices4() {
	myutils.H1("slices4")
	myutils.Wait()

	n := make([]int, 0)

	for i := 1; i <= 10; i++ {
		n = append(n, myutils.GetRandom2(100, 999))
	}
	myutils.ShowSlice("n", n, cap(n), len(n))

	myutils.Comment("sort")
	sort.Ints(n)
	fmt.Println(n)

	myutils.Comment("shuffle")
	iShuffle(n)
	fmt.Println(n)

	myutils.Wait()

	myutils.Comment("numbers as cards")
	hand := n[:5] // the first five
	sort.Ints(hand)
	rest := n[5:] // the rest
	sort.Ints(rest)
	fmt.Println("hand: ", hand)
	fmt.Println("rest: ", rest)
	hand[0] = 0
	rest[0] = 0
	fmt.Println("hand: ", hand)
	fmt.Println("rest: ", rest)
	fmt.Println("all : ", n)
}

func iShuffle(s []int) {
	for i := range s {
		j := myutils.GetRandom2(1, len(s)) - 1
		s[i], s[j] = s[j], s[i]
	}
}
