package main

import (
	"fmt"
	"math"

	"github.com/java-akademie/myutils"
)

func main() {
	myutils.H1("interfaceGeo.main")
	noInterface()
	withInterface()
}

/**
 * Kreis
 */
type kreis struct{ radius int }

func (g kreis) show() {
	fmt.Printf("%10s: %20s: %v \t\t Area: %v \n",
		"Kreis", "Radius",
		g.radius, float64(g.radius*g.radius)*math.Pi)
}

/**
 * Rechteck
 */
type rechteck struct {
	length int
	width  int
}

func (g rechteck) show() {
	fmt.Printf("%10s: %20s: %v/%v \t\t Area: %v \n",
		"Rechteck", "Laenge/Breite",
		g.length, g.width, g.length*g.width)
}
