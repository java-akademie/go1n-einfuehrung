package main

import (
	"github.com/java-akademie/myutils"
)

func doReceiver() {
	receiver1()
	receiver2()
	receiver3()
	receiver4()
}

func receiver1() {
	myutils.H2("receiver1")
	myutils.Wait()

	myutils.Comment("var mike=newPerson()")
	var mike = newPerson(43, "nike")
	mike.print()
	mike.updateName("mike")
	mike.print()
	myutils.ShowObject("mike", mike)

	myutils.Comment("joe:=person{44,'joe'}")
	joe := person{44, "joe"}
	joe.print()
	myutils.ShowObject("joe", joe)

	myutils.Comment("jim:=person{}")
	jim := person{}
	jim.print()
	jim.pnr = 45
	jim.name = "jim"
	jim.print()
	myutils.ShowObject("jim", jim)
}

func receiver2() {
	myutils.H2("receiver2")
	myutils.Wait()

	myutils.Comment("p1:=*new(person)")
	p1 := *new(person)
	p1.print()
	p1.init(111, "jack")
	p1.print()
	myutils.ShowObject("p1", p1)

	myutils.Comment("p2:=new(person)")
	p2 := new(person)
	p2.print()
	p2.init(222, "tony")
	p2.print()
	myutils.ShowObject("p2", p2)
}

func receiver3() {
	myutils.H2("receiver3")
	myutils.Wait()

	//p100 := person{100, "greg"}
	p100 := *new(person)
	p100.init(100, "greg")
	myutils.ShowObject("p100", p100)

	p200 := p100
	myutils.ShowObject("p200", p200)
	p200.init(200, "brad")
	myutils.Br()
	myutils.ShowObject("p100", p100)
	myutils.ShowObject("p200", p200)
}

func receiver4() {
	myutils.H2("receiver4")
	myutils.Wait()

	myutils.Comment("the Flintstones")
	fred := person{}
	fred.init(91, "fred")

	barney := new(person)
	barney.init(92, "barney")

	wilma := *new(person)
	wilma.init(93, "wilma")

	betty := newPerson(94, "betty")

	myutils.ShowObjects(fred, barney, wilma, betty)

}
