package main

import (
	"github.com/java-akademie/myutils"
)

func withInterface() {
	myutils.H2("withInterface")

	k := kreis{2}
	r := rechteck{2, 3}

	show(k)
	show(r)
}

/**
 * alle Typen die eine show()-Methode haben
 * entsprechen dem Interface "geometricFigure"
 */
type geometricFigure interface{ show() }

func show(g geometricFigure) { g.show() }
