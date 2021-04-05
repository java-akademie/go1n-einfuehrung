package main

import (
	"fmt"
)

type person struct {
	pnr  int
	name string
}

func newPerson(pnr int, name string) person {
	return person{pnr, name}
}

func (p *person) updateName(newName string) {
	p.name = newName
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p *person) init(pnr int, name string) {
	p.pnr = pnr
	p.name = name
}
