package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func converting() {
	myutils.H2("converting1")
	initVariables()
	showAllVariables()
	converts1()
	converts2()
	converts3()
	converts4()
}

func converts1() {
	myutils.H2("converts1, OK")
	initVariables()

	j64 = int64(i16)
	fmt.Printf("%v = %v \n", i16, j64)

	j32 = int32(i8)
	fmt.Printf("%v = %v \n", j32, i8)

	text2 = string(r1)
	fmt.Printf("%v = %v \n", text2, r1)

	g64 = float64(h32)
	fmt.Printf("%v == %v is  %v \n", g64, h32, g64 == float64(h32))
	fmt.Printf("%v == %v is  %v \n", g64, h32, float32(g64) == h32)

	myutils.Wait()
}

func converts2() {
	myutils.H2("converts2, probably NOK")
	initVariables()

	j32 = int32(i64)
	fmt.Printf("%v = %v \n", j32, i64)

	j32 = int32(j64)
	fmt.Printf("%v = %v \n", j32, j64)

	u32 = uint32(i32)
	fmt.Printf("%v = %v \n", u32, i32)

	u32 = uint32(j64)
	fmt.Printf("%v = %v \n", u32, j64)

	g32 = float32(h64)
	fmt.Printf("%v == %v is  %v \n", g32, h64, g32 == float32(h64))
	fmt.Printf("%v == %v is  %v \n", g32, h64, float64(g32) == h64)

	myutils.Wait()
}

func converts3() {
	myutils.H2("converts3, your own attempts")
	initVariables()

	u32 = uint32(j16)
	fmt.Printf("%v = %v \n", u32, j16)

	u8 = uint8(j8)
	fmt.Printf("%v = %v \n", u8, j8)

	i64 = int64(u8)
	fmt.Printf("%v = %v \n", i64, u8)

	u16 = uint16(i64)
	fmt.Printf("%v = %v \n", u16, i64)

	myutils.Wait()
}

func converts4() {
	myutils.H2("converts4, apples, pears, fruits")
	initVariables()
	fruits = fruit(int(apples) + int(pears))
	myutils.ShowObjects(apples, pears, fruits)
	myutils.Wait()
}
