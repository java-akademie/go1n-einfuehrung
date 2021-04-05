package main

import (
	"errors"
	"fmt"

	"github.com/java-akademie/myutils"
)

type divtype float64

func doErrors() {
	myutils.H2("Errors")
	doDivide(143, 12)
	doDivide(144, 0)
	doDivideIsPossible(123, 0)
	doDivideIsPossible(123, 2)
	myutils.Wait()
}
func doDivide(dnd divtype, dor divtype) {
	erg, err := divide(dnd, dor)
	if err != nil {
		fmt.Println(err, dnd, dor)
		return
	}
	fmt.Println(erg)
}
func doDivideIsPossible(dnd divtype, dor divtype) {
	_, err := divide(dnd, dor)
	if err != nil {
		fmt.Println("divide  is  not possible", dnd, dor)
		return
	}
	fmt.Println("divide would be possible", dnd, dor)
}
func divide(dnd divtype, dor divtype) (divtype, error) {
	if dor == 0 {
		return 0, errors.New("divisor may not be null")
	}
	return dnd / dor, nil
}
