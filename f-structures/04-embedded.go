package main

import (
	"fmt"

	"github.com/java-akademie/myutils"
)

func doEmbedded() {
	myutils.H2("doEmbedded")
	myutils.Wait()

	type emailZip struct {
		email   string
		zipCode int
	}

	type person struct {
		name    string
		count   int
		price   float64
		active  bool
		contact emailZip // {"" 0}
	}

	p1 := person{}
	fmt.Println(p1)
	fmt.Printf("%+v \n", p1)
	myutils.ShowObjects(p1)
	myutils.Br()

	joe := person{
		name:    "joe",
		contact: emailZip{email: "joe1@gmx.com", zipCode: 44999},
		count:   711,
		price:   995.90,
		active:  false,
	}
	fmt.Println(joe)
	fmt.Printf("%+v \n", joe)
	myutils.ShowObject("joe 1", joe)
	myutils.Br()

	joeNewContact := emailZip{email: "joe2@gmx.com", zipCode: 65777}
	joe.contact = joeNewContact

	fmt.Println(joe)
	fmt.Printf("%+v \n", joe)
	myutils.ShowObject("joe 2", joe)
}
