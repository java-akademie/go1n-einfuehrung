package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func basicLooping() {
	myutils.H2("basic Looping")
	basicLooping1()
	basicLooping2()
}

func basicLooping1() {

	myutils.Comment("for(i=1; i<=5; i++)")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	myutils.Comment("while(i<=5)")
	i1 := 1
	for i1 <= 5 {
		fmt.Println(i1)
		i1++
	}

	myutils.Comment("while(true)")
	i2 := 1
	for {
		if i2 > 5 {
			break
		}
		fmt.Println(i2)
		i2++
	}
	myutils.Wait()
}

func basicLooping2() {

	myutils.Comment("for (i=0; i < len(coll); i++")
	coll := []int{881, 232, 543}
	for i := 0; i < len(coll); i++ {
		fmt.Println(coll[i])
	}
	myutils.Wait()
}
