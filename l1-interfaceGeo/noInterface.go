package main

import (
	"github.com/java-akademie/myutils"
)

func noInterface() {
	myutils.H2("noInterface")

	k := kreis{2}
	r := rechteck{2, 3}

	showKreis(k)
	showRechteck(r)
}

func showKreis(g kreis) {
	g.show()
}
func showRechteck(g rechteck) {
	g.show()
}
