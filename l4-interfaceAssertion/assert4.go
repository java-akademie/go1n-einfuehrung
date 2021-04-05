package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

type dog struct{}
type cat struct{}
type bird struct{}
type butterfly struct{}

func (x dog) bruellen()  { fmt.Println("wauwau") }
func (x cat) bruellen()  { fmt.Println("miau") }
func (x bird) bruellen() { fmt.Println("pieps") }

type brueller interface{ bruellen() }

func doAssert4() {
	myutils.H2("assert4")
	d := dog{}
	c := cat{}
	b := bird{}
	s := butterfly{}
	bruellen(d)
	bruellen(c)
	bruellen(b)
	myutils.Hr()

	check(d)
	check(c)
	check(b)
	check(s)
}

func check(a interface{}) {
	checkBrueller(a)
	checkDog(a)
	checkCat(a)
	checkBird(a)
	fmt.Println()
}

func checkBird(a interface{}) {
	toCheck := a
	v, ok := toCheck.(bird)
	if !ok {
		fmt.Println("not a bird")
		return
	}
	fmt.Print("a bird: ")
	v.bruellen()
}

func checkCat(a interface{}) {
	toCheck := a
	v, ok := toCheck.(cat)
	if !ok {
		fmt.Println("not a cat")
		return
	}
	fmt.Print("a cat: ")
	v.bruellen()
}

func checkDog(a interface{}) {
	toCheck := a
	v, ok := toCheck.(dog)
	if !ok {
		fmt.Println("not a dog")
		return
	}
	fmt.Print("a dog: ")
	v.bruellen()
}

func checkBrueller(a interface{}) {
	toCheck := a
	v, ok := toCheck.(brueller)
	if !ok {
		fmt.Println("not a brueller")
		return
	}
	fmt.Print("a brueller: ")
	v.bruellen()
}

func bruellen(a brueller) {
	a.bruellen()
}
